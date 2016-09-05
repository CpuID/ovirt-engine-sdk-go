/*
Copyright (c) 2015-2016 Red Hat, Inc. / Nathan Sullivan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package org.ovirt.sdk.go;

import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toCollection;
import static java.util.stream.Collectors.toList;

import java.io.File;
import java.io.IOException;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.List;
import java.util.Optional;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.ListType;
import org.ovirt.api.metamodel.concepts.Locator;
import org.ovirt.api.metamodel.concepts.Method;
import org.ovirt.api.metamodel.concepts.Model;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.NameParser;
import org.ovirt.api.metamodel.concepts.Parameter;
import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.Service;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;
import org.ovirt.api.metamodel.tool.SchemaNames;

/**
 * This class is responsible for generating the classes that represent the services of the model.
 */
public class ServicesGenerator implements GoGenerator {
    // Well known method names:
    private static final Name ADD = NameParser.parseUsingCase("Add");
    private static final Name GET = NameParser.parseUsingCase("Get");
    private static final Name LIST = NameParser.parseUsingCase("List");
    private static final Name REMOVE = NameParser.parseUsingCase("Remove");
    private static final Name UPDATE = NameParser.parseUsingCase("Update");

    // The directory were the output will be generated:
    protected File out;

    // Reference to the objects used to generate the code:
    @Inject private GoNames goNames;
    @Inject private SchemaNames schemaNames;

    // The buffer used to generate the Go code:
    private GoBuffer buffer;

    /**
     * Set the directory were the output will be generated.
     */
    public void setOut(File newOut) {
        out = newOut;
    }

    public void generate(Model model) {
        // Calculate the file name:
        String fileName = goNames.getPackagePath() + "/services";
        buffer = new GoBuffer();
        buffer.setFileName(fileName);

        // Generate the source:
        generateSource(model);

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new IllegalStateException("Error writing services file \"" + fileName + "\"", exception);
        }
    }

    private void generateSource(Model model) {
        // We don't care too much about declaration order, but let's use inheritance order (if any), for consistency
        // with other SDK implementations.
        Deque<Service> pending = model.services()
            .sorted()
            .collect(toCollection(ArrayDeque::new));
        Deque<Service> sorted = new ArrayDeque<>(pending.size());
        while (!pending.isEmpty()) {
            Service current = pending.removeFirst();
            Service base = current.getBase();
            if (base == null || sorted.contains(base)) {
                sorted.addLast(current);
            }
            else {
                pending.addLast(current);
            }
        }

        sorted.forEach(this::generateService);
    }

    private void generateService(Service service) {
        // Begin struct type:
        generateTypeDeclaration(service);
        buffer.addLine("	connection *ovirt_http.Connection");
        buffer.addLine("	path string");
        buffer.addLine("}");
        buffer.addLine();

        // Generate the constructor:
        buffer.addComment("Creates a new implementation of the service.");
        GoName serviceName = goNames.getServiceName(service);
        buffer.addLine("func New%1$s(connection *ovirt_http.Connection, path string) *%1$s {", serviceName.getClassName(), serviceName.getClassName());
        buffer.addLine("  c.connection = connection");
        buffer.addLine("  c.path = path");
        buffer.addLine("}");
        buffer.addLine();

        // Generate the methods and locators:
        service.methods().sorted().forEach(this::generateMethod);
        service.locators().sorted().forEach(this::generateLocator);
        generatePathLocator(service);

        // Generate other methods that don't correspond to model methods or locators:
        generateToS(service);
    }

    private void generateMethod(Method method) {
        Name name = method.getName();
        if (ADD.equals(name)) {
            generateAddHttpPost(method);
        }
        else if (GET.equals(name) || LIST.equals(name)) {
            generateHttpGet(method);
        }
        else if (REMOVE.equals(name)) {
            generateHttpDelete(method);
        }
        else if (UPDATE.equals(name)) {
            generateHttpPut(method);
        }
        else {
            generateActionHttpPost(method);
        }
    }

    private void generateAddHttpPost(Method method) {
        // Get the main parameter:
        Parameter parameter = method.parameters()
            .filter(x -> x.isIn() && x.isOut())
            .findFirst()
            .orElse(null);

        // Begin function:
        Name methodName = method.getName();
        Type parameterType = parameter.getType();
        Name parameterName = parameter.getName();
        String arg = goNames.getFuncStyleName(parameterName);
        String doc = method.getDoc();
        if (doc == null) {
            doc = String.format("Adds a new `%1$s`.", arg);
        }
        buffer.addComment();
        buffer.addComment(doc);
        buffer.addLine("func (s *%1%s) %2$s(opts map[string]string)", method.getClass(), goNames.getFuncStyleName(methodName));

        // Body:
        generateConvertLiteral(parameterType, arg);
        buffer.addLine("request = Request.new(:method => :POST, :path => @path)");
        generateWriteRequestBody(parameter, arg);
        buffer.addLine("response = @connection.send(request)");
        buffer.addLine("case response.code");
        buffer.addLine("when 201, 202");
        generateReturnResponseBody(parameter);
        buffer.addLine("else");
        buffer.addLine(  "check_fault(response)");
        buffer.addLine("end");

        // End method:
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateActionHttpPost(Method method) {
        // Begin method:
        Name methodName = method.getName();
        String actionName = goNames.getFuncStyleName(methodName);
        String doc = method.getDoc();
        if (doc == null) {
            doc = String.format("Executes the `%1$s` method.", actionName);
        }
        buffer.addComment();
        buffer.addComment(doc);
        buffer.addComment();
        buffer.addLine("def %1$s(opts = {})", actionName);

        // Generate the method:
        buffer.addLine("action = Action.new(opts)");
        buffer.addLine("writer = XmlWriter.new(nil, true)");
        buffer.addLine("ActionWriter.write_one(action, writer)");
        buffer.addLine("body = writer.string");
        buffer.addLine("writer.close");
        buffer.addLine("request = Request.new({");
        buffer.addLine(  ":method => :POST,");
        buffer.addLine(  ":path => \"#{@path}/%1$s\",", getPath(methodName));
        buffer.addLine(  ":body => body,");
        buffer.addLine("})");
        buffer.addLine("response = @connection.send(request)");
        buffer.addLine("case response.code");
        buffer.addLine("when 200");
        buffer.addLine(  "action = check_action(response)");
        method.parameters()
            .filter(Parameter::isOut)
            .findFirst()
            .ifPresent(this::generateActionResponse);
        buffer.addLine("else");
        buffer.addLine(  "check_fault(response)");
        buffer.addLine("end");

        // End method:
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateActionResponse(Parameter parameter) {
        buffer.addLine("return action.%1$s", goNames.getFuncStyleName(parameter.getName()));
    }

    private void generateHttpGet(Method method) {
        // Get input and output parameters:
        List<Parameter> inParameters = method.parameters()
            .filter(Parameter::isIn)
            .sorted()
            .collect(toList());
        List<Parameter> outParameters = method.parameters()
            .filter(Parameter::isOut)
            .sorted()
            .collect(toList());

        // Get the first output parameter:
        Parameter mainParameter = outParameters.stream()
            .findFirst()
            .orElse(null);

        // Begin method:
        buffer.addComment();
        String doc = method.getDoc();
        if (doc == null) {
            doc = "Returns the representation of the object managed by this service.";
        }
        buffer.addComment(doc);
        Name methodName = method.getName();
        buffer.addLine("def %1$s(opts = {})", goNames.getFuncStyleName(methodName));

        // Generate the input parameters:
        buffer.addLine("query = {}");
        inParameters.forEach(this::generateUrlParameter);

        // Body:
        buffer.addLine("request = Request.new(:method => :GET, :path => @path, :query => query)");
        buffer.addLine("response = @connection.send(request)");
        buffer.addLine("case response.code");
        buffer.addLine("when 200");
        generateReturnResponseBody(mainParameter);
        buffer.addLine("else");
        buffer.addLine(  "check_fault(response)");
        buffer.addLine("end");

        // End method:
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateHttpPut(Method method) {
        // Get the main parameter:
        Parameter parameter = method.parameters()
            .filter(x -> x.isIn() && x.isOut())
            .findFirst()
            .orElse(null);

        // Begin method:
        Name methodName = method.getName();
        Name parameterName = parameter.getName();
        Type parameterType = parameter.getType();
        String arg = goNames.getFuncStyleName(parameterName);
        String doc = method.getDoc();
        if (doc == null) {
            doc = "Updates the object managed by this service.";
        }
        buffer.addComment();
        buffer.addComment(doc);
        buffer.addComment();
        buffer.addLine("def %1$s(%2$s)", goNames.getFuncStyleName(methodName), arg);

        // Body:
        generateConvertLiteral(parameterType, arg);
        buffer.addLine("request = Request.new(:method => :PUT, :path => @path)");
        generateWriteRequestBody(parameter, arg);
        buffer.addLine("response = @connection.send(request)");
        buffer.addLine("case response.code");
        buffer.addLine("when 200");
        generateReturnResponseBody(parameter);
        buffer.addLine(  "return result");
        buffer.addLine("else");
        buffer.addLine(  "check_fault(response)");
        buffer.addLine("end");

        // End method:
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateConvertLiteral(Type type, String variable) {
        if (type instanceof StructType) {
            buffer.addLine("if %1$s.is_a?(Hash)", variable);
            buffer.addLine(  "%1$s = %2$s.new(%1$s)", variable, goNames.getTypeName(type));
            buffer.addLine("end");
        }
        else if (type instanceof ListType) {
            ListType listType = (ListType) type;
            Type elementType = listType.getElementType();
            buffer.addLine("if %1$s.is_a?(Array)", variable);
            buffer.addLine(  "%1$s = List.new(%1$s)", variable);
            buffer.addLine(  "%1$s.each_with_index do |value, index|", variable);
            buffer.addLine(    "if value.is_a?(Hash)");
            buffer.addLine(      "%1$s[index] = %2$s.new(value)", variable, goNames.getTypeName(elementType));
            buffer.addLine(    "end");
            buffer.addLine(  "end");
            buffer.addLine("end");
        }
    }

    private void generateWriteRequestBody(Parameter parameter, String variable) {
        Type type = parameter.getType();
        buffer.addLine("begin");
        buffer.addLine(  "writer = XmlWriter.new(nil, true)");
        if (type instanceof StructType) {
            GoName writer = goNames.getWriterName(type);
            buffer.addLine("%1$s.write_one(%2$s, writer)", writer.getClassName(), variable);
        }
        else if (type instanceof ListType) {
            ListType listType = (ListType) type;
            Type elementType = listType.getElementType();
            GoName writer = goNames.getWriterName(elementType);
            buffer.addLine("%1$s.write_many(%2$s, writer)", writer.getClassName(), variable);
        }
        buffer.addLine(  "request.body = writer.string");
        buffer.addLine("ensure");
        buffer.addLine(  "writer.close");
        buffer.addLine("end");
    }

    private void generateReturnResponseBody(Parameter parameter) {
        Type type = parameter.getType();
        buffer.addLine("begin");
        buffer.addLine(  "reader = XmlReader.new(response.body)");
        if (type instanceof StructType) {
            GoName reader = goNames.getReaderName(type);
            buffer.addLine("return %1$s.read_one(reader)", reader.getClassName());
        }
        else if (type instanceof ListType) {
            ListType listType = (ListType) type;
            Type elementType = listType.getElementType();
            GoName reader = goNames.getReaderName(elementType);
            buffer.addLine("return %1$s.read_many(reader)", reader.getClassName());
        }
        buffer.addLine("ensure");
        buffer.addLine(  "reader.close");
        buffer.addLine("end");
    }

    private void generateHttpDelete(Method method) {
        // Get input parameters:
        List<Parameter> inParameters = method.parameters()
            .filter(Parameter::isIn)
            .sorted()
            .collect(toList());

        // Begin method:
        Name name = method.getName();
        String doc = method.getDoc();
        if (doc == null) {
            doc = "Deletes the object managed by this service.";
        }
        buffer.addComment();
        buffer.addComment(doc);
        buffer.addComment();
        buffer.addComment("@param opts [Hash] Additional options.");
        buffer.addLine("def %1$s(opts = {})", goNames.getFuncStyleName(name));

        // Generate the input parameters:
        buffer.addLine("query = {}");
        inParameters.forEach(this::generateUrlParameter);

        // Generate the method:
        buffer.addLine(  "request = Request.new(:method => :DELETE, :path => @path, :query => query)");
        buffer.addLine(  "response = @connection.send(request)");
        buffer.addLine(  "unless response.code == 200");
        buffer.addLine(    "check_fault(response)");
        buffer.addLine(  "end");
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateUrlParameter(Parameter parameter) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        String symbol = goNames.getFuncStyleName(name);
        String tag = schemaNames.getSchemaTagName(name);
        buffer.addLine("value = opts[:%1$s]", symbol);
        buffer.addLine("unless value.nil?");
        if (type instanceof PrimitiveType) {
            Model model = type.getModel();
            if (type == model.getBooleanType()) {
                buffer.addLine("value = Writer.render_boolean(value)", tag);
            }
            else if (type == model.getIntegerType()) {
                buffer.addLine("value = Writer.render_integer(value)", tag);
            }
            else if (type == model.getDecimalType()) {
                buffer.addLine("value = Writer.render_decimal(value)", tag);
            }
            else if (type == model.getDateType()) {
                buffer.addLine("value = Writer.render_date(value)", tag);
            }
        }
        buffer.addLine(  "query['%1$s'] = value", tag);
        buffer.addLine("end");
    }

    private void generateToS(Service service) {
        GoName serviceName = goNames.getServiceName(service);
        buffer.addComment();
        buffer.addComment("Returns an string representation of this service.");
        buffer.addComment();
        buffer.addComment("@return [String]");
        buffer.addComment();
        buffer.addLine("def to_s");
        buffer.addLine(  "return \"#<#{%1$s}:#{@path}>\"", serviceName.getClassName());
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateLocator(Locator locator) {
        Parameter parameter = locator.getParameters().stream().findFirst().orElse(null);
        if (parameter != null) {
            generateLocatorWithParameters(locator);
        }
        else {
            generateLocatorWithoutParameters(locator);
        }
    }

    private void generateLocatorWithParameters(Locator locator) {
        Parameter parameter = locator.parameters().findFirst().get();
        String methodName = goNames.getFuncStyleName(locator.getName());
        String argName = goNames.getFuncStyleName(parameter.getName());
        GoName serviceName = goNames.getServiceName(locator.getService());
        String doc = locator.getDoc();
        if (doc == null) {
            doc = String.format("Locates the `%1$s` service.", methodName);
        }
        buffer.addComment();
        buffer.addComment(doc);
        buffer.addComment();
        buffer.addComment("@param %1$s [String] The identifier of the `%2$s`.", argName, methodName);
        buffer.addComment();
        buffer.addComment("@return [%1$s] A reference to the `%2$s` service.", serviceName.getClassName(), methodName);
        buffer.addComment();
        buffer.addLine("def %1$s_service(%2$s)", methodName, argName);
        buffer.addLine(  "return %1$s.new(@connection, \"#{@path}/#{%2$s}\")", serviceName.getClassName(), argName);
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateLocatorWithoutParameters(Locator locator) {
        String methodName = goNames.getFuncStyleName(locator.getName());
        String urlSegment = getPath(locator.getName());
        GoName serviceName = goNames.getServiceName(locator.getService());
        String doc = locator.getDoc();
        if (doc == null) {
            doc = String.format("Locates the `%1$s` service.", methodName);
        }
        buffer.addComment();
        buffer.addComment(doc);
        buffer.addComment();
        buffer.addComment("@return [%1$s] A reference to `%2$s` service.", serviceName.getClassName(), methodName);
        buffer.addLine("def %1$s_service", methodName);
        buffer.addLine(  "return %1$s.new(@connection, \"#{@path}/%2$s\")", serviceName.getClassName(), urlSegment);
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generatePathLocator(Service service) {
        // Begin method:
        buffer.addComment();
        buffer.addComment("Locates the service corresponding to the given path.");
        buffer.addComment();
        buffer.addComment("@param path [String] The path of the service.");
        buffer.addComment();
        buffer.addComment("@return [Service] A reference to the service.");
        buffer.addComment();
        buffer.addLine("def service(path)");
        buffer.addLine(  "if path.nil? || path == ''");
        buffer.addLine(    "return self");
        buffer.addLine(  "end");

        // Generate the code that checks if the path corresponds to any of the locators without parameters:
        service.locators().filter(x -> x.getParameters().isEmpty()).sorted().forEach(locator -> {
            Name name = locator.getName();
            String segment = getPath(name);
            buffer.addLine("if path == '%1$s'", segment);
            buffer.addLine(  "return %1$s_service", goNames.getFuncStyleName(name));
            buffer.addLine("end");
            buffer.addLine("if path.start_with?('%1$s/')", segment);
            buffer.addLine(
                "return %1$s_service.service(path[%2$d..-1])",
                goNames.getFuncStyleName(name),
                segment.length() + 1
            );
            buffer.addLine("end");
        });

        // If the path doesn't correspond to a locator without parameters, then it will correspond to the locator
        // with parameters, otherwise it is an error:
        Optional<Locator> optional = service.locators().filter(x -> !x.getParameters().isEmpty()).findAny();
        if (optional.isPresent()) {
            Locator locator = optional.get();
            Name name = locator.getName();
            buffer.addLine("index = path.index('/')");
            buffer.addLine("if index.nil?");
            buffer.addLine(  "return %1$s_service(path)", goNames.getFuncStyleName(name));
            buffer.addLine("end");
            buffer.addLine(
                "return %1$s_service(path[0..(index - 1)]).service(path[(index +1)..-1])",
                goNames.getFuncStyleName(name)
            );
        }
        else {
            buffer.addLine("raise Error.new(\"The path \\\"#{path}\\\" doesn't correspond to any service\")");
        }

        // End method:
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateTypeDeclaration(Service service) {
        GoName serviceName = goNames.getServiceName(service);
        buffer.addLine("type %1$s struct{", serviceName.getClassName());
    }

    private String getPath(Name name) {
        return name.words().map(String::toLowerCase).collect(joining());
    }
}

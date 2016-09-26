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

// TODO: vet all @api private references in ruby codebase, verify that function style is private for all use cases.

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
        // Generate the source:
        generateSource(model);
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
        GoName serviceName = goNames.getServiceName(service);

        // Calculate the file name:
        buffer = new GoBuffer();
        buffer.setFileName(serviceName.getFileName());
        buffer.setPackageName("services");

        // Set our imports:
        String repoSdkUrl = goNames.repoSdkUrl();
        buffer.addImport("fmt");
        buffer.addImport(repoSdkUrl + "/http");
        buffer.addImport(repoSdkUrl + "/ov_xml");
        buffer.addImport(repoSdkUrl + "/readers");
        buffer.addImport(repoSdkUrl + "/types");
        buffer.addImport(repoSdkUrl + "/writers");
        buffer.addImport(repoSdkUrl + "/version");

        // Begin struct type:
        generateTypeDeclaration(service);
        buffer.addLine("	connection *ovirt_http.Connection");
        buffer.addLine("	path string");
        buffer.addLine("}");
        buffer.addLine();

        // Generate the constructor:
        buffer.addComment("Creates a new implementation of the service.");
        buffer.addLine("func New%1$s(connection *ovirt_http.Connection, path string) *%1$s {", serviceName.getClassName(), serviceName.getClassName());
        buffer.addLine("  c.connection = connection");
        buffer.addLine("  c.path = path");
        buffer.addLine("}");
        buffer.addLine();

        // Generate the methods and locators:
        service.methods().sorted().forEach(this::generateMethod);
        service.locators().sorted().forEach(this::generateLocator);
        // TODO: how relevant is this OO hierarchial lookup in a non-OO golang implementation...? commented out for now, see how we go.
        //generatePathLocator(service);

        // Generate other functions that don't correspond to model functions or locators:
        generateToS(service);

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new IllegalStateException("Error writing services file \"" + serviceName.getFileName() + "\"", exception);
        }
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
        // TODO: review, parameterName vs methodName?
        String arg = goNames.getPublicStyleName(parameterName);
        String actionName = goNames.getPublicStyleName(methodName);
        String serviceTypeName = getTypeDeclaration(method.getDeclaringService());
        String doc = method.getDoc();
        if (doc == null) {
            doc = String.format("Adds a new `%1$s`.", arg);
        }
        buffer.addComment(doc);
        // TODO: fix input/outputs of below.
        // TODO: %2 should be the same parameterType as %3
        // TODO: %3 should not be serviceTypeName, it should be the parameter type. verify the format of parameterType (translate?), needs to be types.Type
        buffer.addLine("func (s *%1$s) %2$s(somearg TODO, opts map[string]string) *%3$s {", serviceTypeName, actionName, parameterType);

        // Body:
        buffer.addLine();
        generateConvertLiteral(parameterType, arg);
        buffer.addLine();
        buffer.addLine("  request := types.Request{");
        buffer.addLine("    Method: \"GET\",");
        buffer.addLine("    Path: s.path,");
        buffer.addLine("  }");
        generateWriteRequestBody(parameter, arg);
        buffer.addLine("  response, err := s.connection.Send(&request)");
        buffer.addLine("  if err != nil {");
        buffer.addLine("    // TODO: error handling. call checkFault(&response)?");
        buffer.addLine("  }");
        buffer.addLine("  if *response.Code == 201 || *response.Code == 202 {");
        generateReturnResponseBody(parameter);
        buffer.addLine("  } else {");
        buffer.addLine("    checkFault(&response)");
        buffer.addLine("  }");

        // End function:
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateActionHttpPost(Method method) {
        // Begin method:
        Name methodName = method.getName();
        String actionName = goNames.getPublicStyleName(methodName);
        String serviceTypeName = getTypeDeclaration(method.getDeclaringService());
        String doc = method.getDoc();
        if (doc == null) {
            doc = String.format("Executes the `%1$s` method.", actionName);
        }
        buffer.addComment(doc);
        // TODO: %3 should not be serviceTypeName, it should be the parameter type. verify the format of parameterType (translate?), needs to be types.Type
        buffer.addLine("func (s *%1$s) %2$s(opts map[string]string) *%3$s {", serviceTypeName, actionName, serviceTypeName);

        // Generate the function:
        buffer.addLine("  action := types.NewAction(opts)");
        buffer.addLine("  writer := ov_xml.NewXmlWriter(nil, true)");
        // TODO: error handling on writers?
        buffer.addLine("  writers.ActionWriterWriteOne(action, writer)");
        buffer.addLine("  body := writer.String()");
        buffer.addLine("  writer.Close()");
        buffer.addLine("  request := http.Request{");
        buffer.addLine("    Method: \"POST\",");
        buffer.addLine("    Path: fmt.Sprintf(\"%%s/%1$s\", s.path),", getPath(methodName));
        buffer.addLine("    Body: body,");
        buffer.addLine("  }");
        buffer.addLine("  response := s.Connection.send(&request)");
        buffer.addLine("  if *response.Code == 200 {");
        buffer.addLine("    action = checkAction(&response)");
        method.parameters()
            .filter(Parameter::isOut)
            .findFirst()
            .ifPresent(this::generateActionResponse);
        buffer.addLine("  } else {");
        buffer.addLine("    checkFault(&response)");
        buffer.addLine("  }");

        // End function:
        buffer.addLine("}");
        buffer.addLine();
    }

    // TODO: review
    private void generateActionResponse(Parameter parameter) {
        buffer.addLine("return action.%1$s", goNames.getPublicStyleName(parameter.getName()));
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
        String doc = method.getDoc();
        if (doc == null) {
            doc = "Returns the representation of the object managed by this service.";
        }
        buffer.addComment(doc);
        Name methodName = method.getName();
        String serviceTypeName = getTypeDeclaration(method.getDeclaringService());
        // TODO: %3 should not be serviceTypeName, it should be the parameter type. verify the format of parameterType (translate?), needs to be types.Type
        buffer.addLine("func (s *%1$s) %2$s(opts map[string]string) *%3$s {", serviceTypeName, goNames.getPublicStyleName(methodName), serviceTypeName);

        // Generate the input parameters:
        buffer.addLine("  query := make(map[string]string)");
        inParameters.forEach(this::generateUrlParameter);

        // Body:
        buffer.addLine();
        buffer.addLine("  request := types.Request{");
        buffer.addLine("    Method: \"GET\",");
        buffer.addLine("    Path: s.path,");
        buffer.addLine("    Query: query,");
        buffer.addLine("  }");
        buffer.addLine("  response, err := s.connection.Send(&request)");
        buffer.addLine("  if err != nil {");
        buffer.addLine("    // TODO: error handling. call checkFault(&response)?");
        buffer.addLine("  }");
        buffer.addLine("  if *response.Code == 200 {");
        generateReturnResponseBody(mainParameter);
        buffer.addLine("  } else {");
        buffer.addLine("    checkFault(&response)");
        buffer.addLine("  }");

        // End method:
        buffer.addLine("}");
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
        String arg = goNames.getPublicStyleName(parameterName);
        String doc = method.getDoc();
        if (doc == null) {
            doc = "Updates the object managed by this service.";
        }
        buffer.addComment(doc);
        // TODO: fix function args, needs to be named (is i sufficient? or use something more elegant?). it also needs to come from types.arg
        // TODO: add some return type, bool or error most likely
        buffer.addLine("func (s *%1$s) %2$s(i %3$s)", getTypeDeclaration(method.getDeclaringService()), goNames.getPublicStyleName(methodName), arg);

        // Body:
        generateConvertLiteral(parameterType, arg);
        buffer.addLine();
        buffer.addLine("  request := types.Request{");
        buffer.addLine("    Method: \"PUT\",");
        buffer.addLine("    Path: s.path,");
        buffer.addLine("  }");
        generateWriteRequestBody(parameter, arg);
        buffer.addLine("  response, err := s.connection.Send(&request)");
        buffer.addLine("  if err != nil {");
        buffer.addLine("    // TODO: error handling. call checkFault(&response)?");
        buffer.addLine("  }");
        buffer.addLine("  if *response.Code == 200 {");
        generateReturnResponseBody(parameter);
        buffer.addLine("  } else {");
        buffer.addLine("    checkFault(&response)");
        buffer.addLine("  }");

        // End method:
        buffer.addLine("}");
        buffer.addLine();
    }

    // TODO: review
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

    // TODO: review
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

    // TODO: review
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
            doc = "Deletes the object managed by this service. If Async = true, error will always be nil.";
        }
        buffer.addComment(doc);
        buffer.addLine("func (s *%1$s) %2$s(opts RemoveOpts) error {", getTypeDeclaration(method.getDeclaringService()), goNames.getPublicStyleName(name));

        // Generate the input parameters:
        buffer.addLine("  query := make(map[string]string)");
        inParameters.forEach(this::generateUrlParameter);

        // Generate the method:
        buffer.addLine();
        buffer.addLine("  request := types.Request{");
        buffer.addLine("    Method: \"DELETE\",");
        buffer.addLine("    Path: s.path,");
        buffer.addLine("    Query: query,");
        buffer.addLine("  }");
        buffer.addLine("  response, err := s.connection.Send(&request)");
        buffer.addLine("  if err != nil {");
        buffer.addLine("    // TODO: error handling. call checkFault(&response)?");
        buffer.addLine("  }");
        buffer.addLine("  if *response.Code != 200 {");
        buffer.addLine("    checkFault(&response)");
        buffer.addLine("  }");
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateUrlParameter(Parameter parameter) {
        Type type = parameter.getType();
        Name name = parameter.getName();
        // TODO: use a nicer function than this...?
        String symbol = goNames.getPublicStyleName(name);
        String tag = schemaNames.getSchemaTagName(name);
        // TODO: is this suitable for all input types...? or we need to compare against != "" sometimes? or len(opts.%1$s) != 0?
        buffer.addLine("  if opts.%1$s != nil {", symbol);
        if (type instanceof PrimitiveType) {
            Model model = type.getModel();
            if (type == model.getBooleanType()) {
                buffer.addLine("    value := writers.RenderBoolean(opts.%1$s)", symbol);
            }
            else if (type == model.getIntegerType()) {
                buffer.addLine("    value := writers.RenderInteger(opts.%1$s)", symbol);
            }
            else if (type == model.getDecimalType()) {
                buffer.addLine("    value := writers.RenderDecimal(opts.%1$s)", symbol);
            }
            else if (type == model.getDateType()) {
                buffer.addLine("    value := writers.RenderDate(opts.%1$s)", symbol);
            }
        }
        buffer.addLine("    query[\"%1$s\"] = value", tag);
        buffer.addLine("  }");
    }

    private void generateToS(Service service) {
        GoName serviceName = goNames.getServiceName(service);
        buffer.addComment("Returns an string representation of this service.");
        buffer.addLine("func (s *%1$s) ToString () string {", serviceName.getClassName());
        buffer.addLine("  return fmt.Sprintf(\"#<%1$s:%%s>\", s.path)", serviceName.getClassName());
        buffer.addLine("}");
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
        String methodName = goNames.getPublicStyleName(locator.getName());
        String argName = goNames.getPublicStyleName(parameter.getName());
        GoName serviceName = goNames.getServiceName(locator.getService());

        String doc = locator.getDoc();
        if (doc == null) {
            doc = String.format("Locates the `%1$s` service.", methodName);
        }
        buffer.addComment(doc);
        buffer.addComment("Takes a string input of the identifier of the `%1$s`.", methodName);
        buffer.addComment("Returns a reference to the `%1$s` service.", methodName);
        // TODO: the %1$s is incorrect here, needs to be the type of the service its attached to, not the reference service.
        buffer.addLine("func (s *%1$s) %2$sService(%3$s) *%4$s {", getTypeDeclaration(locator.getService()), methodName, argName, serviceName.getClassName());
        buffer.addLine("  return New%1$s(s.connection, fmt.Sprintf(\"%%s/%2$s\", s.path))", serviceName.getClassName(), argName);
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateLocatorWithoutParameters(Locator locator) {
        String methodName = goNames.getPublicStyleName(locator.getName());
        String urlSegment = getPath(locator.getName());
        GoName serviceName = goNames.getServiceName(locator.getService());
        String doc = locator.getDoc();
        if (doc == null) {
            doc = String.format("Locates the `%1$s` service.", methodName);
        }
        buffer.addComment(doc);
        buffer.addComment("Returns a reference to the `%1$s` service.", methodName);
        // TODO: the %1$s is incorrect here, needs to be the type of the service its attached to, not the reference service.
        buffer.addLine("func (s *%1$s) %2$sService () *%3$s {", getTypeDeclaration(locator.getService()), methodName, serviceName.getClassName());
        buffer.addLine("  return New%1$s(s.connection, fmt.Sprintf(\"%%s/%2$s\", s.path))", serviceName.getClassName(), urlSegment);
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generatePathLocator(Service service) {
        // Begin method:
        buffer.addComment("Locates the service corresponding to the given path.");
        buffer.addComment("Takes a string input of the path of the service.");
        buffer.addComment("Returns a reference to the service.");
        buffer.addLine("func (s %1$s) Service(path string) *Service {");
        buffer.addLine("  if len(path) == 0 {");
        buffer.addLine("    return s");
        buffer.addLine("  }");

        // Generate the code that checks if the path corresponds to any of the locators without parameters:
        // TODO: review
        service.locators().filter(x -> x.getParameters().isEmpty()).sorted().forEach(locator -> {
            Name name = locator.getName();
            String segment = getPath(name);
            buffer.addLine("if path == '%1$s'", segment);
            buffer.addLine(  "return %1$s_service", goNames.getPublicStyleName(name));
            buffer.addLine("end");
            buffer.addLine("if path.start_with?('%1$s/')", segment);
            buffer.addLine(
                "return %1$s_service.service(path[%2$d..-1])",
                goNames.getPublicStyleName(name),
                segment.length() + 1
            );
            buffer.addLine("end");
        });

        // If the path doesn't correspond to a locator without parameters, then it will correspond to the locator
        // with parameters, otherwise it is an error:
        // TODO: review
        Optional<Locator> optional = service.locators().filter(x -> !x.getParameters().isEmpty()).findAny();
        if (optional.isPresent()) {
            Locator locator = optional.get();
            Name name = locator.getName();
            buffer.addLine("index = path.index('/')");
            buffer.addLine("if index.nil?");
            buffer.addLine(  "return %1$s_service(path)", goNames.getPublicStyleName(name));
            buffer.addLine("end");
            buffer.addLine(
                "return %1$s_service(path[0..(index - 1)]).service(path[(index +1)..-1])",
                goNames.getPublicStyleName(name)
            );
        }
        else {
            buffer.addLine("raise Error.new(\"The path \\\"#{path}\\\" doesn't correspond to any service\")");
        }

        // End method:
        buffer.addLine("}");
        buffer.addLine();
    }

    private String getTypeDeclaration(Service service) {
        GoName serviceName = goNames.getServiceName(service);
        return serviceName.getClassName();
    }

    private void generateTypeDeclaration(Service service) {
        buffer.addLine("type %1$s struct{", getTypeDeclaration(service));
    }

    private String getPath(Name name) {
        return name.words().map(String::toLowerCase).collect(joining());
    }
}

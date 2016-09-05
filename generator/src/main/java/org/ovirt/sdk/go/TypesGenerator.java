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

import static java.util.stream.Collectors.toCollection;

import java.io.File;
import java.io.IOException;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;
import java.util.stream.Stream;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.EnumType;
import org.ovirt.api.metamodel.concepts.EnumValue;
import org.ovirt.api.metamodel.concepts.ListType;
import org.ovirt.api.metamodel.concepts.Model;
import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.PrimitiveType;
import org.ovirt.api.metamodel.concepts.StructMember;
import org.ovirt.api.metamodel.concepts.StructType;
import org.ovirt.api.metamodel.concepts.Type;

/**
 * This class is responsible for generating the classes that represent the types of the model.
 */
public class TypesGenerator implements GoGenerator {
    // The directory were the output will be generated:
    protected File out;

    // Reference to the objects used to generate the code:
    @Inject private GoNames goNames;

    // The buffer used to generate the Go code:
    private GoBuffer buffer;

    public void setOut(File newOut) {
        out = newOut;
    }

    public void generate(Model model) {
        // Calculate the file name:
        String fileName = goNames.getPackagePath() + "/types/types";
        buffer = new GoBuffer();
        buffer.setFileName(fileName);
        buffer.setPackageName("types");

        // Generate the source:
        generateStructs(model);
        generateEnums(model);

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new IllegalStateException("Error writing types file \"" + fileName + "\"", exception);
        }
    }

    private void generateStructs(Model model) {
        // The declarations of the types need to appear in inheritance order, otherwise some symbols won't be
        // defined and that will produce errors. To order them correctly we need first to sort them by name, and
        // then sort again so that bases are before extensions.
        Deque<StructType> pending = model.types()
            .filter(StructType.class::isInstance)
            .map(StructType.class::cast)
            .sorted()
            .collect(toCollection(ArrayDeque::new));
        Deque<StructType> sorted = new ArrayDeque<>(pending.size());
        while (!pending.isEmpty()) {
            StructType current = pending.removeFirst();
            StructType base = (StructType) current.getBase();
            if (base == null || sorted.contains(base)) {
                sorted.addLast(current);
            }
            else {
                pending.addLast(current);
            }
        }

        // Generate the complete declarations, using the same order:
        sorted.forEach(this::generateStruct);
    }

    private void generateStruct(StructType type) {
        // Begin strcut type:
        generateTypeDeclaration(type);
        buffer.addLine();

        // Attributes and links:
        List<StructMember> members = new ArrayList<>();
        members.addAll(type.getAttributes());
        members.addAll(type.getLinks());
        members.stream().sorted().forEach(this::generateMember);

        // Constructor with a named parameter for each attribute:
        GoName typeName = goNames.getTypeName(type);
        buffer.addComment();
        buffer.addComment("Creates a new instance of the {%1$s} class.", typeName.getClassName());
        buffer.addComment();
        buffer.addComment("@param opts [Hash] A hash containing the attributes of the object. The keys of the hash ");
        buffer.addComment("  should be symbols corresponding to the names of the attributes. The values of the hash ");
        buffer.addComment("  should be the values of the attributes.");
        buffer.addComment();
        members.stream().sorted().forEach(member -> {
            Type memberType = member.getType();
            Name memberName = member.getName();
            String docName = goNames.getFuncStyleName(memberName);
            if (memberType instanceof PrimitiveType || memberType instanceof EnumType) {
                buffer.addComment("@option opts :%1$s The value of attribute `%1$s`.", docName);
                buffer.addComment();
            }
            else if (memberType instanceof StructType) {
                buffer.addComment("@option opts [Hash] :%1$s The value of attribute `%1$s`.", docName);
                buffer.addComment();
            }
            else if (memberType instanceof ListType) {
                buffer.addComment("@option opts [Array<Hash>] :%1$s The values of attribute `%1$s`.", docName);
                buffer.addComment();
            }
        });
        buffer.addComment();
        buffer.addLine("def initialize(opts = {})");
        buffer.addLine(  "super(opts)");
        members.stream().sorted().forEach(member -> {
            String memberName = goNames.getFuncStyleName(member.getName());
            buffer.addLine("self.%1$s = opts[:%1$s]", memberName);
        });
        buffer.addLine("end");
        buffer.addLine();

        // End type:
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateMember(StructMember member) {
        generateGetter(member);
        generateSetter(member);
    }

    private void generateGetter(StructMember member) {
        Name name = member.getName();
        Type type = member.getType();
        String property = goNames.getFuncStyleName(name);
        buffer.addComment();
        buffer.addComment("Returns the value of the `%1$s` attribute.", property);
        buffer.addComment();
        buffer.addLine("def %1$s", property);
        buffer.addLine(  "return @%1$s", property);
        buffer.addLine("end");
        buffer.addLine();
    }

    private void generateSetter(StructMember member) {
        Name name = member.getName();
        Type type = member.getType();
        String property = goNames.getFuncStyleName(name);
        buffer.addComment();
        buffer.addComment("Sets the value of the `%1$s` attribute.", property);
        buffer.addComment();
        if (type instanceof PrimitiveType || type instanceof EnumType) {
            buffer.addLine("def %1$s=(value)", property);
            buffer.addLine(  "@%1$s = value", property);
            buffer.addLine("end");
        }
        else if (type instanceof StructType) {
            GoName typeName = goNames.getTypeName(type);
            buffer.addComment("@param value [Hash]");
            buffer.addComment();
            buffer.addComment("The `value` parameter can be an instance of {%1$s} or a hash.", typeName);
            buffer.addComment("If it is a hash then a new instance will be created passing the hash as the ");
            buffer.addComment("`opts` parameter to the constructor.");
            buffer.addComment();
            buffer.addLine("def %1$s=(value)", property);
            buffer.addLine(  "if value.is_a?(Hash)");
            buffer.addLine(    "value = %1$s.new(value)", typeName.getClassName());
            buffer.addLine(  "end");
            buffer.addLine(  "@%1$s = value", property);
            buffer.addLine("end");
        }
        else if (type instanceof ListType) {
            ListType listType = (ListType) type;
            Type elementType = listType.getElementType();
            if (elementType instanceof PrimitiveType || elementType instanceof EnumType) {
                buffer.addLine("def %1$s=(list)", property);
                buffer.addLine(  "@%1$s = list", property);
                buffer.addLine("end");
            }
            else if (elementType instanceof StructType) {
                GoName elementTypeName = goNames.getTypeName(elementType);
                buffer.addLine("def %1$s=(list)", property);
                buffer.addLine(  "if list.class == Array");
                buffer.addLine(    "list = List.new(list)");
                buffer.addLine(    "list.each_with_index do |value, index|");
                buffer.addLine(      "if value.is_a?(Hash)");
                buffer.addLine(        "list[index] = %1$s.new(value)", elementTypeName.getClassName());
                buffer.addLine(      "end");
                buffer.addLine(    "end");
                buffer.addLine(  "end");
                buffer.addLine(  "@%1$s = list", property);
                buffer.addLine("end");
            }
        }
        buffer.addLine();
    }

    private void generateEnums(Model model) {
        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .sorted()
            .forEach(this::generateEnum);
    }

    private void generateEnum(EnumType type) {
        // Values:
        type.values().sorted().forEach(this::generateEnumValue);
    }

    private void generateEnumValue(EnumValue value) {
        String constantName = goNames.getConstantStyleName(value.getName());
        String constantValue = goNames.getFuncStyleName(value.getName());
        buffer.addLine("%s = '%s'.freeze", constantName, constantValue);
    }

    private void generateTypeDeclaration(StructType type) {
        GoName typeName = goNames.getTypeName(type);
        buffer.addLine("type %1$s struct{", typeName.getClassName());
    }
}


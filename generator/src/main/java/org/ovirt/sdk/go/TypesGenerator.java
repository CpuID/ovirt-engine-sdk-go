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
        // Generate the source:
        generateStructs(model);
        generateEnums(model);
    }

    private void generateStructs(Model model) {
        // Calculate the file name:
        buffer = new GoBuffer();
        String fileName = goNames.getPackagePath() + "/types/types";
        buffer.setFileName(fileName);
        buffer.setPackageName("types");

        // Set our imports:
        // TODO: setup however required
        //String repoSdkUrl = goNames.repoSdkUrl();
        //buffer.addImport("fmt");
        //buffer.addImport(repoSdkUrl + "/http");

        // We don't care too much about declaration order, but let's use inheritance order (if any), for consistency
        // with other SDK implementations.
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

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new IllegalStateException("Error writing types file \"" + fileName + "\"", exception);
        }
    }

    private void generateStruct(StructType type) {
        // Begin struct type:
        generateTypeDeclaration(type);

        // Attributes and links, for fields and member functions:
        List<StructMember> members = new ArrayList<>();
        members.addAll(type.getAttributes());
        members.addAll(type.getLinks());

        // Add struct fields:
        members.stream().sorted().forEach(this::generateStructField);

        // End type:
        buffer.addLine("}");
        buffer.addLine();

        // Struct member functions:
        members.stream().sorted().forEach(this::generateMember);
    }

    private void generateStructField(StructMember member) {
        Name name = member.getName();
        Type type = member.getType();
        // All of these types are local to the types package, no need to prefix with another package name.
        if (type instanceof ListType) {
          ListType listType = (ListType) type;
          Type elementType = listType.getElementType();
          buffer.addLine("  %1$s []%2$s", goNames.getPublicStyleName(name), goNames.getPublicStyleName(elementType.getName()));
        } else if (type instanceof PrimitiveType) {
          buffer.addLine("  %1$s %2$s", goNames.getPublicStyleName(name), convertMetaModelPrimitiveTypeToGoType(type));
        } else {
          buffer.addLine("  %1$s %2$s", goNames.getPublicStyleName(name), goNames.getPublicStyleName(type.getName()));
        }
    }

    private String convertMetaModelPrimitiveTypeToGoType (Type type) {
      if (!(type instanceof PrimitiveType)) {
        throw new RuntimeException("Variable type " + type + " is not an instance of PrimitiveType");
      }
      Model model = type.getModel();
      if (type == model.getBooleanType()) {
        return "bool";
      } else if (type == model.getIntegerType()) {
        return "int";
      } else if (type == model.getDecimalType()) {
        return "float64";
      } else if (type == model.getStringType()) {
        return "string";
      } else if (type == model.getDateType()) {
        return "time.Time";
      } else {
        throw new RuntimeException("Variable type " + type + " cannot be converted to a builtin Go type");
      }
    }

    private void generateMember(StructMember member) {
        //generateGetter(member);
        generateSetter(member);
    }

    /*
     * struct fields will all be public, don't need a getter
    private void generateGetter(StructMember member) {
        Name name = member.getName();
        Type type = member.getType();
        String property = goNames.getPublicStyleName(name);
        buffer.addComment("Returns the value of the `%1$s` attribute.", property);
        buffer.addLine("func (g *TODOtype) Get%1$s () {", property);
        buffer.addLine(  "return @%1$s", property);
        buffer.addLine("end");
        buffer.addLine();
    }
    */

    private void generateSetter(StructMember member) {
        Name name = member.getName();
        Type type = member.getType();
        String property = goNames.getPublicStyleName(name);
        if (!(type instanceof PrimitiveType)) {
          buffer.addComment("Sets the value of the `%1$s` attribute.", property);
          if (type instanceof EnumType) {
              buffer.addLine("func (s *%1$s) Set%2$s (value %3$s) error {", property, "TODO1", "TODO2");
              buffer.addLine("  s.%1$s = value", property);
              buffer.addLine("  return s.%1$s.Validate()", "TODO1");
              buffer.addLine("}");
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
    }

    private void generateEnums(Model model) {
        // One big file of all enums for now.
        // Calculate the file name:
        buffer = new GoBuffer();
        String fileName = goNames.getPackagePath() + "/types/enum";
        buffer.setFileName(fileName);
        buffer.setPackageName("types");

        // Set our imports:
        String repoSdkUrl = goNames.repoSdkUrl();
        buffer.addImport("fmt");
        buffer.addImport("strings");
        buffer.addImport(repoSdkUrl + "/slice");

        model.types()
            .filter(EnumType.class::isInstance)
            .map(EnumType.class::cast)
            .sorted()
            .forEach(this::generateEnum);

        // Write the file:
        try {
            buffer.write(out);
        }
        catch (IOException exception) {
            throw new IllegalStateException("Error writing types file \"" + fileName + "\"", exception);
        }
    }

    private void generateEnum(EnumType type) {
        GoName typeName = goNames.getTypeName(type);
        buffer.addComment("Enum for %1$s", typeName);
        buffer.addLine("type %1$s string", typeName);
        buffer.addLine();
        buffer.addLine("func (e *%1$s) Validate () error {", typeName);
        buffer.addLine("  allowed_values := []string{");
        type.values().sorted().forEach(this::generateEnumValue);
        buffer.addLine("  }");
        buffer.addLine("  if slice.StringInSlice(e, allowed_values) != true {");
        buffer.addLine("    return fmt.Errorf(\"Type '%1$s' cannot have a value of '%%s'. Permitted values: %%s\", e, strings.Join(allowed_values, \",\"))", typeName);
        buffer.addLine("  } else {");
        buffer.addLine("    return nil");
        buffer.addLine("  }");
        buffer.addLine("}");
        buffer.addLine();
    }

    private void generateEnumValue(EnumValue value) {
        buffer.addLine("    \"%1$s\",", goNames.getPublicStyleName(value.getName()));
    }

    private void generateTypeDeclaration(StructType type) {
        GoName typeName = goNames.getTypeName(type);
        buffer.addLine("type %1$s struct {", typeName.getClassName());
    }
}


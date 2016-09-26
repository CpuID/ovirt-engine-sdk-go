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
        members.stream().sorted().forEach(member -> generateSetterMember(member, type));
    }

    private void generateStructField(StructMember member) {
        Name name = member.getName();
        Type type = member.getType();
        // All of these types are local to the types package, no need to prefix with another package name.
        buffer.addLine("  %1$s %2$s", goNames.getPublicStyleName(name), convertMetamodelTypeToGoType(type));
    }

    private String convertMetamodelTypeToGoType (Type type) {
      if (type instanceof PrimitiveType) {
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
      } else if (type instanceof ListType) {
        ListType listType = (ListType) type;
        Type elementType = listType.getElementType();
        return "[]" + goNames.getPublicStyleName(elementType.getName());
      } else {
        return goNames.getPublicStyleName(type.getName());
      }
    }

    private void generateSetterMember(StructMember member, Type parentType) {
        Name name = member.getName();
        Type type = member.getType();
        String property = goNames.getPublicStyleName(name);
        if (!(type instanceof PrimitiveType)) {
          String propertyType = convertMetamodelTypeToGoType(type);
          if (type instanceof EnumType) {
              buffer.addComment("Sets the value of the `%1$s` attribute.", property);
              buffer.addLine("func (s *%1$s) Set%2$s (value %3$s) error {", goNames.getPublicStyleName(parentType.getName()), property, propertyType);
              buffer.addLine("  s.%1$s = value", property);
              buffer.addLine("  return s.%1$s.Validate()", property);
              buffer.addLine("}");
          } else if (type instanceof StructType) {
              buffer.addComment("To set the value of the `%1$s` attribute, just set the exported field directly.", property);
              buffer.addComment("As Golang is statically typed, you cannot use a hash/map input to set the field (other SDK languages do).");
              buffer.addComment("This can be implemented if required (some of the codegen is stubbed out), but incomplete currently.");
              //buffer.addLine("func (s *%1$s) Set%2$s (value %3$s) {", goNames.getPublicStyleName(declaringType.getName()), property, propertyType);
              //buffer.addLine("  s.%1$s = value", property);
              //buffer.addLine("}");
              // TODOLATER: finish this if it becomes required. try use the typed setter above wherever we can, avoids the need for reflection.
              /*
              buffer.addLine();
              buffer.addComment("For completeness, provide a map input mechanism function equivalent to `Set%1$s` above.", property);
              buffer.addLine("func (s *%1$s) Set%2$sWithMap (value map[string]string) {", goNames.getPublicStyleName(declaringType.getName()), property);
              buffer.addLine("  value := %1$s{}", propertyType);
              buffer.addLine("  for k, v := range value {");
              // TODOLATER: finish logic using reflect to convert map to struct field values
              buffer.addLine("  }");
              buffer.addLine("  s.%1$s = value", property);
              buffer.addLine("}");
              */
          } else if (type instanceof ListType) {
              buffer.addComment("To set the value of the `%1$s` attribute, just set the exported field directly using a slice literal.", property);
              /*
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
              */
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


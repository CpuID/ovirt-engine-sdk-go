/*
Copyright (c) 2015 Red Hat, Inc. / Nathan Sullivan

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

import java.io.File;
import java.util.Arrays;
import java.util.List;
import java.util.Set;
import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;

import org.ovirt.api.metamodel.concepts.Name;
import org.ovirt.api.metamodel.concepts.NameParser;
import org.ovirt.api.metamodel.concepts.Service;
import org.ovirt.api.metamodel.concepts.Type;
import org.ovirt.api.metamodel.tool.ReservedWords;
import org.ovirt.api.metamodel.tool.Words;

/**
 * This class contains the rules used to calculate the names of generated Java concepts.
 */
@ApplicationScoped
public class GoNames {
    // The names of the base classes:
    public static final Name ACTION_NAME = NameParser.parseUsingCase("Action");
    public static final Name FAULT_NAME = NameParser.parseUsingCase("Fault");
    public static final Name LIST_NAME = NameParser.parseUsingCase("List");
    public static final Name READER_NAME = NameParser.parseUsingCase("Reader");
    public static final Name SERVICE_NAME = NameParser.parseUsingCase("Service");
    public static final Name STRUCT_NAME = NameParser.parseUsingCase("Struct");
    public static final Name WRITER_NAME = NameParser.parseUsingCase("Writer");

    // The names of the directories:
    public static final Name READERS_DIR = NameParser.parseUsingCase("Readers");
    public static final Name SERVICES_DIR = NameParser.parseUsingCase("Services");
    public static final Name TYPES_DIR = NameParser.parseUsingCase("Types");
    public static final Name WRITERS_DIR = NameParser.parseUsingCase("Writers");

    // Reference to the object used to do computations with words.
    @Inject private Words words;

    // We need the Go reserved words in order to avoid producing names that aren't legal in Java:
    @Inject
    @ReservedWords(language = "go")
    private Set<String> reservedWords;

    // The path of the package:
    private String packagePath = "";

    // The version of the gem:
    private String version;

    /**
     * Get the path of the package.
     */
    public String getPackagePath() {
        return packagePath;
    }

    /*
     * Defines the repo prefix used when defining imports.
     */
    public String repoSdkUrl() {
      return "github.com/CpuID/ovirt-engine-sdk-go/sdk";
    }

    /**
     * Get the version.
     */
    public String getVersion() {
        return version;
    }

    /**
     * Set the version.
     */
    public void setVersion(String newVersion) {
        version = newVersion;
    }

    /**
     * Calculates the Go name of the base class of the struct types.
     */
    public GoName getBaseStructName() {
        return buildName(STRUCT_NAME, null, TYPES_DIR);
    }

    /**
     * Calculates the Go name of the base class of the list types.
     */
    public GoName getBaseListName() {
        return buildName(LIST_NAME, null, TYPES_DIR);
    }

    /**
     * Calculates the Go name that corresponds to the given type.
     */
    public GoName getTypeName(Type type) {
        return buildName(type.getName(), null, TYPES_DIR);
    }

    /**
     * Calculates the Go name of the base class of the services.
     */
    public GoName getBaseServiceName() {
        return buildName(SERVICE_NAME, null, SERVICES_DIR);
    }

    /**
     * Calculates the Go name that corresponds to the given service.
     */
    public GoName getServiceName(Service service) {
        return buildName(service.getName(), SERVICE_NAME, SERVICES_DIR);
    }

    /**
     * Calculates the Go name of the base class of the readers.
     */
    public GoName getBaseReaderName() {
        return buildName(READER_NAME, null, READERS_DIR);
    }

    /**
     * Calculates the Go name of the base class of the writers.
     */
    public GoName getBaseWriterName() {
        return buildName(WRITER_NAME, null, WRITERS_DIR);
    }

    /**
     * Calculates the Go name of the reader for the given type.
     */
    public GoName getReaderName(Type type) {
        return buildName(type.getName(), READER_NAME, READERS_DIR);
    }

    /**
     * Calculates the Go name of the writer for the given type.
     */
    public GoName getWriterName(Type type) {
        return buildName(type.getName(), WRITER_NAME, WRITERS_DIR);
    }

    /**
     * Calculates the Go name of the fault class.
     */
    public GoName getFaultName() {
        return buildName(FAULT_NAME, null, TYPES_DIR);
    }

    /**
     * Calculates the Go name of the fault reader.
     */
    public GoName getFaultReaderName() {
        return buildName(FAULT_NAME, READER_NAME, READERS_DIR);
    }

    /**
     * Calculates the Go name of the action class.
     */
    public GoName getActionName() {
        return buildName(ACTION_NAME, null, TYPES_DIR);
    }

    /**
     * Calculates the Go name of the action reader.
     */
    public GoName getActionReaderName() {
        return buildName(ACTION_NAME, READER_NAME, READERS_DIR);
    }

    /**
     * Builds a Go name from the given base name and suffix, and a directory.
     *
     * The suffix can be {@code null} or empty, in that case then won't be added.
     *
     * If the {@code directory} parameter is given then it will be appended to the file associated to the name. For
     * example, if the {@code directory} parameter is {@code types} then the file will be {@code ovirt/lib/types}.
     *
     * @param base the base name
     * @param suffix the suffix to add to the name
     * @param directory the directory associated to the name
     * @return the calculated Go name
     */
    private GoName buildName(Name base, Name suffix, Name directory) {
        // Calculate class name:
        List<String> words = base.getWords();
        if (suffix != null) {
            words.addAll(suffix.getWords());
        }
        Name name = new Name(words);
        GoName result = new GoName();
        result.setClassName(getPublicStyleName(name));

        // Calculate the file name:
        StringBuilder fileName = new StringBuilder();
        fileName.append(packagePath);
        fileName.append(File.separator);
        if (directory != null) {
            fileName.append(getFileStyleName(directory));
            fileName.append(File.separator);
        }
        fileName.append(getFileStyleName(name));
        result.setFileName(fileName.toString());

        return result;
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go public/exported types/functions.
     */
    public String getPublicStyleName(Name name) {
        return name.words().map(words::capitalize).collect(joining());
    }

    /**
     * Returns a representation of the given name using the capitalization style typically used for Go private/non-exported types/functions.
     */
    public String getPrivateStyleName(Name name) {
        String capitalized = name.words().map(words::capitalize).collect(joining());
        return Character.toLowerCase(capitalized.charAt(0)) + capitalized.substring(1);
    }

    /**
     * Returns a representation of the given name using the snake case style typically used for Go files.
     */
    public String getFileStyleName(Name name) {
        return name.words().map(String::toLowerCase).collect(joining("_"));
    }
}


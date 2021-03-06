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

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStreamWriter;
import java.io.Writer;
import java.nio.charset.StandardCharsets;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Deque;
import java.util.Formatter;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import javax.inject.Inject;

import org.apache.commons.io.FileUtils;

/**
 * This class is a buffer intended to simplify generation of Go source code. It stores the name of the module, the
 * list of imports and the rest of the source separately, so that imports can be added on demand while generating the
 * rest of the source.
 */
public class GoBuffer {
    // Reference to the object used to generate Go names:
    @Inject private GoNames goNames;

    // The name of the file:
    private String fileName;

    // The Go package within the file:
    private String packageName;

    // The things to be imported:
    private List<String> imports = new ArrayList<>();

    // The lines of the body of the class:
    private List<String> lines = new ArrayList<>();

    // The current indentation level:
    private int level;

    /**
     * Sets the file name.
     */
    public void setFileName(String newFileName) {
        fileName = newFileName;
    }

    /**
     * Sets the package name.
     */
    public void setPackageName(String newPackageName) {
        packageName = newPackageName;
    }

    /**
     * Adds an import to the list for this file/buffer.
     */
    public void addImport(String importPath) {
      imports.add(importPath);
    }

    /**
     * Adds a line to the file. If the line contains new line characters then it will be broken and multiple lines and
     * each one will be processed in sequence. For example, if these lines aren't indented, the result will be indented
     * anyhow.
     */
    public void addLine(String line) {
        if (line != null) {
            String[] parts = line.split("\\n");
            for (String part : parts) {
                addLineNoSplit(part);
            }
        }
    }

    /**
     * Adds a line to the file without taking into account new line characters.
     */
    private void addLineNoSplit(String line) {
        // Check of the line is the begin or end of a block:
        // TODO: adjust block boundary definitions for golang
        boolean isBegin =
            line.endsWith("(") ||
            line.endsWith("[") ||
            line.endsWith("|") ||
            line.equals("begin") ||
            line.equals("else") ||
            line.equals("ensure") ||
            line.startsWith("case ") ||
            line.startsWith("class ") ||
            line.startsWith("def ") ||
            line.startsWith("if ") ||
            line.startsWith("loop ") ||
            line.startsWith("module ") ||
            line.startsWith("unless ") ||
            line.startsWith("when ") ||
            line.startsWith("while ");
        boolean isEnd =
            line.equals(")") ||
            line.equals("]") ||
            line.equals("else") ||
            line.equals("else") ||
            line.equals("end") ||
            line.equals("ensure") ||
            line.startsWith("when ");

        // Decrease the indentation if the line is the end of a block:
        if (isEnd) {
            if (level > 0) {
                level--;
            }
        }

        // Indent the line and add it to the list:
        StringBuilder buffer = new StringBuilder(level * 2 + line.length());
        for (int i = 0; i < level; i++) {
            buffer.append("  ");
        }
        buffer.append(line);
        line = buffer.toString();
        lines.add(line);

        // Increase the indentation if the line is the begin of a block:
        if (isBegin) {
            level++;
        }
    }

    /**
     * Adds an empty line to the body of the class.
     */
    public void addLine() {
        addLine("");
    }

    /**
     * Adds a formatted line to the file. The given {@code args} are formatted using the provided {@code format} using
     * the {@link String#format(String, Object...)} method.
     */
    public void addLine(String format, Object ... args) {
        StringBuilder buffer = new StringBuilder();
        Formatter formatter = new Formatter(buffer);
        formatter.format(format, args);
        String line = buffer.toString();
        addLine(line);
    }

    /**
     * Adds a comment to the file. If the line contains new line characters then it will be broken and multiple lines
     * and each one will be processed in sequence. For example, if these lines aren't indented, the result will be
     * indented anyhow.
     */
    public void addComment(String line) {
        if (line != null) {
            String[] parts = line.split("\\n");
            for (String part : parts) {
                addCommentNoSplit(part);
            }
        }
    }

    /**
     * Adds a comment to the file without taking into account new line characters.
     */
    private void addCommentNoSplit(String line) {
        StringBuilder buffer = new StringBuilder(2 + level * 2 + line.length());
        for (int i = 0; i < level; i++) {
            buffer.append("  ");
        }
        buffer.append("// ");
        buffer.append(line);
        line = buffer.toString();
        lines.add(line);
    }

    /**
     * Adds an empty comment to the file.
     */
    public void addComment() {
        addComment("");
    }

    /**
     * Adds a formatted comment to the file. The given {@code args} are formatted using the provided {@code format}
     * using the {@link String#format(String, Object...)} method.
     */
    public void addComment(String format, Object ... args) {
        StringBuilder buffer = new StringBuilder();
        Formatter formatter = new Formatter(buffer);
        formatter.format(format, args);
        String line = buffer.toString();
        addComment(line);
    }

    /**
     * Generates the complete source code of the class.
     */
    public String toString() {
        StringBuilder buffer = new StringBuilder();

        // License:
        buffer.append("//\n");
        buffer.append("// Copyright (c) 2015-2016 Red Hat, Inc. / Nathan Sullivan\n");
        buffer.append("//\n");
        buffer.append("// Licensed under the Apache License, Version 2.0 (the \"License\");\n");
        buffer.append("// you may not use this file except in compliance with the License.\n");
        buffer.append("// You may obtain a copy of the License at\n");
        buffer.append("//\n");
        buffer.append("//   http://www.apache.org/licenses/LICENSE-2.0\n");
        buffer.append("//\n");
        buffer.append("// Unless required by applicable law or agreed to in writing, software\n");
        buffer.append("// distributed under the License is distributed on an \"AS IS\" BASIS,\n");
        buffer.append("// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n");
        buffer.append("// See the License for the specific language governing permissions and\n");
        buffer.append("// limitations under the License.\n");
        buffer.append("//\n");
        buffer.append("\n");
        buffer.append(String.format("package %1$s\n", packageName));
        buffer.append("\n");

        // Imports:
        List<String> importList = new ArrayList<>(imports);
        Collections.sort(importList);
        buffer.append("import (\n");
        for (String importItem : importList) {
            buffer.append("  \"");
            buffer.append(importItem);
            buffer.append("\"\n");
        }
        buffer.append(")\n");
        buffer.append("\n");

        // Body:
        for (String line : lines) {
            buffer.append(line);
            buffer.append("\n");
        }

        return buffer.toString();
    }

    /**
     * Creates a {@code .go} source file and writes the source. The required intermediate directories will be created
     * if they don't exist.
     *
     * @param dir the base directory for the source code
     * @throws IOException if something fails while creating or writing the file
     */
    public void write(File dir) throws IOException {
        // Calculate the complete file name:
        File file = new File(dir, fileName.replace('/', File.separatorChar) + ".go");

        // Create the directory and all its parent if needed:
        File parent = file.getParentFile();
        FileUtils.forceMkdir(parent);

        // Write the file:
        System.out.println("Writing file \"" + file.getAbsolutePath() + "\".");
        try (Writer writer = new OutputStreamWriter(new FileOutputStream(file), StandardCharsets.UTF_8)) {
            writer.write(toString());
        }
    }
}

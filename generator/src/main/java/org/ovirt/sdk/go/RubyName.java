/*
Copyright (c) 2015 Red Hat, Inc.

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

package org.ovirt.sdk.ruby;

/**
 * This class represents the fully qualified name of a Ruby class, composed by the module name, the class name and the
 * name of the file where it should be stored.
 */
public class RubyName {
    private String moduleName;
    private String className;
    private String fileName;

    public String getModuleName() {
        return moduleName;
    }

    public void setModuleName(String newModuleName) {
        moduleName = newModuleName;
    }

    public String getClassName() {
        return className;
    }

    public void setClassName(String newClassName) {
        className = newClassName;
    }

    public String getFileName() {
        return fileName;
    }

    public void setFileName(String newFileName) {
        this.fileName = newFileName;
    }

    @Override
    public String toString() {
        StringBuilder buffer = new StringBuilder();
        buffer.append(moduleName);
        buffer.append("::");
        buffer.append(className);
        return buffer.toString();
    }
}


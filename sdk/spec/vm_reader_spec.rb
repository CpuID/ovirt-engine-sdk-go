#
# Copyright (c) 2015-2016 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

describe SDK::VmReader do

  describe ".read_one" do

    context "when given an empty XML" do

      it "creates an empty VM" do
        reader = SDK::XmlReader.new('<vm/>')
        result = SDK::VmReader.read_one(reader)
        reader.close
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
      end

    end

    context "when given a VM with an id" do

      it "creates a VM with that id" do
        reader = SDK::XmlReader.new('<vm id="123"/>')
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.id).to eql('123')
      end

    end

    context "when given a VM with a name" do

      it "creates a VM with that name" do
        reader = SDK::XmlReader.new('<vm><name>myvm</name></vm>')
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.name).to eql('myvm')
      end

    end

    context "when given a VM with id and name" do

      it "creates a VM with that id and name" do
        reader = SDK::XmlReader.new('<vm id="123"><name>myvm</name></vm>')
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.id).to eql('123')
        expect(result.name).to eql('myvm')
      end

    end

    context "when given an alternative tag" do

      it "ignores it and reads the attributes correctly" do
        reader = SDK::XmlReader.new('<alternative id="123"><name>myvm</name></alternative>')
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.id).to eql('123')
        expect(result.name).to eql('myvm')
      end

    end

    context "when the href attribute has a value" do

      it "the href getter returns its value" do
        reader = SDK::XmlReader.new('<vm href="myhref"></vm>')
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.href).to eql('myhref')
      end

    end

    context "when given a link to a list" do

      it "the corresponding attribute is populated" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link rel="nics" href="/vms/123/nics"/>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.nics).to_not be_nil
        expect(result.nics).to be_a(SDK::List)
        expect(result.nics.href).to eql('/vms/123/nics')
      end

      it "the next attribute is read correctly" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link rel="nics" href="/vms/123/nics"/>' +
            '<name>myvm</name>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.name).to eql('myvm')
      end

      it "the link is ignored if there is no such link" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link rel="junks" href="/junks"/>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.nics).to be_nil
      end

      it "the link is ignored if it has no rel" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link href="/junks"/>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.nics).to be_nil
      end

      it "the link is ignored if it has no href" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link rel="nics"/>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.nics).to be_nil
      end
    end

    context "when given a multiple links to lists" do

      it "the corresponding attributes are populated" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link rel="nics" href="/vms/123/nics"/>' +
            '<link rel="cdroms" href="/vms/123/cdroms"/>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.nics).to_not be_nil
        expect(result.nics).to be_a(SDK::List)
        expect(result.nics.href).to eql('/vms/123/nics')
        expect(result.cdroms).to_not be_nil
        expect(result.cdroms).to be_a(SDK::List)
        expect(result.cdroms.href).to eql('/vms/123/cdroms')
      end

      it "the next attribute is read correctly" do
        reader = SDK::XmlReader.new(
          '<vm>' +
            '<link rel="nics" href="/vms/123/nics"/>' +
            '<link rel="cdroms" href="/vms/123/cdroms"/>' +
            '<name>myvm</name>' +
          '</vm>'
        )
        result = SDK::VmReader.read_one(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::Vm)
        expect(result.name).to eql('myvm')
      end

    end

  end

  describe ".read_many" do

    context "when given an empty XML" do

      it "creates an empty list" do
        reader = SDK::XmlReader.new('<vms/>')
        result = SDK::VmReader.read_many(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::List)
        expect(result.size).to eql(0)
      end

    end

    context "when given one VM" do

      it "creates an a list with one element" do
        reader = SDK::XmlReader.new('<vms><vm/></vms>')
        result = SDK::VmReader.read_many(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::List)
        expect(result.size).to eql(1)
        expect(result[0]).to_not be_nil
        expect(result[0]).to be_a(SDK::Vm)
      end

    end

    context "when given two VMs" do

      it "creates an a list with two element" do
        reader = SDK::XmlReader.new('<vms><vm/><vm/></vms>')
        result = SDK::VmReader.read_many(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::List)
        expect(result.size).to eql(2)
        expect(result[0]).to_not be_nil
        expect(result[0]).to be_a(SDK::Vm)
        expect(result[1]).to_not be_nil
        expect(result[1]).to be_a(SDK::Vm)
      end

    end

    context "when given alternative tags" do

      it "reads the elements of the list correctly" do
        reader = SDK::XmlReader.new('<myvms><myvm/><myvm/></myvms>')
        result = SDK::VmReader.read_many(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::List)
        expect(result.size).to eql(2)
        expect(result[0]).to_not be_nil
        expect(result[0]).to be_a(SDK::Vm)
        expect(result[1]).to_not be_nil
        expect(result[1]).to be_a(SDK::Vm)
      end

    end

    context "when the href attribute has a value" do

      it "the href getter returns its value" do
        reader = SDK::XmlReader.new('<vms href="myhref"></vms>')
        result = SDK::VmReader.read_many(reader)
        expect(result).to_not be_nil
        expect(result).to be_a(SDK::List)
        expect(result.href).to eql('myhref')
      end

    end

  end

end

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

describe SDK::ClustersService do

  before(:all) do
    start_server
    @connection = test_connection
    @service = @connection.system_service.clusters_service
  end

  after(:all) do
    @connection.close
    stop_server
  end

  describe ".clusters" do

    context "getting the reference to the service" do

      it "doesn't return nil" do
        expect(@service).not_to be_nil
      end

    end

  end

  describe ".list" do

    context "without parameters" do

      it "returns a list, maybe empty" do
        set_xml_response('clusters', 200, '<clusters/>')
        clusters = @service.list
        expect(clusters).not_to be_nil
        expect(clusters).to be_an(Array)
      end

    end

    context "with an unfeasible query" do

      it "returns an empty array" do
        set_xml_response('clusters', 200, '<clusters/>')
        clusters = @service.list(:search => 'name=ugly')
        expect(clusters).to eql([])
      end

    end

  end

end

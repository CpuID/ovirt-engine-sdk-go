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

describe SDK::StorageDomainService do

  before(:all) do
    start_server
    @connection = test_connection
    @service = @connection.system_service.storage_domains_service.storage_domain_service('123')
  end

  after(:all) do
    @connection.close
    stop_server
  end

  describe ".is_attached" do

    context "when invoking the method" do

      it "returns the value of the `is_attached` output parameter" do
        set_xml_response('storagedomains/123/isattached', 200, '<action><is_attached>true</is_attached></action>')
        expect(@service.is_attached).to be(true)
      end

    end

  end

end

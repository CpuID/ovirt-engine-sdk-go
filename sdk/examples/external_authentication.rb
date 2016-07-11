#!/usr/bin/ruby

#
# Copyright (c) 2016 Red Hat, Inc.
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

require 'ovirtsdk4'

# This example shows how to connect to Github server to obtain the OAuth2 access token. The access token won't work
# with engine as ovirt-engine does not currently support external OAuth2 server.

# Create the connection to the server:
connection = OvirtSDK4::Connection.new({
  :url => 'https://engine40.example.com/ovirt-engine/api',
  :sso_url => "https://api.github.com/authorizations",
  :username => 'myusername',
  :password => 'mypassword',
  :ca_file => 'ca.pem',
  :sso_token_name => 'hashed_token',
  :debug => true,
})

# Get the reference to the system service:
api = connection.system_service.get

# Output the full version:
puts(api.product_info.version.full_version)

# Close the connection to the server:
connection.close

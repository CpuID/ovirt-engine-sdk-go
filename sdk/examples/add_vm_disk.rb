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

# This example will connect to the server and attach a disk to an existing
# virtual machine.

# Create the connection to the server:
connection = OvirtSDK4::Connection.new({
  :url => 'https://engine40.example.com/ovirt-engine/api',
  :username => 'admin@internal',
  :password => 'redhat123',
  :ca_file => 'ca.pem',
  :debug => true,
})

# Locate the virtual machines service and use it to find the virtual
# machine:
vms_service = connection.system_service.vms_service
vm = vms_service.list({:search => 'name=myvm'})[0]

# Locate the service that manages the disk attachments of the virtual
# machine:
disk_attachments_service = vms_service.vm_service(vm.id).disk_attachments_service

# Use the "add" method of the disk attachments service to add the disk:
disk_attachment = disk_attachments_service.add(
  OvirtSDK4::DiskAttachment.new(
    :disk => {
      :name => 'mydisk',
      :description => 'My disk',
      :format => OvirtSDK4::DiskFormat::COW,
      :provisioned_size => 1 * 2**20,
      :storage_domains => [{
        :name => 'mydata',
      }],
    },
    :interface => OvirtSDK4::DiskInterface::VIRTIO,
    :bootable => false,
  )
)

# Wait till the disk is OK:
disks_service = connection.system_service.disks_service
disk_service = disks_service.disk_service(disk_attachment.disk.id)
begin
  sleep(5)
  disk = disk_service.get
end while disk.status != OvirtSDK4::DiskStatus::OK

# Close the connection to the server:
connection.close

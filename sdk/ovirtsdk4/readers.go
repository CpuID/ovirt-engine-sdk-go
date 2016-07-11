#
# Copyright (c) 2015-2016 Red Hat, Inc. / Nathan Sullivan
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


package ovirtsdk4

class ActionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Action.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('async')
    object.async = value if not value.nil?
    value = reader.get_attribute('check_connectivity')
    object.check_connectivity = value if not value.nil?
    value = reader.get_attribute('clone')
    object.clone = value if not value.nil?
    value = reader.get_attribute('collapse_snapshots')
    object.collapse_snapshots = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('connectivity_timeout')
    object.connectivity_timeout = value if not value.nil?
    value = reader.get_attribute('deploy_hosted_engine')
    object.deploy_hosted_engine = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('discard_snapshots')
    object.discard_snapshots = value if not value.nil?
    value = reader.get_attribute('exclusive')
    object.exclusive = value if not value.nil?
    value = reader.get_attribute('exclussive')
    object.exclussive = value if not value.nil?
    value = reader.get_attribute('fence_type')
    object.fence_type = value if not value.nil?
    value = reader.get_attribute('filter')
    object.filter = value if not value.nil?
    value = reader.get_attribute('fix_layout')
    object.fix_layout = value if not value.nil?
    value = reader.get_attribute('force')
    object.force = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('image')
    object.image = value if not value.nil?
    value = reader.get_attribute('import_as_template')
    object.import_as_template = value if not value.nil?
    value = reader.get_attribute('is_attached')
    object.is_attached = value if not value.nil?
    value = reader.get_attribute('maintenance_enabled')
    object.maintenance_enabled = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('pause')
    object.pause = value if not value.nil?
    value = reader.get_attribute('reason')
    object.reason = value if not value.nil?
    value = reader.get_attribute('resolution_type')
    object.resolution_type = value if not value.nil?
    value = reader.get_attribute('restore_memory')
    object.restore_memory = value if not value.nil?
    value = reader.get_attribute('root_password')
    object.root_password = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('stop_gluster_service')
    object.stop_gluster_service = value if not value.nil?
    value = reader.get_attribute('succeeded')
    object.succeeded = value if not value.nil?
    value = reader.get_attribute('undeploy_hosted_engine')
    object.undeploy_hosted_engine = value if not value.nil?
    value = reader.get_attribute('use_cloud_init')
    object.use_cloud_init = value if not value.nil?
    value = reader.get_attribute('use_sysprep')
    object.use_sysprep = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'async'
        object.async = Reader.read_boolean(reader)
      when 'bricks'
        object.bricks = GlusterBrickReader.read_many(reader)
      when 'certificates'
        object.certificates = CertificateReader.read_many(reader)
      when 'check_connectivity'
        object.check_connectivity = Reader.read_boolean(reader)
      when 'clone'
        object.clone = Reader.read_boolean(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'collapse_snapshots'
        object.collapse_snapshots = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'connectivity_timeout'
        object.connectivity_timeout = Reader.read_integer(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'deploy_hosted_engine'
        object.deploy_hosted_engine = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'details'
        object.details = GlusterVolumeProfileDetailsReader.read_one(reader)
      when 'discard_snapshots'
        object.discard_snapshots = Reader.read_boolean(reader)
      when 'disk'
        object.disk = DiskReader.read_one(reader)
      when 'disks'
        object.disks = DiskReader.read_many(reader)
      when 'exclusive'
        object.exclusive = Reader.read_boolean(reader)
      when 'exclussive'
        object.exclussive = Reader.read_boolean(reader)
      when 'fault'
        object.fault = FaultReader.read_one(reader)
      when 'fence_type'
        object.fence_type = Reader.read_string(reader)
      when 'filter'
        object.filter = Reader.read_boolean(reader)
      when 'fix_layout'
        object.fix_layout = Reader.read_boolean(reader)
      when 'force'
        object.force = Reader.read_boolean(reader)
      when 'grace_period'
        object.grace_period = GracePeriodReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'image'
        object.image = Reader.read_string(reader)
      when 'import_as_template'
        object.import_as_template = Reader.read_boolean(reader)
      when 'is_attached'
        object.is_attached = Reader.read_boolean(reader)
      when 'iscsi'
        object.iscsi = IscsiDetailsReader.read_one(reader)
      when 'iscsi_targets'
        reader.next_element
      when 'job'
        object.job = JobReader.read_one(reader)
      when 'logical_units'
        object.logical_units = LogicalUnitReader.read_many(reader)
      when 'maintenance_enabled'
        object.maintenance_enabled = Reader.read_boolean(reader)
      when 'modified_bonds'
        object.modified_bonds = HostNicReader.read_many(reader)
      when 'modified_labels'
        object.modified_labels = NetworkLabelReader.read_many(reader)
      when 'modified_network_attachments'
        object.modified_network_attachments = NetworkAttachmentReader.read_many(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'option'
        object.option = OptionReader.read_one(reader)
      when 'pause'
        object.pause = Reader.read_boolean(reader)
      when 'power_management'
        object.power_management = PowerManagementReader.read_one(reader)
      when 'proxy_ticket'
        object.proxy_ticket = ProxyTicketReader.read_one(reader)
      when 'reason'
        object.reason = Reader.read_string(reader)
      when 'removed_bonds'
        object.removed_bonds = HostNicReader.read_many(reader)
      when 'removed_labels'
        object.removed_labels = NetworkLabelReader.read_many(reader)
      when 'removed_network_attachments'
        object.removed_network_attachments = NetworkAttachmentReader.read_many(reader)
      when 'resolution_type'
        object.resolution_type = Reader.read_string(reader)
      when 'restore_memory'
        object.restore_memory = Reader.read_boolean(reader)
      when 'root_password'
        object.root_password = Reader.read_string(reader)
      when 'snapshot'
        object.snapshot = SnapshotReader.read_one(reader)
      when 'ssh'
        object.ssh = SshReader.read_one(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'stop_gluster_service'
        object.stop_gluster_service = Reader.read_boolean(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'storage_domains'
        object.storage_domains = StorageDomainReader.read_many(reader)
      when 'succeeded'
        object.succeeded = Reader.read_boolean(reader)
      when 'synchronized_network_attachments'
        object.synchronized_network_attachments = NetworkAttachmentReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'ticket'
        object.ticket = TicketReader.read_one(reader)
      when 'undeploy_hosted_engine'
        object.undeploy_hosted_engine = Reader.read_boolean(reader)
      when 'use_cloud_init'
        object.use_cloud_init = Reader.read_boolean(reader)
      when 'use_sysprep'
        object.use_sysprep = Reader.read_boolean(reader)
      when 'virtual_functions_configuration'
        object.virtual_functions_configuration = HostNicVirtualFunctionsConfigurationReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class AffinityGroupReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = AffinityGroup.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('enforcing')
    object.enforcing = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('positive')
    object.positive = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'enforcing'
        object.enforcing = Reader.read_boolean(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'positive'
        object.positive = Reader.read_boolean(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class AffinityLabelReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = AffinityLabel.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('read_only')
    object.read_only = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'read_only'
        object.read_only = Reader.read_boolean(reader)
      when 'hosts'
        object.hosts = HostReader.read_many(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'hosts'
        object.hosts = list
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class AgentReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Agent.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('concurrent')
    object.concurrent = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('encrypt_options')
    object.encrypt_options = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('order')
    object.order = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'concurrent'
        object.concurrent = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'encrypt_options'
        object.encrypt_options = Reader.read_boolean(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'options'
        object.options = OptionReader.read_many(reader)
      when 'order'
        object.order = Reader.read_integer(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class AgentConfigurationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = AgentConfiguration.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('broker_type')
    object.broker_type = value if not value.nil?
    value = reader.get_attribute('network_mappings')
    object.network_mappings = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'broker_type'
        object.broker_type = Reader.read_string(reader)
      when 'network_mappings'
        object.network_mappings = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ApiReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Api.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('time')
    object.time = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'product_info'
        object.product_info = ProductInfoReader.read_one(reader)
      when 'special_objects'
        object.special_objects = SpecialObjectsReader.read_one(reader)
      when 'summary'
        object.summary = ApiSummaryReader.read_one(reader)
      when 'time'
        object.time = Reader.read_date(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ApiSummaryReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ApiSummary.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'hosts'
        object.hosts = ApiSummaryItemReader.read_one(reader)
      when 'storage_domains'
        object.storage_domains = ApiSummaryItemReader.read_one(reader)
      when 'users'
        object.users = ApiSummaryItemReader.read_one(reader)
      when 'vms'
        object.vms = ApiSummaryItemReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ApiSummaryItemReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ApiSummaryItem.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active')
    object.active = value if not value.nil?
    value = reader.get_attribute('total')
    object.total = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active'
        object.active = Reader.read_integer(reader)
      when 'total'
        object.total = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ApplicationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Application.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class AuthorizedKeyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = AuthorizedKey.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('key')
    object.key = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'key'
        object.key = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BalanceReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Balance.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'scheduling_policy'
        object.scheduling_policy = SchedulingPolicyReader.read_one(reader)
      when 'scheduling_policy_unit'
        object.scheduling_policy_unit = SchedulingPolicyUnitReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BiosReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Bios.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'boot_menu'
        object.boot_menu = BootMenuReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BlockStatisticReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = BlockStatistic.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BondingReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Bonding.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'options'
        object.options = OptionReader.read_many(reader)
      when 'slaves'
        object.slaves = HostNicReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BookmarkReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Bookmark.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BootReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Boot.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'devices'
        reader.next_element
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BootMenuReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = BootMenu.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class BrickProfileDetailReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = BrickProfileDetail.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'profile_details'
        object.profile_details = ProfileDetailReader.read_many(reader)
      when 'brick'
        object.brick = GlusterBrickReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CdromReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Cdrom.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'file'
        object.file = FileReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class CertificateReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Certificate.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('content')
    object.content = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('organization')
    object.organization = value if not value.nil?
    value = reader.get_attribute('subject')
    object.subject = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'content'
        object.content = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'organization'
        object.organization = Reader.read_string(reader)
      when 'subject'
        object.subject = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CloudInitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = CloudInit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('regenerate_ssh_keys')
    object.regenerate_ssh_keys = value if not value.nil?
    value = reader.get_attribute('timezone')
    object.timezone = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authorized_keys'
        object.authorized_keys = AuthorizedKeyReader.read_many(reader)
      when 'files'
        object.files = FileReader.read_many(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'network_configuration'
        object.network_configuration = NetworkConfigurationReader.read_one(reader)
      when 'regenerate_ssh_keys'
        object.regenerate_ssh_keys = Reader.read_boolean(reader)
      when 'timezone'
        object.timezone = Reader.read_string(reader)
      when 'users'
        object.users = UserReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ClusterReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Cluster.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('ballooning_enabled')
    object.ballooning_enabled = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('gluster_service')
    object.gluster_service = value if not value.nil?
    value = reader.get_attribute('ha_reservation')
    object.ha_reservation = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('maintenance_reason_required')
    object.maintenance_reason_required = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('optional_reason')
    object.optional_reason = value if not value.nil?
    value = reader.get_attribute('switch_type')
    object.switch_type = value if not value.nil?
    value = reader.get_attribute('threads_as_cores')
    object.threads_as_cores = value if not value.nil?
    value = reader.get_attribute('trusted_service')
    object.trusted_service = value if not value.nil?
    value = reader.get_attribute('tunnel_migration')
    object.tunnel_migration = value if not value.nil?
    value = reader.get_attribute('virt_service')
    object.virt_service = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'ballooning_enabled'
        object.ballooning_enabled = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'error_handling'
        object.error_handling = ErrorHandlingReader.read_one(reader)
      when 'fencing_policy'
        object.fencing_policy = FencingPolicyReader.read_one(reader)
      when 'gluster_service'
        object.gluster_service = Reader.read_boolean(reader)
      when 'ha_reservation'
        object.ha_reservation = Reader.read_boolean(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ksm'
        object.ksm = KsmReader.read_one(reader)
      when 'maintenance_reason_required'
        object.maintenance_reason_required = Reader.read_boolean(reader)
      when 'memory_policy'
        object.memory_policy = MemoryPolicyReader.read_one(reader)
      when 'migration'
        object.migration = MigrationOptionsReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'optional_reason'
        object.optional_reason = Reader.read_boolean(reader)
      when 'required_rng_sources'
        reader.next_element
      when 'scheduling_policy'
        object.scheduling_policy = SchedulingPolicyReader.read_one(reader)
      when 'serial_number'
        object.serial_number = SerialNumberReader.read_one(reader)
      when 'supported_versions'
        object.supported_versions = VersionReader.read_many(reader)
      when 'switch_type'
        object.switch_type = Reader.read_string(reader)
      when 'threads_as_cores'
        object.threads_as_cores = Reader.read_boolean(reader)
      when 'trusted_service'
        object.trusted_service = Reader.read_boolean(reader)
      when 'tunnel_migration'
        object.tunnel_migration = Reader.read_boolean(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      when 'virt_service'
        object.virt_service = Reader.read_boolean(reader)
      when 'affinity_groups'
        object.affinity_groups = AffinityGroupReader.read_many(reader)
      when 'cpu_profiles'
        object.cpu_profiles = CpuProfileReader.read_many(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'gluster_hooks'
        object.gluster_hooks = GlusterHookReader.read_many(reader)
      when 'gluster_volumes'
        object.gluster_volumes = GlusterVolumeReader.read_many(reader)
      when 'mac_pool'
        object.mac_pool = MacPoolReader.read_one(reader)
      when 'management_network'
        object.management_network = NetworkReader.read_one(reader)
      when 'network_filters'
        object.network_filters = NetworkFilterReader.read_many(reader)
      when 'networks'
        object.networks = NetworkReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'affinitygroups'
        object.affinity_groups = list
      when 'cpuprofiles'
        object.cpu_profiles = list
      when 'glusterhooks'
        object.gluster_hooks = list
      when 'glustervolumes'
        object.gluster_volumes = list
      when 'networkfilters'
        object.network_filters = list
      when 'networks'
        object.networks = list
      when 'permissions'
        object.permissions = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class ClusterLevelReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ClusterLevel.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'cpu_types'
        object.cpu_types = CpuTypeReader.read_many(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ConfigurationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Configuration.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('data')
    object.data = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'data'
        object.data = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ConsoleReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Console.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CoreReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Core.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('index')
    object.index = value if not value.nil?
    value = reader.get_attribute('socket')
    object.socket = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'index'
        object.index = Reader.read_integer(reader)
      when 'socket'
        object.socket = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CpuReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Cpu.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('architecture')
    object.architecture = value if not value.nil?
    value = reader.get_attribute('level')
    object.level = value if not value.nil?
    value = reader.get_attribute('mode')
    object.mode = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('speed')
    object.speed = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'architecture'
        object.architecture = Reader.read_string(reader)
      when 'cores'
        object.cores = CoreReader.read_many(reader)
      when 'cpu_tune'
        object.cpu_tune = CpuTuneReader.read_one(reader)
      when 'level'
        object.level = Reader.read_integer(reader)
      when 'mode'
        object.mode = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'speed'
        object.speed = Reader.read_decimal(reader)
      when 'topology'
        object.topology = CpuTopologyReader.read_one(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CpuProfileReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = CpuProfile.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'qos'
        object.qos = QosReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class CpuTopologyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = CpuTopology.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('cores')
    object.cores = value if not value.nil?
    value = reader.get_attribute('sockets')
    object.sockets = value if not value.nil?
    value = reader.get_attribute('threads')
    object.threads = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'cores'
        object.cores = Reader.read_integer(reader)
      when 'sockets'
        object.sockets = Reader.read_integer(reader)
      when 'threads'
        object.threads = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CpuTuneReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = CpuTune.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'vcpu_pins'
        object.vcpu_pins = VcpuPinReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CpuTypeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = CpuType.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('architecture')
    object.architecture = value if not value.nil?
    value = reader.get_attribute('level')
    object.level = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'architecture'
        object.architecture = Reader.read_string(reader)
      when 'level'
        object.level = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class CustomPropertyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = CustomProperty.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('regexp')
    object.regexp = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      when 'regexp'
        object.regexp = Reader.read_string(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class DataCenterReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = DataCenter.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('local')
    object.local = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('quota_mode')
    object.quota_mode = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('storage_format')
    object.storage_format = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'local'
        object.local = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'quota_mode'
        object.quota_mode = Reader.read_string(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'storage_format'
        object.storage_format = Reader.read_string(reader)
      when 'supported_versions'
        object.supported_versions = VersionReader.read_many(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      when 'clusters'
        object.clusters = ClusterReader.read_many(reader)
      when 'iscsi_bonds'
        object.iscsi_bonds = IscsiBondReader.read_many(reader)
      when 'mac_pool'
        object.mac_pool = MacPoolReader.read_one(reader)
      when 'networks'
        object.networks = NetworkReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'qoss'
        object.qoss = QosReader.read_many(reader)
      when 'quotas'
        object.quotas = QuotaReader.read_many(reader)
      when 'storage_domains'
        object.storage_domains = StorageDomainReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'clusters'
        object.clusters = list
      when 'iscsibonds'
        object.iscsi_bonds = list
      when 'networks'
        object.networks = list
      when 'permissions'
        object.permissions = list
      when 'qoss'
        object.qoss = list
      when 'quotas'
        object.quotas = list
      when 'storagedomains'
        object.storage_domains = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class DeviceReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Device.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class DiskReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Disk.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active')
    object.active = value if not value.nil?
    value = reader.get_attribute('actual_size')
    object.actual_size = value if not value.nil?
    value = reader.get_attribute('alias')
    object.alias_ = value if not value.nil?
    value = reader.get_attribute('bootable')
    object.bootable = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('format')
    object.format = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('image_id')
    object.image_id = value if not value.nil?
    value = reader.get_attribute('interface')
    object.interface = value if not value.nil?
    value = reader.get_attribute('logical_name')
    object.logical_name = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('propagate_errors')
    object.propagate_errors = value if not value.nil?
    value = reader.get_attribute('provisioned_size')
    object.provisioned_size = value if not value.nil?
    value = reader.get_attribute('read_only')
    object.read_only = value if not value.nil?
    value = reader.get_attribute('sgio')
    object.sgio = value if not value.nil?
    value = reader.get_attribute('shareable')
    object.shareable = value if not value.nil?
    value = reader.get_attribute('sparse')
    object.sparse = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('storage_type')
    object.storage_type = value if not value.nil?
    value = reader.get_attribute('uses_scsi_reservation')
    object.uses_scsi_reservation = value if not value.nil?
    value = reader.get_attribute('wipe_after_delete')
    object.wipe_after_delete = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active'
        object.active = Reader.read_boolean(reader)
      when 'actual_size'
        object.actual_size = Reader.read_integer(reader)
      when 'alias'
        object.alias_ = Reader.read_string(reader)
      when 'bootable'
        object.bootable = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'format'
        object.format = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'image_id'
        object.image_id = Reader.read_string(reader)
      when 'interface'
        object.interface = Reader.read_string(reader)
      when 'logical_name'
        object.logical_name = Reader.read_string(reader)
      when 'lun_storage'
        object.lun_storage = HostStorageReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'propagate_errors'
        object.propagate_errors = Reader.read_boolean(reader)
      when 'provisioned_size'
        object.provisioned_size = Reader.read_integer(reader)
      when 'read_only'
        object.read_only = Reader.read_boolean(reader)
      when 'sgio'
        object.sgio = Reader.read_string(reader)
      when 'shareable'
        object.shareable = Reader.read_boolean(reader)
      when 'sparse'
        object.sparse = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'storage_type'
        object.storage_type = Reader.read_string(reader)
      when 'uses_scsi_reservation'
        object.uses_scsi_reservation = Reader.read_boolean(reader)
      when 'wipe_after_delete'
        object.wipe_after_delete = Reader.read_boolean(reader)
      when 'disk_profile'
        object.disk_profile = DiskProfileReader.read_one(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'openstack_volume_type'
        object.openstack_volume_type = OpenStackVolumeTypeReader.read_one(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'quota'
        object.quota = QuotaReader.read_one(reader)
      when 'snapshot'
        object.snapshot = SnapshotReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'storage_domains'
        object.storage_domains = StorageDomainReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      when 'statistics'
        object.statistics = list
      when 'storagedomains'
        object.storage_domains = list
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class DiskAttachmentReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = DiskAttachment.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active')
    object.active = value if not value.nil?
    value = reader.get_attribute('bootable')
    object.bootable = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('interface')
    object.interface = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('pass_discard')
    object.pass_discard = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active'
        object.active = Reader.read_boolean(reader)
      when 'bootable'
        object.bootable = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'interface'
        object.interface = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'pass_discard'
        object.pass_discard = Reader.read_boolean(reader)
      when 'disk'
        object.disk = DiskReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class DiskProfileReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = DiskProfile.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'qos'
        object.qos = QosReader.read_one(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class DiskSnapshotReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = DiskSnapshot.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active')
    object.active = value if not value.nil?
    value = reader.get_attribute('actual_size')
    object.actual_size = value if not value.nil?
    value = reader.get_attribute('alias')
    object.alias_ = value if not value.nil?
    value = reader.get_attribute('bootable')
    object.bootable = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('format')
    object.format = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('image_id')
    object.image_id = value if not value.nil?
    value = reader.get_attribute('interface')
    object.interface = value if not value.nil?
    value = reader.get_attribute('logical_name')
    object.logical_name = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('propagate_errors')
    object.propagate_errors = value if not value.nil?
    value = reader.get_attribute('provisioned_size')
    object.provisioned_size = value if not value.nil?
    value = reader.get_attribute('read_only')
    object.read_only = value if not value.nil?
    value = reader.get_attribute('sgio')
    object.sgio = value if not value.nil?
    value = reader.get_attribute('shareable')
    object.shareable = value if not value.nil?
    value = reader.get_attribute('sparse')
    object.sparse = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('storage_type')
    object.storage_type = value if not value.nil?
    value = reader.get_attribute('uses_scsi_reservation')
    object.uses_scsi_reservation = value if not value.nil?
    value = reader.get_attribute('wipe_after_delete')
    object.wipe_after_delete = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active'
        object.active = Reader.read_boolean(reader)
      when 'actual_size'
        object.actual_size = Reader.read_integer(reader)
      when 'alias'
        object.alias_ = Reader.read_string(reader)
      when 'bootable'
        object.bootable = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'format'
        object.format = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'image_id'
        object.image_id = Reader.read_string(reader)
      when 'interface'
        object.interface = Reader.read_string(reader)
      when 'logical_name'
        object.logical_name = Reader.read_string(reader)
      when 'lun_storage'
        object.lun_storage = HostStorageReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'propagate_errors'
        object.propagate_errors = Reader.read_boolean(reader)
      when 'provisioned_size'
        object.provisioned_size = Reader.read_integer(reader)
      when 'read_only'
        object.read_only = Reader.read_boolean(reader)
      when 'sgio'
        object.sgio = Reader.read_string(reader)
      when 'shareable'
        object.shareable = Reader.read_boolean(reader)
      when 'sparse'
        object.sparse = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'storage_type'
        object.storage_type = Reader.read_string(reader)
      when 'uses_scsi_reservation'
        object.uses_scsi_reservation = Reader.read_boolean(reader)
      when 'wipe_after_delete'
        object.wipe_after_delete = Reader.read_boolean(reader)
      when 'disk'
        object.disk = DiskReader.read_one(reader)
      when 'disk_profile'
        object.disk_profile = DiskProfileReader.read_one(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'openstack_volume_type'
        object.openstack_volume_type = OpenStackVolumeTypeReader.read_one(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'quota'
        object.quota = QuotaReader.read_one(reader)
      when 'snapshot'
        object.snapshot = SnapshotReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'storage_domains'
        object.storage_domains = StorageDomainReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      when 'statistics'
        object.statistics = list
      when 'storagedomains'
        object.storage_domains = list
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class DisplayReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Display.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('allow_override')
    object.allow_override = value if not value.nil?
    value = reader.get_attribute('copy_paste_enabled')
    object.copy_paste_enabled = value if not value.nil?
    value = reader.get_attribute('disconnect_action')
    object.disconnect_action = value if not value.nil?
    value = reader.get_attribute('file_transfer_enabled')
    object.file_transfer_enabled = value if not value.nil?
    value = reader.get_attribute('keyboard_layout')
    object.keyboard_layout = value if not value.nil?
    value = reader.get_attribute('monitors')
    object.monitors = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('proxy')
    object.proxy = value if not value.nil?
    value = reader.get_attribute('secure_port')
    object.secure_port = value if not value.nil?
    value = reader.get_attribute('single_qxl_pci')
    object.single_qxl_pci = value if not value.nil?
    value = reader.get_attribute('smartcard_enabled')
    object.smartcard_enabled = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'allow_override'
        object.allow_override = Reader.read_boolean(reader)
      when 'certificate'
        object.certificate = CertificateReader.read_one(reader)
      when 'copy_paste_enabled'
        object.copy_paste_enabled = Reader.read_boolean(reader)
      when 'disconnect_action'
        object.disconnect_action = Reader.read_string(reader)
      when 'file_transfer_enabled'
        object.file_transfer_enabled = Reader.read_boolean(reader)
      when 'keyboard_layout'
        object.keyboard_layout = Reader.read_string(reader)
      when 'monitors'
        object.monitors = Reader.read_integer(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'proxy'
        object.proxy = Reader.read_string(reader)
      when 'secure_port'
        object.secure_port = Reader.read_integer(reader)
      when 'single_qxl_pci'
        object.single_qxl_pci = Reader.read_boolean(reader)
      when 'smartcard_enabled'
        object.smartcard_enabled = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class DnsReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Dns.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'search_domains'
        object.search_domains = HostReader.read_many(reader)
      when 'servers'
        object.servers = HostReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class DomainReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Domain.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      when 'groups'
        object.groups = GroupReader.read_many(reader)
      when 'users'
        object.users = UserReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'groups'
        object.groups = list
      when 'users'
        object.users = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class EntityProfileDetailReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = EntityProfileDetail.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'profile_details'
        object.profile_details = ProfileDetailReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ErrorHandlingReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ErrorHandling.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('on_error')
    object.on_error = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'on_error'
        object.on_error = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class EventReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Event.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('code')
    object.code = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('correlation_id')
    object.correlation_id = value if not value.nil?
    value = reader.get_attribute('custom_data')
    object.custom_data = value if not value.nil?
    value = reader.get_attribute('custom_id')
    object.custom_id = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('flood_rate')
    object.flood_rate = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('origin')
    object.origin = value if not value.nil?
    value = reader.get_attribute('severity')
    object.severity = value if not value.nil?
    value = reader.get_attribute('time')
    object.time = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'code'
        object.code = Reader.read_integer(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'correlation_id'
        object.correlation_id = Reader.read_string(reader)
      when 'custom_data'
        object.custom_data = Reader.read_string(reader)
      when 'custom_id'
        object.custom_id = Reader.read_integer(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'flood_rate'
        object.flood_rate = Reader.read_integer(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'origin'
        object.origin = Reader.read_string(reader)
      when 'severity'
        object.severity = Reader.read_string(reader)
      when 'time'
        object.time = Reader.read_date(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ExternalComputeResourceReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ExternalComputeResource.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('provider')
    object.provider = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('user')
    object.user = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'provider'
        object.provider = Reader.read_string(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'user'
        object.user = Reader.read_string(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ExternalDiscoveredHostReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ExternalDiscoveredHost.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('ip')
    object.ip = value if not value.nil?
    value = reader.get_attribute('last_report')
    object.last_report = value if not value.nil?
    value = reader.get_attribute('mac')
    object.mac = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('subnet_name')
    object.subnet_name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ip'
        object.ip = Reader.read_string(reader)
      when 'last_report'
        object.last_report = Reader.read_string(reader)
      when 'mac'
        object.mac = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'subnet_name'
        object.subnet_name = Reader.read_string(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ExternalHostReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ExternalHost.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ExternalHostGroupReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ExternalHostGroup.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('architecture_name')
    object.architecture_name = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('domain_name')
    object.domain_name = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('operating_system_name')
    object.operating_system_name = value if not value.nil?
    value = reader.get_attribute('subnet_name')
    object.subnet_name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'architecture_name'
        object.architecture_name = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'domain_name'
        object.domain_name = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'operating_system_name'
        object.operating_system_name = Reader.read_string(reader)
      when 'subnet_name'
        object.subnet_name = Reader.read_string(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ExternalHostProviderReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ExternalHostProvider.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_url')
    object.authentication_url = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('requires_authentication')
    object.requires_authentication = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authentication_url'
        object.authentication_url = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'requires_authentication'
        object.requires_authentication = Reader.read_boolean(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'certificates'
        object.certificates = CertificateReader.read_many(reader)
      when 'compute_resources'
        object.compute_resources = ExternalComputeResourceReader.read_many(reader)
      when 'discovered_hosts'
        object.discovered_hosts = ExternalDiscoveredHostReader.read_many(reader)
      when 'host_groups'
        object.host_groups = ExternalHostGroupReader.read_many(reader)
      when 'hosts'
        object.hosts = HostReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'certificates'
        object.certificates = list
      when 'computeresources'
        object.compute_resources = list
      when 'discoveredhosts'
        object.discovered_hosts = list
      when 'hostgroups'
        object.host_groups = list
      when 'hosts'
        object.hosts = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class ExternalProviderReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ExternalProvider.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_url')
    object.authentication_url = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('requires_authentication')
    object.requires_authentication = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authentication_url'
        object.authentication_url = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'requires_authentication'
        object.requires_authentication = Reader.read_boolean(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class FaultReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Fault.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('detail')
    object.detail = value if not value.nil?
    value = reader.get_attribute('reason')
    object.reason = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'detail'
        object.detail = Reader.read_string(reader)
      when 'reason'
        object.reason = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class FencingPolicyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = FencingPolicy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'skip_if_connectivity_broken'
        object.skip_if_connectivity_broken = SkipIfConnectivityBrokenReader.read_one(reader)
      when 'skip_if_sd_active'
        object.skip_if_sd_active = SkipIfSdActiveReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class FileReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = File.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('content')
    object.content = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'content'
        object.content = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class FilterReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Filter.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('position')
    object.position = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'position'
        object.position = Reader.read_integer(reader)
      when 'scheduling_policy_unit'
        object.scheduling_policy_unit = SchedulingPolicyUnitReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class FloppyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Floppy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'file'
        object.file = FileReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class FopStatisticReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = FopStatistic.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GlusterBrickReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterBrick.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('brick_dir')
    object.brick_dir = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('device')
    object.device = value if not value.nil?
    value = reader.get_attribute('fs_name')
    object.fs_name = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('mnt_options')
    object.mnt_options = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('pid')
    object.pid = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('server_id')
    object.server_id = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'brick_dir'
        object.brick_dir = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'device'
        object.device = Reader.read_string(reader)
      when 'fs_name'
        object.fs_name = Reader.read_string(reader)
      when 'gluster_clients'
        object.gluster_clients = GlusterClientReader.read_many(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'memory_pools'
        object.memory_pools = GlusterMemoryPoolReader.read_many(reader)
      when 'mnt_options'
        object.mnt_options = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'pid'
        object.pid = Reader.read_integer(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'server_id'
        object.server_id = Reader.read_string(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'gluster_volume'
        object.gluster_volume = GlusterVolumeReader.read_one(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'statistics'
        object.statistics = list
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class GlusterBrickAdvancedDetailsReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterBrickAdvancedDetails.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('device')
    object.device = value if not value.nil?
    value = reader.get_attribute('fs_name')
    object.fs_name = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('mnt_options')
    object.mnt_options = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('pid')
    object.pid = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'device'
        object.device = Reader.read_string(reader)
      when 'fs_name'
        object.fs_name = Reader.read_string(reader)
      when 'gluster_clients'
        object.gluster_clients = GlusterClientReader.read_many(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'memory_pools'
        object.memory_pools = GlusterMemoryPoolReader.read_many(reader)
      when 'mnt_options'
        object.mnt_options = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'pid'
        object.pid = Reader.read_integer(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class GlusterBrickMemoryInfoReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterBrickMemoryInfo.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'memory_pools'
        object.memory_pools = GlusterMemoryPoolReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GlusterClientReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterClient.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('bytes_read')
    object.bytes_read = value if not value.nil?
    value = reader.get_attribute('bytes_written')
    object.bytes_written = value if not value.nil?
    value = reader.get_attribute('client_port')
    object.client_port = value if not value.nil?
    value = reader.get_attribute('host_name')
    object.host_name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bytes_read'
        object.bytes_read = Reader.read_integer(reader)
      when 'bytes_written'
        object.bytes_written = Reader.read_integer(reader)
      when 'client_port'
        object.client_port = Reader.read_integer(reader)
      when 'host_name'
        object.host_name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GlusterHookReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterHook.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('checksum')
    object.checksum = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('conflict_status')
    object.conflict_status = value if not value.nil?
    value = reader.get_attribute('conflicts')
    object.conflicts = value if not value.nil?
    value = reader.get_attribute('content')
    object.content = value if not value.nil?
    value = reader.get_attribute('content_type')
    object.content_type = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('gluster_command')
    object.gluster_command = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('stage')
    object.stage = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'checksum'
        object.checksum = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'conflict_status'
        object.conflict_status = Reader.read_integer(reader)
      when 'conflicts'
        object.conflicts = Reader.read_string(reader)
      when 'content'
        object.content = Reader.read_string(reader)
      when 'content_type'
        object.content_type = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'gluster_command'
        object.gluster_command = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'stage'
        object.stage = Reader.read_string(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'server_hooks'
        object.server_hooks = GlusterServerHookReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'serverhooks'
        object.server_hooks = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class GlusterMemoryPoolReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterMemoryPool.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('alloc_count')
    object.alloc_count = value if not value.nil?
    value = reader.get_attribute('cold_count')
    object.cold_count = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('hot_count')
    object.hot_count = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('max_alloc')
    object.max_alloc = value if not value.nil?
    value = reader.get_attribute('max_stdalloc')
    object.max_stdalloc = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('padded_size')
    object.padded_size = value if not value.nil?
    value = reader.get_attribute('pool_misses')
    object.pool_misses = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'alloc_count'
        object.alloc_count = Reader.read_integer(reader)
      when 'cold_count'
        object.cold_count = Reader.read_integer(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'hot_count'
        object.hot_count = Reader.read_integer(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'max_alloc'
        object.max_alloc = Reader.read_integer(reader)
      when 'max_stdalloc'
        object.max_stdalloc = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'padded_size'
        object.padded_size = Reader.read_integer(reader)
      when 'pool_misses'
        object.pool_misses = Reader.read_integer(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GlusterServerHookReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterServerHook.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('checksum')
    object.checksum = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('content_type')
    object.content_type = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'checksum'
        object.checksum = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'content_type'
        object.content_type = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GlusterVolumeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterVolume.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('disperse_count')
    object.disperse_count = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('redundancy_count')
    object.redundancy_count = value if not value.nil?
    value = reader.get_attribute('replica_count')
    object.replica_count = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('stripe_count')
    object.stripe_count = value if not value.nil?
    value = reader.get_attribute('volume_type')
    object.volume_type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'disperse_count'
        object.disperse_count = Reader.read_integer(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'options'
        object.options = OptionReader.read_many(reader)
      when 'redundancy_count'
        object.redundancy_count = Reader.read_integer(reader)
      when 'replica_count'
        object.replica_count = Reader.read_integer(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'stripe_count'
        object.stripe_count = Reader.read_integer(reader)
      when 'transport_types'
        reader.next_element
      when 'volume_type'
        object.volume_type = Reader.read_string(reader)
      when 'bricks'
        object.bricks = GlusterBrickReader.read_many(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'bricks'
        object.bricks = list
      when 'statistics'
        object.statistics = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class GlusterVolumeProfileDetailsReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GlusterVolumeProfileDetails.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'brick_profile_details'
        object.brick_profile_details = BrickProfileDetailReader.read_many(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'nfs_profile_details'
        object.nfs_profile_details = NfsProfileDetailReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GracePeriodReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GracePeriod.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('expiry')
    object.expiry = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'expiry'
        object.expiry = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GraphicsConsoleReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GraphicsConsole.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('protocol')
    object.protocol = value if not value.nil?
    value = reader.get_attribute('tls_port')
    object.tls_port = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'protocol'
        object.protocol = Reader.read_string(reader)
      when 'tls_port'
        object.tls_port = Reader.read_integer(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class GroupReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Group.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('domain_entry_id')
    object.domain_entry_id = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('namespace')
    object.namespace = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'domain_entry_id'
        object.domain_entry_id = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'namespace'
        object.namespace = Reader.read_string(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'roles'
        object.roles = RoleReader.read_many(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      when 'roles'
        object.roles = list
      when 'tags'
        object.tags = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class GuestOperatingSystemReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = GuestOperatingSystem.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('architecture')
    object.architecture = value if not value.nil?
    value = reader.get_attribute('codename')
    object.codename = value if not value.nil?
    value = reader.get_attribute('distribution')
    object.distribution = value if not value.nil?
    value = reader.get_attribute('family')
    object.family = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'architecture'
        object.architecture = Reader.read_string(reader)
      when 'codename'
        object.codename = Reader.read_string(reader)
      when 'distribution'
        object.distribution = Reader.read_string(reader)
      when 'family'
        object.family = Reader.read_string(reader)
      when 'kernel'
        object.kernel = KernelReader.read_one(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HardwareInformationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HardwareInformation.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('family')
    object.family = value if not value.nil?
    value = reader.get_attribute('manufacturer')
    object.manufacturer = value if not value.nil?
    value = reader.get_attribute('product_name')
    object.product_name = value if not value.nil?
    value = reader.get_attribute('serial_number')
    object.serial_number = value if not value.nil?
    value = reader.get_attribute('uuid')
    object.uuid = value if not value.nil?
    value = reader.get_attribute('version')
    object.version = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'family'
        object.family = Reader.read_string(reader)
      when 'manufacturer'
        object.manufacturer = Reader.read_string(reader)
      when 'product_name'
        object.product_name = Reader.read_string(reader)
      when 'serial_number'
        object.serial_number = Reader.read_string(reader)
      when 'supported_rng_sources'
        reader.next_element
      when 'uuid'
        object.uuid = Reader.read_string(reader)
      when 'version'
        object.version = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HighAvailabilityReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HighAvailability.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    value = reader.get_attribute('priority')
    object.priority = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'priority'
        object.priority = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HookReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Hook.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('event_name')
    object.event_name = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('md5')
    object.md5 = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'event_name'
        object.event_name = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'md5'
        object.md5 = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HostReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Host.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('auto_numa_status')
    object.auto_numa_status = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('external_status')
    object.external_status = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('kdump_status')
    object.kdump_status = value if not value.nil?
    value = reader.get_attribute('max_scheduling_memory')
    object.max_scheduling_memory = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('numa_supported')
    object.numa_supported = value if not value.nil?
    value = reader.get_attribute('override_iptables')
    object.override_iptables = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('protocol')
    object.protocol = value if not value.nil?
    value = reader.get_attribute('root_password')
    object.root_password = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('status_detail')
    object.status_detail = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('update_available')
    object.update_available = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'auto_numa_status'
        object.auto_numa_status = Reader.read_string(reader)
      when 'certificate'
        object.certificate = CertificateReader.read_one(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'device_passthrough'
        object.device_passthrough = HostDevicePassthroughReader.read_one(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'external_status'
        object.external_status = Reader.read_string(reader)
      when 'hardware_information'
        object.hardware_information = HardwareInformationReader.read_one(reader)
      when 'hosted_engine'
        object.hosted_engine = HostedEngineReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'iscsi'
        object.iscsi = IscsiDetailsReader.read_one(reader)
      when 'kdump_status'
        object.kdump_status = Reader.read_string(reader)
      when 'ksm'
        object.ksm = KsmReader.read_one(reader)
      when 'libvirt_version'
        object.libvirt_version = VersionReader.read_one(reader)
      when 'max_scheduling_memory'
        object.max_scheduling_memory = Reader.read_integer(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'numa_supported'
        object.numa_supported = Reader.read_boolean(reader)
      when 'os'
        object.os = OperatingSystemReader.read_one(reader)
      when 'override_iptables'
        object.override_iptables = Reader.read_boolean(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'power_management'
        object.power_management = PowerManagementReader.read_one(reader)
      when 'protocol'
        object.protocol = Reader.read_string(reader)
      when 'root_password'
        object.root_password = Reader.read_string(reader)
      when 'se_linux'
        object.se_linux = SeLinuxReader.read_one(reader)
      when 'spm'
        object.spm = SpmReader.read_one(reader)
      when 'ssh'
        object.ssh = SshReader.read_one(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'status_detail'
        object.status_detail = Reader.read_string(reader)
      when 'summary'
        object.summary = VmSummaryReader.read_one(reader)
      when 'transparent_hugepages'
        object.transparent_huge_pages = TransparentHugePagesReader.read_one(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'update_available'
        object.update_available = Reader.read_boolean(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      when 'affinity_labels'
        object.affinity_labels = AffinityLabelReader.read_many(reader)
      when 'agents'
        object.agents = AgentReader.read_many(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'devices'
        object.devices = DeviceReader.read_many(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      when 'hooks'
        object.hooks = HookReader.read_many(reader)
      when 'katello_errata'
        object.katello_errata = KatelloErratumReader.read_many(reader)
      when 'network_attachments'
        object.network_attachments = NetworkAttachmentReader.read_many(reader)
      when 'nics'
        object.nics = NicReader.read_many(reader)
      when 'host_numa_nodes'
        object.numa_nodes = NumaNodeReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'storage_connection_extensions'
        object.storage_connection_extensions = StorageConnectionExtensionReader.read_many(reader)
      when 'storages'
        object.storages = HostStorageReader.read_many(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'unmanaged_networks'
        object.unmanaged_networks = UnmanagedNetworkReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'affinitylabels'
        object.affinity_labels = list
      when 'agents'
        object.agents = list
      when 'devices'
        object.devices = list
      when 'hooks'
        object.hooks = list
      when 'katelloerrata'
        object.katello_errata = list
      when 'networkattachments'
        object.network_attachments = list
      when 'nics'
        object.nics = list
      when 'numanodes'
        object.numa_nodes = list
      when 'permissions'
        object.permissions = list
      when 'statistics'
        object.statistics = list
      when 'storageconnectionextensions'
        object.storage_connection_extensions = list
      when 'storages'
        object.storages = list
      when 'tags'
        object.tags = list
      when 'unmanagednetworks'
        object.unmanaged_networks = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class HostDeviceReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HostDevice.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('capability')
    object.capability = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('iommu_group')
    object.iommu_group = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('placeholder')
    object.placeholder = value if not value.nil?
    value = reader.get_attribute('virtual_functions')
    object.virtual_functions = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'capability'
        object.capability = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'iommu_group'
        object.iommu_group = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'physical_function'
        object.physical_function = HostDeviceReader.read_one(reader)
      when 'placeholder'
        object.placeholder = Reader.read_boolean(reader)
      when 'product'
        object.product = ProductReader.read_one(reader)
      when 'vendor'
        object.vendor = VendorReader.read_one(reader)
      when 'virtual_functions'
        object.virtual_functions = Reader.read_integer(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'parent_device'
        object.parent_device = HostDeviceReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HostDevicePassthroughReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HostDevicePassthrough.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HostNicReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HostNic.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('base_interface')
    object.base_interface = value if not value.nil?
    value = reader.get_attribute('boot_protocol')
    object.boot_protocol = value if not value.nil?
    value = reader.get_attribute('bridged')
    object.bridged = value if not value.nil?
    value = reader.get_attribute('check_connectivity')
    object.check_connectivity = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('custom_configuration')
    object.custom_configuration = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('ipv6_boot_protocol')
    object.ipv6_boot_protocol = value if not value.nil?
    value = reader.get_attribute('mtu')
    object.mtu = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('override_configuration')
    object.override_configuration = value if not value.nil?
    value = reader.get_attribute('speed')
    object.speed = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'base_interface'
        object.base_interface = Reader.read_string(reader)
      when 'bonding'
        object.bonding = BondingReader.read_one(reader)
      when 'boot_protocol'
        object.boot_protocol = Reader.read_string(reader)
      when 'bridged'
        object.bridged = Reader.read_boolean(reader)
      when 'check_connectivity'
        object.check_connectivity = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'custom_configuration'
        object.custom_configuration = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ip'
        object.ip = IpReader.read_one(reader)
      when 'ipv6'
        object.ipv6 = IpReader.read_one(reader)
      when 'ipv6_boot_protocol'
        object.ipv6_boot_protocol = Reader.read_string(reader)
      when 'mac'
        object.mac = MacReader.read_one(reader)
      when 'mtu'
        object.mtu = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'network_labels'
        object.network_labels = NetworkLabelReader.read_many(reader)
      when 'override_configuration'
        object.override_configuration = Reader.read_boolean(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'speed'
        object.speed = Reader.read_integer(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'virtual_functions_configuration'
        object.virtual_functions_configuration = HostNicVirtualFunctionsConfigurationReader.read_one(reader)
      when 'vlan'
        object.vlan = VlanReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'network'
        object.network = NetworkReader.read_one(reader)
      when 'physical_function'
        object.physical_function = HostNicReader.read_one(reader)
      when 'qos'
        object.qos = QosReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HostNicVirtualFunctionsConfigurationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HostNicVirtualFunctionsConfiguration.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('all_networks_allowed')
    object.all_networks_allowed = value if not value.nil?
    value = reader.get_attribute('max_number_of_virtual_functions')
    object.max_number_of_virtual_functions = value if not value.nil?
    value = reader.get_attribute('number_of_virtual_functions')
    object.number_of_virtual_functions = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'all_networks_allowed'
        object.all_networks_allowed = Reader.read_boolean(reader)
      when 'max_number_of_virtual_functions'
        object.max_number_of_virtual_functions = Reader.read_integer(reader)
      when 'number_of_virtual_functions'
        object.number_of_virtual_functions = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HostStorageReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HostStorage.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('mount_options')
    object.mount_options = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('nfs_retrans')
    object.nfs_retrans = value if not value.nil?
    value = reader.get_attribute('nfs_timeo')
    object.nfs_timeo = value if not value.nil?
    value = reader.get_attribute('nfs_version')
    object.nfs_version = value if not value.nil?
    value = reader.get_attribute('override_luns')
    object.override_luns = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('path')
    object.path = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('portal')
    object.portal = value if not value.nil?
    value = reader.get_attribute('target')
    object.target = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    value = reader.get_attribute('vfs_type')
    object.vfs_type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'logical_units'
        object.logical_units = LogicalUnitReader.read_many(reader)
      when 'mount_options'
        object.mount_options = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'nfs_retrans'
        object.nfs_retrans = Reader.read_integer(reader)
      when 'nfs_timeo'
        object.nfs_timeo = Reader.read_integer(reader)
      when 'nfs_version'
        object.nfs_version = Reader.read_string(reader)
      when 'override_luns'
        object.override_luns = Reader.read_boolean(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'path'
        object.path = Reader.read_string(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'portal'
        object.portal = Reader.read_string(reader)
      when 'target'
        object.target = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'vfs_type'
        object.vfs_type = Reader.read_string(reader)
      when 'volume_group'
        object.volume_group = VolumeGroupReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class HostedEngineReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = HostedEngine.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active')
    object.active = value if not value.nil?
    value = reader.get_attribute('configured')
    object.configured = value if not value.nil?
    value = reader.get_attribute('global_maintenance')
    object.global_maintenance = value if not value.nil?
    value = reader.get_attribute('local_maintenance')
    object.local_maintenance = value if not value.nil?
    value = reader.get_attribute('score')
    object.score = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active'
        object.active = Reader.read_boolean(reader)
      when 'configured'
        object.configured = Reader.read_boolean(reader)
      when 'global_maintenance'
        object.global_maintenance = Reader.read_boolean(reader)
      when 'local_maintenance'
        object.local_maintenance = Reader.read_boolean(reader)
      when 'score'
        object.score = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class IconReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Icon.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('data')
    object.data = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('media_type')
    object.media_type = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'data'
        object.data = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'media_type'
        object.media_type = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class IdentifiedReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Identified.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ImageReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Image.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class InitializationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Initialization.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active_directory_ou')
    object.active_directory_ou = value if not value.nil?
    value = reader.get_attribute('authorized_ssh_keys')
    object.authorized_ssh_keys = value if not value.nil?
    value = reader.get_attribute('custom_script')
    object.custom_script = value if not value.nil?
    value = reader.get_attribute('dns_search')
    object.dns_search = value if not value.nil?
    value = reader.get_attribute('dns_servers')
    object.dns_servers = value if not value.nil?
    value = reader.get_attribute('domain')
    object.domain = value if not value.nil?
    value = reader.get_attribute('host_name')
    object.host_name = value if not value.nil?
    value = reader.get_attribute('input_locale')
    object.input_locale = value if not value.nil?
    value = reader.get_attribute('org_name')
    object.org_name = value if not value.nil?
    value = reader.get_attribute('regenerate_ids')
    object.regenerate_ids = value if not value.nil?
    value = reader.get_attribute('regenerate_ssh_keys')
    object.regenerate_ssh_keys = value if not value.nil?
    value = reader.get_attribute('root_password')
    object.root_password = value if not value.nil?
    value = reader.get_attribute('system_locale')
    object.system_locale = value if not value.nil?
    value = reader.get_attribute('timezone')
    object.timezone = value if not value.nil?
    value = reader.get_attribute('ui_language')
    object.ui_language = value if not value.nil?
    value = reader.get_attribute('user_locale')
    object.user_locale = value if not value.nil?
    value = reader.get_attribute('user_name')
    object.user_name = value if not value.nil?
    value = reader.get_attribute('windows_license_key')
    object.windows_license_key = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active_directory_ou'
        object.active_directory_ou = Reader.read_string(reader)
      when 'authorized_ssh_keys'
        object.authorized_ssh_keys = Reader.read_string(reader)
      when 'cloud_init'
        object.cloud_init = CloudInitReader.read_one(reader)
      when 'configuration'
        object.configuration = ConfigurationReader.read_one(reader)
      when 'custom_script'
        object.custom_script = Reader.read_string(reader)
      when 'dns_search'
        object.dns_search = Reader.read_string(reader)
      when 'dns_servers'
        object.dns_servers = Reader.read_string(reader)
      when 'domain'
        object.domain = Reader.read_string(reader)
      when 'host_name'
        object.host_name = Reader.read_string(reader)
      when 'input_locale'
        object.input_locale = Reader.read_string(reader)
      when 'nic_configurations'
        object.nic_configurations = NicConfigurationReader.read_many(reader)
      when 'org_name'
        object.org_name = Reader.read_string(reader)
      when 'regenerate_ids'
        object.regenerate_ids = Reader.read_boolean(reader)
      when 'regenerate_ssh_keys'
        object.regenerate_ssh_keys = Reader.read_boolean(reader)
      when 'root_password'
        object.root_password = Reader.read_string(reader)
      when 'system_locale'
        object.system_locale = Reader.read_string(reader)
      when 'timezone'
        object.timezone = Reader.read_string(reader)
      when 'ui_language'
        object.ui_language = Reader.read_string(reader)
      when 'user_locale'
        object.user_locale = Reader.read_string(reader)
      when 'user_name'
        object.user_name = Reader.read_string(reader)
      when 'windows_license_key'
        object.windows_license_key = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class InstanceTypeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = InstanceType.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('cpu_shares')
    object.cpu_shares = value if not value.nil?
    value = reader.get_attribute('creation_time')
    object.creation_time = value if not value.nil?
    value = reader.get_attribute('custom_cpu_model')
    object.custom_cpu_model = value if not value.nil?
    value = reader.get_attribute('custom_emulated_machine')
    object.custom_emulated_machine = value if not value.nil?
    value = reader.get_attribute('delete_protected')
    object.delete_protected = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('migration_downtime')
    object.migration_downtime = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('origin')
    object.origin = value if not value.nil?
    value = reader.get_attribute('soundcard_enabled')
    object.soundcard_enabled = value if not value.nil?
    value = reader.get_attribute('start_paused')
    object.start_paused = value if not value.nil?
    value = reader.get_attribute('stateless')
    object.stateless = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('tunnel_migration')
    object.tunnel_migration = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bios'
        object.bios = BiosReader.read_one(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'console'
        object.console = ConsoleReader.read_one(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'cpu_shares'
        object.cpu_shares = Reader.read_integer(reader)
      when 'creation_time'
        object.creation_time = Reader.read_date(reader)
      when 'custom_compatibility_version'
        object.custom_compatibility_version = VersionReader.read_one(reader)
      when 'custom_cpu_model'
        object.custom_cpu_model = Reader.read_string(reader)
      when 'custom_emulated_machine'
        object.custom_emulated_machine = Reader.read_string(reader)
      when 'custom_properties'
        object.custom_properties = CustomPropertyReader.read_many(reader)
      when 'delete_protected'
        object.delete_protected = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'high_availability'
        object.high_availability = HighAvailabilityReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'initialization'
        object.initialization = InitializationReader.read_one(reader)
      when 'io'
        object.io = IoReader.read_one(reader)
      when 'large_icon'
        object.large_icon = IconReader.read_one(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'memory_policy'
        object.memory_policy = MemoryPolicyReader.read_one(reader)
      when 'migration'
        object.migration = MigrationOptionsReader.read_one(reader)
      when 'migration_downtime'
        object.migration_downtime = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'origin'
        object.origin = Reader.read_string(reader)
      when 'os'
        object.os = OperatingSystemReader.read_one(reader)
      when 'rng_device'
        object.rng_device = RngDeviceReader.read_one(reader)
      when 'serial_number'
        object.serial_number = SerialNumberReader.read_one(reader)
      when 'small_icon'
        object.small_icon = IconReader.read_one(reader)
      when 'soundcard_enabled'
        object.soundcard_enabled = Reader.read_boolean(reader)
      when 'sso'
        object.sso = SsoReader.read_one(reader)
      when 'start_paused'
        object.start_paused = Reader.read_boolean(reader)
      when 'stateless'
        object.stateless = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'time_zone'
        object.time_zone = TimeZoneReader.read_one(reader)
      when 'tunnel_migration'
        object.tunnel_migration = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'usb'
        object.usb = UsbReader.read_one(reader)
      when 'version'
        object.version = TemplateVersionReader.read_one(reader)
      when 'virtio_scsi'
        object.virtio_scsi = VirtioScsiReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'cdroms'
        object.cdroms = CdromReader.read_many(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'cpu_profile'
        object.cpu_profile = CpuProfileReader.read_one(reader)
      when 'disk_attachments'
        object.disk_attachments = DiskAttachmentReader.read_many(reader)
      when 'graphics_consoles'
        object.graphics_consoles = GraphicsConsoleReader.read_many(reader)
      when 'nics'
        object.nics = NicReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'watchdogs'
        object.watchdogs = WatchdogReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'cdroms'
        object.cdroms = list
      when 'diskattachments'
        object.disk_attachments = list
      when 'graphicsconsoles'
        object.graphics_consoles = list
      when 'nics'
        object.nics = list
      when 'permissions'
        object.permissions = list
      when 'tags'
        object.tags = list
      when 'watchdogs'
        object.watchdogs = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class IoReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Io.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('threads')
    object.threads = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'threads'
        object.threads = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class IpReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Ip.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('gateway')
    object.gateway = value if not value.nil?
    value = reader.get_attribute('netmask')
    object.netmask = value if not value.nil?
    value = reader.get_attribute('version')
    object.version = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'gateway'
        object.gateway = Reader.read_string(reader)
      when 'netmask'
        object.netmask = Reader.read_string(reader)
      when 'version'
        object.version = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class IpAddressAssignmentReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = IpAddressAssignment.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('assignment_method')
    object.assignment_method = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'assignment_method'
        object.assignment_method = Reader.read_string(reader)
      when 'ip'
        object.ip = IpReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class IscsiBondReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = IscsiBond.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'networks'
        object.networks = NetworkReader.read_many(reader)
      when 'storage_connections'
        object.storage_connections = StorageConnectionReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'networks'
        object.networks = list
      when 'storageconnections'
        object.storage_connections = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class IscsiDetailsReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = IscsiDetails.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('disk_id')
    object.disk_id = value if not value.nil?
    value = reader.get_attribute('initiator')
    object.initiator = value if not value.nil?
    value = reader.get_attribute('lun_mapping')
    object.lun_mapping = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('paths')
    object.paths = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('portal')
    object.portal = value if not value.nil?
    value = reader.get_attribute('product_id')
    object.product_id = value if not value.nil?
    value = reader.get_attribute('serial')
    object.serial = value if not value.nil?
    value = reader.get_attribute('size')
    object.size = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('storage_domain_id')
    object.storage_domain_id = value if not value.nil?
    value = reader.get_attribute('target')
    object.target = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    value = reader.get_attribute('vendor_id')
    object.vendor_id = value if not value.nil?
    value = reader.get_attribute('volume_group_id')
    object.volume_group_id = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'disk_id'
        object.disk_id = Reader.read_string(reader)
      when 'initiator'
        object.initiator = Reader.read_string(reader)
      when 'lun_mapping'
        object.lun_mapping = Reader.read_integer(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'paths'
        object.paths = Reader.read_integer(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'portal'
        object.portal = Reader.read_string(reader)
      when 'product_id'
        object.product_id = Reader.read_string(reader)
      when 'serial'
        object.serial = Reader.read_string(reader)
      when 'size'
        object.size = Reader.read_integer(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'storage_domain_id'
        object.storage_domain_id = Reader.read_string(reader)
      when 'target'
        object.target = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'vendor_id'
        object.vendor_id = Reader.read_string(reader)
      when 'volume_group_id'
        object.volume_group_id = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class JobReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Job.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('auto_cleared')
    object.auto_cleared = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('end_time')
    object.end_time = value if not value.nil?
    value = reader.get_attribute('external')
    object.external = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('last_updated')
    object.last_updated = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('start_time')
    object.start_time = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'auto_cleared'
        object.auto_cleared = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'end_time'
        object.end_time = Reader.read_date(reader)
      when 'external'
        object.external = Reader.read_boolean(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'last_updated'
        object.last_updated = Reader.read_date(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'start_time'
        object.start_time = Reader.read_date(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'owner'
        object.owner = UserReader.read_one(reader)
      when 'steps'
        object.steps = StepReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'steps'
        object.steps = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class KatelloErratumReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = KatelloErratum.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('issued')
    object.issued = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('severity')
    object.severity = value if not value.nil?
    value = reader.get_attribute('solution')
    object.solution = value if not value.nil?
    value = reader.get_attribute('summary')
    object.summary = value if not value.nil?
    value = reader.get_attribute('title')
    object.title = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'issued'
        object.issued = Reader.read_date(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'packages'
        object.packages = PackageReader.read_many(reader)
      when 'severity'
        object.severity = Reader.read_string(reader)
      when 'solution'
        object.solution = Reader.read_string(reader)
      when 'summary'
        object.summary = Reader.read_string(reader)
      when 'title'
        object.title = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class KernelReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Kernel.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'version'
        object.version = VersionReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class KsmReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Ksm.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    value = reader.get_attribute('merge_across_nodes')
    object.merge_across_nodes = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'merge_across_nodes'
        object.merge_across_nodes = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class LogicalUnitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = LogicalUnit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('disk_id')
    object.disk_id = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('lun_mapping')
    object.lun_mapping = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('paths')
    object.paths = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('portal')
    object.portal = value if not value.nil?
    value = reader.get_attribute('product_id')
    object.product_id = value if not value.nil?
    value = reader.get_attribute('serial')
    object.serial = value if not value.nil?
    value = reader.get_attribute('size')
    object.size = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('storage_domain_id')
    object.storage_domain_id = value if not value.nil?
    value = reader.get_attribute('target')
    object.target = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    value = reader.get_attribute('vendor_id')
    object.vendor_id = value if not value.nil?
    value = reader.get_attribute('volume_group_id')
    object.volume_group_id = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'disk_id'
        object.disk_id = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'lun_mapping'
        object.lun_mapping = Reader.read_integer(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'paths'
        object.paths = Reader.read_integer(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'portal'
        object.portal = Reader.read_string(reader)
      when 'product_id'
        object.product_id = Reader.read_string(reader)
      when 'serial'
        object.serial = Reader.read_string(reader)
      when 'size'
        object.size = Reader.read_integer(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'storage_domain_id'
        object.storage_domain_id = Reader.read_string(reader)
      when 'target'
        object.target = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'vendor_id'
        object.vendor_id = Reader.read_string(reader)
      when 'volume_group_id'
        object.volume_group_id = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MacReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Mac.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MacPoolReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = MacPool.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('allow_duplicates')
    object.allow_duplicates = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('default_pool')
    object.default_pool = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'allow_duplicates'
        object.allow_duplicates = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'default_pool'
        object.default_pool = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'ranges'
        object.ranges = RangeReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MemoryOverCommitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = MemoryOverCommit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('percent')
    object.percent = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'percent'
        object.percent = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MemoryPolicyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = MemoryPolicy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('ballooning')
    object.ballooning = value if not value.nil?
    value = reader.get_attribute('guaranteed')
    object.guaranteed = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'ballooning'
        object.ballooning = Reader.read_boolean(reader)
      when 'guaranteed'
        object.guaranteed = Reader.read_integer(reader)
      when 'over_commit'
        object.over_commit = MemoryOverCommitReader.read_one(reader)
      when 'transparent_hugepages'
        object.transparent_huge_pages = TransparentHugePagesReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MethodReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Method.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'id'
        object.id = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MigrationBandwidthReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = MigrationBandwidth.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('assignment_method')
    object.assignment_method = value if not value.nil?
    value = reader.get_attribute('custom_value')
    object.custom_value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'assignment_method'
        object.assignment_method = Reader.read_string(reader)
      when 'custom_value'
        object.custom_value = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MigrationOptionsReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = MigrationOptions.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('auto_converge')
    object.auto_converge = value if not value.nil?
    value = reader.get_attribute('compressed')
    object.compressed = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'auto_converge'
        object.auto_converge = Reader.read_string(reader)
      when 'bandwidth'
        object.bandwidth = MigrationBandwidthReader.read_one(reader)
      when 'compressed'
        object.compressed = Reader.read_string(reader)
      when 'policy'
        object.policy = MigrationPolicyReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class MigrationPolicyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = MigrationPolicy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NetworkReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Network.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('display')
    object.display = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('mtu')
    object.mtu = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('profile_required')
    object.profile_required = value if not value.nil?
    value = reader.get_attribute('required')
    object.required = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('stp')
    object.stp = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = Reader.read_boolean(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ip'
        object.ip = IpReader.read_one(reader)
      when 'mtu'
        object.mtu = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'profile_required'
        object.profile_required = Reader.read_boolean(reader)
      when 'required'
        object.required = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'stp'
        object.stp = Reader.read_boolean(reader)
      when 'usages'
        reader.next_element
      when 'vlan'
        object.vlan = VlanReader.read_one(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'network_labels'
        object.network_labels = NetworkLabelReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'qos'
        object.qos = QosReader.read_one(reader)
      when 'vnic_profiles'
        object.vnic_profiles = VnicProfileReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'networklabels'
        object.network_labels = list
      when 'permissions'
        object.permissions = list
      when 'vnicprofiles'
        object.vnic_profiles = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class NetworkAttachmentReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NetworkAttachment.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('in_sync')
    object.in_sync = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'in_sync'
        object.in_sync = Reader.read_boolean(reader)
      when 'ip_address_assignments'
        object.ip_address_assignments = IpAddressAssignmentReader.read_many(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'reported_configurations'
        object.reported_configurations = ReportedConfigurationReader.read_many(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'host_nic'
        object.host_nic = HostNicReader.read_one(reader)
      when 'network'
        object.network = NetworkReader.read_one(reader)
      when 'qos'
        object.qos = QosReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NetworkConfigurationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NetworkConfiguration.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'dns'
        object.dns = DnsReader.read_one(reader)
      when 'nics'
        object.nics = NicReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NetworkFilterReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NetworkFilter.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NetworkLabelReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NetworkLabel.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'host_nic'
        object.host_nic = HostNicReader.read_one(reader)
      when 'network'
        object.network = NetworkReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NfsProfileDetailReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NfsProfileDetail.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('nfs_server_ip')
    object.nfs_server_ip = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'nfs_server_ip'
        object.nfs_server_ip = Reader.read_string(reader)
      when 'profile_details'
        object.profile_details = ProfileDetailReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NicReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Nic.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('boot_protocol')
    object.boot_protocol = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('interface')
    object.interface = value if not value.nil?
    value = reader.get_attribute('linked')
    object.linked = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('on_boot')
    object.on_boot = value if not value.nil?
    value = reader.get_attribute('plugged')
    object.plugged = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'boot_protocol'
        object.boot_protocol = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'interface'
        object.interface = Reader.read_string(reader)
      when 'linked'
        object.linked = Reader.read_boolean(reader)
      when 'mac'
        object.mac = MacReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'on_boot'
        object.on_boot = Reader.read_boolean(reader)
      when 'plugged'
        object.plugged = Reader.read_boolean(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'network'
        object.network = NetworkReader.read_one(reader)
      when 'network_attachments'
        object.network_attachments = NetworkAttachmentReader.read_many(reader)
      when 'network_labels'
        object.network_labels = NetworkLabelReader.read_many(reader)
      when 'reported_devices'
        object.reported_devices = ReportedDeviceReader.read_many(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'virtual_function_allowed_labels'
        object.virtual_function_allowed_labels = NetworkLabelReader.read_many(reader)
      when 'virtual_function_allowed_networks'
        object.virtual_function_allowed_networks = NetworkReader.read_many(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'vnic_profile'
        object.vnic_profile = VnicProfileReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'networkattachments'
        object.network_attachments = list
      when 'networklabels'
        object.network_labels = list
      when 'reporteddevices'
        object.reported_devices = list
      when 'statistics'
        object.statistics = list
      when 'virtualfunctionallowedlabels'
        object.virtual_function_allowed_labels = list
      when 'virtualfunctionallowednetworks'
        object.virtual_function_allowed_networks = list
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class NicConfigurationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NicConfiguration.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('boot_protocol')
    object.boot_protocol = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('on_boot')
    object.on_boot = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'boot_protocol'
        object.boot_protocol = Reader.read_string(reader)
      when 'ip'
        object.ip = IpReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'on_boot'
        object.on_boot = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class NumaNodeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NumaNode.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('index')
    object.index = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('node_distance')
    object.node_distance = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'index'
        object.index = Reader.read_integer(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'node_distance'
        object.node_distance = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'statistics'
        object.statistics = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class NumaNodePinReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = NumaNodePin.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('index')
    object.index = value if not value.nil?
    value = reader.get_attribute('pinned')
    object.pinned = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'host_numa_node'
        object.host_numa_node = NumaNodeReader.read_one(reader)
      when 'index'
        object.index = Reader.read_integer(reader)
      when 'pinned'
        object.pinned = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OpenStackImageReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackImage.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'openstack_image_provider'
        object.openstack_image_provider = OpenStackImageProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OpenStackImageProviderReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackImageProvider.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_url')
    object.authentication_url = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('requires_authentication')
    object.requires_authentication = value if not value.nil?
    value = reader.get_attribute('tenant_name')
    object.tenant_name = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authentication_url'
        object.authentication_url = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'requires_authentication'
        object.requires_authentication = Reader.read_boolean(reader)
      when 'tenant_name'
        object.tenant_name = Reader.read_string(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'certificates'
        object.certificates = CertificateReader.read_many(reader)
      when 'images'
        object.images = OpenStackImageReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'certificates'
        object.certificates = list
      when 'images'
        object.images = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class OpenStackNetworkReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackNetwork.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'openstack_network_provider'
        object.openstack_network_provider = OpenStackNetworkProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OpenStackNetworkProviderReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackNetworkProvider.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_url')
    object.authentication_url = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('plugin_type')
    object.plugin_type = value if not value.nil?
    value = reader.get_attribute('read_only')
    object.read_only = value if not value.nil?
    value = reader.get_attribute('requires_authentication')
    object.requires_authentication = value if not value.nil?
    value = reader.get_attribute('tenant_name')
    object.tenant_name = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'agent_configuration'
        object.agent_configuration = AgentConfigurationReader.read_one(reader)
      when 'authentication_url'
        object.authentication_url = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'plugin_type'
        object.plugin_type = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'read_only'
        object.read_only = Reader.read_boolean(reader)
      when 'requires_authentication'
        object.requires_authentication = Reader.read_boolean(reader)
      when 'tenant_name'
        object.tenant_name = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'certificates'
        object.certificates = CertificateReader.read_many(reader)
      when 'networks'
        object.networks = OpenStackNetworkReader.read_many(reader)
      when 'subnets'
        object.subnets = OpenStackSubnetReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'certificates'
        object.certificates = list
      when 'networks'
        object.networks = list
      when 'subnets'
        object.subnets = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class OpenStackProviderReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackProvider.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_url')
    object.authentication_url = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('requires_authentication')
    object.requires_authentication = value if not value.nil?
    value = reader.get_attribute('tenant_name')
    object.tenant_name = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authentication_url'
        object.authentication_url = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'requires_authentication'
        object.requires_authentication = Reader.read_boolean(reader)
      when 'tenant_name'
        object.tenant_name = Reader.read_string(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OpenStackSubnetReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackSubnet.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('cidr')
    object.cidr = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('gateway')
    object.gateway = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('ip_version')
    object.ip_version = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'cidr'
        object.cidr = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'dns_servers'
        reader.next_element
      when 'gateway'
        object.gateway = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ip_version'
        object.ip_version = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'openstack_network'
        object.openstack_network = OpenStackNetworkReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OpenStackVolumeProviderReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackVolumeProvider.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_url')
    object.authentication_url = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('requires_authentication')
    object.requires_authentication = value if not value.nil?
    value = reader.get_attribute('tenant_name')
    object.tenant_name = value if not value.nil?
    value = reader.get_attribute('url')
    object.url = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authentication_url'
        object.authentication_url = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'requires_authentication'
        object.requires_authentication = Reader.read_boolean(reader)
      when 'tenant_name'
        object.tenant_name = Reader.read_string(reader)
      when 'url'
        object.url = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'authentication_keys'
        object.authentication_keys = OpenstackVolumeAuthenticationKeyReader.read_many(reader)
      when 'certificates'
        object.certificates = CertificateReader.read_many(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'volume_types'
        object.volume_types = OpenStackVolumeTypeReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'authenticationkeys'
        object.authentication_keys = list
      when 'certificates'
        object.certificates = list
      when 'volumetypes'
        object.volume_types = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class OpenStackVolumeTypeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenStackVolumeType.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'openstack_volume_provider'
        object.openstack_volume_provider = OpenStackVolumeProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OpenstackVolumeAuthenticationKeyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OpenstackVolumeAuthenticationKey.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('creation_date')
    object.creation_date = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('usage_type')
    object.usage_type = value if not value.nil?
    value = reader.get_attribute('uuid')
    object.uuid = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'creation_date'
        object.creation_date = Reader.read_date(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'usage_type'
        object.usage_type = Reader.read_string(reader)
      when 'uuid'
        object.uuid = Reader.read_string(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      when 'openstack_volume_provider'
        object.openstack_volume_provider = OpenStackVolumeProviderReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OperatingSystemReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OperatingSystem.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('cmdline')
    object.cmdline = value if not value.nil?
    value = reader.get_attribute('custom_kernel_cmdline')
    object.custom_kernel_cmdline = value if not value.nil?
    value = reader.get_attribute('initrd')
    object.initrd = value if not value.nil?
    value = reader.get_attribute('kernel')
    object.kernel = value if not value.nil?
    value = reader.get_attribute('reported_kernel_cmdline')
    object.reported_kernel_cmdline = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'boot'
        object.boot = BootReader.read_one(reader)
      when 'cmdline'
        object.cmdline = Reader.read_string(reader)
      when 'custom_kernel_cmdline'
        object.custom_kernel_cmdline = Reader.read_string(reader)
      when 'initrd'
        object.initrd = Reader.read_string(reader)
      when 'kernel'
        object.kernel = Reader.read_string(reader)
      when 'reported_kernel_cmdline'
        object.reported_kernel_cmdline = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OperatingSystemInfoReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = OperatingSystemInfo.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'large_icon'
        object.large_icon = IconReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'small_icon'
        object.small_icon = IconReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class OptionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Option.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PackageReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Package.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PayloadReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Payload.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('volume_id')
    object.volume_id = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'files'
        object.files = FileReader.read_many(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'volume_id'
        object.volume_id = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PermissionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Permission.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'disk'
        object.disk = DiskReader.read_one(reader)
      when 'group'
        object.group = GroupReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'role'
        object.role = RoleReader.read_one(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vm_pool'
        object.vm_pool = VmPoolReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PermitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Permit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('administrative')
    object.administrative = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'administrative'
        object.administrative = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'role'
        object.role = RoleReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PmProxyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = PmProxy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PortMirroringReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = PortMirroring.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    reader.next_element
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PowerManagementReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = PowerManagement.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('automatic_pm_enabled')
    object.automatic_pm_enabled = value if not value.nil?
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    value = reader.get_attribute('kdump_detection')
    object.kdump_detection = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'agents'
        object.agents = AgentReader.read_many(reader)
      when 'automatic_pm_enabled'
        object.automatic_pm_enabled = Reader.read_boolean(reader)
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'kdump_detection'
        object.kdump_detection = Reader.read_boolean(reader)
      when 'options'
        object.options = OptionReader.read_many(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'pm_proxies'
        object.pm_proxies = PmProxyReader.read_many(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ProductReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Product.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ProductInfoReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ProductInfo.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('vendor')
    object.vendor = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      when 'vendor'
        object.vendor = Reader.read_string(reader)
      when 'version'
        object.version = VersionReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ProfileDetailReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ProfileDetail.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('duration')
    object.duration = value if not value.nil?
    value = reader.get_attribute('profile_type')
    object.profile_type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'block_statistics'
        object.block_statistics = BlockStatisticReader.read_many(reader)
      when 'duration'
        object.duration = Reader.read_integer(reader)
      when 'fop_statistics'
        object.fop_statistics = FopStatisticReader.read_many(reader)
      when 'profile_type'
        object.profile_type = Reader.read_string(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class PropertyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Property.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ProxyTicketReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ProxyTicket.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class QosReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Qos.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('cpu_limit')
    object.cpu_limit = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('inbound_average')
    object.inbound_average = value if not value.nil?
    value = reader.get_attribute('inbound_burst')
    object.inbound_burst = value if not value.nil?
    value = reader.get_attribute('inbound_peak')
    object.inbound_peak = value if not value.nil?
    value = reader.get_attribute('max_iops')
    object.max_iops = value if not value.nil?
    value = reader.get_attribute('max_read_iops')
    object.max_read_iops = value if not value.nil?
    value = reader.get_attribute('max_read_throughput')
    object.max_read_throughput = value if not value.nil?
    value = reader.get_attribute('max_throughput')
    object.max_throughput = value if not value.nil?
    value = reader.get_attribute('max_write_iops')
    object.max_write_iops = value if not value.nil?
    value = reader.get_attribute('max_write_throughput')
    object.max_write_throughput = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('outbound_average')
    object.outbound_average = value if not value.nil?
    value = reader.get_attribute('outbound_average_linkshare')
    object.outbound_average_linkshare = value if not value.nil?
    value = reader.get_attribute('outbound_average_realtime')
    object.outbound_average_realtime = value if not value.nil?
    value = reader.get_attribute('outbound_average_upperlimit')
    object.outbound_average_upperlimit = value if not value.nil?
    value = reader.get_attribute('outbound_burst')
    object.outbound_burst = value if not value.nil?
    value = reader.get_attribute('outbound_peak')
    object.outbound_peak = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'cpu_limit'
        object.cpu_limit = Reader.read_integer(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'inbound_average'
        object.inbound_average = Reader.read_integer(reader)
      when 'inbound_burst'
        object.inbound_burst = Reader.read_integer(reader)
      when 'inbound_peak'
        object.inbound_peak = Reader.read_integer(reader)
      when 'max_iops'
        object.max_iops = Reader.read_integer(reader)
      when 'max_read_iops'
        object.max_read_iops = Reader.read_integer(reader)
      when 'max_read_throughput'
        object.max_read_throughput = Reader.read_integer(reader)
      when 'max_throughput'
        object.max_throughput = Reader.read_integer(reader)
      when 'max_write_iops'
        object.max_write_iops = Reader.read_integer(reader)
      when 'max_write_throughput'
        object.max_write_throughput = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'outbound_average'
        object.outbound_average = Reader.read_integer(reader)
      when 'outbound_average_linkshare'
        object.outbound_average_linkshare = Reader.read_integer(reader)
      when 'outbound_average_realtime'
        object.outbound_average_realtime = Reader.read_integer(reader)
      when 'outbound_average_upperlimit'
        object.outbound_average_upperlimit = Reader.read_integer(reader)
      when 'outbound_burst'
        object.outbound_burst = Reader.read_integer(reader)
      when 'outbound_peak'
        object.outbound_peak = Reader.read_integer(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class QuotaReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Quota.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('cluster_hard_limit_pct')
    object.cluster_hard_limit_pct = value if not value.nil?
    value = reader.get_attribute('cluster_soft_limit_pct')
    object.cluster_soft_limit_pct = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('storage_hard_limit_pct')
    object.storage_hard_limit_pct = value if not value.nil?
    value = reader.get_attribute('storage_soft_limit_pct')
    object.storage_soft_limit_pct = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'cluster_hard_limit_pct'
        object.cluster_hard_limit_pct = Reader.read_integer(reader)
      when 'cluster_soft_limit_pct'
        object.cluster_soft_limit_pct = Reader.read_integer(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'disks'
        object.disks = DiskReader.read_many(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'storage_hard_limit_pct'
        object.storage_hard_limit_pct = Reader.read_integer(reader)
      when 'storage_soft_limit_pct'
        object.storage_soft_limit_pct = Reader.read_integer(reader)
      when 'users'
        object.users = UserReader.read_many(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'quota_cluster_limits'
        object.quota_cluster_limits = QuotaClusterLimitReader.read_many(reader)
      when 'quota_storage_limits'
        object.quota_storage_limits = QuotaStorageLimitReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      when 'quotaclusterlimits'
        object.quota_cluster_limits = list
      when 'quotastoragelimits'
        object.quota_storage_limits = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class QuotaClusterLimitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = QuotaClusterLimit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('memory_limit')
    object.memory_limit = value if not value.nil?
    value = reader.get_attribute('memory_usage')
    object.memory_usage = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('vcpu_limit')
    object.vcpu_limit = value if not value.nil?
    value = reader.get_attribute('vcpu_usage')
    object.vcpu_usage = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'memory_limit'
        object.memory_limit = Reader.read_decimal(reader)
      when 'memory_usage'
        object.memory_usage = Reader.read_decimal(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'vcpu_limit'
        object.vcpu_limit = Reader.read_integer(reader)
      when 'vcpu_usage'
        object.vcpu_usage = Reader.read_integer(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'quota'
        object.quota = QuotaReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class QuotaStorageLimitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = QuotaStorageLimit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('limit')
    object.limit = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('usage')
    object.usage = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'limit'
        object.limit = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'usage'
        object.usage = Reader.read_decimal(reader)
      when 'quota'
        object.quota = QuotaReader.read_one(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class RangeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Range.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('from')
    object.from = value if not value.nil?
    value = reader.get_attribute('to')
    object.to = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'from'
        object.from = Reader.read_string(reader)
      when 'to'
        object.to = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class RateReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Rate.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('bytes')
    object.bytes = value if not value.nil?
    value = reader.get_attribute('period')
    object.period = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bytes'
        object.bytes = Reader.read_integer(reader)
      when 'period'
        object.period = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ReportedConfigurationReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ReportedConfiguration.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('actual_value')
    object.actual_value = value if not value.nil?
    value = reader.get_attribute('expected_value')
    object.expected_value = value if not value.nil?
    value = reader.get_attribute('in_sync')
    object.in_sync = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'actual_value'
        object.actual_value = Reader.read_string(reader)
      when 'expected_value'
        object.expected_value = Reader.read_string(reader)
      when 'in_sync'
        object.in_sync = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class ReportedDeviceReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = ReportedDevice.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ips'
        object.ips = IpReader.read_many(reader)
      when 'mac'
        object.mac = MacReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class RngDeviceReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = RngDevice.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('source')
    object.source = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'rate'
        object.rate = RateReader.read_one(reader)
      when 'source'
        object.source = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class RoleReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Role.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('administrative')
    object.administrative = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('mutable')
    object.mutable = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'administrative'
        object.administrative = Reader.read_boolean(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'mutable'
        object.mutable = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'permits'
        object.permits = PermitReader.read_many(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permits'
        object.permits = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class SchedulingPolicyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SchedulingPolicy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('default_policy')
    object.default_policy = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('locked')
    object.locked = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'default_policy'
        object.default_policy = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'locked'
        object.locked = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'balances'
        object.balances = BalanceReader.read_many(reader)
      when 'filters'
        object.filters = FilterReader.read_many(reader)
      when 'weight'
        object.weight = WeightReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'balances'
        object.balances = list
      when 'filters'
        object.filters = list
      when 'weight'
        object.weight = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class SchedulingPolicyUnitReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SchedulingPolicyUnit.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('internal')
    object.internal = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'internal'
        object.internal = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'properties'
        object.properties = PropertyReader.read_many(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SeLinuxReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SeLinux.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('mode')
    object.mode = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'mode'
        object.mode = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SerialNumberReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SerialNumber.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('policy')
    object.policy = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'policy'
        object.policy = Reader.read_string(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SessionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Session.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('console_user')
    object.console_user = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('protocol')
    object.protocol = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'console_user'
        object.console_user = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'ip'
        object.ip = IpReader.read_one(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'protocol'
        object.protocol = Reader.read_string(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SkipIfConnectivityBrokenReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SkipIfConnectivityBroken.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    value = reader.get_attribute('threshold')
    object.threshold = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'threshold'
        object.threshold = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SkipIfSdActiveReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SkipIfSdActive.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SnapshotReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Snapshot.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('cpu_shares')
    object.cpu_shares = value if not value.nil?
    value = reader.get_attribute('creation_time')
    object.creation_time = value if not value.nil?
    value = reader.get_attribute('custom_cpu_model')
    object.custom_cpu_model = value if not value.nil?
    value = reader.get_attribute('custom_emulated_machine')
    object.custom_emulated_machine = value if not value.nil?
    value = reader.get_attribute('date')
    object.date = value if not value.nil?
    value = reader.get_attribute('delete_protected')
    object.delete_protected = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('fqdn')
    object.fqdn = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('migration_downtime')
    object.migration_downtime = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('next_run_configuration_exists')
    object.next_run_configuration_exists = value if not value.nil?
    value = reader.get_attribute('numa_tune_mode')
    object.numa_tune_mode = value if not value.nil?
    value = reader.get_attribute('origin')
    object.origin = value if not value.nil?
    value = reader.get_attribute('persist_memorystate')
    object.persist_memorystate = value if not value.nil?
    value = reader.get_attribute('run_once')
    object.run_once = value if not value.nil?
    value = reader.get_attribute('snapshot_status')
    object.snapshot_status = value if not value.nil?
    value = reader.get_attribute('snapshot_type')
    object.snapshot_type = value if not value.nil?
    value = reader.get_attribute('soundcard_enabled')
    object.soundcard_enabled = value if not value.nil?
    value = reader.get_attribute('start_paused')
    object.start_paused = value if not value.nil?
    value = reader.get_attribute('start_time')
    object.start_time = value if not value.nil?
    value = reader.get_attribute('stateless')
    object.stateless = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('status_detail')
    object.status_detail = value if not value.nil?
    value = reader.get_attribute('stop_reason')
    object.stop_reason = value if not value.nil?
    value = reader.get_attribute('stop_time')
    object.stop_time = value if not value.nil?
    value = reader.get_attribute('tunnel_migration')
    object.tunnel_migration = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('use_latest_template_version')
    object.use_latest_template_version = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bios'
        object.bios = BiosReader.read_one(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'console'
        object.console = ConsoleReader.read_one(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'cpu_shares'
        object.cpu_shares = Reader.read_integer(reader)
      when 'creation_time'
        object.creation_time = Reader.read_date(reader)
      when 'custom_compatibility_version'
        object.custom_compatibility_version = VersionReader.read_one(reader)
      when 'custom_cpu_model'
        object.custom_cpu_model = Reader.read_string(reader)
      when 'custom_emulated_machine'
        object.custom_emulated_machine = Reader.read_string(reader)
      when 'custom_properties'
        object.custom_properties = CustomPropertyReader.read_many(reader)
      when 'date'
        object.date = Reader.read_date(reader)
      when 'delete_protected'
        object.delete_protected = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'fqdn'
        object.fqdn = Reader.read_string(reader)
      when 'guest_operating_system'
        object.guest_operating_system = GuestOperatingSystemReader.read_one(reader)
      when 'guest_time_zone'
        object.guest_time_zone = TimeZoneReader.read_one(reader)
      when 'high_availability'
        object.high_availability = HighAvailabilityReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'initialization'
        object.initialization = InitializationReader.read_one(reader)
      when 'io'
        object.io = IoReader.read_one(reader)
      when 'large_icon'
        object.large_icon = IconReader.read_one(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'memory_policy'
        object.memory_policy = MemoryPolicyReader.read_one(reader)
      when 'migration'
        object.migration = MigrationOptionsReader.read_one(reader)
      when 'migration_downtime'
        object.migration_downtime = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'next_run_configuration_exists'
        object.next_run_configuration_exists = Reader.read_boolean(reader)
      when 'numa_tune_mode'
        object.numa_tune_mode = Reader.read_string(reader)
      when 'origin'
        object.origin = Reader.read_string(reader)
      when 'os'
        object.os = OperatingSystemReader.read_one(reader)
      when 'payloads'
        object.payloads = PayloadReader.read_many(reader)
      when 'persist_memorystate'
        object.persist_memorystate = Reader.read_boolean(reader)
      when 'placement_policy'
        object.placement_policy = VmPlacementPolicyReader.read_one(reader)
      when 'rng_device'
        object.rng_device = RngDeviceReader.read_one(reader)
      when 'run_once'
        object.run_once = Reader.read_boolean(reader)
      when 'serial_number'
        object.serial_number = SerialNumberReader.read_one(reader)
      when 'small_icon'
        object.small_icon = IconReader.read_one(reader)
      when 'snapshot_status'
        object.snapshot_status = Reader.read_string(reader)
      when 'snapshot_type'
        object.snapshot_type = Reader.read_string(reader)
      when 'soundcard_enabled'
        object.soundcard_enabled = Reader.read_boolean(reader)
      when 'sso'
        object.sso = SsoReader.read_one(reader)
      when 'start_paused'
        object.start_paused = Reader.read_boolean(reader)
      when 'start_time'
        object.start_time = Reader.read_date(reader)
      when 'stateless'
        object.stateless = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'status_detail'
        object.status_detail = Reader.read_string(reader)
      when 'stop_reason'
        object.stop_reason = Reader.read_string(reader)
      when 'stop_time'
        object.stop_time = Reader.read_date(reader)
      when 'time_zone'
        object.time_zone = TimeZoneReader.read_one(reader)
      when 'tunnel_migration'
        object.tunnel_migration = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'usb'
        object.usb = UsbReader.read_one(reader)
      when 'use_latest_template_version'
        object.use_latest_template_version = Reader.read_boolean(reader)
      when 'virtio_scsi'
        object.virtio_scsi = VirtioScsiReader.read_one(reader)
      when 'affinity_labels'
        object.affinity_labels = AffinityLabelReader.read_many(reader)
      when 'applications'
        object.applications = ApplicationReader.read_many(reader)
      when 'cdroms'
        object.cdroms = CdromReader.read_many(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'cpu_profile'
        object.cpu_profile = CpuProfileReader.read_one(reader)
      when 'disk_attachments'
        object.disk_attachments = DiskAttachmentReader.read_many(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      when 'floppies'
        object.floppies = FloppyReader.read_many(reader)
      when 'graphics_consoles'
        object.graphics_consoles = GraphicsConsoleReader.read_many(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'host_devices'
        object.host_devices = HostDeviceReader.read_many(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'katello_errata'
        object.katello_errata = KatelloErratumReader.read_many(reader)
      when 'nics'
        object.nics = NicReader.read_many(reader)
      when 'host_numa_nodes'
        object.numa_nodes = NumaNodeReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'quota'
        object.quota = QuotaReader.read_one(reader)
      when 'reported_devices'
        object.reported_devices = ReportedDeviceReader.read_many(reader)
      when 'sessions'
        object.sessions = SessionReader.read_many(reader)
      when 'snapshots'
        object.snapshots = SnapshotReader.read_many(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vm_pool'
        object.vm_pool = VmPoolReader.read_one(reader)
      when 'watchdogs'
        object.watchdogs = WatchdogReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'affinitylabels'
        object.affinity_labels = list
      when 'applications'
        object.applications = list
      when 'cdroms'
        object.cdroms = list
      when 'diskattachments'
        object.disk_attachments = list
      when 'floppies'
        object.floppies = list
      when 'graphicsconsoles'
        object.graphics_consoles = list
      when 'hostdevices'
        object.host_devices = list
      when 'katelloerrata'
        object.katello_errata = list
      when 'nics'
        object.nics = list
      when 'numanodes'
        object.numa_nodes = list
      when 'permissions'
        object.permissions = list
      when 'reporteddevices'
        object.reported_devices = list
      when 'sessions'
        object.sessions = list
      when 'snapshots'
        object.snapshots = list
      when 'statistics'
        object.statistics = list
      when 'tags'
        object.tags = list
      when 'watchdogs'
        object.watchdogs = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class SpecialObjectsReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SpecialObjects.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'blank_template'
        object.blank_template = TemplateReader.read_one(reader)
      when 'root_tag'
        object.root_tag = TagReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SpmReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Spm.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('priority')
    object.priority = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'priority'
        object.priority = Reader.read_integer(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SshReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Ssh.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('authentication_method')
    object.authentication_method = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('fingerprint')
    object.fingerprint = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'authentication_method'
        object.authentication_method = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'fingerprint'
        object.fingerprint = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SshPublicKeyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = SshPublicKey.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('content')
    object.content = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'content'
        object.content = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class SsoReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Sso.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'methods'
        object.methods = MethodReader.read_many(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class StatisticReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Statistic.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('kind')
    object.kind = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('unit')
    object.unit = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'kind'
        object.kind = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'unit'
        object.unit = Reader.read_string(reader)
      when 'values'
        object.values = ValueReader.read_many(reader)
      when 'brick'
        object.brick = GlusterBrickReader.read_one(reader)
      when 'disk'
        object.disk = DiskReader.read_one(reader)
      when 'gluster_volume'
        object.gluster_volume = GlusterVolumeReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'host_nic'
        object.host_nic = HostNicReader.read_one(reader)
      when 'host_numa_node'
        object.host_numa_node = NumaNodeReader.read_one(reader)
      when 'nic'
        object.nic = NicReader.read_one(reader)
      when 'step'
        object.step = StepReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class StepReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Step.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('end_time')
    object.end_time = value if not value.nil?
    value = reader.get_attribute('external')
    object.external = value if not value.nil?
    value = reader.get_attribute('external_type')
    object.external_type = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('number')
    object.number = value if not value.nil?
    value = reader.get_attribute('start_time')
    object.start_time = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'end_time'
        object.end_time = Reader.read_date(reader)
      when 'external'
        object.external = Reader.read_boolean(reader)
      when 'external_type'
        object.external_type = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'number'
        object.number = Reader.read_integer(reader)
      when 'start_time'
        object.start_time = Reader.read_date(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'job'
        object.job = JobReader.read_one(reader)
      when 'parent_step'
        object.parent_step = StepReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'statistics'
        object.statistics = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class StorageConnectionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = StorageConnection.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('address')
    object.address = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('mount_options')
    object.mount_options = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('nfs_retrans')
    object.nfs_retrans = value if not value.nil?
    value = reader.get_attribute('nfs_timeo')
    object.nfs_timeo = value if not value.nil?
    value = reader.get_attribute('nfs_version')
    object.nfs_version = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('path')
    object.path = value if not value.nil?
    value = reader.get_attribute('port')
    object.port = value if not value.nil?
    value = reader.get_attribute('portal')
    object.portal = value if not value.nil?
    value = reader.get_attribute('target')
    object.target = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    value = reader.get_attribute('vfs_type')
    object.vfs_type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'address'
        object.address = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'mount_options'
        object.mount_options = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'nfs_retrans'
        object.nfs_retrans = Reader.read_integer(reader)
      when 'nfs_timeo'
        object.nfs_timeo = Reader.read_integer(reader)
      when 'nfs_version'
        object.nfs_version = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'path'
        object.path = Reader.read_string(reader)
      when 'port'
        object.port = Reader.read_integer(reader)
      when 'portal'
        object.portal = Reader.read_string(reader)
      when 'target'
        object.target = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'vfs_type'
        object.vfs_type = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class StorageConnectionExtensionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = StorageConnectionExtension.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('target')
    object.target = value if not value.nil?
    value = reader.get_attribute('username')
    object.username = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'target'
        object.target = Reader.read_string(reader)
      when 'username'
        object.username = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class StorageDomainReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = StorageDomain.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('available')
    object.available = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('committed')
    object.committed = value if not value.nil?
    value = reader.get_attribute('critical_space_action_blocker')
    object.critical_space_action_blocker = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('external_status')
    object.external_status = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('import')
    object.import = value if not value.nil?
    value = reader.get_attribute('master')
    object.master = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('storage_format')
    object.storage_format = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('used')
    object.used = value if not value.nil?
    value = reader.get_attribute('warning_low_space_indicator')
    object.warning_low_space_indicator = value if not value.nil?
    value = reader.get_attribute('wipe_after_delete')
    object.wipe_after_delete = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'available'
        object.available = Reader.read_integer(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'committed'
        object.committed = Reader.read_integer(reader)
      when 'critical_space_action_blocker'
        object.critical_space_action_blocker = Reader.read_integer(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'external_status'
        object.external_status = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'import'
        object.import = Reader.read_boolean(reader)
      when 'master'
        object.master = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'storage'
        object.storage = HostStorageReader.read_one(reader)
      when 'storage_format'
        object.storage_format = Reader.read_string(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'used'
        object.used = Reader.read_integer(reader)
      when 'warning_low_space_indicator'
        object.warning_low_space_indicator = Reader.read_integer(reader)
      when 'wipe_after_delete'
        object.wipe_after_delete = Reader.read_boolean(reader)
      when 'data_center'
        object.data_center = DataCenterReader.read_one(reader)
      when 'data_centers'
        object.data_centers = DataCenterReader.read_many(reader)
      when 'disk_profiles'
        object.disk_profiles = DiskProfileReader.read_many(reader)
      when 'disk_snapshots'
        object.disk_snapshots = DiskSnapshotReader.read_many(reader)
      when 'disks'
        object.disks = DiskReader.read_many(reader)
      when 'files'
        object.files = FileReader.read_many(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'images'
        object.images = ImageReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'storage_connections'
        object.storage_connections = StorageConnectionReader.read_many(reader)
      when 'templates'
        object.templates = TemplateReader.read_many(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'datacenters'
        object.data_centers = list
      when 'diskprofiles'
        object.disk_profiles = list
      when 'disksnapshots'
        object.disk_snapshots = list
      when 'disks'
        object.disks = list
      when 'files'
        object.files = list
      when 'images'
        object.images = list
      when 'permissions'
        object.permissions = list
      when 'storageconnections'
        object.storage_connections = list
      when 'templates'
        object.templates = list
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class TagReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Tag.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'group'
        object.group = GroupReader.read_one(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'parent'
        object.parent = TagReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'user'
        object.user = UserReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class TemplateReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Template.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('cpu_shares')
    object.cpu_shares = value if not value.nil?
    value = reader.get_attribute('creation_time')
    object.creation_time = value if not value.nil?
    value = reader.get_attribute('custom_cpu_model')
    object.custom_cpu_model = value if not value.nil?
    value = reader.get_attribute('custom_emulated_machine')
    object.custom_emulated_machine = value if not value.nil?
    value = reader.get_attribute('delete_protected')
    object.delete_protected = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('migration_downtime')
    object.migration_downtime = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('origin')
    object.origin = value if not value.nil?
    value = reader.get_attribute('soundcard_enabled')
    object.soundcard_enabled = value if not value.nil?
    value = reader.get_attribute('start_paused')
    object.start_paused = value if not value.nil?
    value = reader.get_attribute('stateless')
    object.stateless = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('tunnel_migration')
    object.tunnel_migration = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bios'
        object.bios = BiosReader.read_one(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'console'
        object.console = ConsoleReader.read_one(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'cpu_shares'
        object.cpu_shares = Reader.read_integer(reader)
      when 'creation_time'
        object.creation_time = Reader.read_date(reader)
      when 'custom_compatibility_version'
        object.custom_compatibility_version = VersionReader.read_one(reader)
      when 'custom_cpu_model'
        object.custom_cpu_model = Reader.read_string(reader)
      when 'custom_emulated_machine'
        object.custom_emulated_machine = Reader.read_string(reader)
      when 'custom_properties'
        object.custom_properties = CustomPropertyReader.read_many(reader)
      when 'delete_protected'
        object.delete_protected = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'high_availability'
        object.high_availability = HighAvailabilityReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'initialization'
        object.initialization = InitializationReader.read_one(reader)
      when 'io'
        object.io = IoReader.read_one(reader)
      when 'large_icon'
        object.large_icon = IconReader.read_one(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'memory_policy'
        object.memory_policy = MemoryPolicyReader.read_one(reader)
      when 'migration'
        object.migration = MigrationOptionsReader.read_one(reader)
      when 'migration_downtime'
        object.migration_downtime = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'origin'
        object.origin = Reader.read_string(reader)
      when 'os'
        object.os = OperatingSystemReader.read_one(reader)
      when 'rng_device'
        object.rng_device = RngDeviceReader.read_one(reader)
      when 'serial_number'
        object.serial_number = SerialNumberReader.read_one(reader)
      when 'small_icon'
        object.small_icon = IconReader.read_one(reader)
      when 'soundcard_enabled'
        object.soundcard_enabled = Reader.read_boolean(reader)
      when 'sso'
        object.sso = SsoReader.read_one(reader)
      when 'start_paused'
        object.start_paused = Reader.read_boolean(reader)
      when 'stateless'
        object.stateless = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'time_zone'
        object.time_zone = TimeZoneReader.read_one(reader)
      when 'tunnel_migration'
        object.tunnel_migration = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'usb'
        object.usb = UsbReader.read_one(reader)
      when 'version'
        object.version = TemplateVersionReader.read_one(reader)
      when 'virtio_scsi'
        object.virtio_scsi = VirtioScsiReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'cdroms'
        object.cdroms = CdromReader.read_many(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'cpu_profile'
        object.cpu_profile = CpuProfileReader.read_one(reader)
      when 'disk_attachments'
        object.disk_attachments = DiskAttachmentReader.read_many(reader)
      when 'graphics_consoles'
        object.graphics_consoles = GraphicsConsoleReader.read_many(reader)
      when 'nics'
        object.nics = NicReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'watchdogs'
        object.watchdogs = WatchdogReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'cdroms'
        object.cdroms = list
      when 'diskattachments'
        object.disk_attachments = list
      when 'graphicsconsoles'
        object.graphics_consoles = list
      when 'nics'
        object.nics = list
      when 'permissions'
        object.permissions = list
      when 'tags'
        object.tags = list
      when 'watchdogs'
        object.watchdogs = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class TemplateVersionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = TemplateVersion.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('version_name')
    object.version_name = value if not value.nil?
    value = reader.get_attribute('version_number')
    object.version_number = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'version_name'
        object.version_name = Reader.read_string(reader)
      when 'version_number'
        object.version_number = Reader.read_integer(reader)
      when 'base_template'
        object.base_template = TemplateReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class TicketReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Ticket.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('expiry')
    object.expiry = value if not value.nil?
    value = reader.get_attribute('value')
    object.value = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'expiry'
        object.expiry = Reader.read_integer(reader)
      when 'value'
        object.value = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class TimeZoneReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = TimeZone.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('utc_offset')
    object.utc_offset = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'name'
        object.name = Reader.read_string(reader)
      when 'utc_offset'
        object.utc_offset = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class TransparentHugePagesReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = TransparentHugePages.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class UnmanagedNetworkReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = UnmanagedNetwork.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'host_nic'
        object.host_nic = HostNicReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class UsbReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Usb.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class UserReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = User.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('department')
    object.department = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('domain_entry_id')
    object.domain_entry_id = value if not value.nil?
    value = reader.get_attribute('email')
    object.email = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('last_name')
    object.last_name = value if not value.nil?
    value = reader.get_attribute('logged_in')
    object.logged_in = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('namespace')
    object.namespace = value if not value.nil?
    value = reader.get_attribute('password')
    object.password = value if not value.nil?
    value = reader.get_attribute('principal')
    object.principal = value if not value.nil?
    value = reader.get_attribute('user_name')
    object.user_name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'department'
        object.department = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'domain_entry_id'
        object.domain_entry_id = Reader.read_string(reader)
      when 'email'
        object.email = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'last_name'
        object.last_name = Reader.read_string(reader)
      when 'logged_in'
        object.logged_in = Reader.read_boolean(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'namespace'
        object.namespace = Reader.read_string(reader)
      when 'password'
        object.password = Reader.read_string(reader)
      when 'principal'
        object.principal = Reader.read_string(reader)
      when 'user_name'
        object.user_name = Reader.read_string(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'groups'
        object.groups = GroupReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'roles'
        object.roles = RoleReader.read_many(reader)
      when 'ssh_public_keys'
        object.ssh_public_keys = SshPublicKeyReader.read_many(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'groups'
        object.groups = list
      when 'permissions'
        object.permissions = list
      when 'roles'
        object.roles = list
      when 'sshpublickeys'
        object.ssh_public_keys = list
      when 'tags'
        object.tags = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class ValueReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Value.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('datum')
    object.datum = value if not value.nil?
    value = reader.get_attribute('detail')
    object.detail = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'datum'
        object.datum = Reader.read_decimal(reader)
      when 'detail'
        object.detail = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VcpuPinReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VcpuPin.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('cpu_set')
    object.cpu_set = value if not value.nil?
    value = reader.get_attribute('vcpu')
    object.vcpu = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'cpu_set'
        object.cpu_set = Reader.read_string(reader)
      when 'vcpu'
        object.vcpu = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VendorReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Vendor.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VersionReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Version.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('build')
    object.build = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('full_version')
    object.full_version = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('major')
    object.major = value if not value.nil?
    value = reader.get_attribute('minor')
    object.minor = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('revision')
    object.revision = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'build'
        object.build = Reader.read_integer(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'full_version'
        object.full_version = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'major'
        object.major = Reader.read_integer(reader)
      when 'minor'
        object.minor = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'revision'
        object.revision = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VirtioScsiReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VirtioScsi.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('enabled')
    object.enabled = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'enabled'
        object.enabled = Reader.read_boolean(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VirtualNumaNodeReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VirtualNumaNode.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('index')
    object.index = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('node_distance')
    object.node_distance = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'index'
        object.index = Reader.read_integer(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'node_distance'
        object.node_distance = Reader.read_string(reader)
      when 'numa_node_pins'
        object.numa_node_pins = NumaNodePinReader.read_many(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'statistics'
        object.statistics = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class VlanReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Vlan.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'id'
        object.id = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VmReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Vm.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('cpu_shares')
    object.cpu_shares = value if not value.nil?
    value = reader.get_attribute('creation_time')
    object.creation_time = value if not value.nil?
    value = reader.get_attribute('custom_cpu_model')
    object.custom_cpu_model = value if not value.nil?
    value = reader.get_attribute('custom_emulated_machine')
    object.custom_emulated_machine = value if not value.nil?
    value = reader.get_attribute('delete_protected')
    object.delete_protected = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('fqdn')
    object.fqdn = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('migration_downtime')
    object.migration_downtime = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('next_run_configuration_exists')
    object.next_run_configuration_exists = value if not value.nil?
    value = reader.get_attribute('numa_tune_mode')
    object.numa_tune_mode = value if not value.nil?
    value = reader.get_attribute('origin')
    object.origin = value if not value.nil?
    value = reader.get_attribute('run_once')
    object.run_once = value if not value.nil?
    value = reader.get_attribute('soundcard_enabled')
    object.soundcard_enabled = value if not value.nil?
    value = reader.get_attribute('start_paused')
    object.start_paused = value if not value.nil?
    value = reader.get_attribute('start_time')
    object.start_time = value if not value.nil?
    value = reader.get_attribute('stateless')
    object.stateless = value if not value.nil?
    value = reader.get_attribute('status')
    object.status = value if not value.nil?
    value = reader.get_attribute('status_detail')
    object.status_detail = value if not value.nil?
    value = reader.get_attribute('stop_reason')
    object.stop_reason = value if not value.nil?
    value = reader.get_attribute('stop_time')
    object.stop_time = value if not value.nil?
    value = reader.get_attribute('tunnel_migration')
    object.tunnel_migration = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('use_latest_template_version')
    object.use_latest_template_version = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bios'
        object.bios = BiosReader.read_one(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'console'
        object.console = ConsoleReader.read_one(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'cpu_shares'
        object.cpu_shares = Reader.read_integer(reader)
      when 'creation_time'
        object.creation_time = Reader.read_date(reader)
      when 'custom_compatibility_version'
        object.custom_compatibility_version = VersionReader.read_one(reader)
      when 'custom_cpu_model'
        object.custom_cpu_model = Reader.read_string(reader)
      when 'custom_emulated_machine'
        object.custom_emulated_machine = Reader.read_string(reader)
      when 'custom_properties'
        object.custom_properties = CustomPropertyReader.read_many(reader)
      when 'delete_protected'
        object.delete_protected = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'fqdn'
        object.fqdn = Reader.read_string(reader)
      when 'guest_operating_system'
        object.guest_operating_system = GuestOperatingSystemReader.read_one(reader)
      when 'guest_time_zone'
        object.guest_time_zone = TimeZoneReader.read_one(reader)
      when 'high_availability'
        object.high_availability = HighAvailabilityReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'initialization'
        object.initialization = InitializationReader.read_one(reader)
      when 'io'
        object.io = IoReader.read_one(reader)
      when 'large_icon'
        object.large_icon = IconReader.read_one(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'memory_policy'
        object.memory_policy = MemoryPolicyReader.read_one(reader)
      when 'migration'
        object.migration = MigrationOptionsReader.read_one(reader)
      when 'migration_downtime'
        object.migration_downtime = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'next_run_configuration_exists'
        object.next_run_configuration_exists = Reader.read_boolean(reader)
      when 'numa_tune_mode'
        object.numa_tune_mode = Reader.read_string(reader)
      when 'origin'
        object.origin = Reader.read_string(reader)
      when 'os'
        object.os = OperatingSystemReader.read_one(reader)
      when 'payloads'
        object.payloads = PayloadReader.read_many(reader)
      when 'placement_policy'
        object.placement_policy = VmPlacementPolicyReader.read_one(reader)
      when 'rng_device'
        object.rng_device = RngDeviceReader.read_one(reader)
      when 'run_once'
        object.run_once = Reader.read_boolean(reader)
      when 'serial_number'
        object.serial_number = SerialNumberReader.read_one(reader)
      when 'small_icon'
        object.small_icon = IconReader.read_one(reader)
      when 'soundcard_enabled'
        object.soundcard_enabled = Reader.read_boolean(reader)
      when 'sso'
        object.sso = SsoReader.read_one(reader)
      when 'start_paused'
        object.start_paused = Reader.read_boolean(reader)
      when 'start_time'
        object.start_time = Reader.read_date(reader)
      when 'stateless'
        object.stateless = Reader.read_boolean(reader)
      when 'status'
        object.status = Reader.read_string(reader)
      when 'status_detail'
        object.status_detail = Reader.read_string(reader)
      when 'stop_reason'
        object.stop_reason = Reader.read_string(reader)
      when 'stop_time'
        object.stop_time = Reader.read_date(reader)
      when 'time_zone'
        object.time_zone = TimeZoneReader.read_one(reader)
      when 'tunnel_migration'
        object.tunnel_migration = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'usb'
        object.usb = UsbReader.read_one(reader)
      when 'use_latest_template_version'
        object.use_latest_template_version = Reader.read_boolean(reader)
      when 'virtio_scsi'
        object.virtio_scsi = VirtioScsiReader.read_one(reader)
      when 'affinity_labels'
        object.affinity_labels = AffinityLabelReader.read_many(reader)
      when 'applications'
        object.applications = ApplicationReader.read_many(reader)
      when 'cdroms'
        object.cdroms = CdromReader.read_many(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'cpu_profile'
        object.cpu_profile = CpuProfileReader.read_one(reader)
      when 'disk_attachments'
        object.disk_attachments = DiskAttachmentReader.read_many(reader)
      when 'external_host_provider'
        object.external_host_provider = ExternalHostProviderReader.read_one(reader)
      when 'floppies'
        object.floppies = FloppyReader.read_many(reader)
      when 'graphics_consoles'
        object.graphics_consoles = GraphicsConsoleReader.read_many(reader)
      when 'host'
        object.host = HostReader.read_one(reader)
      when 'host_devices'
        object.host_devices = HostDeviceReader.read_many(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'katello_errata'
        object.katello_errata = KatelloErratumReader.read_many(reader)
      when 'nics'
        object.nics = NicReader.read_many(reader)
      when 'host_numa_nodes'
        object.numa_nodes = NumaNodeReader.read_many(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'quota'
        object.quota = QuotaReader.read_one(reader)
      when 'reported_devices'
        object.reported_devices = ReportedDeviceReader.read_many(reader)
      when 'sessions'
        object.sessions = SessionReader.read_many(reader)
      when 'snapshots'
        object.snapshots = SnapshotReader.read_many(reader)
      when 'statistics'
        object.statistics = StatisticReader.read_many(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      when 'tags'
        object.tags = TagReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm_pool'
        object.vm_pool = VmPoolReader.read_one(reader)
      when 'watchdogs'
        object.watchdogs = WatchdogReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'affinitylabels'
        object.affinity_labels = list
      when 'applications'
        object.applications = list
      when 'cdroms'
        object.cdroms = list
      when 'diskattachments'
        object.disk_attachments = list
      when 'floppies'
        object.floppies = list
      when 'graphicsconsoles'
        object.graphics_consoles = list
      when 'hostdevices'
        object.host_devices = list
      when 'katelloerrata'
        object.katello_errata = list
      when 'nics'
        object.nics = list
      when 'numanodes'
        object.numa_nodes = list
      when 'permissions'
        object.permissions = list
      when 'reporteddevices'
        object.reported_devices = list
      when 'sessions'
        object.sessions = list
      when 'snapshots'
        object.snapshots = list
      when 'statistics'
        object.statistics = list
      when 'tags'
        object.tags = list
      when 'watchdogs'
        object.watchdogs = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class VmBaseReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VmBase.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('cpu_shares')
    object.cpu_shares = value if not value.nil?
    value = reader.get_attribute('creation_time')
    object.creation_time = value if not value.nil?
    value = reader.get_attribute('custom_cpu_model')
    object.custom_cpu_model = value if not value.nil?
    value = reader.get_attribute('custom_emulated_machine')
    object.custom_emulated_machine = value if not value.nil?
    value = reader.get_attribute('delete_protected')
    object.delete_protected = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('memory')
    object.memory = value if not value.nil?
    value = reader.get_attribute('migration_downtime')
    object.migration_downtime = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('origin')
    object.origin = value if not value.nil?
    value = reader.get_attribute('soundcard_enabled')
    object.soundcard_enabled = value if not value.nil?
    value = reader.get_attribute('start_paused')
    object.start_paused = value if not value.nil?
    value = reader.get_attribute('stateless')
    object.stateless = value if not value.nil?
    value = reader.get_attribute('tunnel_migration')
    object.tunnel_migration = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'bios'
        object.bios = BiosReader.read_one(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'console'
        object.console = ConsoleReader.read_one(reader)
      when 'cpu'
        object.cpu = CpuReader.read_one(reader)
      when 'cpu_shares'
        object.cpu_shares = Reader.read_integer(reader)
      when 'creation_time'
        object.creation_time = Reader.read_date(reader)
      when 'custom_compatibility_version'
        object.custom_compatibility_version = VersionReader.read_one(reader)
      when 'custom_cpu_model'
        object.custom_cpu_model = Reader.read_string(reader)
      when 'custom_emulated_machine'
        object.custom_emulated_machine = Reader.read_string(reader)
      when 'custom_properties'
        object.custom_properties = CustomPropertyReader.read_many(reader)
      when 'delete_protected'
        object.delete_protected = Reader.read_boolean(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'domain'
        object.domain = DomainReader.read_one(reader)
      when 'high_availability'
        object.high_availability = HighAvailabilityReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'initialization'
        object.initialization = InitializationReader.read_one(reader)
      when 'io'
        object.io = IoReader.read_one(reader)
      when 'large_icon'
        object.large_icon = IconReader.read_one(reader)
      when 'memory'
        object.memory = Reader.read_integer(reader)
      when 'memory_policy'
        object.memory_policy = MemoryPolicyReader.read_one(reader)
      when 'migration'
        object.migration = MigrationOptionsReader.read_one(reader)
      when 'migration_downtime'
        object.migration_downtime = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'origin'
        object.origin = Reader.read_string(reader)
      when 'os'
        object.os = OperatingSystemReader.read_one(reader)
      when 'rng_device'
        object.rng_device = RngDeviceReader.read_one(reader)
      when 'serial_number'
        object.serial_number = SerialNumberReader.read_one(reader)
      when 'small_icon'
        object.small_icon = IconReader.read_one(reader)
      when 'soundcard_enabled'
        object.soundcard_enabled = Reader.read_boolean(reader)
      when 'sso'
        object.sso = SsoReader.read_one(reader)
      when 'start_paused'
        object.start_paused = Reader.read_boolean(reader)
      when 'stateless'
        object.stateless = Reader.read_boolean(reader)
      when 'time_zone'
        object.time_zone = TimeZoneReader.read_one(reader)
      when 'tunnel_migration'
        object.tunnel_migration = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'usb'
        object.usb = UsbReader.read_one(reader)
      when 'virtio_scsi'
        object.virtio_scsi = VirtioScsiReader.read_one(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'cpu_profile'
        object.cpu_profile = CpuProfileReader.read_one(reader)
      when 'storage_domain'
        object.storage_domain = StorageDomainReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VmPlacementPolicyReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VmPlacementPolicy.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('affinity')
    object.affinity = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'affinity'
        object.affinity = Reader.read_string(reader)
      when 'hosts'
        object.hosts = HostReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'hosts'
        object.hosts = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class VmPoolReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VmPool.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('max_user_vms')
    object.max_user_vms = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('prestarted_vms')
    object.prestarted_vms = value if not value.nil?
    value = reader.get_attribute('size')
    object.size = value if not value.nil?
    value = reader.get_attribute('soundcard_enabled')
    object.soundcard_enabled = value if not value.nil?
    value = reader.get_attribute('stateful')
    object.stateful = value if not value.nil?
    value = reader.get_attribute('type')
    object.type = value if not value.nil?
    value = reader.get_attribute('use_latest_template_version')
    object.use_latest_template_version = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'display'
        object.display = DisplayReader.read_one(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'max_user_vms'
        object.max_user_vms = Reader.read_integer(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'prestarted_vms'
        object.prestarted_vms = Reader.read_integer(reader)
      when 'rng_device'
        object.rng_device = RngDeviceReader.read_one(reader)
      when 'size'
        object.size = Reader.read_integer(reader)
      when 'soundcard_enabled'
        object.soundcard_enabled = Reader.read_boolean(reader)
      when 'stateful'
        object.stateful = Reader.read_boolean(reader)
      when 'type'
        object.type = Reader.read_string(reader)
      when 'use_latest_template_version'
        object.use_latest_template_version = Reader.read_boolean(reader)
      when 'cluster'
        object.cluster = ClusterReader.read_one(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class VmSummaryReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VmSummary.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('active')
    object.active = value if not value.nil?
    value = reader.get_attribute('migrating')
    object.migrating = value if not value.nil?
    value = reader.get_attribute('total')
    object.total = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'active'
        object.active = Reader.read_integer(reader)
      when 'migrating'
        object.migrating = Reader.read_integer(reader)
      when 'total'
        object.total = Reader.read_integer(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VnicPassThroughReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VnicPassThrough.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('mode')
    object.mode = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'mode'
        object.mode = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class VnicProfileReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VnicProfile.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    value = reader.get_attribute('port_mirroring')
    object.port_mirroring = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'custom_properties'
        object.custom_properties = CustomPropertyReader.read_many(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'pass_through'
        object.pass_through = VnicPassThroughReader.read_one(reader)
      when 'port_mirroring'
        object.port_mirroring = Reader.read_boolean(reader)
      when 'network'
        object.network = NetworkReader.read_one(reader)
      when 'network_filter'
        object.network_filter = NetworkFilterReader.read_one(reader)
      when 'permissions'
        object.permissions = PermissionReader.read_many(reader)
      when 'qos'
        object.qos = QosReader.read_one(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'permissions'
        object.permissions = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class VolumeGroupReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = VolumeGroup.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'id'
        object.id = Reader.read_string(reader)
      when 'logical_units'
        object.logical_units = LogicalUnitReader.read_many(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end

class WatchdogReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Watchdog.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('action')
    object.action = value if not value.nil?
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('model')
    object.model = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'action'
        object.action = Reader.read_string(reader)
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'model'
        object.model = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'instance_type'
        object.instance_type = InstanceTypeReader.read_one(reader)
      when 'template'
        object.template = TemplateReader.read_one(reader)
      when 'vm'
        object.vm = VmReader.read_one(reader)
      when 'vms'
        object.vms = VmReader.read_many(reader)
      when 'link'
        read_link(reader, object)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
  def self.read_link(reader, object)
    # Process the attributes:
    rel = reader.get_attribute('rel')
    href = reader.get_attribute('href')
    if rel && href
      list = List.new
      list.href = href
      case rel
      when 'vms'
        object.vms = list
      end
    end
    
    # Discard the rest of the element:
    reader.next_element
  end
  
end

class WeightReader < Reader # :nodoc:
  
  def self.read_one(reader)
    # Do nothing if there aren't more tags:
    return nil unless reader.forward
    
    # Create the object:
    object = Weight.new
    
    # Process the attributes:
    object.href = reader.get_attribute('href')
    value = reader.get_attribute('comment')
    object.comment = value if not value.nil?
    value = reader.get_attribute('description')
    object.description = value if not value.nil?
    value = reader.get_attribute('factor')
    object.factor = value if not value.nil?
    value = reader.get_attribute('id')
    object.id = value if not value.nil?
    value = reader.get_attribute('name')
    object.name = value if not value.nil?
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return object if empty
    
    # Process the inner elements:
    while reader.forward do
      case reader.node_name
      when 'comment'
        object.comment = Reader.read_string(reader)
      when 'description'
        object.description = Reader.read_string(reader)
      when 'factor'
        object.factor = Reader.read_integer(reader)
      when 'id'
        object.id = Reader.read_string(reader)
      when 'name'
        object.name = Reader.read_string(reader)
      when 'scheduling_policy'
        object.scheduling_policy = SchedulingPolicyReader.read_one(reader)
      when 'scheduling_policy_unit'
        object.scheduling_policy_unit = SchedulingPolicyUnitReader.read_one(reader)
      else
        reader.next_element
      end
    end
    
    # Discard the end tag:
    reader.read
    
    return object
  end
  
  
  def self.read_many(reader)
    # Do nothing if there aren't more tags:
    list = List.new
    return list unless reader.forward
    
    # Process the attributes:
    list.href = reader.get_attribute('href')
    
    # Discard the start tag:
    empty = reader.empty_element?
    reader.read
    return list if empty
    
    # Process the inner elements:
    while reader.forward do
      list << read_one(reader)
    end
    
    # Discard the end tag:
    reader.read
    
    return list
  end
  
end


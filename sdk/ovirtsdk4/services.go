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


# 
# These forward declarations are required in order to avoid circular dependencies.
# 
package ovirtsdk4

class AffinityGroupService < Service
end

class AffinityGroupVmService < Service
end

class AffinityGroupVmsService < Service
end

class AffinityGroupsService < Service
end

class AffinityLabelService < Service
end

class AffinityLabelHostService < Service
end

class AffinityLabelHostsService < Service
end

class AffinityLabelVmService < Service
end

class AffinityLabelVmsService < Service
end

class AffinityLabelsService < Service
end

class AssignedAffinityLabelService < Service
end

class AssignedAffinityLabelsService < Service
end

class AssignedCpuProfileService < Service
end

class AssignedCpuProfilesService < Service
end

class AssignedDiskProfileService < Service
end

class AssignedDiskProfilesService < Service
end

class AssignedNetworkService < Service
end

class AssignedNetworksService < Service
end

class AssignedPermissionsService < Service
end

class AssignedRolesService < Service
end

class AssignedTagService < Service
end

class AssignedTagsService < Service
end

class AssignedVnicProfileService < Service
end

class AssignedVnicProfilesService < Service
end

class AttachedStorageDomainService < Service
end

class AttachedStorageDomainsService < Service
end

class BalanceService < Service
end

class BalancesService < Service
end

class BookmarkService < Service
end

class BookmarksService < Service
end

class ClusterService < Service
end

class ClusterLevelService < Service
end

class ClusterLevelsService < Service
end

class ClustersService < Service
end

class CopyableService < Service
end

class CpuProfileService < Service
end

class CpuProfilesService < Service
end

class DataCenterService < Service
end

class DataCentersService < Service
end

class DiskAttachmentService < Service
end

class DiskAttachmentsService < Service
end

class DiskProfileService < Service
end

class DiskProfilesService < Service
end

class DiskSnapshotService < Service
end

class DiskSnapshotsService < Service
end

class DisksService < Service
end

class DomainService < Service
end

class DomainGroupService < Service
end

class DomainGroupsService < Service
end

class DomainUserService < Service
end

class DomainUsersService < Service
end

class DomainsService < Service
end

class EventService < Service
end

class EventsService < Service
end

class ExternalComputeResourceService < Service
end

class ExternalComputeResourcesService < Service
end

class ExternalDiscoveredHostService < Service
end

class ExternalDiscoveredHostsService < Service
end

class ExternalHostService < Service
end

class ExternalHostGroupService < Service
end

class ExternalHostGroupsService < Service
end

class ExternalHostProvidersService < Service
end

class ExternalHostsService < Service
end

class ExternalProviderService < Service
end

class ExternalProviderCertificateService < Service
end

class ExternalProviderCertificatesService < Service
end

class FenceAgentService < Service
end

class FenceAgentsService < Service
end

class FileService < Service
end

class FilesService < Service
end

class FilterService < Service
end

class FiltersService < Service
end

class GlusterBricksService < Service
end

class GlusterHookService < Service
end

class GlusterHooksService < Service
end

class GlusterVolumesService < Service
end

class GraphicsConsoleService < Service
end

class GraphicsConsolesService < Service
end

class GroupService < Service
end

class GroupsService < Service
end

class HostDeviceService < Service
end

class HostDevicesService < Service
end

class HostHookService < Service
end

class HostHooksService < Service
end

class HostNicsService < Service
end

class HostNumaNodesService < Service
end

class HostStorageService < Service
end

class HostsService < Service
end

class IconService < Service
end

class IconsService < Service
end

class ImageService < Service
end

class ImagesService < Service
end

class InstanceTypeService < Service
end

class InstanceTypeNicService < Service
end

class InstanceTypeNicsService < Service
end

class InstanceTypeWatchdogService < Service
end

class InstanceTypeWatchdogsService < Service
end

class InstanceTypesService < Service
end

class IscsiBondService < Service
end

class IscsiBondsService < Service
end

class JobService < Service
end

class JobsService < Service
end

class KatelloErrataService < Service
end

class KatelloErratumService < Service
end

class MacPoolService < Service
end

class MacPoolsService < Service
end

class MeasurableService < Service
end

class MoveableService < Service
end

class NetworkService < Service
end

class NetworkAttachmentService < Service
end

class NetworkAttachmentsService < Service
end

class NetworkFilterService < Service
end

class NetworkFiltersService < Service
end

class NetworkLabelService < Service
end

class NetworkLabelsService < Service
end

class NetworksService < Service
end

class OpenstackImageService < Service
end

class OpenstackImageProviderService < ExternalProviderService
end

class OpenstackImageProvidersService < Service
end

class OpenstackImagesService < Service
end

class OpenstackNetworkService < Service
end

class OpenstackNetworkProviderService < ExternalProviderService
end

class OpenstackNetworkProvidersService < Service
end

class OpenstackNetworksService < Service
end

class OpenstackSubnetService < Service
end

class OpenstackSubnetsService < Service
end

class OpenstackVolumeAuthenticationKeyService < Service
end

class OpenstackVolumeAuthenticationKeysService < Service
end

class OpenstackVolumeProviderService < ExternalProviderService
end

class OpenstackVolumeProvidersService < Service
end

class OpenstackVolumeTypeService < Service
end

class OpenstackVolumeTypesService < Service
end

class OperatingSystemService < Service
end

class OperatingSystemsService < Service
end

class PermissionService < Service
end

class PermitService < Service
end

class PermitsService < Service
end

class QosService < Service
end

class QossService < Service
end

class QuotaService < Service
end

class QuotaClusterLimitService < Service
end

class QuotaClusterLimitsService < Service
end

class QuotaStorageLimitService < Service
end

class QuotaStorageLimitsService < Service
end

class QuotasService < Service
end

class RoleService < Service
end

class RolesService < Service
end

class SchedulingPoliciesService < Service
end

class SchedulingPolicyService < Service
end

class SchedulingPolicyUnitService < Service
end

class SchedulingPolicyUnitsService < Service
end

class SnapshotService < Service
end

class SnapshotCdromService < Service
end

class SnapshotCdromsService < Service
end

class SnapshotDiskService < Service
end

class SnapshotDisksService < Service
end

class SnapshotNicService < Service
end

class SnapshotNicsService < Service
end

class SnapshotsService < Service
end

class SshPublicKeyService < Service
end

class SshPublicKeysService < Service
end

class StatisticService < Service
end

class StatisticsService < Service
end

class StepService < MeasurableService
end

class StepsService < Service
end

class StorageService < Service
end

class StorageDomainService < Service
end

class StorageDomainContentDiskService < Service
end

class StorageDomainContentDisksService < Service
end

class StorageDomainServerConnectionService < Service
end

class StorageDomainServerConnectionsService < Service
end

class StorageDomainTemplateService < Service
end

class StorageDomainTemplatesService < Service
end

class StorageDomainVmService < Service
end

class StorageDomainVmsService < Service
end

class StorageDomainsService < Service
end

class StorageServerConnectionService < Service
end

class StorageServerConnectionExtensionService < Service
end

class StorageServerConnectionExtensionsService < Service
end

class StorageServerConnectionsService < Service
end

class SystemService < Service
end

class SystemPermissionsService < AssignedPermissionsService
end

class TagService < Service
end

class TagsService < Service
end

class TemplateService < Service
end

class TemplateCdromService < Service
end

class TemplateCdromsService < Service
end

class TemplateDiskService < Service
end

class TemplateDiskAttachmentService < Service
end

class TemplateDiskAttachmentsService < Service
end

class TemplateDisksService < Service
end

class TemplateNicService < Service
end

class TemplateNicsService < Service
end

class TemplateWatchdogService < Service
end

class TemplateWatchdogsService < Service
end

class TemplatesService < Service
end

class UnmanagedNetworkService < Service
end

class UnmanagedNetworksService < Service
end

class UserService < Service
end

class UsersService < Service
end

class VirtualFunctionAllowedNetworkService < Service
end

class VirtualFunctionAllowedNetworksService < Service
end

class VmService < MeasurableService
end

class VmApplicationService < Service
end

class VmApplicationsService < Service
end

class VmCdromService < Service
end

class VmCdromsService < Service
end

class VmDiskService < MeasurableService
end

class VmDisksService < Service
end

class VmGraphicsConsoleService < GraphicsConsoleService
end

class VmHostDeviceService < Service
end

class VmHostDevicesService < Service
end

class VmNicService < MeasurableService
end

class VmNicsService < Service
end

class VmNumaNodeService < Service
end

class VmNumaNodesService < Service
end

class VmPoolService < Service
end

class VmPoolsService < Service
end

class VmReportedDeviceService < Service
end

class VmReportedDevicesService < Service
end

class VmSessionService < Service
end

class VmSessionsService < Service
end

class VmWatchdogService < Service
end

class VmWatchdogsService < Service
end

class VmsService < Service
end

class VnicProfileService < Service
end

class VnicProfilesService < Service
end

class WeightService < Service
end

class WeightsService < Service
end

class DiskService < MeasurableService
end

class EngineKatelloErrataService < KatelloErrataService
end

class ExternalHostProviderService < ExternalProviderService
end

class GlusterBrickService < MeasurableService
end

class GlusterVolumeService < MeasurableService
end

class HostService < MeasurableService
end

class HostNicService < MeasurableService
end

class HostNumaNodeService < MeasurableService
end

class AffinityGroupService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityGroupReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(group)
    if group.is_a?(Hash)
      group = ovirtsdk4::AffinityGroup.new(group)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AffinityGroupWriter.write_one(group, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityGroupReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `vms` service.
  # 
  # @return [AffinityGroupVmsService] A reference to `vms` service.
  def vms_service
    return AffinityGroupVmsService.new(@connection, "#{@path}/vms")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'vms'
      return vms_service
    end
    if path.start_with?('vms/')
      return vms_service.service(path[4..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityGroupService}:#{@path}>"
  end
  
end

class AffinityGroupVmService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityGroupVmService}:#{@path}>"
  end
  
end

class AffinityGroupVmsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `vm`.
  # 
  def add(vm, opts = {})
    if vm.is_a?(Hash)
      vm = ovirtsdk4::Vm.new(vm)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VmWriter.write_one(vm, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `vm` service.
  # 
  # @param id [String] The identifier of the `vm`.
  # 
  # @return [AffinityGroupVmService] A reference to the `vm` service.
  # 
  def vm_service(id)
    return AffinityGroupVmService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return vm_service(path)
    end
    return vm_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityGroupVmsService}:#{@path}>"
  end
  
end

class AffinityGroupsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `group`.
  # 
  def add(group, opts = {})
    if group.is_a?(Hash)
      group = ovirtsdk4::AffinityGroup.new(group)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AffinityGroupWriter.write_one(group, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return AffinityGroupReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityGroupReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `group` service.
  # 
  # @param id [String] The identifier of the `group`.
  # 
  # @return [AffinityGroupService] A reference to the `group` service.
  # 
  def group_service(id)
    return AffinityGroupService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return group_service(path)
    end
    return group_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityGroupsService}:#{@path}>"
  end
  
end

class AffinityLabelService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Retrieves details about a label.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes a label from system and clears all assignments
  # of the removed label.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates a label.
  # 
  # This call will update all metadata like name
  # or description.
  # 
  def update(label)
    if label.is_a?(Hash)
      label = ovirtsdk4::AffinityLabel.new(label)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AffinityLabelWriter.write_one(label, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # List all Hosts with this label.
  # 
  # @return [AffinityLabelHostsService] A reference to `hosts` service.
  def hosts_service
    return AffinityLabelHostsService.new(@connection, "#{@path}/hosts")
  end
  
  # 
  # List all VMs with this label.
  # 
  # @return [AffinityLabelVmsService] A reference to `vms` service.
  def vms_service
    return AffinityLabelVmsService.new(@connection, "#{@path}/vms")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'hosts'
      return hosts_service
    end
    if path.start_with?('hosts/')
      return hosts_service.service(path[6..-1])
    end
    if path == 'vms'
      return vms_service
    end
    if path.start_with?('vms/')
      return vms_service.service(path[4..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityLabelService}:#{@path}>"
  end
  
end

class AffinityLabelHostService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Retrieves details about a host that has this label assigned.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Remove a label from a host.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityLabelHostService}:#{@path}>"
  end
  
end

class AffinityLabelHostsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Add a label to a host.
  # 
  def add(host, opts = {})
    if host.is_a?(Hash)
      host = ovirtsdk4::Host.new(host)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      HostWriter.write_one(host, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # List all hosts with the label.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # A link to the specific label-host assignment to
  # allow label removal.
  # 
  # @param id [String] The identifier of the `host`.
  # 
  # @return [AffinityLabelHostService] A reference to the `host` service.
  # 
  def host_service(id)
    return AffinityLabelHostService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return host_service(path)
    end
    return host_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityLabelHostsService}:#{@path}>"
  end
  
end

class AffinityLabelVmService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Retrieves details about a vm that has this label assigned.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Remove a label from a vm.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityLabelVmService}:#{@path}>"
  end
  
end

class AffinityLabelVmsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Add a label to a vm.
  # 
  def add(vm, opts = {})
    if vm.is_a?(Hash)
      vm = ovirtsdk4::Vm.new(vm)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VmWriter.write_one(vm, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # List all vms with the label.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # A link to the specific label-vm assignment to
  # allow label removal.
  # 
  # @param id [String] The identifier of the `vm`.
  # 
  # @return [AffinityLabelVmService] A reference to the `vm` service.
  # 
  def vm_service(id)
    return AffinityLabelVmService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return vm_service(path)
    end
    return vm_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityLabelVmsService}:#{@path}>"
  end
  
end

class AffinityLabelsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Creates a new label. The label is automatically attached
  # to all entities mentioned in the vms or hosts lists.
  # 
  def add(label, opts = {})
    if label.is_a?(Hash)
      label = ovirtsdk4::AffinityLabel.new(label)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AffinityLabelWriter.write_one(label, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Lists all labels present in the system.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Link to a single label details.
  # 
  # @param id [String] The identifier of the `label`.
  # 
  # @return [AffinityLabelService] A reference to the `label` service.
  # 
  def label_service(id)
    return AffinityLabelService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return label_service(path)
    end
    return label_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AffinityLabelsService}:#{@path}>"
  end
  
end

class AssignedAffinityLabelService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Retrieves details about the attached label.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes the label from an entity. Does not touch the label itself.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedAffinityLabelService}:#{@path}>"
  end
  
end

class AssignedAffinityLabelsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Attaches a label to an entity.
  # 
  def add(label, opts = {})
    if label.is_a?(Hash)
      label = ovirtsdk4::AffinityLabel.new(label)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AffinityLabelWriter.write_one(label, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Lists all labels that are attached to an entity.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AffinityLabelReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Link to the specific entity-label assignment to allow
  # removal.
  # 
  # @param id [String] The identifier of the `label`.
  # 
  # @return [AssignedAffinityLabelService] A reference to the `label` service.
  # 
  def label_service(id)
    return AssignedAffinityLabelService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return label_service(path)
    end
    return label_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedAffinityLabelsService}:#{@path}>"
  end
  
end

class AssignedCpuProfileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedCpuProfileService}:#{@path}>"
  end
  
end

class AssignedCpuProfilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `profile`.
  # 
  def add(profile, opts = {})
    if profile.is_a?(Hash)
      profile = ovirtsdk4::CpuProfile.new(profile)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      CpuProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `profile` service.
  # 
  # @param id [String] The identifier of the `profile`.
  # 
  # @return [AssignedCpuProfileService] A reference to the `profile` service.
  # 
  def profile_service(id)
    return AssignedCpuProfileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return profile_service(path)
    end
    return profile_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedCpuProfilesService}:#{@path}>"
  end
  
end

class AssignedDiskProfileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedDiskProfileService}:#{@path}>"
  end
  
end

class AssignedDiskProfilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `profile`.
  # 
  def add(profile, opts = {})
    if profile.is_a?(Hash)
      profile = ovirtsdk4::DiskProfile.new(profile)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `profile` service.
  # 
  # @param id [String] The identifier of the `profile`.
  # 
  # @return [AssignedDiskProfileService] A reference to the `profile` service.
  # 
  def profile_service(id)
    return AssignedDiskProfileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return profile_service(path)
    end
    return profile_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedDiskProfilesService}:#{@path}>"
  end
  
end

class AssignedNetworkService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(network)
    if network.is_a?(Hash)
      network = ovirtsdk4::Network.new(network)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkWriter.write_one(network, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedNetworkService}:#{@path}>"
  end
  
end

class AssignedNetworksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `network`.
  # 
  def add(network, opts = {})
    if network.is_a?(Hash)
      network = ovirtsdk4::Network.new(network)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkWriter.write_one(network, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network` service.
  # 
  # @param id [String] The identifier of the `network`.
  # 
  # @return [AssignedNetworkService] A reference to the `network` service.
  # 
  def network_service(id)
    return AssignedNetworkService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return network_service(path)
    end
    return network_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedNetworksService}:#{@path}>"
  end
  
end

class AssignedPermissionsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `permission`.
  # 
  def add(permission, opts = {})
    if permission.is_a?(Hash)
      permission = ovirtsdk4::Permission.new(permission)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      PermissionWriter.write_one(permission, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return PermissionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return PermissionReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Sub-resource locator method, returns individual permission resource on which the remainder of the URI is
  # dispatched.
  # 
  # @param id [String] The identifier of the `permission`.
  # 
  # @return [PermissionService] A reference to the `permission` service.
  # 
  def permission_service(id)
    return PermissionService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return permission_service(path)
    end
    return permission_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedPermissionsService}:#{@path}>"
  end
  
end

class AssignedRolesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return RoleReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Sub-resource locator method, returns individual role resource on which the remainder of the URI is dispatched.
  # 
  # @param id [String] The identifier of the `role`.
  # 
  # @return [RoleService] A reference to the `role` service.
  # 
  def role_service(id)
    return RoleService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return role_service(path)
    end
    return role_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedRolesService}:#{@path}>"
  end
  
end

class AssignedTagService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedTagService}:#{@path}>"
  end
  
end

class AssignedTagsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `tag`.
  # 
  def add(tag, opts = {})
    if tag.is_a?(Hash)
      tag = ovirtsdk4::Tag.new(tag)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      TagWriter.write_one(tag, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `tag` service.
  # 
  # @param id [String] The identifier of the `tag`.
  # 
  # @return [AssignedTagService] A reference to the `tag` service.
  # 
  def tag_service(id)
    return AssignedTagService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return tag_service(path)
    end
    return tag_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedTagsService}:#{@path}>"
  end
  
end

class AssignedVnicProfileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedVnicProfileService}:#{@path}>"
  end
  
end

class AssignedVnicProfilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `profile`.
  # 
  def add(profile, opts = {})
    if profile.is_a?(Hash)
      profile = ovirtsdk4::VnicProfile.new(profile)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VnicProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `profile` service.
  # 
  # @param id [String] The identifier of the `profile`.
  # 
  # @return [AssignedVnicProfileService] A reference to the `profile` service.
  # 
  def profile_service(id)
    return AssignedVnicProfileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return profile_service(path)
    end
    return profile_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AssignedVnicProfilesService}:#{@path}>"
  end
  
end

class AttachedStorageDomainService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `activate` method.
  # 
  def activate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/activate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `deactivate` method.
  # 
  def deactivate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/deactivate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disks` service.
  # 
  # @return [DisksService] A reference to `disks` service.
  def disks_service
    return DisksService.new(@connection, "#{@path}/disks")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'disks'
      return disks_service
    end
    if path.start_with?('disks/')
      return disks_service.service(path[6..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AttachedStorageDomainService}:#{@path}>"
  end
  
end

class AttachedStorageDomainsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `storage_domain`.
  # 
  def add(storage_domain, opts = {})
    if storage_domain.is_a?(Hash)
      storage_domain = ovirtsdk4::StorageDomain.new(storage_domain)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageDomainWriter.write_one(storage_domain, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `storage_domain` service.
  # 
  # @param id [String] The identifier of the `storage_domain`.
  # 
  # @return [AttachedStorageDomainService] A reference to the `storage_domain` service.
  # 
  def storage_domain_service(id)
    return AttachedStorageDomainService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return storage_domain_service(path)
    end
    return storage_domain_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{AttachedStorageDomainsService}:#{@path}>"
  end
  
end

class BalanceService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return BalanceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{BalanceService}:#{@path}>"
  end
  
end

class BalancesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `balance`.
  # 
  def add(balance, opts = {})
    if balance.is_a?(Hash)
      balance = ovirtsdk4::Balance.new(balance)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      BalanceWriter.write_one(balance, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return BalanceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return BalanceReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `balance` service.
  # 
  # @param id [String] The identifier of the `balance`.
  # 
  # @return [BalanceService] A reference to the `balance` service.
  # 
  def balance_service(id)
    return BalanceService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return balance_service(path)
    end
    return balance_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{BalancesService}:#{@path}>"
  end
  
end

class BookmarkService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return BookmarkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(bookmark)
    if bookmark.is_a?(Hash)
      bookmark = ovirtsdk4::Bookmark.new(bookmark)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      BookmarkWriter.write_one(bookmark, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return BookmarkReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{BookmarkService}:#{@path}>"
  end
  
end

class BookmarksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `bookmark`.
  # 
  def add(bookmark, opts = {})
    if bookmark.is_a?(Hash)
      bookmark = ovirtsdk4::Bookmark.new(bookmark)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      BookmarkWriter.write_one(bookmark, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return BookmarkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return BookmarkReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `bookmark` service.
  # 
  # @param id [String] The identifier of the `bookmark`.
  # 
  # @return [BookmarkService] A reference to the `bookmark` service.
  # 
  def bookmark_service(id)
    return BookmarkService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return bookmark_service(path)
    end
    return bookmark_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{BookmarksService}:#{@path}>"
  end
  
end

class ClusterService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ClusterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `reset_emulated_machine` method.
  # 
  def reset_emulated_machine(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/resetemulatedmachine",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(cluster)
    if cluster.is_a?(Hash)
      cluster = ovirtsdk4::Cluster.new(cluster)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      ClusterWriter.write_one(cluster, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ClusterReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `affinity_groups` service.
  # 
  # @return [AffinityGroupsService] A reference to `affinity_groups` service.
  def affinity_groups_service
    return AffinityGroupsService.new(@connection, "#{@path}/affinitygroups")
  end
  
  # 
  # Locates the `cpu_profiles` service.
  # 
  # @return [AssignedCpuProfilesService] A reference to `cpu_profiles` service.
  def cpu_profiles_service
    return AssignedCpuProfilesService.new(@connection, "#{@path}/cpuprofiles")
  end
  
  # 
  # Locates the `gluster_hooks` service.
  # 
  # @return [GlusterHooksService] A reference to `gluster_hooks` service.
  def gluster_hooks_service
    return GlusterHooksService.new(@connection, "#{@path}/glusterhooks")
  end
  
  # 
  # Locates the `gluster_volumes` service.
  # 
  # @return [GlusterVolumesService] A reference to `gluster_volumes` service.
  def gluster_volumes_service
    return GlusterVolumesService.new(@connection, "#{@path}/glustervolumes")
  end
  
  # 
  # A sub collection with all the supported network filters for this cluster.
  # 
  # @return [NetworkFiltersService] A reference to `network_filters` service.
  def network_filters_service
    return NetworkFiltersService.new(@connection, "#{@path}/networkfilters")
  end
  
  # 
  # Locates the `networks` service.
  # 
  # @return [AssignedNetworksService] A reference to `networks` service.
  def networks_service
    return AssignedNetworksService.new(@connection, "#{@path}/networks")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'affinitygroups'
      return affinity_groups_service
    end
    if path.start_with?('affinitygroups/')
      return affinity_groups_service.service(path[15..-1])
    end
    if path == 'cpuprofiles'
      return cpu_profiles_service
    end
    if path.start_with?('cpuprofiles/')
      return cpu_profiles_service.service(path[12..-1])
    end
    if path == 'glusterhooks'
      return gluster_hooks_service
    end
    if path.start_with?('glusterhooks/')
      return gluster_hooks_service.service(path[13..-1])
    end
    if path == 'glustervolumes'
      return gluster_volumes_service
    end
    if path.start_with?('glustervolumes/')
      return gluster_volumes_service.service(path[15..-1])
    end
    if path == 'networkfilters'
      return network_filters_service
    end
    if path.start_with?('networkfilters/')
      return network_filters_service.service(path[15..-1])
    end
    if path == 'networks'
      return networks_service
    end
    if path.start_with?('networks/')
      return networks_service.service(path[9..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ClusterService}:#{@path}>"
  end
  
end

class ClusterLevelService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Provides the information about the capabilities of the specific cluster level managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ClusterLevelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ClusterLevelService}:#{@path}>"
  end
  
end

class ClusterLevelsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Lists the cluster levels supported by the system.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ClusterLevelReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Reference to the service that provides information about an specific cluster level.
  # 
  # @param id [String] The identifier of the `level`.
  # 
  # @return [ClusterLevelService] A reference to the `level` service.
  # 
  def level_service(id)
    return ClusterLevelService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return level_service(path)
    end
    return level_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ClusterLevelsService}:#{@path}>"
  end
  
end

class ClustersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `cluster`.
  # 
  def add(cluster, opts = {})
    if cluster.is_a?(Hash)
      cluster = ovirtsdk4::Cluster.new(cluster)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      ClusterWriter.write_one(cluster, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return ClusterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ClusterReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `cluster` service.
  # 
  # @param id [String] The identifier of the `cluster`.
  # 
  # @return [ClusterService] A reference to the `cluster` service.
  # 
  def cluster_service(id)
    return ClusterService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return cluster_service(path)
    end
    return cluster_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ClustersService}:#{@path}>"
  end
  
end

class CopyableService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `copy` method.
  # 
  def copy(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/copy",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{CopyableService}:#{@path}>"
  end
  
end

class CpuProfileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(profile)
    if profile.is_a?(Hash)
      profile = ovirtsdk4::CpuProfile.new(profile)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      CpuProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{CpuProfileService}:#{@path}>"
  end
  
end

class CpuProfilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `profile`.
  # 
  def add(profile, opts = {})
    if profile.is_a?(Hash)
      profile = ovirtsdk4::CpuProfile.new(profile)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      CpuProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CpuProfileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `profile` service.
  # 
  # @param id [String] The identifier of the `profile`.
  # 
  # @return [CpuProfileService] A reference to the `profile` service.
  # 
  def profile_service(id)
    return CpuProfileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return profile_service(path)
    end
    return profile_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{CpuProfilesService}:#{@path}>"
  end
  
end

class DataCenterService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DataCenterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(data_center)
    if data_center.is_a?(Hash)
      data_center = ovirtsdk4::DataCenter.new(data_center)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DataCenterWriter.write_one(data_center, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DataCenterReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `clusters` service.
  # 
  # @return [ClustersService] A reference to `clusters` service.
  def clusters_service
    return ClustersService.new(@connection, "#{@path}/clusters")
  end
  
  # 
  # Locates the `iscsi_bonds` service.
  # 
  # @return [IscsiBondsService] A reference to `iscsi_bonds` service.
  def iscsi_bonds_service
    return IscsiBondsService.new(@connection, "#{@path}/iscsibonds")
  end
  
  # 
  # Locates the `networks` service.
  # 
  # @return [NetworksService] A reference to `networks` service.
  def networks_service
    return NetworksService.new(@connection, "#{@path}/networks")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `qoss` service.
  # 
  # @return [QossService] A reference to `qoss` service.
  def qoss_service
    return QossService.new(@connection, "#{@path}/qoss")
  end
  
  # 
  # Locates the `quotas` service.
  # 
  # @return [QuotasService] A reference to `quotas` service.
  def quotas_service
    return QuotasService.new(@connection, "#{@path}/quotas")
  end
  
  # 
  # Locates the `storage_domains` service.
  # 
  # @return [AttachedStorageDomainsService] A reference to `storage_domains` service.
  def storage_domains_service
    return AttachedStorageDomainsService.new(@connection, "#{@path}/storagedomains")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'clusters'
      return clusters_service
    end
    if path.start_with?('clusters/')
      return clusters_service.service(path[9..-1])
    end
    if path == 'iscsibonds'
      return iscsi_bonds_service
    end
    if path.start_with?('iscsibonds/')
      return iscsi_bonds_service.service(path[11..-1])
    end
    if path == 'networks'
      return networks_service
    end
    if path.start_with?('networks/')
      return networks_service.service(path[9..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'qoss'
      return qoss_service
    end
    if path.start_with?('qoss/')
      return qoss_service.service(path[5..-1])
    end
    if path == 'quotas'
      return quotas_service
    end
    if path.start_with?('quotas/')
      return quotas_service.service(path[7..-1])
    end
    if path == 'storagedomains'
      return storage_domains_service
    end
    if path.start_with?('storagedomains/')
      return storage_domains_service.service(path[15..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DataCenterService}:#{@path}>"
  end
  
end

class DataCentersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `data_center`.
  # 
  def add(data_center, opts = {})
    if data_center.is_a?(Hash)
      data_center = ovirtsdk4::DataCenter.new(data_center)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DataCenterWriter.write_one(data_center, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return DataCenterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DataCenterReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `data_center` service.
  # 
  # @param id [String] The identifier of the `data_center`.
  # 
  # @return [DataCenterService] A reference to the `data_center` service.
  # 
  def data_center_service(id)
    return DataCenterService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return data_center_service(path)
    end
    return data_center_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DataCentersService}:#{@path}>"
  end
  
end

class DiskAttachmentService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the details of the attachment, including the bootable flag and link to the disk.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes the disk attachment. This will only detach the disk from the virtual machine, but won't remove it from
  # the system, unless the `detach_only` parameter is `false`.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:detach_only]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['detach_only'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Update the disk attachment and the disk properties within it.
  # 
  # [source]
  # ----
  # PUT /vms/{vm:id}/disksattachments/{attachment:id}
  # <disk_attachment>
  #   <bootable>true</bootable>
  #   <interface>ide</interface>
  #   <disk>
  #     <name>mydisk</name>
  #     <provisioned_size>1024</provisioned_size>
  #     ...
  #   </disk>
  # </disk_attachment>
  # ----
  # 
  def update(disk_attachment)
    if disk_attachment.is_a?(Hash)
      disk_attachment = ovirtsdk4::DiskAttachment.new(disk_attachment)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskAttachmentWriter.write_one(disk_attachment, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskAttachmentService}:#{@path}>"
  end
  
end

class DiskAttachmentsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new disk attachment to the virtual machine. The `attachment` parameter can contain just a reference, if
  # the disk already exists:
  # 
  # [source,xml]
  # ----
  # <disk_attachment>
  #   <bootable>true</bootable>
  #   <pass_discard>true</pass_discard>
  #   <interface>ide</interface>
  #   <disk id="123"/>
  # </disk_attachment>
  # ----
  # 
  # Or it can contain the complete representation of the disk, if the disk doesn't exist yet:
  # 
  # [source,xml]
  # ----
  # <disk_attachment>
  #   <bootable>true</bootable>
  #   <pass_discard>true</pass_discard>
  #   <interface>ide</interface>
  #   <disk>
  #     <name>mydisk</name>
  #     <provisioned_size>1024</provisioned_size>
  #     ...
  #   </disk>
  # </disk_attachment>
  # ----
  # 
  # In this case the disk will be created and then attached to the virtual machine.
  # 
  def add(attachment, opts = {})
    if attachment.is_a?(Hash)
      attachment = ovirtsdk4::DiskAttachment.new(attachment)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskAttachmentWriter.write_one(attachment, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return DiskAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # List the disk that are attached to the virtual machine.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskAttachmentReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Reference to the service that manages a specific attachment.
  # 
  # @param id [String] The identifier of the `attachment`.
  # 
  # @return [DiskAttachmentService] A reference to the `attachment` service.
  # 
  def attachment_service(id)
    return DiskAttachmentService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return attachment_service(path)
    end
    return attachment_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskAttachmentsService}:#{@path}>"
  end
  
end

class DiskProfileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(profile)
    if profile.is_a?(Hash)
      profile = ovirtsdk4::DiskProfile.new(profile)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskProfileService}:#{@path}>"
  end
  
end

class DiskProfilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `profile`.
  # 
  def add(profile, opts = {})
    if profile.is_a?(Hash)
      profile = ovirtsdk4::DiskProfile.new(profile)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskProfileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk_profile` service.
  # 
  # @param id [String] The identifier of the `disk_profile`.
  # 
  # @return [DiskProfileService] A reference to the `disk_profile` service.
  # 
  def disk_profile_service(id)
    return DiskProfileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return disk_profile_service(path)
    end
    return disk_profile_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskProfilesService}:#{@path}>"
  end
  
end

class DiskSnapshotService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskSnapshotReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskSnapshotService}:#{@path}>"
  end
  
end

class DiskSnapshotsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskSnapshotReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `snapshot` service.
  # 
  # @param id [String] The identifier of the `snapshot`.
  # 
  # @return [DiskSnapshotService] A reference to the `snapshot` service.
  # 
  def snapshot_service(id)
    return DiskSnapshotService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return snapshot_service(path)
    end
    return snapshot_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskSnapshotsService}:#{@path}>"
  end
  
end

class DisksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `disk`.
  # 
  def add(disk, opts = {})
    if disk.is_a?(Hash)
      disk = ovirtsdk4::Disk.new(disk)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskWriter.write_one(disk, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk` service.
  # 
  # @param id [String] The identifier of the `disk`.
  # 
  # @return [DiskService] A reference to the `disk` service.
  # 
  def disk_service(id)
    return DiskService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return disk_service(path)
    end
    return disk_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DisksService}:#{@path}>"
  end
  
end

class DomainService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DomainReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `groups` service.
  # 
  # @return [DomainGroupsService] A reference to `groups` service.
  def groups_service
    return DomainGroupsService.new(@connection, "#{@path}/groups")
  end
  
  # 
  # Locates the `users` service.
  # 
  # @return [DomainUsersService] A reference to `users` service.
  def users_service
    return DomainUsersService.new(@connection, "#{@path}/users")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'groups'
      return groups_service
    end
    if path.start_with?('groups/')
      return groups_service.service(path[7..-1])
    end
    if path == 'users'
      return users_service
    end
    if path.start_with?('users/')
      return users_service.service(path[6..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DomainService}:#{@path}>"
  end
  
end

class DomainGroupService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GroupReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DomainGroupService}:#{@path}>"
  end
  
end

class DomainGroupsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GroupReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `group` service.
  # 
  # @param id [String] The identifier of the `group`.
  # 
  # @return [DomainGroupService] A reference to the `group` service.
  # 
  def group_service(id)
    return DomainGroupService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return group_service(path)
    end
    return group_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DomainGroupsService}:#{@path}>"
  end
  
end

class DomainUserService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return UserReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DomainUserService}:#{@path}>"
  end
  
end

class DomainUsersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return UserReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `user` service.
  # 
  # @param id [String] The identifier of the `user`.
  # 
  # @return [DomainUserService] A reference to the `user` service.
  # 
  def user_service(id)
    return DomainUserService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return user_service(path)
    end
    return user_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DomainUsersService}:#{@path}>"
  end
  
end

class DomainsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DomainReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `domain` service.
  # 
  # @param id [String] The identifier of the `domain`.
  # 
  # @return [DomainService] A reference to the `domain` service.
  # 
  def domain_service(id)
    return DomainService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return domain_service(path)
    end
    return domain_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DomainsService}:#{@path}>"
  end
  
end

class EventService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return EventReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{EventService}:#{@path}>"
  end
  
end

class EventsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds an external event to the internal audit log.
  # 
  # This is intended for integration with external systems that detect or produce events relevant for the
  # administrator of the system. For example, an external monitoring tool may be able to detect that a file system
  # is full inside the guest operating system of a virtual machine. This event can be added to the internal audit
  # log sending a request like this:
  # 
  # [source]
  # ----
  # POST /ovirt-engine/api/events
  # <event>
  #   <description>File system /home is full</description>
  #   <severity>alert</severity>
  #   <origin>mymonitor</origin>
  #   <custom_id>1467879754</custom_id>
  # </event>
  # ----
  # 
  # Events can also be linked to specific objects. For example, the above event could be linked to the specific
  # virtual machine where it happened, using the `vm` link:
  # 
  # [source]
  # ----
  # POST /ovirt-engine/api/events
  # <event>
  #   <description>File system /home is full</description>
  #   <severity>alert</severity>
  #   <origin>mymonitor</origin>
  #   <custom_id>1467879754</custom_id>
  #   <vm id="aae98225-5b73-490d-a252-899209af17e9"/>
  # </event>
  # ----
  # 
  # NOTE: When using links, like the `vm` in the previous example, only the `id` attribute is accepted. The `name`
  # attribute, if provided, is simply ignored.
  # 
  def add(event, opts = {})
    if event.is_a?(Hash)
      event = ovirtsdk4::Event.new(event)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      EventWriter.write_one(event, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return EventReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:from]
    unless value.nil?
      value = Writer.render_integer(value)
      query['from'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return EventReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `undelete` method.
  # 
  def undelete(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/undelete",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `event` service.
  # 
  # @param id [String] The identifier of the `event`.
  # 
  # @return [EventService] A reference to the `event` service.
  # 
  def event_service(id)
    return EventService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return event_service(path)
    end
    return event_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{EventsService}:#{@path}>"
  end
  
end

class ExternalComputeResourceService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalComputeResourceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalComputeResourceService}:#{@path}>"
  end
  
end

class ExternalComputeResourcesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalComputeResourceReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `resource` service.
  # 
  # @param id [String] The identifier of the `resource`.
  # 
  # @return [ExternalComputeResourceService] A reference to the `resource` service.
  # 
  def resource_service(id)
    return ExternalComputeResourceService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return resource_service(path)
    end
    return resource_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalComputeResourcesService}:#{@path}>"
  end
  
end

class ExternalDiscoveredHostService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalDiscoveredHostReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalDiscoveredHostService}:#{@path}>"
  end
  
end

class ExternalDiscoveredHostsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalDiscoveredHostReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `host` service.
  # 
  # @param id [String] The identifier of the `host`.
  # 
  # @return [ExternalDiscoveredHostService] A reference to the `host` service.
  # 
  def host_service(id)
    return ExternalDiscoveredHostService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return host_service(path)
    end
    return host_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalDiscoveredHostsService}:#{@path}>"
  end
  
end

class ExternalHostService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalHostService}:#{@path}>"
  end
  
end

class ExternalHostGroupService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostGroupReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalHostGroupService}:#{@path}>"
  end
  
end

class ExternalHostGroupsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostGroupReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `group` service.
  # 
  # @param id [String] The identifier of the `group`.
  # 
  # @return [ExternalHostGroupService] A reference to the `group` service.
  # 
  def group_service(id)
    return ExternalHostGroupService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return group_service(path)
    end
    return group_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalHostGroupsService}:#{@path}>"
  end
  
end

class ExternalHostProvidersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `provider`.
  # 
  def add(provider, opts = {})
    if provider.is_a?(Hash)
      provider = ovirtsdk4::ExternalHostProvider.new(provider)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      ExternalHostProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostProviderReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `provider` service.
  # 
  # @param id [String] The identifier of the `provider`.
  # 
  # @return [ExternalHostProviderService] A reference to the `provider` service.
  # 
  def provider_service(id)
    return ExternalHostProviderService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return provider_service(path)
    end
    return provider_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalHostProvidersService}:#{@path}>"
  end
  
end

class ExternalHostsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `host` service.
  # 
  # @param id [String] The identifier of the `host`.
  # 
  # @return [ExternalHostService] A reference to the `host` service.
  # 
  def host_service(id)
    return ExternalHostService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return host_service(path)
    end
    return host_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalHostsService}:#{@path}>"
  end
  
end

class ExternalProviderService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `import_certificates` method.
  # 
  def import_certificates(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/importcertificates",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `test_connectivity` method.
  # 
  def test_connectivity(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/testconnectivity",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `certificates` service.
  # 
  # @return [ExternalProviderCertificatesService] A reference to `certificates` service.
  def certificates_service
    return ExternalProviderCertificatesService.new(@connection, "#{@path}/certificates")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'certificates'
      return certificates_service
    end
    if path.start_with?('certificates/')
      return certificates_service.service(path[13..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalProviderService}:#{@path}>"
  end
  
end

class ExternalProviderCertificateService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CertificateReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalProviderCertificateService}:#{@path}>"
  end
  
end

class ExternalProviderCertificatesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CertificateReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `certificate` service.
  # 
  # @param id [String] The identifier of the `certificate`.
  # 
  # @return [ExternalProviderCertificateService] A reference to the `certificate` service.
  # 
  def certificate_service(id)
    return ExternalProviderCertificateService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return certificate_service(path)
    end
    return certificate_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalProviderCertificatesService}:#{@path}>"
  end
  
end

class FenceAgentService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AgentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(agent)
    if agent.is_a?(Hash)
      agent = ovirtsdk4::Agent.new(agent)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AgentWriter.write_one(agent, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AgentReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{FenceAgentService}:#{@path}>"
  end
  
end

class FenceAgentsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `agent`.
  # 
  def add(agent, opts = {})
    if agent.is_a?(Hash)
      agent = ovirtsdk4::Agent.new(agent)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      AgentWriter.write_one(agent, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return AgentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return AgentReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `agent` service.
  # 
  # @param id [String] The identifier of the `agent`.
  # 
  # @return [FenceAgentService] A reference to the `agent` service.
  # 
  def agent_service(id)
    return FenceAgentService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return agent_service(path)
    end
    return agent_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{FenceAgentsService}:#{@path}>"
  end
  
end

class FileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return FileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{FileService}:#{@path}>"
  end
  
end

class FilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return FileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `file` service.
  # 
  # @param id [String] The identifier of the `file`.
  # 
  # @return [FileService] A reference to the `file` service.
  # 
  def file_service(id)
    return FileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return file_service(path)
    end
    return file_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{FilesService}:#{@path}>"
  end
  
end

class FilterService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return FilterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{FilterService}:#{@path}>"
  end
  
end

class FiltersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `filter`.
  # 
  def add(filter, opts = {})
    if filter.is_a?(Hash)
      filter = ovirtsdk4::Filter.new(filter)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      FilterWriter.write_one(filter, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return FilterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return FilterReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `filter` service.
  # 
  # @param id [String] The identifier of the `filter`.
  # 
  # @return [FilterService] A reference to the `filter` service.
  # 
  def filter_service(id)
    return FilterService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return filter_service(path)
    end
    return filter_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{FiltersService}:#{@path}>"
  end
  
end

class GlusterBricksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `activate` method.
  # 
  def activate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/activate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Adds given list of bricks to the volume, and updates the database accordingly. The properties `serverId` and
  # `brickDir`are required.
  # 
  def add(bricks, opts = {})
    if bricks.is_a?(Array)
      bricks = List.new(bricks)
      bricks.each_with_index do |value, index|
        if value.is_a?(Hash)
          bricks[index] = ovirtsdk4::GlusterBrick.new(value)
        end
      end
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      GlusterBrickWriter.write_many(bricks, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return GlusterBrickReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GlusterBrickReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `migrate` method.
  # 
  def migrate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/migrate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes the given list of bricks brick from the volume and deletes them from the database.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    value = opts[:bricks]
    unless value.nil?
      query['bricks'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `stop_migrate` method.
  # 
  def stop_migrate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/stopmigrate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `brick` service.
  # 
  # @param id [String] The identifier of the `brick`.
  # 
  # @return [GlusterBrickService] A reference to the `brick` service.
  # 
  def brick_service(id)
    return GlusterBrickService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return brick_service(path)
    end
    return brick_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GlusterBricksService}:#{@path}>"
  end
  
end

class GlusterHookService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Resolves status conflict of hook among servers in cluster by disabling Gluster hook in all servers of the
  # cluster. This updates the hook status to `DISABLED` in database.
  # 
  def disable(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/disable",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Resolves status conflict of hook among servers in cluster by disabling Gluster hook in all servers of the
  # cluster. This updates the hook status to `DISABLED` in database.
  # 
  def enable(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/enable",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GlusterHookReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes the this Gluster hook from all servers in cluster and deletes it from the database.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Resolves missing hook conflict depending on the resolution type.
  # 
  # For `ADD` resolves by copying hook stored in engine database to all servers where the hook is missing. The
  # engine maintains a list of all servers where hook is missing.
  # 
  # For `COPY` resolves conflict in hook content by copying hook stored in engine database to all servers where
  # the hook is missing. The engine maintains a list of all servers where the content is conflicting. If a host
  # id is passed as parameter, the hook content from the server is used as the master to copy to other servers
  # in cluster.
  # 
  def resolve(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/resolve",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GlusterHookService}:#{@path}>"
  end
  
end

class GlusterHooksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GlusterHookReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `hook` service.
  # 
  # @param id [String] The identifier of the `hook`.
  # 
  # @return [GlusterHookService] A reference to the `hook` service.
  # 
  def hook_service(id)
    return GlusterHookService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return hook_service(path)
    end
    return hook_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GlusterHooksService}:#{@path}>"
  end
  
end

class GlusterVolumesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Creates a new Gluster volume and adds it to the database. The volume is created based on properties of the
  # `volume` parameter. The properties `name`, `volumeType` and `bricks` are required.
  # 
  def add(volume, opts = {})
    if volume.is_a?(Hash)
      volume = ovirtsdk4::GlusterVolume.new(volume)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      GlusterVolumeWriter.write_one(volume, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return GlusterVolumeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GlusterVolumeReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `volume` service.
  # 
  # @param id [String] The identifier of the `volume`.
  # 
  # @return [GlusterVolumeService] A reference to the `volume` service.
  # 
  def volume_service(id)
    return GlusterVolumeService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return volume_service(path)
    end
    return volume_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GlusterVolumesService}:#{@path}>"
  end
  
end

class GraphicsConsoleService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GraphicsConsoleReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GraphicsConsoleService}:#{@path}>"
  end
  
end

class GraphicsConsolesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `console`.
  # 
  def add(console, opts = {})
    if console.is_a?(Hash)
      console = ovirtsdk4::GraphicsConsole.new(console)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      GraphicsConsoleWriter.write_one(console, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return GraphicsConsoleReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GraphicsConsoleReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `console` service.
  # 
  # @param id [String] The identifier of the `console`.
  # 
  # @return [GraphicsConsoleService] A reference to the `console` service.
  # 
  def console_service(id)
    return GraphicsConsoleService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return console_service(path)
    end
    return console_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GraphicsConsolesService}:#{@path}>"
  end
  
end

class GroupService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GroupReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `roles` service.
  # 
  # @return [AssignedRolesService] A reference to `roles` service.
  def roles_service
    return AssignedRolesService.new(@connection, "#{@path}/roles")
  end
  
  # 
  # Locates the `tags` service.
  # 
  # @return [AssignedTagsService] A reference to `tags` service.
  def tags_service
    return AssignedTagsService.new(@connection, "#{@path}/tags")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'roles'
      return roles_service
    end
    if path.start_with?('roles/')
      return roles_service.service(path[6..-1])
    end
    if path == 'tags'
      return tags_service
    end
    if path.start_with?('tags/')
      return tags_service.service(path[5..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GroupService}:#{@path}>"
  end
  
end

class GroupsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `group`.
  # 
  def add(group, opts = {})
    if group.is_a?(Hash)
      group = ovirtsdk4::Group.new(group)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      GroupWriter.write_one(group, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return GroupReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GroupReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `group` service.
  # 
  # @param id [String] The identifier of the `group`.
  # 
  # @return [GroupService] A reference to the `group` service.
  # 
  def group_service(id)
    return GroupService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return group_service(path)
    end
    return group_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GroupsService}:#{@path}>"
  end
  
end

class HostDeviceService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostDeviceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostDeviceService}:#{@path}>"
  end
  
end

class HostDevicesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostDeviceReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `device` service.
  # 
  # @param id [String] The identifier of the `device`.
  # 
  # @return [HostDeviceService] A reference to the `device` service.
  # 
  def device_service(id)
    return HostDeviceService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return device_service(path)
    end
    return device_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostDevicesService}:#{@path}>"
  end
  
end

class HostHookService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HookReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostHookService}:#{@path}>"
  end
  
end

class HostHooksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HookReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `hook` service.
  # 
  # @param id [String] The identifier of the `hook`.
  # 
  # @return [HostHookService] A reference to the `hook` service.
  # 
  def hook_service(id)
    return HostHookService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return hook_service(path)
    end
    return hook_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostHooksService}:#{@path}>"
  end
  
end

class HostNicsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostNicReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `nic` service.
  # 
  # @param id [String] The identifier of the `nic`.
  # 
  # @return [HostNicService] A reference to the `nic` service.
  # 
  def nic_service(id)
    return HostNicService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return nic_service(path)
    end
    return nic_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostNicsService}:#{@path}>"
  end
  
end

class HostNumaNodesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NumaNodeReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `node` service.
  # 
  # @param id [String] The identifier of the `node`.
  # 
  # @return [HostNumaNodeService] A reference to the `node` service.
  # 
  def node_service(id)
    return HostNumaNodeService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return node_service(path)
    end
    return node_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostNumaNodesService}:#{@path}>"
  end
  
end

class HostStorageService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:report_status]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['report_status'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostStorageReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `storage` service.
  # 
  # @param id [String] The identifier of the `storage`.
  # 
  # @return [StorageService] A reference to the `storage` service.
  # 
  def storage_service(id)
    return StorageService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return storage_service(path)
    end
    return storage_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostStorageService}:#{@path}>"
  end
  
end

class HostsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Creates a new host and adds it to the database. The host is created based on the properties of the `host`
  # parameter. The `name`, `address` `rootPassword` properties are required.
  # 
  def add(host, opts = {})
    if host.is_a?(Hash)
      host = ovirtsdk4::Host.new(host)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      HostWriter.write_one(host, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `host` service.
  # 
  # @param id [String] The identifier of the `host`.
  # 
  # @return [HostService] A reference to the `host` service.
  # 
  def host_service(id)
    return HostService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return host_service(path)
    end
    return host_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostsService}:#{@path}>"
  end
  
end

class IconService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return IconReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{IconService}:#{@path}>"
  end
  
end

class IconsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return IconReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `icon` service.
  # 
  # @param id [String] The identifier of the `icon`.
  # 
  # @return [IconService] A reference to the `icon` service.
  # 
  def icon_service(id)
    return IconService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return icon_service(path)
    end
    return icon_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{IconsService}:#{@path}>"
  end
  
end

class ImageService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ImageReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import` method.
  # 
  def import(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/import",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ImageService}:#{@path}>"
  end
  
end

class ImagesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ImageReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `image` service.
  # 
  # @param id [String] The identifier of the `image`.
  # 
  # @return [ImageService] A reference to the `image` service.
  # 
  def image_service(id)
    return ImageService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return image_service(path)
    end
    return image_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ImagesService}:#{@path}>"
  end
  
end

class InstanceTypeService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return InstanceTypeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(instance_type)
    if instance_type.is_a?(Hash)
      instance_type = ovirtsdk4::InstanceType.new(instance_type)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      InstanceTypeWriter.write_one(instance_type, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return InstanceTypeReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `graphics_consoles` service.
  # 
  # @return [GraphicsConsolesService] A reference to `graphics_consoles` service.
  def graphics_consoles_service
    return GraphicsConsolesService.new(@connection, "#{@path}/graphicsconsoles")
  end
  
  # 
  # Locates the `nics` service.
  # 
  # @return [InstanceTypeNicsService] A reference to `nics` service.
  def nics_service
    return InstanceTypeNicsService.new(@connection, "#{@path}/nics")
  end
  
  # 
  # Locates the `watchdogs` service.
  # 
  # @return [InstanceTypeWatchdogsService] A reference to `watchdogs` service.
  def watchdogs_service
    return InstanceTypeWatchdogsService.new(@connection, "#{@path}/watchdogs")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'graphicsconsoles'
      return graphics_consoles_service
    end
    if path.start_with?('graphicsconsoles/')
      return graphics_consoles_service.service(path[17..-1])
    end
    if path == 'nics'
      return nics_service
    end
    if path.start_with?('nics/')
      return nics_service.service(path[5..-1])
    end
    if path == 'watchdogs'
      return watchdogs_service
    end
    if path.start_with?('watchdogs/')
      return watchdogs_service.service(path[10..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{InstanceTypeService}:#{@path}>"
  end
  
end

class InstanceTypeNicService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(nic)
    if nic.is_a?(Hash)
      nic = ovirtsdk4::Nic.new(nic)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NicWriter.write_one(nic, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{InstanceTypeNicService}:#{@path}>"
  end
  
end

class InstanceTypeNicsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `nic`.
  # 
  def add(nic, opts = {})
    if nic.is_a?(Hash)
      nic = ovirtsdk4::Nic.new(nic)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NicWriter.write_one(nic, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `nic` service.
  # 
  # @param id [String] The identifier of the `nic`.
  # 
  # @return [InstanceTypeNicService] A reference to the `nic` service.
  # 
  def nic_service(id)
    return InstanceTypeNicService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return nic_service(path)
    end
    return nic_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{InstanceTypeNicsService}:#{@path}>"
  end
  
end

class InstanceTypeWatchdogService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(watchdog)
    if watchdog.is_a?(Hash)
      watchdog = ovirtsdk4::Watchdog.new(watchdog)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WatchdogWriter.write_one(watchdog, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{InstanceTypeWatchdogService}:#{@path}>"
  end
  
end

class InstanceTypeWatchdogsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `watchdog`.
  # 
  def add(watchdog, opts = {})
    if watchdog.is_a?(Hash)
      watchdog = ovirtsdk4::Watchdog.new(watchdog)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WatchdogWriter.write_one(watchdog, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `watchdog` service.
  # 
  # @param id [String] The identifier of the `watchdog`.
  # 
  # @return [InstanceTypeWatchdogService] A reference to the `watchdog` service.
  # 
  def watchdog_service(id)
    return InstanceTypeWatchdogService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return watchdog_service(path)
    end
    return watchdog_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{InstanceTypeWatchdogsService}:#{@path}>"
  end
  
end

class InstanceTypesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `instance_type`.
  # 
  def add(instance_type, opts = {})
    if instance_type.is_a?(Hash)
      instance_type = ovirtsdk4::InstanceType.new(instance_type)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      InstanceTypeWriter.write_one(instance_type, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return InstanceTypeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return InstanceTypeReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `instance_type` service.
  # 
  # @param id [String] The identifier of the `instance_type`.
  # 
  # @return [InstanceTypeService] A reference to the `instance_type` service.
  # 
  def instance_type_service(id)
    return InstanceTypeService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return instance_type_service(path)
    end
    return instance_type_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{InstanceTypesService}:#{@path}>"
  end
  
end

class IscsiBondService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return IscsiBondReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(bond)
    if bond.is_a?(Hash)
      bond = ovirtsdk4::IscsiBond.new(bond)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      IscsiBondWriter.write_one(bond, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return IscsiBondReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `networks` service.
  # 
  # @return [NetworksService] A reference to `networks` service.
  def networks_service
    return NetworksService.new(@connection, "#{@path}/networks")
  end
  
  # 
  # Locates the `storage_server_connections` service.
  # 
  # @return [StorageServerConnectionsService] A reference to `storage_server_connections` service.
  def storage_server_connections_service
    return StorageServerConnectionsService.new(@connection, "#{@path}/storageserverconnections")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'networks'
      return networks_service
    end
    if path.start_with?('networks/')
      return networks_service.service(path[9..-1])
    end
    if path == 'storageserverconnections'
      return storage_server_connections_service
    end
    if path.start_with?('storageserverconnections/')
      return storage_server_connections_service.service(path[25..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{IscsiBondService}:#{@path}>"
  end
  
end

class IscsiBondsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `bond`.
  # 
  def add(bond, opts = {})
    if bond.is_a?(Hash)
      bond = ovirtsdk4::IscsiBond.new(bond)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      IscsiBondWriter.write_one(bond, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return IscsiBondReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return IscsiBondReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `iscsi_bond` service.
  # 
  # @param id [String] The identifier of the `iscsi_bond`.
  # 
  # @return [IscsiBondService] A reference to the `iscsi_bond` service.
  # 
  def iscsi_bond_service(id)
    return IscsiBondService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return iscsi_bond_service(path)
    end
    return iscsi_bond_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{IscsiBondsService}:#{@path}>"
  end
  
end

class JobService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `clear` method.
  # 
  def clear(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/clear",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `end_` method.
  # 
  def end_(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/end",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return JobReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `steps` service.
  # 
  # @return [StepsService] A reference to `steps` service.
  def steps_service
    return StepsService.new(@connection, "#{@path}/steps")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'steps'
      return steps_service
    end
    if path.start_with?('steps/')
      return steps_service.service(path[6..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{JobService}:#{@path}>"
  end
  
end

class JobsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `job`.
  # 
  def add(job, opts = {})
    if job.is_a?(Hash)
      job = ovirtsdk4::Job.new(job)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      JobWriter.write_one(job, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return JobReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return JobReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `job` service.
  # 
  # @param id [String] The identifier of the `job`.
  # 
  # @return [JobService] A reference to the `job` service.
  # 
  def job_service(id)
    return JobService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return job_service(path)
    end
    return job_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{JobsService}:#{@path}>"
  end
  
end

class KatelloErrataService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return KatelloErratumReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `katello_erratum` service.
  # 
  # @param id [String] The identifier of the `katello_erratum`.
  # 
  # @return [KatelloErratumService] A reference to the `katello_erratum` service.
  # 
  def katello_erratum_service(id)
    return KatelloErratumService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return katello_erratum_service(path)
    end
    return katello_erratum_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{KatelloErrataService}:#{@path}>"
  end
  
end

class KatelloErratumService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return KatelloErratumReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{KatelloErratumService}:#{@path}>"
  end
  
end

class MacPoolService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return MacPoolReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(pool)
    if pool.is_a?(Hash)
      pool = ovirtsdk4::MacPool.new(pool)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      MacPoolWriter.write_one(pool, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return MacPoolReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{MacPoolService}:#{@path}>"
  end
  
end

class MacPoolsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `pool`.
  # 
  def add(pool, opts = {})
    if pool.is_a?(Hash)
      pool = ovirtsdk4::MacPool.new(pool)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      MacPoolWriter.write_one(pool, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return MacPoolReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return MacPoolReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `mac_pool` service.
  # 
  # @param id [String] The identifier of the `mac_pool`.
  # 
  # @return [MacPoolService] A reference to the `mac_pool` service.
  # 
  def mac_pool_service(id)
    return MacPoolService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return mac_pool_service(path)
    end
    return mac_pool_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{MacPoolsService}:#{@path}>"
  end
  
end

class MeasurableService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{MeasurableService}:#{@path}>"
  end
  
end

class MoveableService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `move` method.
  # 
  def move(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/move",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{MoveableService}:#{@path}>"
  end
  
end

class NetworkService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(network)
    if network.is_a?(Hash)
      network = ovirtsdk4::Network.new(network)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkWriter.write_one(network, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network_labels` service.
  # 
  # @return [NetworkLabelsService] A reference to `network_labels` service.
  def network_labels_service
    return NetworkLabelsService.new(@connection, "#{@path}/networklabels")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `vnic_profiles` service.
  # 
  # @return [AssignedVnicProfilesService] A reference to `vnic_profiles` service.
  def vnic_profiles_service
    return AssignedVnicProfilesService.new(@connection, "#{@path}/vnicprofiles")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'networklabels'
      return network_labels_service
    end
    if path.start_with?('networklabels/')
      return network_labels_service.service(path[14..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'vnicprofiles'
      return vnic_profiles_service
    end
    if path.start_with?('vnicprofiles/')
      return vnic_profiles_service.service(path[13..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkService}:#{@path}>"
  end
  
end

class NetworkAttachmentService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(attachment)
    if attachment.is_a?(Hash)
      attachment = ovirtsdk4::NetworkAttachment.new(attachment)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkAttachmentWriter.write_one(attachment, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkAttachmentService}:#{@path}>"
  end
  
end

class NetworkAttachmentsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `attachment`.
  # 
  def add(attachment, opts = {})
    if attachment.is_a?(Hash)
      attachment = ovirtsdk4::NetworkAttachment.new(attachment)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkAttachmentWriter.write_one(attachment, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NetworkAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkAttachmentReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `attachment` service.
  # 
  # @param id [String] The identifier of the `attachment`.
  # 
  # @return [NetworkAttachmentService] A reference to the `attachment` service.
  # 
  def attachment_service(id)
    return NetworkAttachmentService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return attachment_service(path)
    end
    return attachment_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkAttachmentsService}:#{@path}>"
  end
  
end

class NetworkFilterService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Retrieves a representation of the network filter.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkFilterReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkFilterService}:#{@path}>"
  end
  
end

class NetworkFiltersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Retrieves the representations of the network filters.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkFilterReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network_filter` service.
  # 
  # @param id [String] The identifier of the `network_filter`.
  # 
  # @return [NetworkFilterService] A reference to the `network_filter` service.
  # 
  def network_filter_service(id)
    return NetworkFilterService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return network_filter_service(path)
    end
    return network_filter_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkFiltersService}:#{@path}>"
  end
  
end

class NetworkLabelService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkLabelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkLabelService}:#{@path}>"
  end
  
end

class NetworkLabelsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `label`.
  # 
  def add(label, opts = {})
    if label.is_a?(Hash)
      label = ovirtsdk4::NetworkLabel.new(label)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkLabelWriter.write_one(label, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NetworkLabelReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkLabelReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `label` service.
  # 
  # @param id [String] The identifier of the `label`.
  # 
  # @return [NetworkLabelService] A reference to the `label` service.
  # 
  def label_service(id)
    return NetworkLabelService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return label_service(path)
    end
    return label_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworkLabelsService}:#{@path}>"
  end
  
end

class NetworksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `network`.
  # 
  def add(network, opts = {})
    if network.is_a?(Hash)
      network = ovirtsdk4::Network.new(network)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkWriter.write_one(network, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network` service.
  # 
  # @param id [String] The identifier of the `network`.
  # 
  # @return [NetworkService] A reference to the `network` service.
  # 
  def network_service(id)
    return NetworkService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return network_service(path)
    end
    return network_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{NetworksService}:#{@path}>"
  end
  
end

class OpenstackImageService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackImageReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import` method.
  # 
  def import(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/import",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackImageService}:#{@path}>"
  end
  
end

class OpenstackImageProviderService < ExternalProviderService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackImageProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import_certificates` method.
  # 
  def import_certificates(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/importcertificates",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `test_connectivity` method.
  # 
  def test_connectivity(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/testconnectivity",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(provider)
    if provider.is_a?(Hash)
      provider = ovirtsdk4::OpenStackImageProvider.new(provider)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackImageProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackImageProviderReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `certificates` service.
  # 
  # @return [ExternalProviderCertificatesService] A reference to `certificates` service.
  def certificates_service
    return ExternalProviderCertificatesService.new(@connection, "#{@path}/certificates")
  end
  
  # 
  # Locates the `images` service.
  # 
  # @return [OpenstackImagesService] A reference to `images` service.
  def images_service
    return OpenstackImagesService.new(@connection, "#{@path}/images")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'certificates'
      return certificates_service
    end
    if path.start_with?('certificates/')
      return certificates_service.service(path[13..-1])
    end
    if path == 'images'
      return images_service
    end
    if path.start_with?('images/')
      return images_service.service(path[7..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackImageProviderService}:#{@path}>"
  end
  
end

class OpenstackImageProvidersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `provider`.
  # 
  def add(provider, opts = {})
    if provider.is_a?(Hash)
      provider = ovirtsdk4::OpenStackImageProvider.new(provider)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackImageProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return OpenStackImageProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackImageProviderReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `provider` service.
  # 
  # @param id [String] The identifier of the `provider`.
  # 
  # @return [OpenstackImageProviderService] A reference to the `provider` service.
  # 
  def provider_service(id)
    return OpenstackImageProviderService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return provider_service(path)
    end
    return provider_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackImageProvidersService}:#{@path}>"
  end
  
end

class OpenstackImagesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackImageReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `image` service.
  # 
  # @param id [String] The identifier of the `image`.
  # 
  # @return [OpenstackImageService] A reference to the `image` service.
  # 
  def image_service(id)
    return OpenstackImageService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return image_service(path)
    end
    return image_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackImagesService}:#{@path}>"
  end
  
end

class OpenstackNetworkService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackNetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # This operation imports an external network into oVirt.
  # The network will be added to the data center specified.
  # 
  def import(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/import",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `subnets` service.
  # 
  # @return [OpenstackSubnetsService] A reference to `subnets` service.
  def subnets_service
    return OpenstackSubnetsService.new(@connection, "#{@path}/subnets")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'subnets'
      return subnets_service
    end
    if path.start_with?('subnets/')
      return subnets_service.service(path[8..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackNetworkService}:#{@path}>"
  end
  
end

class OpenstackNetworkProviderService < ExternalProviderService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackNetworkProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import_certificates` method.
  # 
  def import_certificates(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/importcertificates",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `test_connectivity` method.
  # 
  def test_connectivity(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/testconnectivity",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(provider)
    if provider.is_a?(Hash)
      provider = ovirtsdk4::OpenStackNetworkProvider.new(provider)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackNetworkProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackNetworkProviderReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `certificates` service.
  # 
  # @return [ExternalProviderCertificatesService] A reference to `certificates` service.
  def certificates_service
    return ExternalProviderCertificatesService.new(@connection, "#{@path}/certificates")
  end
  
  # 
  # Locates the `networks` service.
  # 
  # @return [OpenstackNetworksService] A reference to `networks` service.
  def networks_service
    return OpenstackNetworksService.new(@connection, "#{@path}/networks")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'certificates'
      return certificates_service
    end
    if path.start_with?('certificates/')
      return certificates_service.service(path[13..-1])
    end
    if path == 'networks'
      return networks_service
    end
    if path.start_with?('networks/')
      return networks_service.service(path[9..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackNetworkProviderService}:#{@path}>"
  end
  
end

class OpenstackNetworkProvidersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # The operation adds a new network provider to the system.
  # If the `type` property is not present, a default value of `NEUTRON` will be used.
  # 
  def add(provider, opts = {})
    if provider.is_a?(Hash)
      provider = ovirtsdk4::OpenStackNetworkProvider.new(provider)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackNetworkProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return OpenStackNetworkProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackNetworkProviderReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `provider` service.
  # 
  # @param id [String] The identifier of the `provider`.
  # 
  # @return [OpenstackNetworkProviderService] A reference to the `provider` service.
  # 
  def provider_service(id)
    return OpenstackNetworkProviderService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return provider_service(path)
    end
    return provider_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackNetworkProvidersService}:#{@path}>"
  end
  
end

class OpenstackNetworksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackNetworkReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network` service.
  # 
  # @param id [String] The identifier of the `network`.
  # 
  # @return [OpenstackNetworkService] A reference to the `network` service.
  # 
  def network_service(id)
    return OpenstackNetworkService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return network_service(path)
    end
    return network_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackNetworksService}:#{@path}>"
  end
  
end

class OpenstackSubnetService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackSubnetReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackSubnetService}:#{@path}>"
  end
  
end

class OpenstackSubnetsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `subnet`.
  # 
  def add(subnet, opts = {})
    if subnet.is_a?(Hash)
      subnet = ovirtsdk4::OpenStackSubnet.new(subnet)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackSubnetWriter.write_one(subnet, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return OpenStackSubnetReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackSubnetReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `subnet` service.
  # 
  # @param id [String] The identifier of the `subnet`.
  # 
  # @return [OpenstackSubnetService] A reference to the `subnet` service.
  # 
  def subnet_service(id)
    return OpenstackSubnetService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return subnet_service(path)
    end
    return subnet_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackSubnetsService}:#{@path}>"
  end
  
end

class OpenstackVolumeAuthenticationKeyService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenstackVolumeAuthenticationKeyReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(key)
    if key.is_a?(Hash)
      key = ovirtsdk4::OpenstackVolumeAuthenticationKey.new(key)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenstackVolumeAuthenticationKeyWriter.write_one(key, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenstackVolumeAuthenticationKeyReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackVolumeAuthenticationKeyService}:#{@path}>"
  end
  
end

class OpenstackVolumeAuthenticationKeysService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `key`.
  # 
  def add(key, opts = {})
    if key.is_a?(Hash)
      key = ovirtsdk4::OpenstackVolumeAuthenticationKey.new(key)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenstackVolumeAuthenticationKeyWriter.write_one(key, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return OpenstackVolumeAuthenticationKeyReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenstackVolumeAuthenticationKeyReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `key` service.
  # 
  # @param id [String] The identifier of the `key`.
  # 
  # @return [OpenstackVolumeAuthenticationKeyService] A reference to the `key` service.
  # 
  def key_service(id)
    return OpenstackVolumeAuthenticationKeyService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return key_service(path)
    end
    return key_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackVolumeAuthenticationKeysService}:#{@path}>"
  end
  
end

class OpenstackVolumeProviderService < ExternalProviderService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackVolumeProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import_certificates` method.
  # 
  def import_certificates(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/importcertificates",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `test_connectivity` method.
  # 
  def test_connectivity(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/testconnectivity",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(provider)
    if provider.is_a?(Hash)
      provider = ovirtsdk4::OpenStackVolumeProvider.new(provider)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackVolumeProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackVolumeProviderReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `authentication_keys` service.
  # 
  # @return [OpenstackVolumeAuthenticationKeysService] A reference to `authentication_keys` service.
  def authentication_keys_service
    return OpenstackVolumeAuthenticationKeysService.new(@connection, "#{@path}/authenticationkeys")
  end
  
  # 
  # Locates the `certificates` service.
  # 
  # @return [ExternalProviderCertificatesService] A reference to `certificates` service.
  def certificates_service
    return ExternalProviderCertificatesService.new(@connection, "#{@path}/certificates")
  end
  
  # 
  # Locates the `volume_types` service.
  # 
  # @return [OpenstackVolumeTypesService] A reference to `volume_types` service.
  def volume_types_service
    return OpenstackVolumeTypesService.new(@connection, "#{@path}/volumetypes")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'authenticationkeys'
      return authentication_keys_service
    end
    if path.start_with?('authenticationkeys/')
      return authentication_keys_service.service(path[19..-1])
    end
    if path == 'certificates'
      return certificates_service
    end
    if path.start_with?('certificates/')
      return certificates_service.service(path[13..-1])
    end
    if path == 'volumetypes'
      return volume_types_service
    end
    if path.start_with?('volumetypes/')
      return volume_types_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackVolumeProviderService}:#{@path}>"
  end
  
end

class OpenstackVolumeProvidersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `provider`.
  # 
  def add(provider, opts = {})
    if provider.is_a?(Hash)
      provider = ovirtsdk4::OpenStackVolumeProvider.new(provider)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      OpenStackVolumeProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return OpenStackVolumeProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackVolumeProviderReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `provider` service.
  # 
  # @param id [String] The identifier of the `provider`.
  # 
  # @return [OpenstackVolumeProviderService] A reference to the `provider` service.
  # 
  def provider_service(id)
    return OpenstackVolumeProviderService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return provider_service(path)
    end
    return provider_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackVolumeProvidersService}:#{@path}>"
  end
  
end

class OpenstackVolumeTypeService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackVolumeTypeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackVolumeTypeService}:#{@path}>"
  end
  
end

class OpenstackVolumeTypesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OpenStackVolumeTypeReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `type` service.
  # 
  # @param id [String] The identifier of the `type`.
  # 
  # @return [OpenstackVolumeTypeService] A reference to the `type` service.
  # 
  def type_service(id)
    return OpenstackVolumeTypeService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return type_service(path)
    end
    return type_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OpenstackVolumeTypesService}:#{@path}>"
  end
  
end

class OperatingSystemService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OperatingSystemInfoReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OperatingSystemService}:#{@path}>"
  end
  
end

class OperatingSystemsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return OperatingSystemInfoReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `operating_system` service.
  # 
  # @param id [String] The identifier of the `operating_system`.
  # 
  # @return [OperatingSystemService] A reference to the `operating_system` service.
  # 
  def operating_system_service(id)
    return OperatingSystemService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return operating_system_service(path)
    end
    return operating_system_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{OperatingSystemsService}:#{@path}>"
  end
  
end

class PermissionService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return PermissionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{PermissionService}:#{@path}>"
  end
  
end

class PermitService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return PermitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{PermitService}:#{@path}>"
  end
  
end

class PermitsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a permit to the set aggregated by parent role. The permit must be one retrieved from the capabilities
  # resource.
  # 
  def add(permit, opts = {})
    if permit.is_a?(Hash)
      permit = ovirtsdk4::Permit.new(permit)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      PermitWriter.write_one(permit, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return PermitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return PermitReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Sub-resource locator method, returns individual permit resource on which the remainder of the URI is dispatched.
  # 
  # @param id [String] The identifier of the `permit`.
  # 
  # @return [PermitService] A reference to the `permit` service.
  # 
  def permit_service(id)
    return PermitService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return permit_service(path)
    end
    return permit_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{PermitsService}:#{@path}>"
  end
  
end

class QosService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QosReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(qos)
    if qos.is_a?(Hash)
      qos = ovirtsdk4::Qos.new(qos)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      QosWriter.write_one(qos, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QosReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QosService}:#{@path}>"
  end
  
end

class QossService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `qos`.
  # 
  def add(qos, opts = {})
    if qos.is_a?(Hash)
      qos = ovirtsdk4::Qos.new(qos)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      QosWriter.write_one(qos, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return QosReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QosReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `qos` service.
  # 
  # @param id [String] The identifier of the `qos`.
  # 
  # @return [QosService] A reference to the `qos` service.
  # 
  def qos_service(id)
    return QosService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return qos_service(path)
    end
    return qos_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QossService}:#{@path}>"
  end
  
end

class QuotaService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(quota)
    if quota.is_a?(Hash)
      quota = ovirtsdk4::Quota.new(quota)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      QuotaWriter.write_one(quota, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `quota_cluster_limits` service.
  # 
  # @return [QuotaClusterLimitsService] A reference to `quota_cluster_limits` service.
  def quota_cluster_limits_service
    return QuotaClusterLimitsService.new(@connection, "#{@path}/quotaclusterlimits")
  end
  
  # 
  # Locates the `quota_storage_limits` service.
  # 
  # @return [QuotaStorageLimitsService] A reference to `quota_storage_limits` service.
  def quota_storage_limits_service
    return QuotaStorageLimitsService.new(@connection, "#{@path}/quotastoragelimits")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'quotaclusterlimits'
      return quota_cluster_limits_service
    end
    if path.start_with?('quotaclusterlimits/')
      return quota_cluster_limits_service.service(path[19..-1])
    end
    if path == 'quotastoragelimits'
      return quota_storage_limits_service
    end
    if path.start_with?('quotastoragelimits/')
      return quota_storage_limits_service.service(path[19..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QuotaService}:#{@path}>"
  end
  
end

class QuotaClusterLimitService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaClusterLimitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QuotaClusterLimitService}:#{@path}>"
  end
  
end

class QuotaClusterLimitsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `limit`.
  # 
  def add(limit, opts = {})
    if limit.is_a?(Hash)
      limit = ovirtsdk4::QuotaClusterLimit.new(limit)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      QuotaClusterLimitWriter.write_one(limit, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return QuotaClusterLimitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaClusterLimitReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `limit` service.
  # 
  # @param id [String] The identifier of the `limit`.
  # 
  # @return [QuotaClusterLimitService] A reference to the `limit` service.
  # 
  def limit_service(id)
    return QuotaClusterLimitService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return limit_service(path)
    end
    return limit_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QuotaClusterLimitsService}:#{@path}>"
  end
  
end

class QuotaStorageLimitService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaStorageLimitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QuotaStorageLimitService}:#{@path}>"
  end
  
end

class QuotaStorageLimitsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `limit`.
  # 
  def add(limit, opts = {})
    if limit.is_a?(Hash)
      limit = ovirtsdk4::QuotaStorageLimit.new(limit)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      QuotaStorageLimitWriter.write_one(limit, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return QuotaStorageLimitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaStorageLimitReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `limit` service.
  # 
  # @param id [String] The identifier of the `limit`.
  # 
  # @return [QuotaStorageLimitService] A reference to the `limit` service.
  # 
  def limit_service(id)
    return QuotaStorageLimitService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return limit_service(path)
    end
    return limit_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QuotaStorageLimitsService}:#{@path}>"
  end
  
end

class QuotasService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `quota`.
  # 
  def add(quota, opts = {})
    if quota.is_a?(Hash)
      quota = ovirtsdk4::Quota.new(quota)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      QuotaWriter.write_one(quota, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return QuotaReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return QuotaReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `quota` service.
  # 
  # @param id [String] The identifier of the `quota`.
  # 
  # @return [QuotaService] A reference to the `quota` service.
  # 
  def quota_service(id)
    return QuotaService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return quota_service(path)
    end
    return quota_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{QuotasService}:#{@path}>"
  end
  
end

class RoleService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return RoleReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(role)
    if role.is_a?(Hash)
      role = ovirtsdk4::Role.new(role)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      RoleWriter.write_one(role, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return RoleReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permits` service.
  # 
  # @return [PermitsService] A reference to `permits` service.
  def permits_service
    return PermitsService.new(@connection, "#{@path}/permits")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permits'
      return permits_service
    end
    if path.start_with?('permits/')
      return permits_service.service(path[8..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{RoleService}:#{@path}>"
  end
  
end

class RolesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `role`.
  # 
  def add(role, opts = {})
    if role.is_a?(Hash)
      role = ovirtsdk4::Role.new(role)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      RoleWriter.write_one(role, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return RoleReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return RoleReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Sub-resource locator method, returns individual role resource on which the remainder of the URI is dispatched.
  # 
  # @param id [String] The identifier of the `role`.
  # 
  # @return [RoleService] A reference to the `role` service.
  # 
  def role_service(id)
    return RoleService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return role_service(path)
    end
    return role_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{RolesService}:#{@path}>"
  end
  
end

class SchedulingPoliciesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `policy`.
  # 
  def add(policy, opts = {})
    if policy.is_a?(Hash)
      policy = ovirtsdk4::SchedulingPolicy.new(policy)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      SchedulingPolicyWriter.write_one(policy, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return SchedulingPolicyReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SchedulingPolicyReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `policy` service.
  # 
  # @param id [String] The identifier of the `policy`.
  # 
  # @return [SchedulingPolicyService] A reference to the `policy` service.
  # 
  def policy_service(id)
    return SchedulingPolicyService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return policy_service(path)
    end
    return policy_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SchedulingPoliciesService}:#{@path}>"
  end
  
end

class SchedulingPolicyService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SchedulingPolicyReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(policy)
    if policy.is_a?(Hash)
      policy = ovirtsdk4::SchedulingPolicy.new(policy)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      SchedulingPolicyWriter.write_one(policy, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SchedulingPolicyReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `balances` service.
  # 
  # @return [BalancesService] A reference to `balances` service.
  def balances_service
    return BalancesService.new(@connection, "#{@path}/balances")
  end
  
  # 
  # Locates the `filters` service.
  # 
  # @return [FiltersService] A reference to `filters` service.
  def filters_service
    return FiltersService.new(@connection, "#{@path}/filters")
  end
  
  # 
  # Locates the `weights` service.
  # 
  # @return [WeightsService] A reference to `weights` service.
  def weights_service
    return WeightsService.new(@connection, "#{@path}/weights")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'balances'
      return balances_service
    end
    if path.start_with?('balances/')
      return balances_service.service(path[9..-1])
    end
    if path == 'filters'
      return filters_service
    end
    if path.start_with?('filters/')
      return filters_service.service(path[8..-1])
    end
    if path == 'weights'
      return weights_service
    end
    if path.start_with?('weights/')
      return weights_service.service(path[8..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SchedulingPolicyService}:#{@path}>"
  end
  
end

class SchedulingPolicyUnitService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SchedulingPolicyUnitReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SchedulingPolicyUnitService}:#{@path}>"
  end
  
end

class SchedulingPolicyUnitsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SchedulingPolicyUnitReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `unit` service.
  # 
  # @param id [String] The identifier of the `unit`.
  # 
  # @return [SchedulingPolicyUnitService] A reference to the `unit` service.
  # 
  def unit_service(id)
    return SchedulingPolicyUnitService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return unit_service(path)
    end
    return unit_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SchedulingPolicyUnitsService}:#{@path}>"
  end
  
end

class SnapshotService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SnapshotReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `restore` method.
  # 
  def restore(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/restore",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `cdroms` service.
  # 
  # @return [SnapshotCdromsService] A reference to `cdroms` service.
  def cdroms_service
    return SnapshotCdromsService.new(@connection, "#{@path}/cdroms")
  end
  
  # 
  # Locates the `disks` service.
  # 
  # @return [SnapshotDisksService] A reference to `disks` service.
  def disks_service
    return SnapshotDisksService.new(@connection, "#{@path}/disks")
  end
  
  # 
  # Locates the `nics` service.
  # 
  # @return [SnapshotNicsService] A reference to `nics` service.
  def nics_service
    return SnapshotNicsService.new(@connection, "#{@path}/nics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'cdroms'
      return cdroms_service
    end
    if path.start_with?('cdroms/')
      return cdroms_service.service(path[7..-1])
    end
    if path == 'disks'
      return disks_service
    end
    if path.start_with?('disks/')
      return disks_service.service(path[6..-1])
    end
    if path == 'nics'
      return nics_service
    end
    if path.start_with?('nics/')
      return nics_service.service(path[5..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotService}:#{@path}>"
  end
  
end

class SnapshotCdromService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotCdromService}:#{@path}>"
  end
  
end

class SnapshotCdromsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `cdrom` service.
  # 
  # @param id [String] The identifier of the `cdrom`.
  # 
  # @return [SnapshotCdromService] A reference to the `cdrom` service.
  # 
  def cdrom_service(id)
    return SnapshotCdromService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return cdrom_service(path)
    end
    return cdrom_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotCdromsService}:#{@path}>"
  end
  
end

class SnapshotDiskService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotDiskService}:#{@path}>"
  end
  
end

class SnapshotDisksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk` service.
  # 
  # @param id [String] The identifier of the `disk`.
  # 
  # @return [SnapshotDiskService] A reference to the `disk` service.
  # 
  def disk_service(id)
    return SnapshotDiskService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return disk_service(path)
    end
    return disk_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotDisksService}:#{@path}>"
  end
  
end

class SnapshotNicService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotNicService}:#{@path}>"
  end
  
end

class SnapshotNicsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `nic` service.
  # 
  # @param id [String] The identifier of the `nic`.
  # 
  # @return [SnapshotNicService] A reference to the `nic` service.
  # 
  def nic_service(id)
    return SnapshotNicService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return nic_service(path)
    end
    return nic_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotNicsService}:#{@path}>"
  end
  
end

class SnapshotsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `snapshot`.
  # 
  def add(snapshot, opts = {})
    if snapshot.is_a?(Hash)
      snapshot = ovirtsdk4::Snapshot.new(snapshot)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      SnapshotWriter.write_one(snapshot, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return SnapshotReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SnapshotReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `snapshot` service.
  # 
  # @param id [String] The identifier of the `snapshot`.
  # 
  # @return [SnapshotService] A reference to the `snapshot` service.
  # 
  def snapshot_service(id)
    return SnapshotService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return snapshot_service(path)
    end
    return snapshot_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SnapshotsService}:#{@path}>"
  end
  
end

class SshPublicKeyService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SshPublicKeyReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(key)
    if key.is_a?(Hash)
      key = ovirtsdk4::SshPublicKey.new(key)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      SshPublicKeyWriter.write_one(key, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SshPublicKeyReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SshPublicKeyService}:#{@path}>"
  end
  
end

class SshPublicKeysService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `key`.
  # 
  def add(key, opts = {})
    if key.is_a?(Hash)
      key = ovirtsdk4::SshPublicKey.new(key)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      SshPublicKeyWriter.write_one(key, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return SshPublicKeyReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SshPublicKeyReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `key` service.
  # 
  # @param id [String] The identifier of the `key`.
  # 
  # @return [SshPublicKeyService] A reference to the `key` service.
  # 
  def key_service(id)
    return SshPublicKeyService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return key_service(path)
    end
    return key_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SshPublicKeysService}:#{@path}>"
  end
  
end

class StatisticService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:statistic]
    unless value.nil?
      query['statistic'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StatisticReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StatisticService}:#{@path}>"
  end
  
end

class StatisticsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StatisticReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `statistic` service.
  # 
  # @param id [String] The identifier of the `statistic`.
  # 
  # @return [StatisticService] A reference to the `statistic` service.
  # 
  def statistic_service(id)
    return StatisticService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return statistic_service(path)
    end
    return statistic_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StatisticsService}:#{@path}>"
  end
  
end

class StepService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `end_` method.
  # 
  def end_(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/end",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StepReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StepService}:#{@path}>"
  end
  
end

class StepsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `step`.
  # 
  def add(step, opts = {})
    if step.is_a?(Hash)
      step = ovirtsdk4::Step.new(step)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StepWriter.write_one(step, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return StepReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StepReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `step` service.
  # 
  # @param id [String] The identifier of the `step`.
  # 
  # @return [StepService] A reference to the `step` service.
  # 
  def step_service(id)
    return StepService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return step_service(path)
    end
    return step_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StepsService}:#{@path}>"
  end
  
end

class StorageService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:report_status]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['report_status'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostStorageReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageService}:#{@path}>"
  end
  
end

class StorageDomainService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `is_attached` method.
  # 
  def is_attached(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/isattached",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.is_attached
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `refresh_luns` method.
  # 
  def refresh_luns(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/refreshluns",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes the storage domain.
  # 
  # Without any special parameters, the storage domain is detached from the system and removed fro the database. The
  # storage domain can then be imported to the same or different setup, with all the data on it. If the storage isn't
  # accessible the operation will fail.
  # 
  # If the `destroy` parameter is `true` then the operation will always succeed, even if the storage isn't
  # accessible, the failure is just ignored and the storage domain is removed from the database anyway.
  # 
  # If the `format` parameter is `true` then the actual storage is formatted, and the metadata is removed from the
  # LUN or directory, so it can no longer be imported to the same or a different setup.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    value = opts[:destroy]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['destroy'] = value
    end
    value = opts[:format]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['format'] = value
    end
    value = opts[:host]
    unless value.nil?
      query['host'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(storage_domain)
    if storage_domain.is_a?(Hash)
      storage_domain = ovirtsdk4::StorageDomain.new(storage_domain)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageDomainWriter.write_one(storage_domain, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # This operation forces the update of the `OVF_STORE`
  # of this storage domain.
  # 
  # The `OVF_STORE` is a disk image that contains the meta-data
  # of virtual machines and disks that reside in the
  # storage domain. This meta-data is used in case the
  # domain is imported or exported to or from a different
  # data center or a different installation.
  # 
  # By default the `OVF_STORE` is updated periodically
  # (set by default to 60 minutes) but users might want to force an
  # update after an important change, or when the they believe the
  # `OVF_STORE` is corrupt.
  # 
  # When initiated by the user, `OVF_STORE` update will be performed whether
  # an update is needed or not.
  # 
  def update_ovf_store(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/updateovfstore",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk_profiles` service.
  # 
  # @return [AssignedDiskProfilesService] A reference to `disk_profiles` service.
  def disk_profiles_service
    return AssignedDiskProfilesService.new(@connection, "#{@path}/diskprofiles")
  end
  
  # 
  # Locates the `disk_snapshots` service.
  # 
  # @return [DiskSnapshotsService] A reference to `disk_snapshots` service.
  def disk_snapshots_service
    return DiskSnapshotsService.new(@connection, "#{@path}/disksnapshots")
  end
  
  # 
  # Locates the `disks` service.
  # 
  # @return [DisksService] A reference to `disks` service.
  def disks_service
    return DisksService.new(@connection, "#{@path}/disks")
  end
  
  # 
  # Locates the `files` service.
  # 
  # @return [FilesService] A reference to `files` service.
  def files_service
    return FilesService.new(@connection, "#{@path}/files")
  end
  
  # 
  # Locates the `images` service.
  # 
  # @return [ImagesService] A reference to `images` service.
  def images_service
    return ImagesService.new(@connection, "#{@path}/images")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `storage_connections` service.
  # 
  # @return [StorageDomainServerConnectionsService] A reference to `storage_connections` service.
  def storage_connections_service
    return StorageDomainServerConnectionsService.new(@connection, "#{@path}/storageconnections")
  end
  
  # 
  # Locates the `templates` service.
  # 
  # @return [StorageDomainTemplatesService] A reference to `templates` service.
  def templates_service
    return StorageDomainTemplatesService.new(@connection, "#{@path}/templates")
  end
  
  # 
  # Locates the `vms` service.
  # 
  # @return [StorageDomainVmsService] A reference to `vms` service.
  def vms_service
    return StorageDomainVmsService.new(@connection, "#{@path}/vms")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'diskprofiles'
      return disk_profiles_service
    end
    if path.start_with?('diskprofiles/')
      return disk_profiles_service.service(path[13..-1])
    end
    if path == 'disksnapshots'
      return disk_snapshots_service
    end
    if path.start_with?('disksnapshots/')
      return disk_snapshots_service.service(path[14..-1])
    end
    if path == 'disks'
      return disks_service
    end
    if path.start_with?('disks/')
      return disks_service.service(path[6..-1])
    end
    if path == 'files'
      return files_service
    end
    if path.start_with?('files/')
      return files_service.service(path[6..-1])
    end
    if path == 'images'
      return images_service
    end
    if path.start_with?('images/')
      return images_service.service(path[7..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'storageconnections'
      return storage_connections_service
    end
    if path.start_with?('storageconnections/')
      return storage_connections_service.service(path[19..-1])
    end
    if path == 'templates'
      return templates_service
    end
    if path.start_with?('templates/')
      return templates_service.service(path[10..-1])
    end
    if path == 'vms'
      return vms_service
    end
    if path.start_with?('vms/')
      return vms_service.service(path[4..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainService}:#{@path}>"
  end
  
end

class StorageDomainContentDiskService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainContentDiskService}:#{@path}>"
  end
  
end

class StorageDomainContentDisksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk` service.
  # 
  # @param id [String] The identifier of the `disk`.
  # 
  # @return [StorageDomainContentDiskService] A reference to the `disk` service.
  # 
  def disk_service(id)
    return StorageDomainContentDiskService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return disk_service(path)
    end
    return disk_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainContentDisksService}:#{@path}>"
  end
  
end

class StorageDomainServerConnectionService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainServerConnectionService}:#{@path}>"
  end
  
end

class StorageDomainServerConnectionsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `connection`.
  # 
  def add(connection, opts = {})
    if connection.is_a?(Hash)
      connection = ovirtsdk4::StorageConnection.new(connection)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageConnectionWriter.write_one(connection, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `connection` service.
  # 
  # @param id [String] The identifier of the `connection`.
  # 
  # @return [StorageDomainServerConnectionService] A reference to the `connection` service.
  # 
  def connection_service(id)
    return StorageDomainServerConnectionService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return connection_service(path)
    end
    return connection_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainServerConnectionsService}:#{@path}>"
  end
  
end

class StorageDomainTemplateService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TemplateReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import` method.
  # 
  def import(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/import",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `register` method.
  # 
  def register(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/register",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disks` service.
  # 
  # @return [StorageDomainContentDisksService] A reference to `disks` service.
  def disks_service
    return StorageDomainContentDisksService.new(@connection, "#{@path}/disks")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'disks'
      return disks_service
    end
    if path.start_with?('disks/')
      return disks_service.service(path[6..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainTemplateService}:#{@path}>"
  end
  
end

class StorageDomainTemplatesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TemplateReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `template` service.
  # 
  # @param id [String] The identifier of the `template`.
  # 
  # @return [StorageDomainTemplateService] A reference to the `template` service.
  # 
  def template_service(id)
    return StorageDomainTemplateService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return template_service(path)
    end
    return template_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainTemplatesService}:#{@path}>"
  end
  
end

class StorageDomainVmService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import` method.
  # 
  def import(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/import",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `register` method.
  # 
  def register(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/register",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disks` service.
  # 
  # @return [StorageDomainContentDisksService] A reference to `disks` service.
  def disks_service
    return StorageDomainContentDisksService.new(@connection, "#{@path}/disks")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'disks'
      return disks_service
    end
    if path.start_with?('disks/')
      return disks_service.service(path[6..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainVmService}:#{@path}>"
  end
  
end

class StorageDomainVmsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `vm` service.
  # 
  # @param id [String] The identifier of the `vm`.
  # 
  # @return [StorageDomainVmService] A reference to the `vm` service.
  # 
  def vm_service(id)
    return StorageDomainVmService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return vm_service(path)
    end
    return vm_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainVmsService}:#{@path}>"
  end
  
end

class StorageDomainsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `storage_domain`.
  # 
  def add(storage_domain, opts = {})
    if storage_domain.is_a?(Hash)
      storage_domain = ovirtsdk4::StorageDomain.new(storage_domain)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageDomainWriter.write_one(storage_domain, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageDomainReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `storage_domain` service.
  # 
  # @param id [String] The identifier of the `storage_domain`.
  # 
  # @return [StorageDomainService] A reference to the `storage_domain` service.
  # 
  def storage_domain_service(id)
    return StorageDomainService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return storage_domain_service(path)
    end
    return storage_domain_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageDomainsService}:#{@path}>"
  end
  
end

class StorageServerConnectionService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(connection)
    if connection.is_a?(Hash)
      connection = ovirtsdk4::StorageConnection.new(connection)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageConnectionWriter.write_one(connection, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageServerConnectionService}:#{@path}>"
  end
  
end

class StorageServerConnectionExtensionService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionExtensionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(extension)
    if extension.is_a?(Hash)
      extension = ovirtsdk4::StorageConnectionExtension.new(extension)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageConnectionExtensionWriter.write_one(extension, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionExtensionReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageServerConnectionExtensionService}:#{@path}>"
  end
  
end

class StorageServerConnectionExtensionsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `extension`.
  # 
  def add(extension, opts = {})
    if extension.is_a?(Hash)
      extension = ovirtsdk4::StorageConnectionExtension.new(extension)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageConnectionExtensionWriter.write_one(extension, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionExtensionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionExtensionReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `storage_connection_extension` service.
  # 
  # @param id [String] The identifier of the `storage_connection_extension`.
  # 
  # @return [StorageServerConnectionExtensionService] A reference to the `storage_connection_extension` service.
  # 
  def storage_connection_extension_service(id)
    return StorageServerConnectionExtensionService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return storage_connection_extension_service(path)
    end
    return storage_connection_extension_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageServerConnectionExtensionsService}:#{@path}>"
  end
  
end

class StorageServerConnectionsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `connection`.
  # 
  def add(connection, opts = {})
    if connection.is_a?(Hash)
      connection = ovirtsdk4::StorageConnection.new(connection)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      StorageConnectionWriter.write_one(connection, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return StorageConnectionReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `storage_connection` service.
  # 
  # @param id [String] The identifier of the `storage_connection`.
  # 
  # @return [StorageServerConnectionService] A reference to the `storage_connection` service.
  # 
  def storage_connection_service(id)
    return StorageServerConnectionService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return storage_connection_service(path)
    end
    return storage_connection_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{StorageServerConnectionsService}:#{@path}>"
  end
  
end

class SystemService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns basic information describing the API, like the product name, the version number and a summary of the
  # number of relevant objects.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ApiReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `reload_configurations` method.
  # 
  def reload_configurations(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/reloadconfigurations",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # List all known affinity labels.
  # 
  # @return [AffinityLabelsService] A reference to `affinity_labels` service.
  def affinity_labels_service
    return AffinityLabelsService.new(@connection, "#{@path}/affinitylabels")
  end
  
  # 
  # Locates the `bookmarks` service.
  # 
  # @return [BookmarksService] A reference to `bookmarks` service.
  def bookmarks_service
    return BookmarksService.new(@connection, "#{@path}/bookmarks")
  end
  
  # 
  # Reference to the service that provides information about the cluster levels supported by the system.
  # 
  # @return [ClusterLevelsService] A reference to `cluster_levels` service.
  def cluster_levels_service
    return ClusterLevelsService.new(@connection, "#{@path}/clusterlevels")
  end
  
  # 
  # Locates the `clusters` service.
  # 
  # @return [ClustersService] A reference to `clusters` service.
  def clusters_service
    return ClustersService.new(@connection, "#{@path}/clusters")
  end
  
  # 
  # Locates the `cpu_profiles` service.
  # 
  # @return [CpuProfilesService] A reference to `cpu_profiles` service.
  def cpu_profiles_service
    return CpuProfilesService.new(@connection, "#{@path}/cpuprofiles")
  end
  
  # 
  # Locates the `data_centers` service.
  # 
  # @return [DataCentersService] A reference to `data_centers` service.
  def data_centers_service
    return DataCentersService.new(@connection, "#{@path}/datacenters")
  end
  
  # 
  # Locates the `disk_profiles` service.
  # 
  # @return [DiskProfilesService] A reference to `disk_profiles` service.
  def disk_profiles_service
    return DiskProfilesService.new(@connection, "#{@path}/diskprofiles")
  end
  
  # 
  # Locates the `disks` service.
  # 
  # @return [DisksService] A reference to `disks` service.
  def disks_service
    return DisksService.new(@connection, "#{@path}/disks")
  end
  
  # 
  # Locates the `domains` service.
  # 
  # @return [DomainsService] A reference to `domains` service.
  def domains_service
    return DomainsService.new(@connection, "#{@path}/domains")
  end
  
  # 
  # Locates the `events` service.
  # 
  # @return [EventsService] A reference to `events` service.
  def events_service
    return EventsService.new(@connection, "#{@path}/events")
  end
  
  # 
  # Locates the `external_host_providers` service.
  # 
  # @return [ExternalHostProvidersService] A reference to `external_host_providers` service.
  def external_host_providers_service
    return ExternalHostProvidersService.new(@connection, "#{@path}/externalhostproviders")
  end
  
  # 
  # Locates the `groups` service.
  # 
  # @return [GroupsService] A reference to `groups` service.
  def groups_service
    return GroupsService.new(@connection, "#{@path}/groups")
  end
  
  # 
  # Locates the `hosts` service.
  # 
  # @return [HostsService] A reference to `hosts` service.
  def hosts_service
    return HostsService.new(@connection, "#{@path}/hosts")
  end
  
  # 
  # Locates the `icons` service.
  # 
  # @return [IconsService] A reference to `icons` service.
  def icons_service
    return IconsService.new(@connection, "#{@path}/icons")
  end
  
  # 
  # Locates the `instance_types` service.
  # 
  # @return [InstanceTypesService] A reference to `instance_types` service.
  def instance_types_service
    return InstanceTypesService.new(@connection, "#{@path}/instancetypes")
  end
  
  # 
  # Locates the `jobs` service.
  # 
  # @return [JobsService] A reference to `jobs` service.
  def jobs_service
    return JobsService.new(@connection, "#{@path}/jobs")
  end
  
  # 
  # Locates the `katello_errata` service.
  # 
  # @return [EngineKatelloErrataService] A reference to `katello_errata` service.
  def katello_errata_service
    return EngineKatelloErrataService.new(@connection, "#{@path}/katelloerrata")
  end
  
  # 
  # Locates the `mac_pools` service.
  # 
  # @return [MacPoolsService] A reference to `mac_pools` service.
  def mac_pools_service
    return MacPoolsService.new(@connection, "#{@path}/macpools")
  end
  
  # 
  # Network filters will enhance the admin ability to manage the network packets traffic from/to the participated
  # VMs.
  # 
  # @return [NetworkFiltersService] A reference to `network_filters` service.
  def network_filters_service
    return NetworkFiltersService.new(@connection, "#{@path}/networkfilters")
  end
  
  # 
  # Locates the `networks` service.
  # 
  # @return [NetworksService] A reference to `networks` service.
  def networks_service
    return NetworksService.new(@connection, "#{@path}/networks")
  end
  
  # 
  # Locates the `openstack_image_providers` service.
  # 
  # @return [OpenstackImageProvidersService] A reference to `openstack_image_providers` service.
  def openstack_image_providers_service
    return OpenstackImageProvidersService.new(@connection, "#{@path}/openstackimageproviders")
  end
  
  # 
  # Locates the `openstack_network_providers` service.
  # 
  # @return [OpenstackNetworkProvidersService] A reference to `openstack_network_providers` service.
  def openstack_network_providers_service
    return OpenstackNetworkProvidersService.new(@connection, "#{@path}/openstacknetworkproviders")
  end
  
  # 
  # Locates the `openstack_volume_providers` service.
  # 
  # @return [OpenstackVolumeProvidersService] A reference to `openstack_volume_providers` service.
  def openstack_volume_providers_service
    return OpenstackVolumeProvidersService.new(@connection, "#{@path}/openstackvolumeproviders")
  end
  
  # 
  # Locates the `operating_systems` service.
  # 
  # @return [OperatingSystemsService] A reference to `operating_systems` service.
  def operating_systems_service
    return OperatingSystemsService.new(@connection, "#{@path}/operatingsystems")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [SystemPermissionsService] A reference to `permissions` service.
  def permissions_service
    return SystemPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `roles` service.
  # 
  # @return [RolesService] A reference to `roles` service.
  def roles_service
    return RolesService.new(@connection, "#{@path}/roles")
  end
  
  # 
  # Locates the `scheduling_policies` service.
  # 
  # @return [SchedulingPoliciesService] A reference to `scheduling_policies` service.
  def scheduling_policies_service
    return SchedulingPoliciesService.new(@connection, "#{@path}/schedulingpolicies")
  end
  
  # 
  # Locates the `scheduling_policy_units` service.
  # 
  # @return [SchedulingPolicyUnitsService] A reference to `scheduling_policy_units` service.
  def scheduling_policy_units_service
    return SchedulingPolicyUnitsService.new(@connection, "#{@path}/schedulingpolicyunits")
  end
  
  # 
  # Locates the `storage_connections` service.
  # 
  # @return [StorageServerConnectionsService] A reference to `storage_connections` service.
  def storage_connections_service
    return StorageServerConnectionsService.new(@connection, "#{@path}/storageconnections")
  end
  
  # 
  # Locates the `storage_domains` service.
  # 
  # @return [StorageDomainsService] A reference to `storage_domains` service.
  def storage_domains_service
    return StorageDomainsService.new(@connection, "#{@path}/storagedomains")
  end
  
  # 
  # Locates the `tags` service.
  # 
  # @return [TagsService] A reference to `tags` service.
  def tags_service
    return TagsService.new(@connection, "#{@path}/tags")
  end
  
  # 
  # Locates the `templates` service.
  # 
  # @return [TemplatesService] A reference to `templates` service.
  def templates_service
    return TemplatesService.new(@connection, "#{@path}/templates")
  end
  
  # 
  # Locates the `users` service.
  # 
  # @return [UsersService] A reference to `users` service.
  def users_service
    return UsersService.new(@connection, "#{@path}/users")
  end
  
  # 
  # Locates the `vm_pools` service.
  # 
  # @return [VmPoolsService] A reference to `vm_pools` service.
  def vm_pools_service
    return VmPoolsService.new(@connection, "#{@path}/vmpools")
  end
  
  # 
  # Locates the `vms` service.
  # 
  # @return [VmsService] A reference to `vms` service.
  def vms_service
    return VmsService.new(@connection, "#{@path}/vms")
  end
  
  # 
  # Locates the `vnic_profiles` service.
  # 
  # @return [VnicProfilesService] A reference to `vnic_profiles` service.
  def vnic_profiles_service
    return VnicProfilesService.new(@connection, "#{@path}/vnicprofiles")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'affinitylabels'
      return affinity_labels_service
    end
    if path.start_with?('affinitylabels/')
      return affinity_labels_service.service(path[15..-1])
    end
    if path == 'bookmarks'
      return bookmarks_service
    end
    if path.start_with?('bookmarks/')
      return bookmarks_service.service(path[10..-1])
    end
    if path == 'clusterlevels'
      return cluster_levels_service
    end
    if path.start_with?('clusterlevels/')
      return cluster_levels_service.service(path[14..-1])
    end
    if path == 'clusters'
      return clusters_service
    end
    if path.start_with?('clusters/')
      return clusters_service.service(path[9..-1])
    end
    if path == 'cpuprofiles'
      return cpu_profiles_service
    end
    if path.start_with?('cpuprofiles/')
      return cpu_profiles_service.service(path[12..-1])
    end
    if path == 'datacenters'
      return data_centers_service
    end
    if path.start_with?('datacenters/')
      return data_centers_service.service(path[12..-1])
    end
    if path == 'diskprofiles'
      return disk_profiles_service
    end
    if path.start_with?('diskprofiles/')
      return disk_profiles_service.service(path[13..-1])
    end
    if path == 'disks'
      return disks_service
    end
    if path.start_with?('disks/')
      return disks_service.service(path[6..-1])
    end
    if path == 'domains'
      return domains_service
    end
    if path.start_with?('domains/')
      return domains_service.service(path[8..-1])
    end
    if path == 'events'
      return events_service
    end
    if path.start_with?('events/')
      return events_service.service(path[7..-1])
    end
    if path == 'externalhostproviders'
      return external_host_providers_service
    end
    if path.start_with?('externalhostproviders/')
      return external_host_providers_service.service(path[22..-1])
    end
    if path == 'groups'
      return groups_service
    end
    if path.start_with?('groups/')
      return groups_service.service(path[7..-1])
    end
    if path == 'hosts'
      return hosts_service
    end
    if path.start_with?('hosts/')
      return hosts_service.service(path[6..-1])
    end
    if path == 'icons'
      return icons_service
    end
    if path.start_with?('icons/')
      return icons_service.service(path[6..-1])
    end
    if path == 'instancetypes'
      return instance_types_service
    end
    if path.start_with?('instancetypes/')
      return instance_types_service.service(path[14..-1])
    end
    if path == 'jobs'
      return jobs_service
    end
    if path.start_with?('jobs/')
      return jobs_service.service(path[5..-1])
    end
    if path == 'katelloerrata'
      return katello_errata_service
    end
    if path.start_with?('katelloerrata/')
      return katello_errata_service.service(path[14..-1])
    end
    if path == 'macpools'
      return mac_pools_service
    end
    if path.start_with?('macpools/')
      return mac_pools_service.service(path[9..-1])
    end
    if path == 'networkfilters'
      return network_filters_service
    end
    if path.start_with?('networkfilters/')
      return network_filters_service.service(path[15..-1])
    end
    if path == 'networks'
      return networks_service
    end
    if path.start_with?('networks/')
      return networks_service.service(path[9..-1])
    end
    if path == 'openstackimageproviders'
      return openstack_image_providers_service
    end
    if path.start_with?('openstackimageproviders/')
      return openstack_image_providers_service.service(path[24..-1])
    end
    if path == 'openstacknetworkproviders'
      return openstack_network_providers_service
    end
    if path.start_with?('openstacknetworkproviders/')
      return openstack_network_providers_service.service(path[26..-1])
    end
    if path == 'openstackvolumeproviders'
      return openstack_volume_providers_service
    end
    if path.start_with?('openstackvolumeproviders/')
      return openstack_volume_providers_service.service(path[25..-1])
    end
    if path == 'operatingsystems'
      return operating_systems_service
    end
    if path.start_with?('operatingsystems/')
      return operating_systems_service.service(path[17..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'roles'
      return roles_service
    end
    if path.start_with?('roles/')
      return roles_service.service(path[6..-1])
    end
    if path == 'schedulingpolicies'
      return scheduling_policies_service
    end
    if path.start_with?('schedulingpolicies/')
      return scheduling_policies_service.service(path[19..-1])
    end
    if path == 'schedulingpolicyunits'
      return scheduling_policy_units_service
    end
    if path.start_with?('schedulingpolicyunits/')
      return scheduling_policy_units_service.service(path[22..-1])
    end
    if path == 'storageconnections'
      return storage_connections_service
    end
    if path.start_with?('storageconnections/')
      return storage_connections_service.service(path[19..-1])
    end
    if path == 'storagedomains'
      return storage_domains_service
    end
    if path.start_with?('storagedomains/')
      return storage_domains_service.service(path[15..-1])
    end
    if path == 'tags'
      return tags_service
    end
    if path.start_with?('tags/')
      return tags_service.service(path[5..-1])
    end
    if path == 'templates'
      return templates_service
    end
    if path.start_with?('templates/')
      return templates_service.service(path[10..-1])
    end
    if path == 'users'
      return users_service
    end
    if path.start_with?('users/')
      return users_service.service(path[6..-1])
    end
    if path == 'vmpools'
      return vm_pools_service
    end
    if path.start_with?('vmpools/')
      return vm_pools_service.service(path[8..-1])
    end
    if path == 'vms'
      return vms_service
    end
    if path.start_with?('vms/')
      return vms_service.service(path[4..-1])
    end
    if path == 'vnicprofiles'
      return vnic_profiles_service
    end
    if path.start_with?('vnicprofiles/')
      return vnic_profiles_service.service(path[13..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SystemService}:#{@path}>"
  end
  
end

class SystemPermissionsService < AssignedPermissionsService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `permission`.
  # 
  def add(permission, opts = {})
    if permission.is_a?(Hash)
      permission = ovirtsdk4::Permission.new(permission)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      PermissionWriter.write_one(permission, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return PermissionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return PermissionReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Sub-resource locator method, returns individual permission resource on which the remainder of the URI is
  # dispatched.
  # 
  # @param id [String] The identifier of the `permission`.
  # 
  # @return [PermissionService] A reference to the `permission` service.
  # 
  def permission_service(id)
    return PermissionService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return permission_service(path)
    end
    return permission_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{SystemPermissionsService}:#{@path}>"
  end
  
end

class TagService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(tag)
    if tag.is_a?(Hash)
      tag = ovirtsdk4::Tag.new(tag)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      TagWriter.write_one(tag, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TagService}:#{@path}>"
  end
  
end

class TagsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `tag`.
  # 
  def add(tag, opts = {})
    if tag.is_a?(Hash)
      tag = ovirtsdk4::Tag.new(tag)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      TagWriter.write_one(tag, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TagReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `tag` service.
  # 
  # @param id [String] The identifier of the `tag`.
  # 
  # @return [TagService] A reference to the `tag` service.
  # 
  def tag_service(id)
    return TagService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return tag_service(path)
    end
    return tag_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TagsService}:#{@path}>"
  end
  
end

class TemplateService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `export` method.
  # 
  def export(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/export",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TemplateReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(template)
    if template.is_a?(Hash)
      template = ovirtsdk4::Template.new(template)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      TemplateWriter.write_one(template, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TemplateReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `cdroms` service.
  # 
  # @return [TemplateCdromsService] A reference to `cdroms` service.
  def cdroms_service
    return TemplateCdromsService.new(@connection, "#{@path}/cdroms")
  end
  
  # 
  # Reference to the service that manages a specific
  # disk attachment of the template.
  # 
  # @return [TemplateDiskAttachmentsService] A reference to `disk_attachments` service.
  def disk_attachments_service
    return TemplateDiskAttachmentsService.new(@connection, "#{@path}/diskattachments")
  end
  
  # 
  # Locates the `graphics_consoles` service.
  # 
  # @return [GraphicsConsolesService] A reference to `graphics_consoles` service.
  def graphics_consoles_service
    return GraphicsConsolesService.new(@connection, "#{@path}/graphicsconsoles")
  end
  
  # 
  # Locates the `nics` service.
  # 
  # @return [TemplateNicsService] A reference to `nics` service.
  def nics_service
    return TemplateNicsService.new(@connection, "#{@path}/nics")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `tags` service.
  # 
  # @return [AssignedTagsService] A reference to `tags` service.
  def tags_service
    return AssignedTagsService.new(@connection, "#{@path}/tags")
  end
  
  # 
  # Locates the `watchdogs` service.
  # 
  # @return [TemplateWatchdogsService] A reference to `watchdogs` service.
  def watchdogs_service
    return TemplateWatchdogsService.new(@connection, "#{@path}/watchdogs")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'cdroms'
      return cdroms_service
    end
    if path.start_with?('cdroms/')
      return cdroms_service.service(path[7..-1])
    end
    if path == 'diskattachments'
      return disk_attachments_service
    end
    if path.start_with?('diskattachments/')
      return disk_attachments_service.service(path[16..-1])
    end
    if path == 'graphicsconsoles'
      return graphics_consoles_service
    end
    if path.start_with?('graphicsconsoles/')
      return graphics_consoles_service.service(path[17..-1])
    end
    if path == 'nics'
      return nics_service
    end
    if path.start_with?('nics/')
      return nics_service.service(path[5..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'tags'
      return tags_service
    end
    if path.start_with?('tags/')
      return tags_service.service(path[5..-1])
    end
    if path == 'watchdogs'
      return watchdogs_service
    end
    if path.start_with?('watchdogs/')
      return watchdogs_service.service(path[10..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateService}:#{@path}>"
  end
  
end

class TemplateCdromService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateCdromService}:#{@path}>"
  end
  
end

class TemplateCdromsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `cdrom` service.
  # 
  # @param id [String] The identifier of the `cdrom`.
  # 
  # @return [TemplateCdromService] A reference to the `cdrom` service.
  # 
  def cdrom_service(id)
    return TemplateCdromService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return cdrom_service(path)
    end
    return cdrom_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateCdromsService}:#{@path}>"
  end
  
end

class TemplateDiskService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `copy` method.
  # 
  def copy(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/copy",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `export` method.
  # 
  def export(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/export",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateDiskService}:#{@path}>"
  end
  
end

class TemplateDiskAttachmentService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the details of the attachment.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskAttachmentReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes the disk from the template. The disk will only be removed if there are other existing copies of the
  # disk on other storage domains.
  # 
  # A storage domain has to be specified to determine which of the copies should be removed (template disks can
  # have copies on multiple storage domains).
  # 
  # [source]
  # ----
  # DELETE /ovirt-engine/api/templates/{template:id}/diskattachments/{attachment:id}?storage_domain=072fbaa1-08f3-4a40-9f34-a5ca22dd1d74
  # ----
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:force]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['force'] = value
    end
    value = opts[:storage_domain]
    unless value.nil?
      query['storage_domain'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateDiskAttachmentService}:#{@path}>"
  end
  
end

class TemplateDiskAttachmentsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # List the disks that are attached to the template.
  def list(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskAttachmentReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Reference to the service that manages a specific attachment.
  # 
  # @param id [String] The identifier of the `attachment`.
  # 
  # @return [TemplateDiskAttachmentService] A reference to the `attachment` service.
  # 
  def attachment_service(id)
    return TemplateDiskAttachmentService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return attachment_service(path)
    end
    return attachment_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateDiskAttachmentsService}:#{@path}>"
  end
  
end

class TemplateDisksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk` service.
  # 
  # @param id [String] The identifier of the `disk`.
  # 
  # @return [TemplateDiskService] A reference to the `disk` service.
  # 
  def disk_service(id)
    return TemplateDiskService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return disk_service(path)
    end
    return disk_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateDisksService}:#{@path}>"
  end
  
end

class TemplateNicService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(nic)
    if nic.is_a?(Hash)
      nic = ovirtsdk4::Nic.new(nic)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NicWriter.write_one(nic, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateNicService}:#{@path}>"
  end
  
end

class TemplateNicsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `nic`.
  # 
  def add(nic, opts = {})
    if nic.is_a?(Hash)
      nic = ovirtsdk4::Nic.new(nic)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NicWriter.write_one(nic, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `nic` service.
  # 
  # @param id [String] The identifier of the `nic`.
  # 
  # @return [TemplateNicService] A reference to the `nic` service.
  # 
  def nic_service(id)
    return TemplateNicService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return nic_service(path)
    end
    return nic_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateNicsService}:#{@path}>"
  end
  
end

class TemplateWatchdogService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(watchdog)
    if watchdog.is_a?(Hash)
      watchdog = ovirtsdk4::Watchdog.new(watchdog)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WatchdogWriter.write_one(watchdog, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateWatchdogService}:#{@path}>"
  end
  
end

class TemplateWatchdogsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `watchdog`.
  # 
  def add(watchdog, opts = {})
    if watchdog.is_a?(Hash)
      watchdog = ovirtsdk4::Watchdog.new(watchdog)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WatchdogWriter.write_one(watchdog, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `watchdog` service.
  # 
  # @param id [String] The identifier of the `watchdog`.
  # 
  # @return [TemplateWatchdogService] A reference to the `watchdog` service.
  # 
  def watchdog_service(id)
    return TemplateWatchdogService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return watchdog_service(path)
    end
    return watchdog_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplateWatchdogsService}:#{@path}>"
  end
  
end

class TemplatesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `template`.
  # 
  def add(template, opts = {})
    if template.is_a?(Hash)
      template = ovirtsdk4::Template.new(template)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      TemplateWriter.write_one(template, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return TemplateReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return TemplateReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `template` service.
  # 
  # @param id [String] The identifier of the `template`.
  # 
  # @return [TemplateService] A reference to the `template` service.
  # 
  def template_service(id)
    return TemplateService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return template_service(path)
    end
    return template_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{TemplatesService}:#{@path}>"
  end
  
end

class UnmanagedNetworkService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return UnmanagedNetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{UnmanagedNetworkService}:#{@path}>"
  end
  
end

class UnmanagedNetworksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return UnmanagedNetworkReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `unmanaged_network` service.
  # 
  # @param id [String] The identifier of the `unmanaged_network`.
  # 
  # @return [UnmanagedNetworkService] A reference to the `unmanaged_network` service.
  # 
  def unmanaged_network_service(id)
    return UnmanagedNetworkService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return unmanaged_network_service(path)
    end
    return unmanaged_network_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{UnmanagedNetworksService}:#{@path}>"
  end
  
end

class UserService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return UserReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `roles` service.
  # 
  # @return [AssignedRolesService] A reference to `roles` service.
  def roles_service
    return AssignedRolesService.new(@connection, "#{@path}/roles")
  end
  
  # 
  # Locates the `ssh_public_keys` service.
  # 
  # @return [SshPublicKeysService] A reference to `ssh_public_keys` service.
  def ssh_public_keys_service
    return SshPublicKeysService.new(@connection, "#{@path}/sshpublickeys")
  end
  
  # 
  # Locates the `tags` service.
  # 
  # @return [AssignedTagsService] A reference to `tags` service.
  def tags_service
    return AssignedTagsService.new(@connection, "#{@path}/tags")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'roles'
      return roles_service
    end
    if path.start_with?('roles/')
      return roles_service.service(path[6..-1])
    end
    if path == 'sshpublickeys'
      return ssh_public_keys_service
    end
    if path.start_with?('sshpublickeys/')
      return ssh_public_keys_service.service(path[14..-1])
    end
    if path == 'tags'
      return tags_service
    end
    if path.start_with?('tags/')
      return tags_service.service(path[5..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{UserService}:#{@path}>"
  end
  
end

class UsersService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `user`.
  # 
  def add(user, opts = {})
    if user.is_a?(Hash)
      user = ovirtsdk4::User.new(user)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      UserWriter.write_one(user, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return UserReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return UserReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `user` service.
  # 
  # @param id [String] The identifier of the `user`.
  # 
  # @return [UserService] A reference to the `user` service.
  # 
  def user_service(id)
    return UserService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return user_service(path)
    end
    return user_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{UsersService}:#{@path}>"
  end
  
end

class VirtualFunctionAllowedNetworkService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VirtualFunctionAllowedNetworkService}:#{@path}>"
  end
  
end

class VirtualFunctionAllowedNetworksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `network`.
  # 
  def add(network, opts = {})
    if network.is_a?(Hash)
      network = ovirtsdk4::Network.new(network)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NetworkWriter.write_one(network, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NetworkReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network` service.
  # 
  # @param id [String] The identifier of the `network`.
  # 
  # @return [VirtualFunctionAllowedNetworkService] A reference to the `network` service.
  # 
  def network_service(id)
    return VirtualFunctionAllowedNetworkService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return network_service(path)
    end
    return network_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VirtualFunctionAllowedNetworksService}:#{@path}>"
  end
  
end

class VmService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `cancel_migration` method.
  # 
  def cancel_migration(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/cancelmigration",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `clone` method.
  # 
  def clone(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/clone",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `commit_snapshot` method.
  # 
  def commit_snapshot(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/commitsnapshot",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `detach` method.
  # 
  def detach(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/detach",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `export` method.
  # 
  def export(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/export",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `freeze_filesystems` method.
  # 
  def freeze_filesystems(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/freezefilesystems",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Retrieves the description of the virtual machine.
  # 
  # Note that some elements of the description of the virtual machine won't be returned unless the `All-Content`
  # header is present in the request and has the value `true`. The elements that aren't currently returned are
  # the following:
  # 
  # - `console`
  # - `initialization.configuration.data` - The OVF document describing the virtual machine.
  # - `rng_source`
  # - `soundcard`
  # - `virtio_scsi`
  # 
  # With the Python SDK the `All-Content` header can be set using the `all_content` parameter of the `get`
  # method:
  # 
  # [source,python]
  # ----
  # api.vms.get(name="myvm", all_content=True)
  # ----
  # 
  # Note that the reason for not including these elements is performance: they are seldom used and they require
  # additional queries in the server. So try to use the `All-Content` header only when it is really needed.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:next_run]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['next_run'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `logon` method.
  # 
  def logon(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/logon",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `maintenance` method.
  # 
  def maintenance(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/maintenance",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `migrate` method.
  # 
  def migrate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/migrate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `preview_snapshot` method.
  # 
  def preview_snapshot(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/previewsnapshot",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `reboot` method.
  # 
  def reboot(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/reboot",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `reorder_mac_addresses` method.
  # 
  def reorder_mac_addresses(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/reordermacaddresses",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `shutdown` method.
  # 
  def shutdown(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/shutdown",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `start` method.
  # 
  def start(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/start",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `stop` method.
  # 
  def stop(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/stop",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `suspend` method.
  # 
  def suspend(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/suspend",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `thaw_filesystems` method.
  # 
  def thaw_filesystems(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/thawfilesystems",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `ticket` method.
  # 
  def ticket(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/ticket",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.ticket
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `undo_snapshot` method.
  # 
  def undo_snapshot(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/undosnapshot",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(vm)
    if vm.is_a?(Hash)
      vm = ovirtsdk4::Vm.new(vm)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VmWriter.write_one(vm, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # List of scheduling labels assigned to this VM.
  # 
  # @return [AssignedAffinityLabelsService] A reference to `affinity_labels` service.
  def affinity_labels_service
    return AssignedAffinityLabelsService.new(@connection, "#{@path}/affinitylabels")
  end
  
  # 
  # Locates the `applications` service.
  # 
  # @return [VmApplicationsService] A reference to `applications` service.
  def applications_service
    return VmApplicationsService.new(@connection, "#{@path}/applications")
  end
  
  # 
  # Locates the `cdroms` service.
  # 
  # @return [VmCdromsService] A reference to `cdroms` service.
  def cdroms_service
    return VmCdromsService.new(@connection, "#{@path}/cdroms")
  end
  
  # 
  # List of disks attached to this virtual machine.
  # 
  # @return [DiskAttachmentsService] A reference to `disk_attachments` service.
  def disk_attachments_service
    return DiskAttachmentsService.new(@connection, "#{@path}/diskattachments")
  end
  
  # 
  # Locates the `graphics_consoles` service.
  # 
  # @return [GraphicsConsolesService] A reference to `graphics_consoles` service.
  def graphics_consoles_service
    return GraphicsConsolesService.new(@connection, "#{@path}/graphicsconsoles")
  end
  
  # 
  # Locates the `host_devices` service.
  # 
  # @return [VmHostDevicesService] A reference to `host_devices` service.
  def host_devices_service
    return VmHostDevicesService.new(@connection, "#{@path}/hostdevices")
  end
  
  # 
  # Locates the `katello_errata` service.
  # 
  # @return [KatelloErrataService] A reference to `katello_errata` service.
  def katello_errata_service
    return KatelloErrataService.new(@connection, "#{@path}/katelloerrata")
  end
  
  # 
  # Locates the `nics` service.
  # 
  # @return [VmNicsService] A reference to `nics` service.
  def nics_service
    return VmNicsService.new(@connection, "#{@path}/nics")
  end
  
  # 
  # Locates the `numa_nodes` service.
  # 
  # @return [VmNumaNodesService] A reference to `numa_nodes` service.
  def numa_nodes_service
    return VmNumaNodesService.new(@connection, "#{@path}/numanodes")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `reported_devices` service.
  # 
  # @return [VmReportedDevicesService] A reference to `reported_devices` service.
  def reported_devices_service
    return VmReportedDevicesService.new(@connection, "#{@path}/reporteddevices")
  end
  
  # 
  # Locates the `sessions` service.
  # 
  # @return [VmSessionsService] A reference to `sessions` service.
  def sessions_service
    return VmSessionsService.new(@connection, "#{@path}/sessions")
  end
  
  # 
  # Locates the `snapshots` service.
  # 
  # @return [SnapshotsService] A reference to `snapshots` service.
  def snapshots_service
    return SnapshotsService.new(@connection, "#{@path}/snapshots")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the `tags` service.
  # 
  # @return [AssignedTagsService] A reference to `tags` service.
  def tags_service
    return AssignedTagsService.new(@connection, "#{@path}/tags")
  end
  
  # 
  # Locates the `watchdogs` service.
  # 
  # @return [VmWatchdogsService] A reference to `watchdogs` service.
  def watchdogs_service
    return VmWatchdogsService.new(@connection, "#{@path}/watchdogs")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'affinitylabels'
      return affinity_labels_service
    end
    if path.start_with?('affinitylabels/')
      return affinity_labels_service.service(path[15..-1])
    end
    if path == 'applications'
      return applications_service
    end
    if path.start_with?('applications/')
      return applications_service.service(path[13..-1])
    end
    if path == 'cdroms'
      return cdroms_service
    end
    if path.start_with?('cdroms/')
      return cdroms_service.service(path[7..-1])
    end
    if path == 'diskattachments'
      return disk_attachments_service
    end
    if path.start_with?('diskattachments/')
      return disk_attachments_service.service(path[16..-1])
    end
    if path == 'graphicsconsoles'
      return graphics_consoles_service
    end
    if path.start_with?('graphicsconsoles/')
      return graphics_consoles_service.service(path[17..-1])
    end
    if path == 'hostdevices'
      return host_devices_service
    end
    if path.start_with?('hostdevices/')
      return host_devices_service.service(path[12..-1])
    end
    if path == 'katelloerrata'
      return katello_errata_service
    end
    if path.start_with?('katelloerrata/')
      return katello_errata_service.service(path[14..-1])
    end
    if path == 'nics'
      return nics_service
    end
    if path.start_with?('nics/')
      return nics_service.service(path[5..-1])
    end
    if path == 'numanodes'
      return numa_nodes_service
    end
    if path.start_with?('numanodes/')
      return numa_nodes_service.service(path[10..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'reporteddevices'
      return reported_devices_service
    end
    if path.start_with?('reporteddevices/')
      return reported_devices_service.service(path[16..-1])
    end
    if path == 'sessions'
      return sessions_service
    end
    if path.start_with?('sessions/')
      return sessions_service.service(path[9..-1])
    end
    if path == 'snapshots'
      return snapshots_service
    end
    if path.start_with?('snapshots/')
      return snapshots_service.service(path[10..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    if path == 'tags'
      return tags_service
    end
    if path.start_with?('tags/')
      return tags_service.service(path[5..-1])
    end
    if path == 'watchdogs'
      return watchdogs_service
    end
    if path.start_with?('watchdogs/')
      return watchdogs_service.service(path[10..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmService}:#{@path}>"
  end
  
end

class VmApplicationService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ApplicationReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmApplicationService}:#{@path}>"
  end
  
end

class VmApplicationsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ApplicationReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `application` service.
  # 
  # @param id [String] The identifier of the `application`.
  # 
  # @return [VmApplicationService] A reference to the `application` service.
  # 
  def application_service(id)
    return VmApplicationService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return application_service(path)
    end
    return application_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmApplicationsService}:#{@path}>"
  end
  
end

class VmCdromService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(cdorm)
    if cdorm.is_a?(Hash)
      cdorm = ovirtsdk4::Cdrom.new(cdorm)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      CdromWriter.write_one(cdorm, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmCdromService}:#{@path}>"
  end
  
end

class VmCdromsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `cdrom`.
  # 
  def add(cdrom, opts = {})
    if cdrom.is_a?(Hash)
      cdrom = ovirtsdk4::Cdrom.new(cdrom)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      CdromWriter.write_one(cdrom, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return CdromReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `cdrom` service.
  # 
  # @param id [String] The identifier of the `cdrom`.
  # 
  # @return [VmCdromService] A reference to the `cdrom` service.
  # 
  def cdrom_service(id)
    return VmCdromService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return cdrom_service(path)
    end
    return cdrom_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmCdromsService}:#{@path}>"
  end
  
end

class VmDiskService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `activate` method.
  # 
  def activate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/activate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `deactivate` method.
  # 
  def deactivate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/deactivate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `export` method.
  # 
  def export(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/export",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `move` method.
  # 
  def move(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/move",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Detach the disk from the virtual machine.
  # 
  # NOTE: In version 3 of the API this used to also remove the disk completely from the system, but starting with
  # version 4 it doesn't. If you need to remove it completely use the <<services/disk/methods/remove,remove
  # method of the top level disk service>>.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(disk)
    if disk.is_a?(Hash)
      disk = ovirtsdk4::Disk.new(disk)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskWriter.write_one(disk, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmDiskService}:#{@path}>"
  end
  
end

class VmDisksService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `disk`.
  # 
  def add(disk, opts = {})
    if disk.is_a?(Hash)
      disk = ovirtsdk4::Disk.new(disk)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      DiskWriter.write_one(disk, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `disk` service.
  # 
  # @param id [String] The identifier of the `disk`.
  # 
  # @return [VmDiskService] A reference to the `disk` service.
  # 
  def disk_service(id)
    return VmDiskService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return disk_service(path)
    end
    return disk_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmDisksService}:#{@path}>"
  end
  
end

class VmGraphicsConsoleService < GraphicsConsoleService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GraphicsConsoleReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `proxy_ticket` method.
  # 
  def proxy_ticket(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/proxyticket",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.proxy_ticket
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmGraphicsConsoleService}:#{@path}>"
  end
  
end

class VmHostDeviceService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostDeviceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmHostDeviceService}:#{@path}>"
  end
  
end

class VmHostDevicesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `device`.
  # 
  def add(device, opts = {})
    if device.is_a?(Hash)
      device = ovirtsdk4::HostDevice.new(device)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      HostDeviceWriter.write_one(device, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return HostDeviceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostDeviceReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `device` service.
  # 
  # @param id [String] The identifier of the `device`.
  # 
  # @return [VmHostDeviceService] A reference to the `device` service.
  # 
  def device_service(id)
    return VmHostDeviceService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return device_service(path)
    end
    return device_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmHostDevicesService}:#{@path}>"
  end
  
end

class VmNicService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `activate` method.
  # 
  def activate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/activate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `deactivate` method.
  # 
  def deactivate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/deactivate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(nic)
    if nic.is_a?(Hash)
      nic = ovirtsdk4::Nic.new(nic)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NicWriter.write_one(nic, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `reported_devices` service.
  # 
  # @return [VmReportedDevicesService] A reference to `reported_devices` service.
  def reported_devices_service
    return VmReportedDevicesService.new(@connection, "#{@path}/reporteddevices")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'reporteddevices'
      return reported_devices_service
    end
    if path.start_with?('reporteddevices/')
      return reported_devices_service.service(path[16..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmNicService}:#{@path}>"
  end
  
end

class VmNicsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `nic`.
  # 
  def add(nic, opts = {})
    if nic.is_a?(Hash)
      nic = ovirtsdk4::Nic.new(nic)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      NicWriter.write_one(nic, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NicReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `nic` service.
  # 
  # @param id [String] The identifier of the `nic`.
  # 
  # @return [VmNicService] A reference to the `nic` service.
  # 
  def nic_service(id)
    return VmNicService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return nic_service(path)
    end
    return nic_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmNicsService}:#{@path}>"
  end
  
end

class VmNumaNodeService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VirtualNumaNodeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(node)
    if node.is_a?(Hash)
      node = ovirtsdk4::VirtualNumaNode.new(node)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VirtualNumaNodeWriter.write_one(node, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VirtualNumaNodeReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmNumaNodeService}:#{@path}>"
  end
  
end

class VmNumaNodesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `node`.
  # 
  def add(node, opts = {})
    if node.is_a?(Hash)
      node = ovirtsdk4::VirtualNumaNode.new(node)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VirtualNumaNodeWriter.write_one(node, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VirtualNumaNodeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VirtualNumaNodeReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `node` service.
  # 
  # @param id [String] The identifier of the `node`.
  # 
  # @return [VmNumaNodeService] A reference to the `node` service.
  # 
  def node_service(id)
    return VmNumaNodeService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return node_service(path)
    end
    return node_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmNumaNodesService}:#{@path}>"
  end
  
end

class VmPoolService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `allocate_vm` method.
  # 
  def allocate_vm(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/allocatevm",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmPoolReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(pool)
    if pool.is_a?(Hash)
      pool = ovirtsdk4::VmPool.new(pool)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VmPoolWriter.write_one(pool, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmPoolReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmPoolService}:#{@path}>"
  end
  
end

class VmPoolsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `pool`.
  # 
  def add(pool, opts = {})
    if pool.is_a?(Hash)
      pool = ovirtsdk4::VmPool.new(pool)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VmPoolWriter.write_one(pool, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VmPoolReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmPoolReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `pool` service.
  # 
  # @param id [String] The identifier of the `pool`.
  # 
  # @return [VmPoolService] A reference to the `pool` service.
  # 
  def pool_service(id)
    return VmPoolService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return pool_service(path)
    end
    return pool_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmPoolsService}:#{@path}>"
  end
  
end

class VmReportedDeviceService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ReportedDeviceReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmReportedDeviceService}:#{@path}>"
  end
  
end

class VmReportedDevicesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ReportedDeviceReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `reported_device` service.
  # 
  # @param id [String] The identifier of the `reported_device`.
  # 
  # @return [VmReportedDeviceService] A reference to the `reported_device` service.
  # 
  def reported_device_service(id)
    return VmReportedDeviceService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return reported_device_service(path)
    end
    return reported_device_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmReportedDevicesService}:#{@path}>"
  end
  
end

class VmSessionService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SessionReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmSessionService}:#{@path}>"
  end
  
end

class VmSessionsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return SessionReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `session` service.
  # 
  # @param id [String] The identifier of the `session`.
  # 
  # @return [VmSessionService] A reference to the `session` service.
  # 
  def session_service(id)
    return VmSessionService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return session_service(path)
    end
    return session_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmSessionsService}:#{@path}>"
  end
  
end

class VmWatchdogService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(watchdog)
    if watchdog.is_a?(Hash)
      watchdog = ovirtsdk4::Watchdog.new(watchdog)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WatchdogWriter.write_one(watchdog, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmWatchdogService}:#{@path}>"
  end
  
end

class VmWatchdogsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `watchdog`.
  # 
  def add(watchdog, opts = {})
    if watchdog.is_a?(Hash)
      watchdog = ovirtsdk4::Watchdog.new(watchdog)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WatchdogWriter.write_one(watchdog, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WatchdogReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `watchdog` service.
  # 
  # @param id [String] The identifier of the `watchdog`.
  # 
  # @return [VmWatchdogService] A reference to the `watchdog` service.
  # 
  def watchdog_service(id)
    return VmWatchdogService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return watchdog_service(path)
    end
    return watchdog_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmWatchdogsService}:#{@path}>"
  end
  
end

class VmsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Creates a new virtual machine.
  # 
  # The virtual machine can be created in different ways:
  # 
  # - From a template. In this case the identifier or name of the template must be provided. For example, using a
  #   plain shell script and XML:
  # 
  # [source,bash]
  # ----
  # #!/bin/sh -ex
  # 
  # url="https://engine.example.com/ovirt-engine/api"
  # user="admin@internal"
  # password="..."
  # curl \
  # --verbose \
  # --cacert /etc/pki/ovirt-engine/ca.pem \
  # --user "${user}:${password}" \
  # --request POST \
  # --header "Content-Type: application/xml" \
  # --header "Accept: application/xml" \
  # --data '
  # <vm>
  #   <name>myvm</name>
  #   <template>
  #     <name>Blank</name>
  #   </template>
  #   <cluster>
  #     <name>mycluster</name>
  #   </cluster>
  # </vm>
  # ' \
  # "${url}/vms"
  # ----
  # 
  # - From a snapshot. In this case the identifier of the snapshot has to be provided. For example, using a plain
  #   shel script and XML:
  # 
  # [source,bash]
  # ----
  # #!/bin/sh -ex
  # 
  # url="https://engine.example.com/ovirt-engine/api"
  # user="admin@internal"
  # password="..."
  # curl \
  # --verbose \
  # --cacert /etc/pki/ovirt-engine/ca.pem \
  # --user "${user}:${password}" \
  # --request POST \
  # --header "Content-Type: application/xml" \
  # --header "Accept: application/xml" \
  # --data '
  # <vm>
  #   <name>myvm</name>
  #   <snapshots>
  #     <snapshot id="266742a5-6a65-483c-816d-d2ce49746680"/>
  #   </snapshots>
  #   <cluster>
  #     <name>mycluster</name>
  #   </cluster>
  # </vm>
  # ' \
  # "${url}/vms"
  # ----
  # 
  # When creating a virtual machine from a template or from a snapshot it is usually useful to explicitly indicate
  # in what storage domain to create the disks for the virtual machine. If the virtual machine is created from
  # a template then this is achieved passing a set of `disk` elements that indicate the mapping:
  # 
  # [source,xml]
  # ----
  # <vm>
  #   ...
  #   <disks>
  #     <disk id="8d4bd566-6c86-4592-a4a7-912dbf93c298">
  #       <storage_domains>
  #         <storage_domain id="9cb6cb0a-cf1d-41c2-92ca-5a6d665649c9"/>
  #       </storage_domains>
  #     </disk>
  #   </disks>
  # </vm>
  # ----
  # 
  # When the virtual machine is created from a snapshot this set of disks is sligthly different, it uses the
  # `imageId` attribute instead of `id`.
  # 
  # [source,xml]
  # ----
  # <vm>
  #   ...
  #   <disks>
  #     <disk>
  #       <image_id>8d4bd566-6c86-4592-a4a7-912dbf93c298</image_id>
  #       <storage_domains>
  #         <storage_domain id="9cb6cb0a-cf1d-41c2-92ca-5a6d665649c9"/>
  #       </storage_domains>
  #     </disk>
  #   </disks>
  # </vm>
  # ----
  # 
  # In all cases the name or identifier of the cluster where the virtual machine will be created is mandatory.
  # 
  # This is an example of how creating a virtual machine from a snapshot with the disks in a different storage
  # domain can be done with the Python SDK:
  # 
  # [source,python]
  # ----
  # # Find the VM:
  # vm = api.vms.get(name="myvm")
  # 
  # # Find the snapshot:
  # snapshot = None
  # for current in vm.snapshots.list():
  #   if current.get_description() == 'mysnap':
  #     snapshot = current
  #     break
  # 
  # # Find the identifiers of the disks of the snapshot, as we need them in
  # # order to explicitly indicate that we want them created in a different storage
  # # domain:
  # disk_ids = []
  # for current in snapshot.disks.list():
  #   disk_ids.append(current.get_id())
  # 
  # # Find the storage domain where the disks should be created:
  # sd = api.storagedomains.get(name="yourdata")
  # 
  # # Prepare the list of disks for the operation to create the snapshot,
  # # explicitly indicating for each of them the storage domain where it should be
  # # created:
  # disk_list = []
  # for disk_id in disk_ids:
  #   disk = params.Disk(
  #     image_id=disk_id,
  #     storage_domains=params.StorageDomains(
  #       storage_domain=[
  #         params.StorageDomain(
  #           id=sd.get_id(),
  #         ),
  #       ],
  #     ),
  #   )
  #   disk_list.append(disk)
  # 
  # # Create the VM from the snapshot:
  # api.vms.add(
  #   params.VM(
  #     name="myclone",
  #     cluster=params.Cluster(name="mycluster"),
  #     snapshots=params.Snapshots(
  #       snapshot=[
  #         params.Snapshot(
  #           id=snapshot.get_id(),
  #         ),
  #       ],
  #     ),
  #     disks=params.Disks(
  #       disk=disk_list,
  #     ),
  #   )
  # )
  # ----
  # 
  def add(vm, opts = {})
    if vm.is_a?(Hash)
      vm = ovirtsdk4::Vm.new(vm)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VmWriter.write_one(vm, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:case_sensitive]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['case_sensitive'] = value
    end
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    value = opts[:search]
    unless value.nil?
      query['search'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VmReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `vm` service.
  # 
  # @param id [String] The identifier of the `vm`.
  # 
  # @return [VmService] A reference to the `vm` service.
  # 
  def vm_service(id)
    return VmService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return vm_service(path)
    end
    return vm_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VmsService}:#{@path}>"
  end
  
end

class VnicProfileService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(profile)
    if profile.is_a?(Hash)
      profile = ovirtsdk4::VnicProfile.new(profile)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VnicProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VnicProfileService}:#{@path}>"
  end
  
end

class VnicProfilesService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `profile`.
  # 
  def add(profile, opts = {})
    if profile.is_a?(Hash)
      profile = ovirtsdk4::VnicProfile.new(profile)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      VnicProfileWriter.write_one(profile, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return VnicProfileReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `profile` service.
  # 
  # @param id [String] The identifier of the `profile`.
  # 
  # @return [VnicProfileService] A reference to the `profile` service.
  # 
  def profile_service(id)
    return VnicProfileService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return profile_service(path)
    end
    return profile_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{VnicProfilesService}:#{@path}>"
  end
  
end

class WeightService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WeightReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{WeightService}:#{@path}>"
  end
  
end

class WeightsService < Service
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Adds a new `weight`.
  # 
  def add(weight, opts = {})
    if weight.is_a?(Hash)
      weight = ovirtsdk4::Weight.new(weight)
    end
    request = Request.new(:method => :POST, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      WeightWriter.write_one(weight, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 201, 202
      begin
        reader = XmlReader.new(response.body)
        return WeightReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return WeightReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `weight` service.
  # 
  # @param id [String] The identifier of the `weight`.
  # 
  # @return [WeightService] A reference to the `weight` service.
  # 
  def weight_service(id)
    return WeightService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return weight_service(path)
    end
    return weight_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{WeightsService}:#{@path}>"
  end
  
end

class DiskService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `copy` method.
  # 
  def copy(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/copy",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `export` method.
  # 
  def export(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/export",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return DiskReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `move` method.
  # 
  def move(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/move",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{DiskService}:#{@path}>"
  end
  
end

class EngineKatelloErrataService < KatelloErrataService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def list(opts = {})
    query = {}
    value = opts[:max]
    unless value.nil?
      value = Writer.render_integer(value)
      query['max'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return KatelloErratumReader.read_many(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `katello_erratum` service.
  # 
  # @param id [String] The identifier of the `katello_erratum`.
  # 
  # @return [KatelloErratumService] A reference to the `katello_erratum` service.
  # 
  def katello_erratum_service(id)
    return KatelloErratumService.new(@connection, "#{@path}/#{id}")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    index = path.index('/')
    if index.nil?
      return katello_erratum_service(path)
    end
    return katello_erratum_service(path[0..(index - 1)]).service(path[(index +1)..-1])
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{EngineKatelloErrataService}:#{@path}>"
  end
  
end

class ExternalHostProviderService < ExternalProviderService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostProviderReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `import_certificates` method.
  # 
  def import_certificates(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/importcertificates",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `test_connectivity` method.
  # 
  def test_connectivity(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/testconnectivity",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(provider)
    if provider.is_a?(Hash)
      provider = ovirtsdk4::ExternalHostProvider.new(provider)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      ExternalHostProviderWriter.write_one(provider, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return ExternalHostProviderReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `certificates` service.
  # 
  # @return [ExternalProviderCertificatesService] A reference to `certificates` service.
  def certificates_service
    return ExternalProviderCertificatesService.new(@connection, "#{@path}/certificates")
  end
  
  # 
  # Locates the `compute_resources` service.
  # 
  # @return [ExternalComputeResourcesService] A reference to `compute_resources` service.
  def compute_resources_service
    return ExternalComputeResourcesService.new(@connection, "#{@path}/computeresources")
  end
  
  # 
  # Locates the `discovered_hosts` service.
  # 
  # @return [ExternalDiscoveredHostsService] A reference to `discovered_hosts` service.
  def discovered_hosts_service
    return ExternalDiscoveredHostsService.new(@connection, "#{@path}/discoveredhosts")
  end
  
  # 
  # Locates the `host_groups` service.
  # 
  # @return [ExternalHostGroupsService] A reference to `host_groups` service.
  def host_groups_service
    return ExternalHostGroupsService.new(@connection, "#{@path}/hostgroups")
  end
  
  # 
  # Locates the `hosts` service.
  # 
  # @return [ExternalHostsService] A reference to `hosts` service.
  def hosts_service
    return ExternalHostsService.new(@connection, "#{@path}/hosts")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'certificates'
      return certificates_service
    end
    if path.start_with?('certificates/')
      return certificates_service.service(path[13..-1])
    end
    if path == 'computeresources'
      return compute_resources_service
    end
    if path.start_with?('computeresources/')
      return compute_resources_service.service(path[17..-1])
    end
    if path == 'discoveredhosts'
      return discovered_hosts_service
    end
    if path.start_with?('discoveredhosts/')
      return discovered_hosts_service.service(path[16..-1])
    end
    if path == 'hostgroups'
      return host_groups_service
    end
    if path.start_with?('hostgroups/')
      return host_groups_service.service(path[11..-1])
    end
    if path == 'hosts'
      return hosts_service
    end
    if path.start_with?('hosts/')
      return hosts_service.service(path[6..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{ExternalHostProviderService}:#{@path}>"
  end
  
end

class GlusterBrickService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GlusterBrickReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Removes this brick from the volume and deletes it from the database.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Replaces this brick with a new one. The property `brick` is required.
  # 
  def replace(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/replace",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GlusterBrickService}:#{@path}>"
  end
  
end

class GlusterVolumeService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return GlusterVolumeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `get_profile_statistics` method.
  # 
  def get_profile_statistics(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/getprofilestatistics",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.details
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `rebalance` method.
  # 
  def rebalance(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/rebalance",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # Executes the `reset_all_options` method.
  # 
  def reset_all_options(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/resetalloptions",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `reset_option` method.
  # 
  def reset_option(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/resetoption",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `set_option` method.
  # 
  def set_option(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/setoption",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `start` method.
  # 
  def start(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/start",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `start_profile` method.
  # 
  def start_profile(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/startprofile",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `stop` method.
  # 
  def stop(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/stop",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `stop_profile` method.
  # 
  def stop_profile(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/stopprofile",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `stop_rebalance` method.
  # 
  def stop_rebalance(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/stoprebalance",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `gluster_bricks` service.
  # 
  # @return [GlusterBricksService] A reference to `gluster_bricks` service.
  def gluster_bricks_service
    return GlusterBricksService.new(@connection, "#{@path}/glusterbricks")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'glusterbricks'
      return gluster_bricks_service
    end
    if path.start_with?('glusterbricks/')
      return gluster_bricks_service.service(path[14..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{GlusterVolumeService}:#{@path}>"
  end
  
end

class HostService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Executes the `activate` method.
  # 
  def activate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/activate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `approve` method.
  # 
  def approve(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/approve",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `commit_net_config` method.
  # 
  def commit_net_config(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/commitnetconfig",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `deactivate` method.
  # 
  def deactivate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/deactivate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `enroll_certificate` method.
  # 
  def enroll_certificate(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/enrollcertificate",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `fence` method.
  # 
  def fence(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/fence",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.power_management
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `force_select_spm` method.
  # 
  def force_select_spm(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/forceselectspm",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    value = opts[:filter]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['filter'] = value
    end
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `install` method.
  # 
  def install(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/install",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `iscsi_discover` method.
  # 
  def iscsi_discover(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/iscsidiscover",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.iscsi_targets
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `iscsi_login` method.
  # 
  def iscsi_login(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/iscsilogin",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `refresh` method.
  # 
  def refresh(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/refresh",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Deletes the object managed by this service.
  # 
  # @param opts [Hash] Additional options.
  def remove(opts = {})
    query = {}
    value = opts[:async]
    unless value.nil?
      value = Writer.render_boolean(value)
      query['async'] = value
    end
    request = Request.new(:method => :DELETE, :path => @path, :query => query)
    response = @connection.send(request)
    unless response.code == 200
      check_fault(response)
    end
  end
  
  # 
  # This method is used to change the configuration of the network interfaces of a host.
  # 
  # For example, lets assume that you have a host with three network interfaces `eth0`, `eth1` and `eth2` and that
  # you want to configure a new bond using `eth0` and `eth1`, and put a VLAN on top of it. Using a simple shell
  # script and the `curl` command line HTTP client that can be done as follows:
  # 
  # [source]
  # ----
  # #!/bin/sh -ex
  # 
  # url="https://engine.example.com/ovirt-engine/api"
  # user="admin@internal"
  # password="..."
  # 
  # curl \
  # --verbose \
  # --cacert /etc/pki/ovirt-engine/ca.pem \
  # --user "${user}:${password}" \
  # --request POST \
  # --header "Content-Type: application/xml" \
  # --header "Accept: application/xml" \
  # --data '
  # <action>
  #   <modified_bonds>
  #     <host_nic>
  #       <name>bond0</name>
  #       <bonding>
  #         <options>
  #           <option>
  #             <name>mode</name>
  #             <value>4</value>
  #           </option>
  #           <option>
  #             <name>miimon</name>
  #             <value>100</value>
  #           </option>
  #         </options>
  #         <slaves>
  #           <host_nic>
  #             <name>eth1</name>
  #           </host_nic>
  #           <host_nic>
  #             <name>eth2</name>
  #           </host_nic>
  #         </slaves>
  #       </bonding>
  #     </host_nic>
  #   </modified_bonds>
  #   <modified_network_attachments>
  #     <network_attachment>
  #       <network>
  #         <name>myvlan</name>
  #       </network>
  #       <host_nic>
  #         <name>bond0</name>
  #       </host_nic>
  #       <ip_address_assignments>
  #         <assignment_method>static</assignment_method>
  #         <ip_address_assignment>
  #           <ip>
  #             <address>192.168.122.10</address>
  #             <netmask>255.255.255.0</netmask>
  #           </ip>
  #         </ip_address_assignment>
  #       </ip_address_assignments>
  #     </network_attachment>
  #   </modified_network_attachments>
  #  </action>
  # ' \
  # "${url}/hosts/1ff7a191-2f3b-4eff-812b-9f91a30c3acc/setupnetworks"
  # ----
  # 
  # Note that this is valid for version 4 of the API. In previous versions some elements were represented as XML
  # attributes instead of XML elements. In particular the `options` and `ip` elements were represented as follows:
  # 
  # [source,xml]
  # ----
  # <options name="mode" value="4"/>
  # <options name="miimon" value="100"/>
  # <ip address="192.168.122.10" netmask="255.255.255.0"/>
  # ----
  # 
  # Using the Python SDK the same can be done with the following code:
  # 
  # [source,python]
  # ----
  # host.setupnetworks(
  #   params.Action(
  #     modified_bonds=params.HostNics(
  #       host_nic=[
  #         params.HostNIC(
  #           name="bond0",
  #           bonding=params.Bonding(
  #             options=params.Options(
  #               option=[
  #                 params.Option(name="mode", value="4"),
  #                 params.Option(name="miimon", value="100"),
  #               ],
  #             ),
  #             slaves=params.Slaves(
  #               host_nic=[
  #                 params.HostNIC(name="eth1"),
  #                 params.HostNIC(name="eth2"),
  #               ],
  #             ),
  #           ),
  #         ),
  #       ],
  #     ),
  #     modified_network_attachments=params.NetworkAttachments(
  #       network_attachment=[
  #         params.NetworkAttachment(
  #           network=params.Network(name="myvlan"),
  #           host_nic=params.HostNIC(name="bond0"),
  #           ip_address_assignments=params.IpAddressAssignments(
  #             ip_address_assignment=[
  #               params.IpAddressAssignment(
  #                 assignment_method="static",
  #                 ip=params.IP(
  #                   address="192.168.122.10",
  #                   netmask="255.255.255.0",
  #                 ),
  #               ),
  #             ],
  #           ),
  #         ),
  #       ],
  #     ),
  #   ),
  # )
  # ----
  # 
  def setup_networks(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/setupnetworks",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `unregistered_storage_domains_discover` method.
  # 
  def unregistered_storage_domains_discover(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/unregisteredstoragedomainsdiscover",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
      return action.storage_domains
    else
      check_fault(response)
    end
  end
  
  # 
  # Updates the object managed by this service.
  # 
  def update(host)
    if host.is_a?(Hash)
      host = ovirtsdk4::Host.new(host)
    end
    request = Request.new(:method => :PUT, :path => @path)
    begin
      writer = XmlWriter.new(nil, true)
      HostWriter.write_one(host, writer)
      request.body = writer.string
    ensure
      writer.close
    end
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostReader.read_one(reader)
      ensure
        reader.close
      end
      return result
    else
      check_fault(response)
    end
  end
  
  # 
  # Executes the `upgrade` method.
  # 
  def upgrade(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/upgrade",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # List of scheduling labels assigned to this host.
  # 
  # @return [AssignedAffinityLabelsService] A reference to `affinity_labels` service.
  def affinity_labels_service
    return AssignedAffinityLabelsService.new(@connection, "#{@path}/affinitylabels")
  end
  
  # 
  # Locates the `devices` service.
  # 
  # @return [HostDevicesService] A reference to `devices` service.
  def devices_service
    return HostDevicesService.new(@connection, "#{@path}/devices")
  end
  
  # 
  # Locates the `fence_agents` service.
  # 
  # @return [FenceAgentsService] A reference to `fence_agents` service.
  def fence_agents_service
    return FenceAgentsService.new(@connection, "#{@path}/fenceagents")
  end
  
  # 
  # Locates the `hooks` service.
  # 
  # @return [HostHooksService] A reference to `hooks` service.
  def hooks_service
    return HostHooksService.new(@connection, "#{@path}/hooks")
  end
  
  # 
  # Locates the `katello_errata` service.
  # 
  # @return [KatelloErrataService] A reference to `katello_errata` service.
  def katello_errata_service
    return KatelloErrataService.new(@connection, "#{@path}/katelloerrata")
  end
  
  # 
  # Locates the `network_attachments` service.
  # 
  # @return [NetworkAttachmentsService] A reference to `network_attachments` service.
  def network_attachments_service
    return NetworkAttachmentsService.new(@connection, "#{@path}/networkattachments")
  end
  
  # 
  # Locates the `nics` service.
  # 
  # @return [HostNicsService] A reference to `nics` service.
  def nics_service
    return HostNicsService.new(@connection, "#{@path}/nics")
  end
  
  # 
  # Locates the `numa_nodes` service.
  # 
  # @return [HostNumaNodesService] A reference to `numa_nodes` service.
  def numa_nodes_service
    return HostNumaNodesService.new(@connection, "#{@path}/numanodes")
  end
  
  # 
  # Locates the `permissions` service.
  # 
  # @return [AssignedPermissionsService] A reference to `permissions` service.
  def permissions_service
    return AssignedPermissionsService.new(@connection, "#{@path}/permissions")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the `storage` service.
  # 
  # @return [HostStorageService] A reference to `storage` service.
  def storage_service
    return HostStorageService.new(@connection, "#{@path}/storage")
  end
  
  # 
  # Locates the `storage_connection_extensions` service.
  # 
  # @return [StorageServerConnectionExtensionsService] A reference to `storage_connection_extensions` service.
  def storage_connection_extensions_service
    return StorageServerConnectionExtensionsService.new(@connection, "#{@path}/storageconnectionextensions")
  end
  
  # 
  # Locates the `tags` service.
  # 
  # @return [AssignedTagsService] A reference to `tags` service.
  def tags_service
    return AssignedTagsService.new(@connection, "#{@path}/tags")
  end
  
  # 
  # Locates the `unmanaged_networks` service.
  # 
  # @return [UnmanagedNetworksService] A reference to `unmanaged_networks` service.
  def unmanaged_networks_service
    return UnmanagedNetworksService.new(@connection, "#{@path}/unmanagednetworks")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'affinitylabels'
      return affinity_labels_service
    end
    if path.start_with?('affinitylabels/')
      return affinity_labels_service.service(path[15..-1])
    end
    if path == 'devices'
      return devices_service
    end
    if path.start_with?('devices/')
      return devices_service.service(path[8..-1])
    end
    if path == 'fenceagents'
      return fence_agents_service
    end
    if path.start_with?('fenceagents/')
      return fence_agents_service.service(path[12..-1])
    end
    if path == 'hooks'
      return hooks_service
    end
    if path.start_with?('hooks/')
      return hooks_service.service(path[6..-1])
    end
    if path == 'katelloerrata'
      return katello_errata_service
    end
    if path.start_with?('katelloerrata/')
      return katello_errata_service.service(path[14..-1])
    end
    if path == 'networkattachments'
      return network_attachments_service
    end
    if path.start_with?('networkattachments/')
      return network_attachments_service.service(path[19..-1])
    end
    if path == 'nics'
      return nics_service
    end
    if path.start_with?('nics/')
      return nics_service.service(path[5..-1])
    end
    if path == 'numanodes'
      return numa_nodes_service
    end
    if path.start_with?('numanodes/')
      return numa_nodes_service.service(path[10..-1])
    end
    if path == 'permissions'
      return permissions_service
    end
    if path.start_with?('permissions/')
      return permissions_service.service(path[12..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    if path == 'storage'
      return storage_service
    end
    if path.start_with?('storage/')
      return storage_service.service(path[8..-1])
    end
    if path == 'storageconnectionextensions'
      return storage_connection_extensions_service
    end
    if path.start_with?('storageconnectionextensions/')
      return storage_connection_extensions_service.service(path[28..-1])
    end
    if path == 'tags'
      return tags_service
    end
    if path.start_with?('tags/')
      return tags_service.service(path[5..-1])
    end
    if path == 'unmanagednetworks'
      return unmanaged_networks_service
    end
    if path.start_with?('unmanagednetworks/')
      return unmanaged_networks_service.service(path[18..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostService}:#{@path}>"
  end
  
end

class HostNicService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return HostNicReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # The action updates virtual function configuration in case the current resource represents an SR-IOV enabled NIC.
  # The input should be consisted of at least one of the following properties:
  # 
  # - `allNetworksAllowed`
  # - `numberOfVirtualFunctions`
  # 
  # Please see the `HostNicVirtualFunctionsConfiguration` type for the meaning of the properties.
  # 
  def update_virtual_functions_configuration(opts = {})
    action = Action.new(opts)
    writer = XmlWriter.new(nil, true)
    ActionWriter.write_one(action, writer)
    body = writer.string
    writer.close
    request = Request.new({
    :method => :POST,
    :path => "#{@path}/updatevirtualfunctionsconfiguration",
    :body => body,
    })
    response = @connection.send(request)
    case response.code
    when 200
      action = check_action(response)
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `network_attachments` service.
  # 
  # @return [NetworkAttachmentsService] A reference to `network_attachments` service.
  def network_attachments_service
    return NetworkAttachmentsService.new(@connection, "#{@path}/networkattachments")
  end
  
  # 
  # Locates the `network_labels` service.
  # 
  # @return [NetworkLabelsService] A reference to `network_labels` service.
  def network_labels_service
    return NetworkLabelsService.new(@connection, "#{@path}/networklabels")
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Retrieves sub-collection resource of network labels that are allowed on an the virtual functions
  # in case that the current resource represents an SR-IOV physical function NIC.
  # 
  # @return [NetworkLabelsService] A reference to `virtual_function_allowed_labels` service.
  def virtual_function_allowed_labels_service
    return NetworkLabelsService.new(@connection, "#{@path}/virtualfunctionallowedlabels")
  end
  
  # 
  # Retrieves sub-collection resource of networks that are allowed on an the virtual functions
  # in case that the current resource represents an SR-IOV physical function NIC.
  # 
  # @return [VirtualFunctionAllowedNetworksService] A reference to `virtual_function_allowed_networks` service.
  def virtual_function_allowed_networks_service
    return VirtualFunctionAllowedNetworksService.new(@connection, "#{@path}/virtualfunctionallowednetworks")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'networkattachments'
      return network_attachments_service
    end
    if path.start_with?('networkattachments/')
      return network_attachments_service.service(path[19..-1])
    end
    if path == 'networklabels'
      return network_labels_service
    end
    if path.start_with?('networklabels/')
      return network_labels_service.service(path[14..-1])
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    if path == 'virtualfunctionallowedlabels'
      return virtual_function_allowed_labels_service
    end
    if path.start_with?('virtualfunctionallowedlabels/')
      return virtual_function_allowed_labels_service.service(path[29..-1])
    end
    if path == 'virtualfunctionallowednetworks'
      return virtual_function_allowed_networks_service
    end
    if path.start_with?('virtualfunctionallowednetworks/')
      return virtual_function_allowed_networks_service.service(path[31..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostNicService}:#{@path}>"
  end
  
end

class HostNumaNodeService < MeasurableService
  
  # 
  # Creates a new implementation of the service.
  # 
  # @param connection [Connection] The connection to be used by this service.
  # 
  # @param path [String] The relative path of this service, for example `vms/123/disks`.
  # 
  # @api private
  # 
  def initialize(connection, path)
    @connection = connection
    @path = path
  end
  
  # 
  # Returns the representation of the object managed by this service.
  def get(opts = {})
    query = {}
    request = Request.new(:method => :GET, :path => @path, :query => query)
    response = @connection.send(request)
    case response.code
    when 200
      begin
        reader = XmlReader.new(response.body)
        return NumaNodeReader.read_one(reader)
      ensure
        reader.close
      end
    else
      check_fault(response)
    end
  end
  
  # 
  # Locates the `statistics` service.
  # 
  # @return [StatisticsService] A reference to `statistics` service.
  def statistics_service
    return StatisticsService.new(@connection, "#{@path}/statistics")
  end
  
  # 
  # Locates the service corresponding to the given path.
  # 
  # @param path [String] The path of the service.
  # 
  # @return [Service] A reference to the service.
  # 
  def service(path)
    if path.nil? || path == ''
      return self
    end
    if path == 'statistics'
      return statistics_service
    end
    if path.start_with?('statistics/')
      return statistics_service.service(path[11..-1])
    end
    raise Error.new("The path \"#{path}\" doesn't correspond to any service")
  end
  
  # 
  # Returns an string representation of this service.
  # 
  # @return [String]
  # 
  def to_s
    return "#<#{HostNumaNodeService}:#{@path}>"
  end
  
end


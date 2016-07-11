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

# 
# These forward declarations are required in order to avoid circular dependencies.
# 
class AgentConfiguration < Struct
end

class Api < Struct
end

class ApiSummary < Struct
end

class ApiSummaryItem < Struct
end

class Bios < Struct
end

class BlockStatistic < Struct
end

class Bonding < Struct
end

class Boot < Struct
end

class BootMenu < Struct
end

class CloudInit < Struct
end

class Configuration < Struct
end

class Console < Struct
end

class Core < Struct
end

class Cpu < Struct
end

class CpuTopology < Struct
end

class CpuTune < Struct
end

class CpuType < Struct
end

class CustomProperty < Struct
end

class Display < Struct
end

class Dns < Struct
end

class EntityProfileDetail < Struct
end

class ErrorHandling < Struct
end

class Fault < Struct
end

class FencingPolicy < Struct
end

class FopStatistic < Struct
end

class GlusterBrickMemoryInfo < Struct
end

class GlusterClient < Struct
end

class GracePeriod < Struct
end

class GuestOperatingSystem < Struct
end

class HardwareInformation < Struct
end

class HighAvailability < Struct
end

class HostDevicePassthrough < Struct
end

class HostNicVirtualFunctionsConfiguration < Struct
end

class HostedEngine < Struct
end

class Identified < Struct
end

class Image < Identified
end

class Initialization < Struct
end

class Io < Struct
end

class Ip < Struct
end

class IpAddressAssignment < Struct
end

class IscsiBond < Identified
end

class IscsiDetails < Struct
end

class Job < Identified
end

class KatelloErratum < Identified
end

class Kernel < Struct
end

class Ksm < Struct
end

class LogicalUnit < Struct
end

class Mac < Struct
end

class MacPool < Identified
end

class MemoryOverCommit < Struct
end

class MemoryPolicy < Struct
end

class Method < Struct
end

class MigrationBandwidth < Struct
end

class MigrationOptions < Struct
end

class MigrationPolicy < Identified
end

class Network < Identified
end

class NetworkAttachment < Identified
end

class NetworkConfiguration < Struct
end

class NetworkFilter < Identified
end

class NetworkLabel < Identified
end

class NfsProfileDetail < EntityProfileDetail
end

class NicConfiguration < Struct
end

class NumaNode < Identified
end

class NumaNodePin < Struct
end

class OpenStackImage < Identified
end

class OpenStackNetwork < Identified
end

class OpenStackSubnet < Identified
end

class OpenStackVolumeType < Identified
end

class OpenstackVolumeAuthenticationKey < Identified
end

class OperatingSystem < Struct
end

class OperatingSystemInfo < Identified
end

class Option < Struct
end

class Package < Struct
end

class Payload < Struct
end

class Permission < Identified
end

class Permit < Identified
end

class PmProxy < Struct
end

class PortMirroring < Struct
end

class PowerManagement < Struct
end

class Product < Identified
end

class ProductInfo < Struct
end

class ProfileDetail < Struct
end

class Property < Struct
end

class ProxyTicket < Struct
end

class Qos < Identified
end

class Quota < Identified
end

class QuotaClusterLimit < Identified
end

class QuotaStorageLimit < Identified
end

class Range < Struct
end

class Rate < Struct
end

class ReportedConfiguration < Struct
end

class ReportedDevice < Identified
end

class RngDevice < Struct
end

class Role < Identified
end

class SchedulingPolicy < Identified
end

class SchedulingPolicyUnit < Identified
end

class SeLinux < Struct
end

class SerialNumber < Struct
end

class Session < Identified
end

class SkipIfConnectivityBroken < Struct
end

class SkipIfSdActive < Struct
end

class SpecialObjects < Struct
end

class Spm < Struct
end

class Ssh < Identified
end

class SshPublicKey < Identified
end

class Sso < Struct
end

class Statistic < Identified
end

class Step < Identified
end

class StorageConnection < Identified
end

class StorageConnectionExtension < Identified
end

class StorageDomain < Identified
end

class Tag < Identified
end

class TemplateVersion < Struct
end

class Ticket < Struct
end

class TimeZone < Struct
end

class TransparentHugePages < Struct
end

class UnmanagedNetwork < Identified
end

class Usb < Struct
end

class User < Identified
end

class Value < Struct
end

class VcpuPin < Struct
end

class Vendor < Identified
end

class Version < Identified
end

class VirtioScsi < Struct
end

class VirtualNumaNode < NumaNode
end

class Vlan < Struct
end

class VmBase < Identified
end

class VmPlacementPolicy < Struct
end

class VmPool < Identified
end

class VmSummary < Struct
end

class VnicPassThrough < Struct
end

class VnicProfile < Identified
end

class VolumeGroup < Struct
end

class Weight < Identified
end

class Action < Identified
end

class AffinityGroup < Identified
end

class AffinityLabel < Identified
end

class Agent < Identified
end

class Application < Identified
end

class AuthorizedKey < Identified
end

class Balance < Identified
end

class Bookmark < Identified
end

class BrickProfileDetail < EntityProfileDetail
end

class Certificate < Identified
end

class Cluster < Identified
end

class ClusterLevel < Identified
end

class CpuProfile < Identified
end

class DataCenter < Identified
end

class Device < Identified
end

class Disk < Device
end

class DiskAttachment < Identified
end

class DiskProfile < Identified
end

class DiskSnapshot < Disk
end

class Domain < Identified
end

class Event < Identified
end

class ExternalComputeResource < Identified
end

class ExternalDiscoveredHost < Identified
end

class ExternalHost < Identified
end

class ExternalHostGroup < Identified
end

class ExternalProvider < Identified
end

class File < Identified
end

class Filter < Identified
end

class Floppy < Device
end

class GlusterBrickAdvancedDetails < Device
end

class GlusterHook < Identified
end

class GlusterMemoryPool < Identified
end

class GlusterServerHook < Identified
end

class GlusterVolume < Identified
end

class GlusterVolumeProfileDetails < Identified
end

class GraphicsConsole < Identified
end

class Group < Identified
end

class Hook < Identified
end

class Host < Identified
end

class HostDevice < Identified
end

class HostNic < Identified
end

class HostStorage < Identified
end

class Icon < Identified
end

class Nic < Device
end

class OpenStackProvider < ExternalProvider
end

class OpenStackVolumeProvider < OpenStackProvider
end

class Template < VmBase
end

class Vm < VmBase
end

class Watchdog < Device
end

class Cdrom < Device
end

class ExternalHostProvider < ExternalProvider
end

class GlusterBrick < GlusterBrickAdvancedDetails
end

class InstanceType < Template
end

class OpenStackImageProvider < OpenStackProvider
end

class OpenStackNetworkProvider < OpenStackProvider
end

class Snapshot < Vm
end

class AgentConfiguration < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `broker_type` attribute.
  # 
  def broker_type
    return @broker_type
  end
  
  # 
  # Sets the value of the `broker_type` attribute.
  # 
  def broker_type=(value)
    @broker_type = value
  end
  
  # 
  # Returns the value of the `network_mappings` attribute.
  # 
  def network_mappings
    return @network_mappings
  end
  
  # 
  # Sets the value of the `network_mappings` attribute.
  # 
  def network_mappings=(value)
    @network_mappings = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {AgentConfiguration} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :broker_type The value of attribute `broker_type`.
  # 
  # @option opts :network_mappings The value of attribute `network_mappings`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.broker_type = opts[:broker_type]
    self.network_mappings = opts[:network_mappings]
    self.password = opts[:password]
    self.port = opts[:port]
    self.username = opts[:username]
  end
  
end

class Api < Struct
  
  # 
  # Returns the value of the `product_info` attribute.
  # 
  def product_info
    return @product_info
  end
  
  # 
  # Sets the value of the `product_info` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ProductInfo} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def product_info=(value)
    if value.is_a?(Hash)
      value = ProductInfo.new(value)
    end
    @product_info = value
  end
  
  # 
  # Returns the value of the `special_objects` attribute.
  # 
  def special_objects
    return @special_objects
  end
  
  # 
  # Sets the value of the `special_objects` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SpecialObjects} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def special_objects=(value)
    if value.is_a?(Hash)
      value = SpecialObjects.new(value)
    end
    @special_objects = value
  end
  
  # 
  # Returns the value of the `summary` attribute.
  # 
  def summary
    return @summary
  end
  
  # 
  # Sets the value of the `summary` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ApiSummary} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def summary=(value)
    if value.is_a?(Hash)
      value = ApiSummary.new(value)
    end
    @summary = value
  end
  
  # 
  # Returns the value of the `time` attribute.
  # 
  def time
    return @time
  end
  
  # 
  # Sets the value of the `time` attribute.
  # 
  def time=(value)
    @time = value
  end
  
  # 
  # Creates a new instance of the {Api} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :product_info The value of attribute `product_info`.
  # 
  # @option opts [Hash] :special_objects The value of attribute `special_objects`.
  # 
  # @option opts [Hash] :summary The value of attribute `summary`.
  # 
  # @option opts :time The value of attribute `time`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.product_info = opts[:product_info]
    self.special_objects = opts[:special_objects]
    self.summary = opts[:summary]
    self.time = opts[:time]
  end
  
end

class ApiSummary < Struct
  
  # 
  # Returns the value of the `hosts` attribute.
  # 
  def hosts
    return @hosts
  end
  
  # 
  # Sets the value of the `hosts` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ApiSummaryItem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def hosts=(value)
    if value.is_a?(Hash)
      value = ApiSummaryItem.new(value)
    end
    @hosts = value
  end
  
  # 
  # Returns the value of the `storage_domains` attribute.
  # 
  def storage_domains
    return @storage_domains
  end
  
  # 
  # Sets the value of the `storage_domains` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ApiSummaryItem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domains=(value)
    if value.is_a?(Hash)
      value = ApiSummaryItem.new(value)
    end
    @storage_domains = value
  end
  
  # 
  # Returns the value of the `users` attribute.
  # 
  def users
    return @users
  end
  
  # 
  # Sets the value of the `users` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ApiSummaryItem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def users=(value)
    if value.is_a?(Hash)
      value = ApiSummaryItem.new(value)
    end
    @users = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ApiSummaryItem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vms=(value)
    if value.is_a?(Hash)
      value = ApiSummaryItem.new(value)
    end
    @vms = value
  end
  
  # 
  # Creates a new instance of the {ApiSummary} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :hosts The value of attribute `hosts`.
  # 
  # @option opts [Hash] :storage_domains The value of attribute `storage_domains`.
  # 
  # @option opts [Hash] :users The value of attribute `users`.
  # 
  # @option opts [Hash] :vms The value of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.hosts = opts[:hosts]
    self.storage_domains = opts[:storage_domains]
    self.users = opts[:users]
    self.vms = opts[:vms]
  end
  
end

class ApiSummaryItem < Struct
  
  # 
  # Returns the value of the `active` attribute.
  # 
  def active
    return @active
  end
  
  # 
  # Sets the value of the `active` attribute.
  # 
  def active=(value)
    @active = value
  end
  
  # 
  # Returns the value of the `total` attribute.
  # 
  def total
    return @total
  end
  
  # 
  # Sets the value of the `total` attribute.
  # 
  def total=(value)
    @total = value
  end
  
  # 
  # Creates a new instance of the {ApiSummaryItem} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active The value of attribute `active`.
  # 
  # @option opts :total The value of attribute `total`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active = opts[:active]
    self.total = opts[:total]
  end
  
end

class Bios < Struct
  
  # 
  # Returns the value of the `boot_menu` attribute.
  # 
  def boot_menu
    return @boot_menu
  end
  
  # 
  # Sets the value of the `boot_menu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::BootMenu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def boot_menu=(value)
    if value.is_a?(Hash)
      value = BootMenu.new(value)
    end
    @boot_menu = value
  end
  
  # 
  # Creates a new instance of the {Bios} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :boot_menu The value of attribute `boot_menu`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.boot_menu = opts[:boot_menu]
  end
  
end

class BlockStatistic < Struct
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Creates a new instance of the {BlockStatistic} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.statistics = opts[:statistics]
  end
  
end

class Bonding < Struct
  
  # 
  # Returns the value of the `options` attribute.
  # 
  def options
    return @options
  end
  
  # 
  # Sets the value of the `options` attribute.
  # 
  def options=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Option.new(value)
        end
      end
    end
    @options = list
  end
  
  # 
  # Returns the value of the `slaves` attribute.
  # 
  def slaves
    return @slaves
  end
  
  # 
  # Sets the value of the `slaves` attribute.
  # 
  def slaves=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = HostNic.new(value)
        end
      end
    end
    @slaves = list
  end
  
  # 
  # Creates a new instance of the {Bonding} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :options The values of attribute `options`.
  # 
  # @option opts [Array<Hash>] :slaves The values of attribute `slaves`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.options = opts[:options]
    self.slaves = opts[:slaves]
  end
  
end

class Boot < Struct
  
  # 
  # Returns the value of the `devices` attribute.
  # 
  def devices
    return @devices
  end
  
  # 
  # Sets the value of the `devices` attribute.
  # 
  def devices=(list)
    @devices = list
  end
  
  # 
  # Creates a new instance of the {Boot} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :devices The values of attribute `devices`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.devices = opts[:devices]
  end
  
end

class BootMenu < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Creates a new instance of the {BootMenu} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
  end
  
end

class CloudInit < Struct
  
  # 
  # Returns the value of the `authorized_keys` attribute.
  # 
  def authorized_keys
    return @authorized_keys
  end
  
  # 
  # Sets the value of the `authorized_keys` attribute.
  # 
  def authorized_keys=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = AuthorizedKey.new(value)
        end
      end
    end
    @authorized_keys = list
  end
  
  # 
  # Returns the value of the `files` attribute.
  # 
  def files
    return @files
  end
  
  # 
  # Sets the value of the `files` attribute.
  # 
  def files=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = File.new(value)
        end
      end
    end
    @files = list
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `network_configuration` attribute.
  # 
  def network_configuration
    return @network_configuration
  end
  
  # 
  # Sets the value of the `network_configuration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::NetworkConfiguration} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network_configuration=(value)
    if value.is_a?(Hash)
      value = NetworkConfiguration.new(value)
    end
    @network_configuration = value
  end
  
  # 
  # Returns the value of the `regenerate_ssh_keys` attribute.
  # 
  def regenerate_ssh_keys
    return @regenerate_ssh_keys
  end
  
  # 
  # Sets the value of the `regenerate_ssh_keys` attribute.
  # 
  def regenerate_ssh_keys=(value)
    @regenerate_ssh_keys = value
  end
  
  # 
  # Returns the value of the `timezone` attribute.
  # 
  def timezone
    return @timezone
  end
  
  # 
  # Sets the value of the `timezone` attribute.
  # 
  def timezone=(value)
    @timezone = value
  end
  
  # 
  # Returns the value of the `users` attribute.
  # 
  def users
    return @users
  end
  
  # 
  # Sets the value of the `users` attribute.
  # 
  def users=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = User.new(value)
        end
      end
    end
    @users = list
  end
  
  # 
  # Creates a new instance of the {CloudInit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :authorized_keys The values of attribute `authorized_keys`.
  # 
  # @option opts [Array<Hash>] :files The values of attribute `files`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts [Hash] :network_configuration The value of attribute `network_configuration`.
  # 
  # @option opts :regenerate_ssh_keys The value of attribute `regenerate_ssh_keys`.
  # 
  # @option opts :timezone The value of attribute `timezone`.
  # 
  # @option opts [Array<Hash>] :users The values of attribute `users`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authorized_keys = opts[:authorized_keys]
    self.files = opts[:files]
    self.host = opts[:host]
    self.network_configuration = opts[:network_configuration]
    self.regenerate_ssh_keys = opts[:regenerate_ssh_keys]
    self.timezone = opts[:timezone]
    self.users = opts[:users]
  end
  
end

class Configuration < Struct
  
  # 
  # Returns the value of the `data` attribute.
  # 
  def data
    return @data
  end
  
  # 
  # Sets the value of the `data` attribute.
  # 
  def data=(value)
    @data = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {Configuration} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :data The value of attribute `data`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.data = opts[:data]
    self.type = opts[:type]
  end
  
end

class Console < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Creates a new instance of the {Console} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
  end
  
end

class Core < Struct
  
  # 
  # Returns the value of the `index` attribute.
  # 
  def index
    return @index
  end
  
  # 
  # Sets the value of the `index` attribute.
  # 
  def index=(value)
    @index = value
  end
  
  # 
  # Returns the value of the `socket` attribute.
  # 
  def socket
    return @socket
  end
  
  # 
  # Sets the value of the `socket` attribute.
  # 
  def socket=(value)
    @socket = value
  end
  
  # 
  # Creates a new instance of the {Core} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :index The value of attribute `index`.
  # 
  # @option opts :socket The value of attribute `socket`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.index = opts[:index]
    self.socket = opts[:socket]
  end
  
end

class Cpu < Struct
  
  # 
  # Returns the value of the `architecture` attribute.
  # 
  def architecture
    return @architecture
  end
  
  # 
  # Sets the value of the `architecture` attribute.
  # 
  def architecture=(value)
    @architecture = value
  end
  
  # 
  # Returns the value of the `cores` attribute.
  # 
  def cores
    return @cores
  end
  
  # 
  # Sets the value of the `cores` attribute.
  # 
  def cores=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Core.new(value)
        end
      end
    end
    @cores = list
  end
  
  # 
  # Returns the value of the `cpu_tune` attribute.
  # 
  def cpu_tune
    return @cpu_tune
  end
  
  # 
  # Sets the value of the `cpu_tune` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuTune} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu_tune=(value)
    if value.is_a?(Hash)
      value = CpuTune.new(value)
    end
    @cpu_tune = value
  end
  
  # 
  # Returns the value of the `level` attribute.
  # 
  def level
    return @level
  end
  
  # 
  # Sets the value of the `level` attribute.
  # 
  def level=(value)
    @level = value
  end
  
  # 
  # Returns the value of the `mode` attribute.
  # 
  def mode
    return @mode
  end
  
  # 
  # Sets the value of the `mode` attribute.
  # 
  def mode=(value)
    @mode = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `speed` attribute.
  # 
  def speed
    return @speed
  end
  
  # 
  # Sets the value of the `speed` attribute.
  # 
  def speed=(value)
    @speed = value
  end
  
  # 
  # Returns the value of the `topology` attribute.
  # 
  def topology
    return @topology
  end
  
  # 
  # Sets the value of the `topology` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuTopology} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def topology=(value)
    if value.is_a?(Hash)
      value = CpuTopology.new(value)
    end
    @topology = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {Cpu} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :architecture The value of attribute `architecture`.
  # 
  # @option opts [Array<Hash>] :cores The values of attribute `cores`.
  # 
  # @option opts [Hash] :cpu_tune The value of attribute `cpu_tune`.
  # 
  # @option opts :level The value of attribute `level`.
  # 
  # @option opts :mode The value of attribute `mode`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :speed The value of attribute `speed`.
  # 
  # @option opts [Hash] :topology The value of attribute `topology`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.architecture = opts[:architecture]
    self.cores = opts[:cores]
    self.cpu_tune = opts[:cpu_tune]
    self.level = opts[:level]
    self.mode = opts[:mode]
    self.name = opts[:name]
    self.speed = opts[:speed]
    self.topology = opts[:topology]
    self.type = opts[:type]
  end
  
end

class CpuTopology < Struct
  
  # 
  # Returns the value of the `cores` attribute.
  # 
  def cores
    return @cores
  end
  
  # 
  # Sets the value of the `cores` attribute.
  # 
  def cores=(value)
    @cores = value
  end
  
  # 
  # Returns the value of the `sockets` attribute.
  # 
  def sockets
    return @sockets
  end
  
  # 
  # Sets the value of the `sockets` attribute.
  # 
  def sockets=(value)
    @sockets = value
  end
  
  # 
  # Returns the value of the `threads` attribute.
  # 
  def threads
    return @threads
  end
  
  # 
  # Sets the value of the `threads` attribute.
  # 
  def threads=(value)
    @threads = value
  end
  
  # 
  # Creates a new instance of the {CpuTopology} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :cores The value of attribute `cores`.
  # 
  # @option opts :sockets The value of attribute `sockets`.
  # 
  # @option opts :threads The value of attribute `threads`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cores = opts[:cores]
    self.sockets = opts[:sockets]
    self.threads = opts[:threads]
  end
  
end

class CpuTune < Struct
  
  # 
  # Returns the value of the `vcpu_pins` attribute.
  # 
  def vcpu_pins
    return @vcpu_pins
  end
  
  # 
  # Sets the value of the `vcpu_pins` attribute.
  # 
  def vcpu_pins=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = VcpuPin.new(value)
        end
      end
    end
    @vcpu_pins = list
  end
  
  # 
  # Creates a new instance of the {CpuTune} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :vcpu_pins The values of attribute `vcpu_pins`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.vcpu_pins = opts[:vcpu_pins]
  end
  
end

class CpuType < Struct
  
  # 
  # Returns the value of the `architecture` attribute.
  # 
  def architecture
    return @architecture
  end
  
  # 
  # Sets the value of the `architecture` attribute.
  # 
  def architecture=(value)
    @architecture = value
  end
  
  # 
  # Returns the value of the `level` attribute.
  # 
  def level
    return @level
  end
  
  # 
  # Sets the value of the `level` attribute.
  # 
  def level=(value)
    @level = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {CpuType} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :architecture The value of attribute `architecture`.
  # 
  # @option opts :level The value of attribute `level`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.architecture = opts[:architecture]
    self.level = opts[:level]
    self.name = opts[:name]
  end
  
end

class CustomProperty < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `regexp` attribute.
  # 
  def regexp
    return @regexp
  end
  
  # 
  # Sets the value of the `regexp` attribute.
  # 
  def regexp=(value)
    @regexp = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {CustomProperty} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :regexp The value of attribute `regexp`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
    self.regexp = opts[:regexp]
    self.value = opts[:value]
  end
  
end

class Display < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `allow_override` attribute.
  # 
  def allow_override
    return @allow_override
  end
  
  # 
  # Sets the value of the `allow_override` attribute.
  # 
  def allow_override=(value)
    @allow_override = value
  end
  
  # 
  # Returns the value of the `certificate` attribute.
  # 
  def certificate
    return @certificate
  end
  
  # 
  # Sets the value of the `certificate` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Certificate} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def certificate=(value)
    if value.is_a?(Hash)
      value = Certificate.new(value)
    end
    @certificate = value
  end
  
  # 
  # Returns the value of the `copy_paste_enabled` attribute.
  # 
  def copy_paste_enabled
    return @copy_paste_enabled
  end
  
  # 
  # Sets the value of the `copy_paste_enabled` attribute.
  # 
  def copy_paste_enabled=(value)
    @copy_paste_enabled = value
  end
  
  # 
  # Returns the value of the `disconnect_action` attribute.
  # 
  def disconnect_action
    return @disconnect_action
  end
  
  # 
  # Sets the value of the `disconnect_action` attribute.
  # 
  def disconnect_action=(value)
    @disconnect_action = value
  end
  
  # 
  # Returns the value of the `file_transfer_enabled` attribute.
  # 
  def file_transfer_enabled
    return @file_transfer_enabled
  end
  
  # 
  # Sets the value of the `file_transfer_enabled` attribute.
  # 
  def file_transfer_enabled=(value)
    @file_transfer_enabled = value
  end
  
  # 
  # Returns the value of the `keyboard_layout` attribute.
  # 
  def keyboard_layout
    return @keyboard_layout
  end
  
  # 
  # Sets the value of the `keyboard_layout` attribute.
  # 
  def keyboard_layout=(value)
    @keyboard_layout = value
  end
  
  # 
  # Returns the value of the `monitors` attribute.
  # 
  def monitors
    return @monitors
  end
  
  # 
  # Sets the value of the `monitors` attribute.
  # 
  def monitors=(value)
    @monitors = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `proxy` attribute.
  # 
  def proxy
    return @proxy
  end
  
  # 
  # Sets the value of the `proxy` attribute.
  # 
  def proxy=(value)
    @proxy = value
  end
  
  # 
  # Returns the value of the `secure_port` attribute.
  # 
  def secure_port
    return @secure_port
  end
  
  # 
  # Sets the value of the `secure_port` attribute.
  # 
  def secure_port=(value)
    @secure_port = value
  end
  
  # 
  # Returns the value of the `single_qxl_pci` attribute.
  # 
  def single_qxl_pci
    return @single_qxl_pci
  end
  
  # 
  # Sets the value of the `single_qxl_pci` attribute.
  # 
  def single_qxl_pci=(value)
    @single_qxl_pci = value
  end
  
  # 
  # Returns the value of the `smartcard_enabled` attribute.
  # 
  def smartcard_enabled
    return @smartcard_enabled
  end
  
  # 
  # Sets the value of the `smartcard_enabled` attribute.
  # 
  def smartcard_enabled=(value)
    @smartcard_enabled = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {Display} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :allow_override The value of attribute `allow_override`.
  # 
  # @option opts [Hash] :certificate The value of attribute `certificate`.
  # 
  # @option opts :copy_paste_enabled The value of attribute `copy_paste_enabled`.
  # 
  # @option opts :disconnect_action The value of attribute `disconnect_action`.
  # 
  # @option opts :file_transfer_enabled The value of attribute `file_transfer_enabled`.
  # 
  # @option opts :keyboard_layout The value of attribute `keyboard_layout`.
  # 
  # @option opts :monitors The value of attribute `monitors`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :proxy The value of attribute `proxy`.
  # 
  # @option opts :secure_port The value of attribute `secure_port`.
  # 
  # @option opts :single_qxl_pci The value of attribute `single_qxl_pci`.
  # 
  # @option opts :smartcard_enabled The value of attribute `smartcard_enabled`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.allow_override = opts[:allow_override]
    self.certificate = opts[:certificate]
    self.copy_paste_enabled = opts[:copy_paste_enabled]
    self.disconnect_action = opts[:disconnect_action]
    self.file_transfer_enabled = opts[:file_transfer_enabled]
    self.keyboard_layout = opts[:keyboard_layout]
    self.monitors = opts[:monitors]
    self.port = opts[:port]
    self.proxy = opts[:proxy]
    self.secure_port = opts[:secure_port]
    self.single_qxl_pci = opts[:single_qxl_pci]
    self.smartcard_enabled = opts[:smartcard_enabled]
    self.type = opts[:type]
  end
  
end

class Dns < Struct
  
  # 
  # Returns the value of the `search_domains` attribute.
  # 
  def search_domains
    return @search_domains
  end
  
  # 
  # Sets the value of the `search_domains` attribute.
  # 
  def search_domains=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Host.new(value)
        end
      end
    end
    @search_domains = list
  end
  
  # 
  # Returns the value of the `servers` attribute.
  # 
  def servers
    return @servers
  end
  
  # 
  # Sets the value of the `servers` attribute.
  # 
  def servers=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Host.new(value)
        end
      end
    end
    @servers = list
  end
  
  # 
  # Creates a new instance of the {Dns} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :search_domains The values of attribute `search_domains`.
  # 
  # @option opts [Array<Hash>] :servers The values of attribute `servers`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.search_domains = opts[:search_domains]
    self.servers = opts[:servers]
  end
  
end

class EntityProfileDetail < Struct
  
  # 
  # Returns the value of the `profile_details` attribute.
  # 
  def profile_details
    return @profile_details
  end
  
  # 
  # Sets the value of the `profile_details` attribute.
  # 
  def profile_details=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ProfileDetail.new(value)
        end
      end
    end
    @profile_details = list
  end
  
  # 
  # Creates a new instance of the {EntityProfileDetail} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :profile_details The values of attribute `profile_details`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.profile_details = opts[:profile_details]
  end
  
end

class ErrorHandling < Struct
  
  # 
  # Returns the value of the `on_error` attribute.
  # 
  def on_error
    return @on_error
  end
  
  # 
  # Sets the value of the `on_error` attribute.
  # 
  def on_error=(value)
    @on_error = value
  end
  
  # 
  # Creates a new instance of the {ErrorHandling} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :on_error The value of attribute `on_error`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.on_error = opts[:on_error]
  end
  
end

class Fault < Struct
  
  # 
  # Returns the value of the `detail` attribute.
  # 
  def detail
    return @detail
  end
  
  # 
  # Sets the value of the `detail` attribute.
  # 
  def detail=(value)
    @detail = value
  end
  
  # 
  # Returns the value of the `reason` attribute.
  # 
  def reason
    return @reason
  end
  
  # 
  # Sets the value of the `reason` attribute.
  # 
  def reason=(value)
    @reason = value
  end
  
  # 
  # Creates a new instance of the {Fault} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :detail The value of attribute `detail`.
  # 
  # @option opts :reason The value of attribute `reason`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.detail = opts[:detail]
    self.reason = opts[:reason]
  end
  
end

class FencingPolicy < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `skip_if_connectivity_broken` attribute.
  # 
  def skip_if_connectivity_broken
    return @skip_if_connectivity_broken
  end
  
  # 
  # Sets the value of the `skip_if_connectivity_broken` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SkipIfConnectivityBroken} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def skip_if_connectivity_broken=(value)
    if value.is_a?(Hash)
      value = SkipIfConnectivityBroken.new(value)
    end
    @skip_if_connectivity_broken = value
  end
  
  # 
  # Returns the value of the `skip_if_sd_active` attribute.
  # 
  def skip_if_sd_active
    return @skip_if_sd_active
  end
  
  # 
  # Sets the value of the `skip_if_sd_active` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SkipIfSdActive} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def skip_if_sd_active=(value)
    if value.is_a?(Hash)
      value = SkipIfSdActive.new(value)
    end
    @skip_if_sd_active = value
  end
  
  # 
  # Creates a new instance of the {FencingPolicy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts [Hash] :skip_if_connectivity_broken The value of attribute `skip_if_connectivity_broken`.
  # 
  # @option opts [Hash] :skip_if_sd_active The value of attribute `skip_if_sd_active`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
    self.skip_if_connectivity_broken = opts[:skip_if_connectivity_broken]
    self.skip_if_sd_active = opts[:skip_if_sd_active]
  end
  
end

class FopStatistic < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Creates a new instance of the {FopStatistic} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
    self.statistics = opts[:statistics]
  end
  
end

class GlusterBrickMemoryInfo < Struct
  
  # 
  # Returns the value of the `memory_pools` attribute.
  # 
  def memory_pools
    return @memory_pools
  end
  
  # 
  # Sets the value of the `memory_pools` attribute.
  # 
  def memory_pools=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterMemoryPool.new(value)
        end
      end
    end
    @memory_pools = list
  end
  
  # 
  # Creates a new instance of the {GlusterBrickMemoryInfo} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :memory_pools The values of attribute `memory_pools`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.memory_pools = opts[:memory_pools]
  end
  
end

class GlusterClient < Struct
  
  # 
  # Returns the value of the `bytes_read` attribute.
  # 
  def bytes_read
    return @bytes_read
  end
  
  # 
  # Sets the value of the `bytes_read` attribute.
  # 
  def bytes_read=(value)
    @bytes_read = value
  end
  
  # 
  # Returns the value of the `bytes_written` attribute.
  # 
  def bytes_written
    return @bytes_written
  end
  
  # 
  # Sets the value of the `bytes_written` attribute.
  # 
  def bytes_written=(value)
    @bytes_written = value
  end
  
  # 
  # Returns the value of the `client_port` attribute.
  # 
  def client_port
    return @client_port
  end
  
  # 
  # Sets the value of the `client_port` attribute.
  # 
  def client_port=(value)
    @client_port = value
  end
  
  # 
  # Returns the value of the `host_name` attribute.
  # 
  def host_name
    return @host_name
  end
  
  # 
  # Sets the value of the `host_name` attribute.
  # 
  def host_name=(value)
    @host_name = value
  end
  
  # 
  # Creates a new instance of the {GlusterClient} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :bytes_read The value of attribute `bytes_read`.
  # 
  # @option opts :bytes_written The value of attribute `bytes_written`.
  # 
  # @option opts :client_port The value of attribute `client_port`.
  # 
  # @option opts :host_name The value of attribute `host_name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.bytes_read = opts[:bytes_read]
    self.bytes_written = opts[:bytes_written]
    self.client_port = opts[:client_port]
    self.host_name = opts[:host_name]
  end
  
end

class GracePeriod < Struct
  
  # 
  # Returns the value of the `expiry` attribute.
  # 
  def expiry
    return @expiry
  end
  
  # 
  # Sets the value of the `expiry` attribute.
  # 
  def expiry=(value)
    @expiry = value
  end
  
  # 
  # Creates a new instance of the {GracePeriod} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :expiry The value of attribute `expiry`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.expiry = opts[:expiry]
  end
  
end

class GuestOperatingSystem < Struct
  
  # 
  # Returns the value of the `architecture` attribute.
  # 
  def architecture
    return @architecture
  end
  
  # 
  # Sets the value of the `architecture` attribute.
  # 
  def architecture=(value)
    @architecture = value
  end
  
  # 
  # Returns the value of the `codename` attribute.
  # 
  def codename
    return @codename
  end
  
  # 
  # Sets the value of the `codename` attribute.
  # 
  def codename=(value)
    @codename = value
  end
  
  # 
  # Returns the value of the `distribution` attribute.
  # 
  def distribution
    return @distribution
  end
  
  # 
  # Sets the value of the `distribution` attribute.
  # 
  def distribution=(value)
    @distribution = value
  end
  
  # 
  # Returns the value of the `family` attribute.
  # 
  def family
    return @family
  end
  
  # 
  # Sets the value of the `family` attribute.
  # 
  def family=(value)
    @family = value
  end
  
  # 
  # Returns the value of the `kernel` attribute.
  # 
  def kernel
    return @kernel
  end
  
  # 
  # Sets the value of the `kernel` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Kernel} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def kernel=(value)
    if value.is_a?(Hash)
      value = Kernel.new(value)
    end
    @kernel = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {GuestOperatingSystem} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :architecture The value of attribute `architecture`.
  # 
  # @option opts :codename The value of attribute `codename`.
  # 
  # @option opts :distribution The value of attribute `distribution`.
  # 
  # @option opts :family The value of attribute `family`.
  # 
  # @option opts [Hash] :kernel The value of attribute `kernel`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.architecture = opts[:architecture]
    self.codename = opts[:codename]
    self.distribution = opts[:distribution]
    self.family = opts[:family]
    self.kernel = opts[:kernel]
    self.version = opts[:version]
  end
  
end

class HardwareInformation < Struct
  
  # 
  # Returns the value of the `family` attribute.
  # 
  def family
    return @family
  end
  
  # 
  # Sets the value of the `family` attribute.
  # 
  def family=(value)
    @family = value
  end
  
  # 
  # Returns the value of the `manufacturer` attribute.
  # 
  def manufacturer
    return @manufacturer
  end
  
  # 
  # Sets the value of the `manufacturer` attribute.
  # 
  def manufacturer=(value)
    @manufacturer = value
  end
  
  # 
  # Returns the value of the `product_name` attribute.
  # 
  def product_name
    return @product_name
  end
  
  # 
  # Sets the value of the `product_name` attribute.
  # 
  def product_name=(value)
    @product_name = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  def serial_number=(value)
    @serial_number = value
  end
  
  # 
  # Returns the value of the `supported_rng_sources` attribute.
  # 
  def supported_rng_sources
    return @supported_rng_sources
  end
  
  # 
  # Sets the value of the `supported_rng_sources` attribute.
  # 
  def supported_rng_sources=(list)
    @supported_rng_sources = list
  end
  
  # 
  # Returns the value of the `uuid` attribute.
  # 
  def uuid
    return @uuid
  end
  
  # 
  # Sets the value of the `uuid` attribute.
  # 
  def uuid=(value)
    @uuid = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  def version=(value)
    @version = value
  end
  
  # 
  # Creates a new instance of the {HardwareInformation} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :family The value of attribute `family`.
  # 
  # @option opts :manufacturer The value of attribute `manufacturer`.
  # 
  # @option opts :product_name The value of attribute `product_name`.
  # 
  # @option opts :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Array<Hash>] :supported_rng_sources The values of attribute `supported_rng_sources`.
  # 
  # @option opts :uuid The value of attribute `uuid`.
  # 
  # @option opts :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.family = opts[:family]
    self.manufacturer = opts[:manufacturer]
    self.product_name = opts[:product_name]
    self.serial_number = opts[:serial_number]
    self.supported_rng_sources = opts[:supported_rng_sources]
    self.uuid = opts[:uuid]
    self.version = opts[:version]
  end
  
end

class HighAvailability < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `priority` attribute.
  # 
  def priority
    return @priority
  end
  
  # 
  # Sets the value of the `priority` attribute.
  # 
  def priority=(value)
    @priority = value
  end
  
  # 
  # Creates a new instance of the {HighAvailability} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts :priority The value of attribute `priority`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
    self.priority = opts[:priority]
  end
  
end

class HostDevicePassthrough < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Creates a new instance of the {HostDevicePassthrough} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
  end
  
end

class HostNicVirtualFunctionsConfiguration < Struct
  
  # 
  # Returns the value of the `all_networks_allowed` attribute.
  # 
  def all_networks_allowed
    return @all_networks_allowed
  end
  
  # 
  # Sets the value of the `all_networks_allowed` attribute.
  # 
  def all_networks_allowed=(value)
    @all_networks_allowed = value
  end
  
  # 
  # Returns the value of the `max_number_of_virtual_functions` attribute.
  # 
  def max_number_of_virtual_functions
    return @max_number_of_virtual_functions
  end
  
  # 
  # Sets the value of the `max_number_of_virtual_functions` attribute.
  # 
  def max_number_of_virtual_functions=(value)
    @max_number_of_virtual_functions = value
  end
  
  # 
  # Returns the value of the `number_of_virtual_functions` attribute.
  # 
  def number_of_virtual_functions
    return @number_of_virtual_functions
  end
  
  # 
  # Sets the value of the `number_of_virtual_functions` attribute.
  # 
  def number_of_virtual_functions=(value)
    @number_of_virtual_functions = value
  end
  
  # 
  # Creates a new instance of the {HostNicVirtualFunctionsConfiguration} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :all_networks_allowed The value of attribute `all_networks_allowed`.
  # 
  # @option opts :max_number_of_virtual_functions The value of attribute `max_number_of_virtual_functions`.
  # 
  # @option opts :number_of_virtual_functions The value of attribute `number_of_virtual_functions`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.all_networks_allowed = opts[:all_networks_allowed]
    self.max_number_of_virtual_functions = opts[:max_number_of_virtual_functions]
    self.number_of_virtual_functions = opts[:number_of_virtual_functions]
  end
  
end

class HostedEngine < Struct
  
  # 
  # Returns the value of the `active` attribute.
  # 
  def active
    return @active
  end
  
  # 
  # Sets the value of the `active` attribute.
  # 
  def active=(value)
    @active = value
  end
  
  # 
  # Returns the value of the `configured` attribute.
  # 
  def configured
    return @configured
  end
  
  # 
  # Sets the value of the `configured` attribute.
  # 
  def configured=(value)
    @configured = value
  end
  
  # 
  # Returns the value of the `global_maintenance` attribute.
  # 
  def global_maintenance
    return @global_maintenance
  end
  
  # 
  # Sets the value of the `global_maintenance` attribute.
  # 
  def global_maintenance=(value)
    @global_maintenance = value
  end
  
  # 
  # Returns the value of the `local_maintenance` attribute.
  # 
  def local_maintenance
    return @local_maintenance
  end
  
  # 
  # Sets the value of the `local_maintenance` attribute.
  # 
  def local_maintenance=(value)
    @local_maintenance = value
  end
  
  # 
  # Returns the value of the `score` attribute.
  # 
  def score
    return @score
  end
  
  # 
  # Sets the value of the `score` attribute.
  # 
  def score=(value)
    @score = value
  end
  
  # 
  # Creates a new instance of the {HostedEngine} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active The value of attribute `active`.
  # 
  # @option opts :configured The value of attribute `configured`.
  # 
  # @option opts :global_maintenance The value of attribute `global_maintenance`.
  # 
  # @option opts :local_maintenance The value of attribute `local_maintenance`.
  # 
  # @option opts :score The value of attribute `score`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active = opts[:active]
    self.configured = opts[:configured]
    self.global_maintenance = opts[:global_maintenance]
    self.local_maintenance = opts[:local_maintenance]
    self.score = opts[:score]
  end
  
end

class Identified < Struct
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {Identified} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class Image < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Creates a new instance of the {Image} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.storage_domain = opts[:storage_domain]
  end
  
end

class Initialization < Struct
  
  # 
  # Returns the value of the `active_directory_ou` attribute.
  # 
  def active_directory_ou
    return @active_directory_ou
  end
  
  # 
  # Sets the value of the `active_directory_ou` attribute.
  # 
  def active_directory_ou=(value)
    @active_directory_ou = value
  end
  
  # 
  # Returns the value of the `authorized_ssh_keys` attribute.
  # 
  def authorized_ssh_keys
    return @authorized_ssh_keys
  end
  
  # 
  # Sets the value of the `authorized_ssh_keys` attribute.
  # 
  def authorized_ssh_keys=(value)
    @authorized_ssh_keys = value
  end
  
  # 
  # Returns the value of the `cloud_init` attribute.
  # 
  def cloud_init
    return @cloud_init
  end
  
  # 
  # Sets the value of the `cloud_init` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CloudInit} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cloud_init=(value)
    if value.is_a?(Hash)
      value = CloudInit.new(value)
    end
    @cloud_init = value
  end
  
  # 
  # Returns the value of the `configuration` attribute.
  # 
  def configuration
    return @configuration
  end
  
  # 
  # Sets the value of the `configuration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Configuration} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def configuration=(value)
    if value.is_a?(Hash)
      value = Configuration.new(value)
    end
    @configuration = value
  end
  
  # 
  # Returns the value of the `custom_script` attribute.
  # 
  def custom_script
    return @custom_script
  end
  
  # 
  # Sets the value of the `custom_script` attribute.
  # 
  def custom_script=(value)
    @custom_script = value
  end
  
  # 
  # Returns the value of the `dns_search` attribute.
  # 
  def dns_search
    return @dns_search
  end
  
  # 
  # Sets the value of the `dns_search` attribute.
  # 
  def dns_search=(value)
    @dns_search = value
  end
  
  # 
  # Returns the value of the `dns_servers` attribute.
  # 
  def dns_servers
    return @dns_servers
  end
  
  # 
  # Sets the value of the `dns_servers` attribute.
  # 
  def dns_servers=(value)
    @dns_servers = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  def domain=(value)
    @domain = value
  end
  
  # 
  # Returns the value of the `host_name` attribute.
  # 
  def host_name
    return @host_name
  end
  
  # 
  # Sets the value of the `host_name` attribute.
  # 
  def host_name=(value)
    @host_name = value
  end
  
  # 
  # Returns the value of the `input_locale` attribute.
  # 
  def input_locale
    return @input_locale
  end
  
  # 
  # Sets the value of the `input_locale` attribute.
  # 
  def input_locale=(value)
    @input_locale = value
  end
  
  # 
  # Returns the value of the `nic_configurations` attribute.
  # 
  def nic_configurations
    return @nic_configurations
  end
  
  # 
  # Sets the value of the `nic_configurations` attribute.
  # 
  def nic_configurations=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NicConfiguration.new(value)
        end
      end
    end
    @nic_configurations = list
  end
  
  # 
  # Returns the value of the `org_name` attribute.
  # 
  def org_name
    return @org_name
  end
  
  # 
  # Sets the value of the `org_name` attribute.
  # 
  def org_name=(value)
    @org_name = value
  end
  
  # 
  # Returns the value of the `regenerate_ids` attribute.
  # 
  def regenerate_ids
    return @regenerate_ids
  end
  
  # 
  # Sets the value of the `regenerate_ids` attribute.
  # 
  def regenerate_ids=(value)
    @regenerate_ids = value
  end
  
  # 
  # Returns the value of the `regenerate_ssh_keys` attribute.
  # 
  def regenerate_ssh_keys
    return @regenerate_ssh_keys
  end
  
  # 
  # Sets the value of the `regenerate_ssh_keys` attribute.
  # 
  def regenerate_ssh_keys=(value)
    @regenerate_ssh_keys = value
  end
  
  # 
  # Returns the value of the `root_password` attribute.
  # 
  def root_password
    return @root_password
  end
  
  # 
  # Sets the value of the `root_password` attribute.
  # 
  def root_password=(value)
    @root_password = value
  end
  
  # 
  # Returns the value of the `system_locale` attribute.
  # 
  def system_locale
    return @system_locale
  end
  
  # 
  # Sets the value of the `system_locale` attribute.
  # 
  def system_locale=(value)
    @system_locale = value
  end
  
  # 
  # Returns the value of the `timezone` attribute.
  # 
  def timezone
    return @timezone
  end
  
  # 
  # Sets the value of the `timezone` attribute.
  # 
  def timezone=(value)
    @timezone = value
  end
  
  # 
  # Returns the value of the `ui_language` attribute.
  # 
  def ui_language
    return @ui_language
  end
  
  # 
  # Sets the value of the `ui_language` attribute.
  # 
  def ui_language=(value)
    @ui_language = value
  end
  
  # 
  # Returns the value of the `user_locale` attribute.
  # 
  def user_locale
    return @user_locale
  end
  
  # 
  # Sets the value of the `user_locale` attribute.
  # 
  def user_locale=(value)
    @user_locale = value
  end
  
  # 
  # Returns the value of the `user_name` attribute.
  # 
  def user_name
    return @user_name
  end
  
  # 
  # Sets the value of the `user_name` attribute.
  # 
  def user_name=(value)
    @user_name = value
  end
  
  # 
  # Returns the value of the `windows_license_key` attribute.
  # 
  def windows_license_key
    return @windows_license_key
  end
  
  # 
  # Sets the value of the `windows_license_key` attribute.
  # 
  def windows_license_key=(value)
    @windows_license_key = value
  end
  
  # 
  # Creates a new instance of the {Initialization} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active_directory_ou The value of attribute `active_directory_ou`.
  # 
  # @option opts :authorized_ssh_keys The value of attribute `authorized_ssh_keys`.
  # 
  # @option opts [Hash] :cloud_init The value of attribute `cloud_init`.
  # 
  # @option opts [Hash] :configuration The value of attribute `configuration`.
  # 
  # @option opts :custom_script The value of attribute `custom_script`.
  # 
  # @option opts :dns_search The value of attribute `dns_search`.
  # 
  # @option opts :dns_servers The value of attribute `dns_servers`.
  # 
  # @option opts :domain The value of attribute `domain`.
  # 
  # @option opts :host_name The value of attribute `host_name`.
  # 
  # @option opts :input_locale The value of attribute `input_locale`.
  # 
  # @option opts [Array<Hash>] :nic_configurations The values of attribute `nic_configurations`.
  # 
  # @option opts :org_name The value of attribute `org_name`.
  # 
  # @option opts :regenerate_ids The value of attribute `regenerate_ids`.
  # 
  # @option opts :regenerate_ssh_keys The value of attribute `regenerate_ssh_keys`.
  # 
  # @option opts :root_password The value of attribute `root_password`.
  # 
  # @option opts :system_locale The value of attribute `system_locale`.
  # 
  # @option opts :timezone The value of attribute `timezone`.
  # 
  # @option opts :ui_language The value of attribute `ui_language`.
  # 
  # @option opts :user_locale The value of attribute `user_locale`.
  # 
  # @option opts :user_name The value of attribute `user_name`.
  # 
  # @option opts :windows_license_key The value of attribute `windows_license_key`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active_directory_ou = opts[:active_directory_ou]
    self.authorized_ssh_keys = opts[:authorized_ssh_keys]
    self.cloud_init = opts[:cloud_init]
    self.configuration = opts[:configuration]
    self.custom_script = opts[:custom_script]
    self.dns_search = opts[:dns_search]
    self.dns_servers = opts[:dns_servers]
    self.domain = opts[:domain]
    self.host_name = opts[:host_name]
    self.input_locale = opts[:input_locale]
    self.nic_configurations = opts[:nic_configurations]
    self.org_name = opts[:org_name]
    self.regenerate_ids = opts[:regenerate_ids]
    self.regenerate_ssh_keys = opts[:regenerate_ssh_keys]
    self.root_password = opts[:root_password]
    self.system_locale = opts[:system_locale]
    self.timezone = opts[:timezone]
    self.ui_language = opts[:ui_language]
    self.user_locale = opts[:user_locale]
    self.user_name = opts[:user_name]
    self.windows_license_key = opts[:windows_license_key]
  end
  
end

class Io < Struct
  
  # 
  # Returns the value of the `threads` attribute.
  # 
  def threads
    return @threads
  end
  
  # 
  # Sets the value of the `threads` attribute.
  # 
  def threads=(value)
    @threads = value
  end
  
  # 
  # Creates a new instance of the {Io} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :threads The value of attribute `threads`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.threads = opts[:threads]
  end
  
end

class Ip < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `gateway` attribute.
  # 
  def gateway
    return @gateway
  end
  
  # 
  # Sets the value of the `gateway` attribute.
  # 
  def gateway=(value)
    @gateway = value
  end
  
  # 
  # Returns the value of the `netmask` attribute.
  # 
  def netmask
    return @netmask
  end
  
  # 
  # Sets the value of the `netmask` attribute.
  # 
  def netmask=(value)
    @netmask = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  def version=(value)
    @version = value
  end
  
  # 
  # Creates a new instance of the {Ip} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :gateway The value of attribute `gateway`.
  # 
  # @option opts :netmask The value of attribute `netmask`.
  # 
  # @option opts :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.gateway = opts[:gateway]
    self.netmask = opts[:netmask]
    self.version = opts[:version]
  end
  
end

class IpAddressAssignment < Struct
  
  # 
  # Returns the value of the `assignment_method` attribute.
  # 
  def assignment_method
    return @assignment_method
  end
  
  # 
  # Sets the value of the `assignment_method` attribute.
  # 
  def assignment_method=(value)
    @assignment_method = value
  end
  
  # 
  # Returns the value of the `ip` attribute.
  # 
  def ip
    return @ip
  end
  
  # 
  # Sets the value of the `ip` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ip} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ip=(value)
    if value.is_a?(Hash)
      value = Ip.new(value)
    end
    @ip = value
  end
  
  # 
  # Creates a new instance of the {IpAddressAssignment} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :assignment_method The value of attribute `assignment_method`.
  # 
  # @option opts [Hash] :ip The value of attribute `ip`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.assignment_method = opts[:assignment_method]
    self.ip = opts[:ip]
  end
  
end

class IscsiBond < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `networks` attribute.
  # 
  def networks
    return @networks
  end
  
  # 
  # Sets the value of the `networks` attribute.
  # 
  def networks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Network.new(value)
        end
      end
    end
    @networks = list
  end
  
  # 
  # Returns the value of the `storage_connections` attribute.
  # 
  def storage_connections
    return @storage_connections
  end
  
  # 
  # Sets the value of the `storage_connections` attribute.
  # 
  def storage_connections=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageConnection.new(value)
        end
      end
    end
    @storage_connections = list
  end
  
  # 
  # Creates a new instance of the {IscsiBond} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :networks The values of attribute `networks`.
  # 
  # @option opts [Array<Hash>] :storage_connections The values of attribute `storage_connections`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.networks = opts[:networks]
    self.storage_connections = opts[:storage_connections]
  end
  
end

class IscsiDetails < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `disk_id` attribute.
  # 
  def disk_id
    return @disk_id
  end
  
  # 
  # Sets the value of the `disk_id` attribute.
  # 
  def disk_id=(value)
    @disk_id = value
  end
  
  # 
  # Returns the value of the `initiator` attribute.
  # 
  def initiator
    return @initiator
  end
  
  # 
  # Sets the value of the `initiator` attribute.
  # 
  def initiator=(value)
    @initiator = value
  end
  
  # 
  # Returns the value of the `lun_mapping` attribute.
  # 
  def lun_mapping
    return @lun_mapping
  end
  
  # 
  # Sets the value of the `lun_mapping` attribute.
  # 
  def lun_mapping=(value)
    @lun_mapping = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `paths` attribute.
  # 
  def paths
    return @paths
  end
  
  # 
  # Sets the value of the `paths` attribute.
  # 
  def paths=(value)
    @paths = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `portal` attribute.
  # 
  def portal
    return @portal
  end
  
  # 
  # Sets the value of the `portal` attribute.
  # 
  def portal=(value)
    @portal = value
  end
  
  # 
  # Returns the value of the `product_id` attribute.
  # 
  def product_id
    return @product_id
  end
  
  # 
  # Sets the value of the `product_id` attribute.
  # 
  def product_id=(value)
    @product_id = value
  end
  
  # 
  # Returns the value of the `serial` attribute.
  # 
  def serial
    return @serial
  end
  
  # 
  # Sets the value of the `serial` attribute.
  # 
  def serial=(value)
    @serial = value
  end
  
  # 
  # Returns the value of the `size` attribute.
  # 
  def size
    return @size
  end
  
  # 
  # Sets the value of the `size` attribute.
  # 
  def size=(value)
    @size = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domain_id` attribute.
  # 
  def storage_domain_id
    return @storage_domain_id
  end
  
  # 
  # Sets the value of the `storage_domain_id` attribute.
  # 
  def storage_domain_id=(value)
    @storage_domain_id = value
  end
  
  # 
  # Returns the value of the `target` attribute.
  # 
  def target
    return @target
  end
  
  # 
  # Sets the value of the `target` attribute.
  # 
  def target=(value)
    @target = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Returns the value of the `vendor_id` attribute.
  # 
  def vendor_id
    return @vendor_id
  end
  
  # 
  # Sets the value of the `vendor_id` attribute.
  # 
  def vendor_id=(value)
    @vendor_id = value
  end
  
  # 
  # Returns the value of the `volume_group_id` attribute.
  # 
  def volume_group_id
    return @volume_group_id
  end
  
  # 
  # Sets the value of the `volume_group_id` attribute.
  # 
  def volume_group_id=(value)
    @volume_group_id = value
  end
  
  # 
  # Creates a new instance of the {IscsiDetails} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :disk_id The value of attribute `disk_id`.
  # 
  # @option opts :initiator The value of attribute `initiator`.
  # 
  # @option opts :lun_mapping The value of attribute `lun_mapping`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :paths The value of attribute `paths`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :portal The value of attribute `portal`.
  # 
  # @option opts :product_id The value of attribute `product_id`.
  # 
  # @option opts :serial The value of attribute `serial`.
  # 
  # @option opts :size The value of attribute `size`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :storage_domain_id The value of attribute `storage_domain_id`.
  # 
  # @option opts :target The value of attribute `target`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # @option opts :vendor_id The value of attribute `vendor_id`.
  # 
  # @option opts :volume_group_id The value of attribute `volume_group_id`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.disk_id = opts[:disk_id]
    self.initiator = opts[:initiator]
    self.lun_mapping = opts[:lun_mapping]
    self.password = opts[:password]
    self.paths = opts[:paths]
    self.port = opts[:port]
    self.portal = opts[:portal]
    self.product_id = opts[:product_id]
    self.serial = opts[:serial]
    self.size = opts[:size]
    self.status = opts[:status]
    self.storage_domain_id = opts[:storage_domain_id]
    self.target = opts[:target]
    self.username = opts[:username]
    self.vendor_id = opts[:vendor_id]
    self.volume_group_id = opts[:volume_group_id]
  end
  
end

class Job < Identified
  
  # 
  # Returns the value of the `auto_cleared` attribute.
  # 
  def auto_cleared
    return @auto_cleared
  end
  
  # 
  # Sets the value of the `auto_cleared` attribute.
  # 
  def auto_cleared=(value)
    @auto_cleared = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `end_time` attribute.
  # 
  def end_time
    return @end_time
  end
  
  # 
  # Sets the value of the `end_time` attribute.
  # 
  def end_time=(value)
    @end_time = value
  end
  
  # 
  # Returns the value of the `external` attribute.
  # 
  def external
    return @external
  end
  
  # 
  # Sets the value of the `external` attribute.
  # 
  def external=(value)
    @external = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `last_updated` attribute.
  # 
  def last_updated
    return @last_updated
  end
  
  # 
  # Sets the value of the `last_updated` attribute.
  # 
  def last_updated=(value)
    @last_updated = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `owner` attribute.
  # 
  def owner
    return @owner
  end
  
  # 
  # Sets the value of the `owner` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def owner=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @owner = value
  end
  
  # 
  # Returns the value of the `start_time` attribute.
  # 
  def start_time
    return @start_time
  end
  
  # 
  # Sets the value of the `start_time` attribute.
  # 
  def start_time=(value)
    @start_time = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `steps` attribute.
  # 
  def steps
    return @steps
  end
  
  # 
  # Sets the value of the `steps` attribute.
  # 
  def steps=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Step.new(value)
        end
      end
    end
    @steps = list
  end
  
  # 
  # Creates a new instance of the {Job} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :auto_cleared The value of attribute `auto_cleared`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :end_time The value of attribute `end_time`.
  # 
  # @option opts :external The value of attribute `external`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :last_updated The value of attribute `last_updated`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :owner The value of attribute `owner`.
  # 
  # @option opts :start_time The value of attribute `start_time`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Array<Hash>] :steps The values of attribute `steps`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.auto_cleared = opts[:auto_cleared]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.end_time = opts[:end_time]
    self.external = opts[:external]
    self.id = opts[:id]
    self.last_updated = opts[:last_updated]
    self.name = opts[:name]
    self.owner = opts[:owner]
    self.start_time = opts[:start_time]
    self.status = opts[:status]
    self.steps = opts[:steps]
  end
  
end

class KatelloErratum < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `issued` attribute.
  # 
  def issued
    return @issued
  end
  
  # 
  # Sets the value of the `issued` attribute.
  # 
  def issued=(value)
    @issued = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `packages` attribute.
  # 
  def packages
    return @packages
  end
  
  # 
  # Sets the value of the `packages` attribute.
  # 
  def packages=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Package.new(value)
        end
      end
    end
    @packages = list
  end
  
  # 
  # Returns the value of the `severity` attribute.
  # 
  def severity
    return @severity
  end
  
  # 
  # Sets the value of the `severity` attribute.
  # 
  def severity=(value)
    @severity = value
  end
  
  # 
  # Returns the value of the `solution` attribute.
  # 
  def solution
    return @solution
  end
  
  # 
  # Sets the value of the `solution` attribute.
  # 
  def solution=(value)
    @solution = value
  end
  
  # 
  # Returns the value of the `summary` attribute.
  # 
  def summary
    return @summary
  end
  
  # 
  # Sets the value of the `summary` attribute.
  # 
  def summary=(value)
    @summary = value
  end
  
  # 
  # Returns the value of the `title` attribute.
  # 
  def title
    return @title
  end
  
  # 
  # Sets the value of the `title` attribute.
  # 
  def title=(value)
    @title = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {KatelloErratum} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :issued The value of attribute `issued`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :packages The values of attribute `packages`.
  # 
  # @option opts :severity The value of attribute `severity`.
  # 
  # @option opts :solution The value of attribute `solution`.
  # 
  # @option opts :summary The value of attribute `summary`.
  # 
  # @option opts :title The value of attribute `title`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.issued = opts[:issued]
    self.name = opts[:name]
    self.packages = opts[:packages]
    self.severity = opts[:severity]
    self.solution = opts[:solution]
    self.summary = opts[:summary]
    self.title = opts[:title]
    self.type = opts[:type]
    self.vm = opts[:vm]
  end
  
end

class Kernel < Struct
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {Kernel} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.version = opts[:version]
  end
  
end

class Ksm < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `merge_across_nodes` attribute.
  # 
  def merge_across_nodes
    return @merge_across_nodes
  end
  
  # 
  # Sets the value of the `merge_across_nodes` attribute.
  # 
  def merge_across_nodes=(value)
    @merge_across_nodes = value
  end
  
  # 
  # Creates a new instance of the {Ksm} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts :merge_across_nodes The value of attribute `merge_across_nodes`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
    self.merge_across_nodes = opts[:merge_across_nodes]
  end
  
end

class LogicalUnit < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `disk_id` attribute.
  # 
  def disk_id
    return @disk_id
  end
  
  # 
  # Sets the value of the `disk_id` attribute.
  # 
  def disk_id=(value)
    @disk_id = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `lun_mapping` attribute.
  # 
  def lun_mapping
    return @lun_mapping
  end
  
  # 
  # Sets the value of the `lun_mapping` attribute.
  # 
  def lun_mapping=(value)
    @lun_mapping = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `paths` attribute.
  # 
  def paths
    return @paths
  end
  
  # 
  # Sets the value of the `paths` attribute.
  # 
  def paths=(value)
    @paths = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `portal` attribute.
  # 
  def portal
    return @portal
  end
  
  # 
  # Sets the value of the `portal` attribute.
  # 
  def portal=(value)
    @portal = value
  end
  
  # 
  # Returns the value of the `product_id` attribute.
  # 
  def product_id
    return @product_id
  end
  
  # 
  # Sets the value of the `product_id` attribute.
  # 
  def product_id=(value)
    @product_id = value
  end
  
  # 
  # Returns the value of the `serial` attribute.
  # 
  def serial
    return @serial
  end
  
  # 
  # Sets the value of the `serial` attribute.
  # 
  def serial=(value)
    @serial = value
  end
  
  # 
  # Returns the value of the `size` attribute.
  # 
  def size
    return @size
  end
  
  # 
  # Sets the value of the `size` attribute.
  # 
  def size=(value)
    @size = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domain_id` attribute.
  # 
  def storage_domain_id
    return @storage_domain_id
  end
  
  # 
  # Sets the value of the `storage_domain_id` attribute.
  # 
  def storage_domain_id=(value)
    @storage_domain_id = value
  end
  
  # 
  # Returns the value of the `target` attribute.
  # 
  def target
    return @target
  end
  
  # 
  # Sets the value of the `target` attribute.
  # 
  def target=(value)
    @target = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Returns the value of the `vendor_id` attribute.
  # 
  def vendor_id
    return @vendor_id
  end
  
  # 
  # Sets the value of the `vendor_id` attribute.
  # 
  def vendor_id=(value)
    @vendor_id = value
  end
  
  # 
  # Returns the value of the `volume_group_id` attribute.
  # 
  def volume_group_id
    return @volume_group_id
  end
  
  # 
  # Sets the value of the `volume_group_id` attribute.
  # 
  def volume_group_id=(value)
    @volume_group_id = value
  end
  
  # 
  # Creates a new instance of the {LogicalUnit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :disk_id The value of attribute `disk_id`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :lun_mapping The value of attribute `lun_mapping`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :paths The value of attribute `paths`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :portal The value of attribute `portal`.
  # 
  # @option opts :product_id The value of attribute `product_id`.
  # 
  # @option opts :serial The value of attribute `serial`.
  # 
  # @option opts :size The value of attribute `size`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :storage_domain_id The value of attribute `storage_domain_id`.
  # 
  # @option opts :target The value of attribute `target`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # @option opts :vendor_id The value of attribute `vendor_id`.
  # 
  # @option opts :volume_group_id The value of attribute `volume_group_id`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.disk_id = opts[:disk_id]
    self.id = opts[:id]
    self.lun_mapping = opts[:lun_mapping]
    self.password = opts[:password]
    self.paths = opts[:paths]
    self.port = opts[:port]
    self.portal = opts[:portal]
    self.product_id = opts[:product_id]
    self.serial = opts[:serial]
    self.size = opts[:size]
    self.status = opts[:status]
    self.storage_domain_id = opts[:storage_domain_id]
    self.target = opts[:target]
    self.username = opts[:username]
    self.vendor_id = opts[:vendor_id]
    self.volume_group_id = opts[:volume_group_id]
  end
  
end

class Mac < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Creates a new instance of the {Mac} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
  end
  
end

class MacPool < Identified
  
  # 
  # Returns the value of the `allow_duplicates` attribute.
  # 
  def allow_duplicates
    return @allow_duplicates
  end
  
  # 
  # Sets the value of the `allow_duplicates` attribute.
  # 
  def allow_duplicates=(value)
    @allow_duplicates = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `default_pool` attribute.
  # 
  def default_pool
    return @default_pool
  end
  
  # 
  # Sets the value of the `default_pool` attribute.
  # 
  def default_pool=(value)
    @default_pool = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `ranges` attribute.
  # 
  def ranges
    return @ranges
  end
  
  # 
  # Sets the value of the `ranges` attribute.
  # 
  def ranges=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Range.new(value)
        end
      end
    end
    @ranges = list
  end
  
  # 
  # Creates a new instance of the {MacPool} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :allow_duplicates The value of attribute `allow_duplicates`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :default_pool The value of attribute `default_pool`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :ranges The values of attribute `ranges`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.allow_duplicates = opts[:allow_duplicates]
    self.comment = opts[:comment]
    self.default_pool = opts[:default_pool]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.ranges = opts[:ranges]
  end
  
end

class MemoryOverCommit < Struct
  
  # 
  # Returns the value of the `percent` attribute.
  # 
  def percent
    return @percent
  end
  
  # 
  # Sets the value of the `percent` attribute.
  # 
  def percent=(value)
    @percent = value
  end
  
  # 
  # Creates a new instance of the {MemoryOverCommit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :percent The value of attribute `percent`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.percent = opts[:percent]
  end
  
end

class MemoryPolicy < Struct
  
  # 
  # Returns the value of the `ballooning` attribute.
  # 
  def ballooning
    return @ballooning
  end
  
  # 
  # Sets the value of the `ballooning` attribute.
  # 
  def ballooning=(value)
    @ballooning = value
  end
  
  # 
  # Returns the value of the `guaranteed` attribute.
  # 
  def guaranteed
    return @guaranteed
  end
  
  # 
  # Sets the value of the `guaranteed` attribute.
  # 
  def guaranteed=(value)
    @guaranteed = value
  end
  
  # 
  # Returns the value of the `over_commit` attribute.
  # 
  def over_commit
    return @over_commit
  end
  
  # 
  # Sets the value of the `over_commit` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryOverCommit} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def over_commit=(value)
    if value.is_a?(Hash)
      value = MemoryOverCommit.new(value)
    end
    @over_commit = value
  end
  
  # 
  # Returns the value of the `transparent_huge_pages` attribute.
  # 
  def transparent_huge_pages
    return @transparent_huge_pages
  end
  
  # 
  # Sets the value of the `transparent_huge_pages` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TransparentHugePages} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def transparent_huge_pages=(value)
    if value.is_a?(Hash)
      value = TransparentHugePages.new(value)
    end
    @transparent_huge_pages = value
  end
  
  # 
  # Creates a new instance of the {MemoryPolicy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :ballooning The value of attribute `ballooning`.
  # 
  # @option opts :guaranteed The value of attribute `guaranteed`.
  # 
  # @option opts [Hash] :over_commit The value of attribute `over_commit`.
  # 
  # @option opts [Hash] :transparent_huge_pages The value of attribute `transparent_huge_pages`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.ballooning = opts[:ballooning]
    self.guaranteed = opts[:guaranteed]
    self.over_commit = opts[:over_commit]
    self.transparent_huge_pages = opts[:transparent_huge_pages]
  end
  
end

class Method < Struct
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Creates a new instance of the {Method} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.id = opts[:id]
  end
  
end

class MigrationBandwidth < Struct
  
  # 
  # Returns the value of the `assignment_method` attribute.
  # 
  def assignment_method
    return @assignment_method
  end
  
  # 
  # Sets the value of the `assignment_method` attribute.
  # 
  def assignment_method=(value)
    @assignment_method = value
  end
  
  # 
  # Returns the value of the `custom_value` attribute.
  # 
  def custom_value
    return @custom_value
  end
  
  # 
  # Sets the value of the `custom_value` attribute.
  # 
  def custom_value=(value)
    @custom_value = value
  end
  
  # 
  # Creates a new instance of the {MigrationBandwidth} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :assignment_method The value of attribute `assignment_method`.
  # 
  # @option opts :custom_value The value of attribute `custom_value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.assignment_method = opts[:assignment_method]
    self.custom_value = opts[:custom_value]
  end
  
end

class MigrationOptions < Struct
  
  # 
  # Returns the value of the `auto_converge` attribute.
  # 
  def auto_converge
    return @auto_converge
  end
  
  # 
  # Sets the value of the `auto_converge` attribute.
  # 
  def auto_converge=(value)
    @auto_converge = value
  end
  
  # 
  # Returns the value of the `bandwidth` attribute.
  # 
  def bandwidth
    return @bandwidth
  end
  
  # 
  # Sets the value of the `bandwidth` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationBandwidth} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bandwidth=(value)
    if value.is_a?(Hash)
      value = MigrationBandwidth.new(value)
    end
    @bandwidth = value
  end
  
  # 
  # Returns the value of the `compressed` attribute.
  # 
  def compressed
    return @compressed
  end
  
  # 
  # Sets the value of the `compressed` attribute.
  # 
  def compressed=(value)
    @compressed = value
  end
  
  # 
  # Returns the value of the `policy` attribute.
  # 
  def policy
    return @policy
  end
  
  # 
  # Sets the value of the `policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def policy=(value)
    if value.is_a?(Hash)
      value = MigrationPolicy.new(value)
    end
    @policy = value
  end
  
  # 
  # Creates a new instance of the {MigrationOptions} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :auto_converge The value of attribute `auto_converge`.
  # 
  # @option opts [Hash] :bandwidth The value of attribute `bandwidth`.
  # 
  # @option opts :compressed The value of attribute `compressed`.
  # 
  # @option opts [Hash] :policy The value of attribute `policy`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.auto_converge = opts[:auto_converge]
    self.bandwidth = opts[:bandwidth]
    self.compressed = opts[:compressed]
    self.policy = opts[:policy]
  end
  
end

class MigrationPolicy < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {MigrationPolicy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class Network < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  def display=(value)
    @display = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ip` attribute.
  # 
  def ip
    return @ip
  end
  
  # 
  # Sets the value of the `ip` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ip} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ip=(value)
    if value.is_a?(Hash)
      value = Ip.new(value)
    end
    @ip = value
  end
  
  # 
  # Returns the value of the `mtu` attribute.
  # 
  def mtu
    return @mtu
  end
  
  # 
  # Sets the value of the `mtu` attribute.
  # 
  def mtu=(value)
    @mtu = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network_labels` attribute.
  # 
  def network_labels
    return @network_labels
  end
  
  # 
  # Sets the value of the `network_labels` attribute.
  # 
  def network_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkLabel.new(value)
        end
      end
    end
    @network_labels = list
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `profile_required` attribute.
  # 
  def profile_required
    return @profile_required
  end
  
  # 
  # Sets the value of the `profile_required` attribute.
  # 
  def profile_required=(value)
    @profile_required = value
  end
  
  # 
  # Returns the value of the `qos` attribute.
  # 
  def qos
    return @qos
  end
  
  # 
  # Sets the value of the `qos` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Qos} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def qos=(value)
    if value.is_a?(Hash)
      value = Qos.new(value)
    end
    @qos = value
  end
  
  # 
  # Returns the value of the `required` attribute.
  # 
  def required
    return @required
  end
  
  # 
  # Sets the value of the `required` attribute.
  # 
  def required=(value)
    @required = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `stp` attribute.
  # 
  def stp
    return @stp
  end
  
  # 
  # Sets the value of the `stp` attribute.
  # 
  def stp=(value)
    @stp = value
  end
  
  # 
  # Returns the value of the `usages` attribute.
  # 
  def usages
    return @usages
  end
  
  # 
  # Sets the value of the `usages` attribute.
  # 
  def usages=(list)
    @usages = list
  end
  
  # 
  # Returns the value of the `vlan` attribute.
  # 
  def vlan
    return @vlan
  end
  
  # 
  # Sets the value of the `vlan` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vlan} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vlan=(value)
    if value.is_a?(Hash)
      value = Vlan.new(value)
    end
    @vlan = value
  end
  
  # 
  # Returns the value of the `vnic_profiles` attribute.
  # 
  def vnic_profiles
    return @vnic_profiles
  end
  
  # 
  # Sets the value of the `vnic_profiles` attribute.
  # 
  def vnic_profiles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = VnicProfile.new(value)
        end
      end
    end
    @vnic_profiles = list
  end
  
  # 
  # Creates a new instance of the {Network} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :display The value of attribute `display`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :ip The value of attribute `ip`.
  # 
  # @option opts :mtu The value of attribute `mtu`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :network_labels The values of attribute `network_labels`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :profile_required The value of attribute `profile_required`.
  # 
  # @option opts [Hash] :qos The value of attribute `qos`.
  # 
  # @option opts :required The value of attribute `required`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :stp The value of attribute `stp`.
  # 
  # @option opts [Array<Hash>] :usages The values of attribute `usages`.
  # 
  # @option opts [Hash] :vlan The value of attribute `vlan`.
  # 
  # @option opts [Array<Hash>] :vnic_profiles The values of attribute `vnic_profiles`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.display = opts[:display]
    self.id = opts[:id]
    self.ip = opts[:ip]
    self.mtu = opts[:mtu]
    self.name = opts[:name]
    self.network_labels = opts[:network_labels]
    self.permissions = opts[:permissions]
    self.profile_required = opts[:profile_required]
    self.qos = opts[:qos]
    self.required = opts[:required]
    self.status = opts[:status]
    self.stp = opts[:stp]
    self.usages = opts[:usages]
    self.vlan = opts[:vlan]
    self.vnic_profiles = opts[:vnic_profiles]
  end
  
end

class NetworkAttachment < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `host_nic` attribute.
  # 
  def host_nic
    return @host_nic
  end
  
  # 
  # Sets the value of the `host_nic` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNic} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host_nic=(value)
    if value.is_a?(Hash)
      value = HostNic.new(value)
    end
    @host_nic = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `in_sync` attribute.
  # 
  def in_sync
    return @in_sync
  end
  
  # 
  # Sets the value of the `in_sync` attribute.
  # 
  def in_sync=(value)
    @in_sync = value
  end
  
  # 
  # Returns the value of the `ip_address_assignments` attribute.
  # 
  def ip_address_assignments
    return @ip_address_assignments
  end
  
  # 
  # Sets the value of the `ip_address_assignments` attribute.
  # 
  def ip_address_assignments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = IpAddressAssignment.new(value)
        end
      end
    end
    @ip_address_assignments = list
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network` attribute.
  # 
  def network
    return @network
  end
  
  # 
  # Sets the value of the `network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Network} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network=(value)
    if value.is_a?(Hash)
      value = Network.new(value)
    end
    @network = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `qos` attribute.
  # 
  def qos
    return @qos
  end
  
  # 
  # Sets the value of the `qos` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Qos} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def qos=(value)
    if value.is_a?(Hash)
      value = Qos.new(value)
    end
    @qos = value
  end
  
  # 
  # Returns the value of the `reported_configurations` attribute.
  # 
  def reported_configurations
    return @reported_configurations
  end
  
  # 
  # Sets the value of the `reported_configurations` attribute.
  # 
  def reported_configurations=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ReportedConfiguration.new(value)
        end
      end
    end
    @reported_configurations = list
  end
  
  # 
  # Creates a new instance of the {NetworkAttachment} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts [Hash] :host_nic The value of attribute `host_nic`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :in_sync The value of attribute `in_sync`.
  # 
  # @option opts [Array<Hash>] :ip_address_assignments The values of attribute `ip_address_assignments`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :network The value of attribute `network`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts [Hash] :qos The value of attribute `qos`.
  # 
  # @option opts [Array<Hash>] :reported_configurations The values of attribute `reported_configurations`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.host_nic = opts[:host_nic]
    self.id = opts[:id]
    self.in_sync = opts[:in_sync]
    self.ip_address_assignments = opts[:ip_address_assignments]
    self.name = opts[:name]
    self.network = opts[:network]
    self.properties = opts[:properties]
    self.qos = opts[:qos]
    self.reported_configurations = opts[:reported_configurations]
  end
  
end

class NetworkConfiguration < Struct
  
  # 
  # Returns the value of the `dns` attribute.
  # 
  def dns
    return @dns
  end
  
  # 
  # Sets the value of the `dns` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Dns} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def dns=(value)
    if value.is_a?(Hash)
      value = Dns.new(value)
    end
    @dns = value
  end
  
  # 
  # Returns the value of the `nics` attribute.
  # 
  def nics
    return @nics
  end
  
  # 
  # Sets the value of the `nics` attribute.
  # 
  def nics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Nic.new(value)
        end
      end
    end
    @nics = list
  end
  
  # 
  # Creates a new instance of the {NetworkConfiguration} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :dns The value of attribute `dns`.
  # 
  # @option opts [Array<Hash>] :nics The values of attribute `nics`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.dns = opts[:dns]
    self.nics = opts[:nics]
  end
  
end

class NetworkFilter < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {NetworkFilter} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.version = opts[:version]
  end
  
end

class NetworkLabel < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host_nic` attribute.
  # 
  def host_nic
    return @host_nic
  end
  
  # 
  # Sets the value of the `host_nic` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNic} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host_nic=(value)
    if value.is_a?(Hash)
      value = HostNic.new(value)
    end
    @host_nic = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network` attribute.
  # 
  def network
    return @network
  end
  
  # 
  # Sets the value of the `network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Network} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network=(value)
    if value.is_a?(Hash)
      value = Network.new(value)
    end
    @network = value
  end
  
  # 
  # Creates a new instance of the {NetworkLabel} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host_nic The value of attribute `host_nic`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :network The value of attribute `network`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host_nic = opts[:host_nic]
    self.id = opts[:id]
    self.name = opts[:name]
    self.network = opts[:network]
  end
  
end

class NfsProfileDetail < EntityProfileDetail
  
  # 
  # Returns the value of the `nfs_server_ip` attribute.
  # 
  def nfs_server_ip
    return @nfs_server_ip
  end
  
  # 
  # Sets the value of the `nfs_server_ip` attribute.
  # 
  def nfs_server_ip=(value)
    @nfs_server_ip = value
  end
  
  # 
  # Returns the value of the `profile_details` attribute.
  # 
  def profile_details
    return @profile_details
  end
  
  # 
  # Sets the value of the `profile_details` attribute.
  # 
  def profile_details=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ProfileDetail.new(value)
        end
      end
    end
    @profile_details = list
  end
  
  # 
  # Creates a new instance of the {NfsProfileDetail} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :nfs_server_ip The value of attribute `nfs_server_ip`.
  # 
  # @option opts [Array<Hash>] :profile_details The values of attribute `profile_details`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.nfs_server_ip = opts[:nfs_server_ip]
    self.profile_details = opts[:profile_details]
  end
  
end

class NicConfiguration < Struct
  
  # 
  # Returns the value of the `boot_protocol` attribute.
  # 
  def boot_protocol
    return @boot_protocol
  end
  
  # 
  # Sets the value of the `boot_protocol` attribute.
  # 
  def boot_protocol=(value)
    @boot_protocol = value
  end
  
  # 
  # Returns the value of the `ip` attribute.
  # 
  def ip
    return @ip
  end
  
  # 
  # Sets the value of the `ip` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ip} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ip=(value)
    if value.is_a?(Hash)
      value = Ip.new(value)
    end
    @ip = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `on_boot` attribute.
  # 
  def on_boot
    return @on_boot
  end
  
  # 
  # Sets the value of the `on_boot` attribute.
  # 
  def on_boot=(value)
    @on_boot = value
  end
  
  # 
  # Creates a new instance of the {NicConfiguration} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :boot_protocol The value of attribute `boot_protocol`.
  # 
  # @option opts [Hash] :ip The value of attribute `ip`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :on_boot The value of attribute `on_boot`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.boot_protocol = opts[:boot_protocol]
    self.ip = opts[:ip]
    self.name = opts[:name]
    self.on_boot = opts[:on_boot]
  end
  
end

class NumaNode < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `index` attribute.
  # 
  def index
    return @index
  end
  
  # 
  # Sets the value of the `index` attribute.
  # 
  def index=(value)
    @index = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `node_distance` attribute.
  # 
  def node_distance
    return @node_distance
  end
  
  # 
  # Sets the value of the `node_distance` attribute.
  # 
  def node_distance=(value)
    @node_distance = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Creates a new instance of the {NumaNode} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :index The value of attribute `index`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :node_distance The value of attribute `node_distance`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.cpu = opts[:cpu]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.index = opts[:index]
    self.memory = opts[:memory]
    self.name = opts[:name]
    self.node_distance = opts[:node_distance]
    self.statistics = opts[:statistics]
  end
  
end

class NumaNodePin < Struct
  
  # 
  # Returns the value of the `host_numa_node` attribute.
  # 
  def host_numa_node
    return @host_numa_node
  end
  
  # 
  # Sets the value of the `host_numa_node` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::NumaNode} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host_numa_node=(value)
    if value.is_a?(Hash)
      value = NumaNode.new(value)
    end
    @host_numa_node = value
  end
  
  # 
  # Returns the value of the `index` attribute.
  # 
  def index
    return @index
  end
  
  # 
  # Sets the value of the `index` attribute.
  # 
  def index=(value)
    @index = value
  end
  
  # 
  # Returns the value of the `pinned` attribute.
  # 
  def pinned
    return @pinned
  end
  
  # 
  # Sets the value of the `pinned` attribute.
  # 
  def pinned=(value)
    @pinned = value
  end
  
  # 
  # Creates a new instance of the {NumaNodePin} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :host_numa_node The value of attribute `host_numa_node`.
  # 
  # @option opts :index The value of attribute `index`.
  # 
  # @option opts :pinned The value of attribute `pinned`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.host_numa_node = opts[:host_numa_node]
    self.index = opts[:index]
    self.pinned = opts[:pinned]
  end
  
end

class OpenStackImage < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_image_provider` attribute.
  # 
  def openstack_image_provider
    return @openstack_image_provider
  end
  
  # 
  # Sets the value of the `openstack_image_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackImageProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_image_provider=(value)
    if value.is_a?(Hash)
      value = OpenStackImageProvider.new(value)
    end
    @openstack_image_provider = value
  end
  
  # 
  # Creates a new instance of the {OpenStackImage} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_image_provider The value of attribute `openstack_image_provider`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.openstack_image_provider = opts[:openstack_image_provider]
  end
  
end

class OpenStackNetwork < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_network_provider` attribute.
  # 
  def openstack_network_provider
    return @openstack_network_provider
  end
  
  # 
  # Sets the value of the `openstack_network_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackNetworkProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_network_provider=(value)
    if value.is_a?(Hash)
      value = OpenStackNetworkProvider.new(value)
    end
    @openstack_network_provider = value
  end
  
  # 
  # Creates a new instance of the {OpenStackNetwork} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_network_provider The value of attribute `openstack_network_provider`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.openstack_network_provider = opts[:openstack_network_provider]
  end
  
end

class OpenStackSubnet < Identified
  
  # 
  # Returns the value of the `cidr` attribute.
  # 
  def cidr
    return @cidr
  end
  
  # 
  # Sets the value of the `cidr` attribute.
  # 
  def cidr=(value)
    @cidr = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `dns_servers` attribute.
  # 
  def dns_servers
    return @dns_servers
  end
  
  # 
  # Sets the value of the `dns_servers` attribute.
  # 
  def dns_servers=(list)
    @dns_servers = list
  end
  
  # 
  # Returns the value of the `gateway` attribute.
  # 
  def gateway
    return @gateway
  end
  
  # 
  # Sets the value of the `gateway` attribute.
  # 
  def gateway=(value)
    @gateway = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ip_version` attribute.
  # 
  def ip_version
    return @ip_version
  end
  
  # 
  # Sets the value of the `ip_version` attribute.
  # 
  def ip_version=(value)
    @ip_version = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_network` attribute.
  # 
  def openstack_network
    return @openstack_network
  end
  
  # 
  # Sets the value of the `openstack_network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackNetwork} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_network=(value)
    if value.is_a?(Hash)
      value = OpenStackNetwork.new(value)
    end
    @openstack_network = value
  end
  
  # 
  # Creates a new instance of the {OpenStackSubnet} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :cidr The value of attribute `cidr`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :dns_servers The values of attribute `dns_servers`.
  # 
  # @option opts :gateway The value of attribute `gateway`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :ip_version The value of attribute `ip_version`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_network The value of attribute `openstack_network`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cidr = opts[:cidr]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.dns_servers = opts[:dns_servers]
    self.gateway = opts[:gateway]
    self.id = opts[:id]
    self.ip_version = opts[:ip_version]
    self.name = opts[:name]
    self.openstack_network = opts[:openstack_network]
  end
  
end

class OpenStackVolumeType < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_volume_provider` attribute.
  # 
  def openstack_volume_provider
    return @openstack_volume_provider
  end
  
  # 
  # Sets the value of the `openstack_volume_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackVolumeProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_volume_provider=(value)
    if value.is_a?(Hash)
      value = OpenStackVolumeProvider.new(value)
    end
    @openstack_volume_provider = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Creates a new instance of the {OpenStackVolumeType} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_volume_provider The value of attribute `openstack_volume_provider`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.openstack_volume_provider = opts[:openstack_volume_provider]
    self.properties = opts[:properties]
  end
  
end

class OpenstackVolumeAuthenticationKey < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `creation_date` attribute.
  # 
  def creation_date
    return @creation_date
  end
  
  # 
  # Sets the value of the `creation_date` attribute.
  # 
  def creation_date=(value)
    @creation_date = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_volume_provider` attribute.
  # 
  def openstack_volume_provider
    return @openstack_volume_provider
  end
  
  # 
  # Sets the value of the `openstack_volume_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackVolumeProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_volume_provider=(value)
    if value.is_a?(Hash)
      value = OpenStackVolumeProvider.new(value)
    end
    @openstack_volume_provider = value
  end
  
  # 
  # Returns the value of the `usage_type` attribute.
  # 
  def usage_type
    return @usage_type
  end
  
  # 
  # Sets the value of the `usage_type` attribute.
  # 
  def usage_type=(value)
    @usage_type = value
  end
  
  # 
  # Returns the value of the `uuid` attribute.
  # 
  def uuid
    return @uuid
  end
  
  # 
  # Sets the value of the `uuid` attribute.
  # 
  def uuid=(value)
    @uuid = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {OpenstackVolumeAuthenticationKey} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :creation_date The value of attribute `creation_date`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_volume_provider The value of attribute `openstack_volume_provider`.
  # 
  # @option opts :usage_type The value of attribute `usage_type`.
  # 
  # @option opts :uuid The value of attribute `uuid`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.creation_date = opts[:creation_date]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.openstack_volume_provider = opts[:openstack_volume_provider]
    self.usage_type = opts[:usage_type]
    self.uuid = opts[:uuid]
    self.value = opts[:value]
  end
  
end

class OperatingSystem < Struct
  
  # 
  # Returns the value of the `boot` attribute.
  # 
  def boot
    return @boot
  end
  
  # 
  # Sets the value of the `boot` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Boot} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def boot=(value)
    if value.is_a?(Hash)
      value = Boot.new(value)
    end
    @boot = value
  end
  
  # 
  # Returns the value of the `cmdline` attribute.
  # 
  def cmdline
    return @cmdline
  end
  
  # 
  # Sets the value of the `cmdline` attribute.
  # 
  def cmdline=(value)
    @cmdline = value
  end
  
  # 
  # Returns the value of the `custom_kernel_cmdline` attribute.
  # 
  def custom_kernel_cmdline
    return @custom_kernel_cmdline
  end
  
  # 
  # Sets the value of the `custom_kernel_cmdline` attribute.
  # 
  def custom_kernel_cmdline=(value)
    @custom_kernel_cmdline = value
  end
  
  # 
  # Returns the value of the `initrd` attribute.
  # 
  def initrd
    return @initrd
  end
  
  # 
  # Sets the value of the `initrd` attribute.
  # 
  def initrd=(value)
    @initrd = value
  end
  
  # 
  # Returns the value of the `kernel` attribute.
  # 
  def kernel
    return @kernel
  end
  
  # 
  # Sets the value of the `kernel` attribute.
  # 
  def kernel=(value)
    @kernel = value
  end
  
  # 
  # Returns the value of the `reported_kernel_cmdline` attribute.
  # 
  def reported_kernel_cmdline
    return @reported_kernel_cmdline
  end
  
  # 
  # Sets the value of the `reported_kernel_cmdline` attribute.
  # 
  def reported_kernel_cmdline=(value)
    @reported_kernel_cmdline = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {OperatingSystem} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :boot The value of attribute `boot`.
  # 
  # @option opts :cmdline The value of attribute `cmdline`.
  # 
  # @option opts :custom_kernel_cmdline The value of attribute `custom_kernel_cmdline`.
  # 
  # @option opts :initrd The value of attribute `initrd`.
  # 
  # @option opts :kernel The value of attribute `kernel`.
  # 
  # @option opts :reported_kernel_cmdline The value of attribute `reported_kernel_cmdline`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.boot = opts[:boot]
    self.cmdline = opts[:cmdline]
    self.custom_kernel_cmdline = opts[:custom_kernel_cmdline]
    self.initrd = opts[:initrd]
    self.kernel = opts[:kernel]
    self.reported_kernel_cmdline = opts[:reported_kernel_cmdline]
    self.type = opts[:type]
    self.version = opts[:version]
  end
  
end

class OperatingSystemInfo < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `large_icon` attribute.
  # 
  def large_icon
    return @large_icon
  end
  
  # 
  # Sets the value of the `large_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def large_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @large_icon = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `small_icon` attribute.
  # 
  def small_icon
    return @small_icon
  end
  
  # 
  # Sets the value of the `small_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def small_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @small_icon = value
  end
  
  # 
  # Creates a new instance of the {OperatingSystemInfo} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :large_icon The value of attribute `large_icon`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :small_icon The value of attribute `small_icon`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.large_icon = opts[:large_icon]
    self.name = opts[:name]
    self.small_icon = opts[:small_icon]
  end
  
end

class Option < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {Option} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
    self.type = opts[:type]
    self.value = opts[:value]
  end
  
end

class Package < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {Package} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
  end
  
end

class Payload < Struct
  
  # 
  # Returns the value of the `files` attribute.
  # 
  def files
    return @files
  end
  
  # 
  # Sets the value of the `files` attribute.
  # 
  def files=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = File.new(value)
        end
      end
    end
    @files = list
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `volume_id` attribute.
  # 
  def volume_id
    return @volume_id
  end
  
  # 
  # Sets the value of the `volume_id` attribute.
  # 
  def volume_id=(value)
    @volume_id = value
  end
  
  # 
  # Creates a new instance of the {Payload} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :files The values of attribute `files`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :volume_id The value of attribute `volume_id`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.files = opts[:files]
    self.type = opts[:type]
    self.volume_id = opts[:volume_id]
  end
  
end

class Permission < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk` attribute.
  # 
  def disk
    return @disk
  end
  
  # 
  # Sets the value of the `disk` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Disk} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk=(value)
    if value.is_a?(Hash)
      value = Disk.new(value)
    end
    @disk = value
  end
  
  # 
  # Returns the value of the `group` attribute.
  # 
  def group
    return @group
  end
  
  # 
  # Sets the value of the `group` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Group} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def group=(value)
    if value.is_a?(Hash)
      value = Group.new(value)
    end
    @group = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `role` attribute.
  # 
  def role
    return @role
  end
  
  # 
  # Sets the value of the `role` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Role} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def role=(value)
    if value.is_a?(Hash)
      value = Role.new(value)
    end
    @role = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vm_pool` attribute.
  # 
  def vm_pool
    return @vm_pool
  end
  
  # 
  # Sets the value of the `vm_pool` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VmPool} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm_pool=(value)
    if value.is_a?(Hash)
      value = VmPool.new(value)
    end
    @vm_pool = value
  end
  
  # 
  # Creates a new instance of the {Permission} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :disk The value of attribute `disk`.
  # 
  # @option opts [Hash] :group The value of attribute `group`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :role The value of attribute `role`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Hash] :vm_pool The value of attribute `vm_pool`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.disk = opts[:disk]
    self.group = opts[:group]
    self.host = opts[:host]
    self.id = opts[:id]
    self.name = opts[:name]
    self.role = opts[:role]
    self.storage_domain = opts[:storage_domain]
    self.template = opts[:template]
    self.user = opts[:user]
    self.vm = opts[:vm]
    self.vm_pool = opts[:vm_pool]
  end
  
end

class Permit < Identified
  
  # 
  # Returns the value of the `administrative` attribute.
  # 
  def administrative
    return @administrative
  end
  
  # 
  # Sets the value of the `administrative` attribute.
  # 
  def administrative=(value)
    @administrative = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `role` attribute.
  # 
  def role
    return @role
  end
  
  # 
  # Sets the value of the `role` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Role} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def role=(value)
    if value.is_a?(Hash)
      value = Role.new(value)
    end
    @role = value
  end
  
  # 
  # Creates a new instance of the {Permit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :administrative The value of attribute `administrative`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :role The value of attribute `role`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.administrative = opts[:administrative]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.role = opts[:role]
  end
  
end

class PmProxy < Struct
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {PmProxy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.type = opts[:type]
  end
  
end

class PortMirroring < Struct
  
  # 
  # Creates a new instance of the {PortMirroring} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # 
  def initialize(opts = {})
    super(opts)
  end
  
end

class PowerManagement < Struct
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `agents` attribute.
  # 
  def agents
    return @agents
  end
  
  # 
  # Sets the value of the `agents` attribute.
  # 
  def agents=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Agent.new(value)
        end
      end
    end
    @agents = list
  end
  
  # 
  # Returns the value of the `automatic_pm_enabled` attribute.
  # 
  def automatic_pm_enabled
    return @automatic_pm_enabled
  end
  
  # 
  # Sets the value of the `automatic_pm_enabled` attribute.
  # 
  def automatic_pm_enabled=(value)
    @automatic_pm_enabled = value
  end
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `kdump_detection` attribute.
  # 
  def kdump_detection
    return @kdump_detection
  end
  
  # 
  # Sets the value of the `kdump_detection` attribute.
  # 
  def kdump_detection=(value)
    @kdump_detection = value
  end
  
  # 
  # Returns the value of the `options` attribute.
  # 
  def options
    return @options
  end
  
  # 
  # Sets the value of the `options` attribute.
  # 
  def options=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Option.new(value)
        end
      end
    end
    @options = list
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `pm_proxies` attribute.
  # 
  def pm_proxies
    return @pm_proxies
  end
  
  # 
  # Sets the value of the `pm_proxies` attribute.
  # 
  def pm_proxies=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = PmProxy.new(value)
        end
      end
    end
    @pm_proxies = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {PowerManagement} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts [Array<Hash>] :agents The values of attribute `agents`.
  # 
  # @option opts :automatic_pm_enabled The value of attribute `automatic_pm_enabled`.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts :kdump_detection The value of attribute `kdump_detection`.
  # 
  # @option opts [Array<Hash>] :options The values of attribute `options`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :pm_proxies The values of attribute `pm_proxies`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.agents = opts[:agents]
    self.automatic_pm_enabled = opts[:automatic_pm_enabled]
    self.enabled = opts[:enabled]
    self.kdump_detection = opts[:kdump_detection]
    self.options = opts[:options]
    self.password = opts[:password]
    self.pm_proxies = opts[:pm_proxies]
    self.status = opts[:status]
    self.type = opts[:type]
    self.username = opts[:username]
  end
  
end

class Product < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {Product} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class ProductInfo < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `vendor` attribute.
  # 
  def vendor
    return @vendor
  end
  
  # 
  # Sets the value of the `vendor` attribute.
  # 
  def vendor=(value)
    @vendor = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {ProductInfo} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :vendor The value of attribute `vendor`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
    self.vendor = opts[:vendor]
    self.version = opts[:version]
  end
  
end

class ProfileDetail < Struct
  
  # 
  # Returns the value of the `block_statistics` attribute.
  # 
  def block_statistics
    return @block_statistics
  end
  
  # 
  # Sets the value of the `block_statistics` attribute.
  # 
  def block_statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = BlockStatistic.new(value)
        end
      end
    end
    @block_statistics = list
  end
  
  # 
  # Returns the value of the `duration` attribute.
  # 
  def duration
    return @duration
  end
  
  # 
  # Sets the value of the `duration` attribute.
  # 
  def duration=(value)
    @duration = value
  end
  
  # 
  # Returns the value of the `fop_statistics` attribute.
  # 
  def fop_statistics
    return @fop_statistics
  end
  
  # 
  # Sets the value of the `fop_statistics` attribute.
  # 
  def fop_statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = FopStatistic.new(value)
        end
      end
    end
    @fop_statistics = list
  end
  
  # 
  # Returns the value of the `profile_type` attribute.
  # 
  def profile_type
    return @profile_type
  end
  
  # 
  # Sets the value of the `profile_type` attribute.
  # 
  def profile_type=(value)
    @profile_type = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Creates a new instance of the {ProfileDetail} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :block_statistics The values of attribute `block_statistics`.
  # 
  # @option opts :duration The value of attribute `duration`.
  # 
  # @option opts [Array<Hash>] :fop_statistics The values of attribute `fop_statistics`.
  # 
  # @option opts :profile_type The value of attribute `profile_type`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.block_statistics = opts[:block_statistics]
    self.duration = opts[:duration]
    self.fop_statistics = opts[:fop_statistics]
    self.profile_type = opts[:profile_type]
    self.statistics = opts[:statistics]
  end
  
end

class Property < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {Property} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
    self.value = opts[:value]
  end
  
end

class ProxyTicket < Struct
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {ProxyTicket} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.value = opts[:value]
  end
  
end

class Qos < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `cpu_limit` attribute.
  # 
  def cpu_limit
    return @cpu_limit
  end
  
  # 
  # Sets the value of the `cpu_limit` attribute.
  # 
  def cpu_limit=(value)
    @cpu_limit = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `inbound_average` attribute.
  # 
  def inbound_average
    return @inbound_average
  end
  
  # 
  # Sets the value of the `inbound_average` attribute.
  # 
  def inbound_average=(value)
    @inbound_average = value
  end
  
  # 
  # Returns the value of the `inbound_burst` attribute.
  # 
  def inbound_burst
    return @inbound_burst
  end
  
  # 
  # Sets the value of the `inbound_burst` attribute.
  # 
  def inbound_burst=(value)
    @inbound_burst = value
  end
  
  # 
  # Returns the value of the `inbound_peak` attribute.
  # 
  def inbound_peak
    return @inbound_peak
  end
  
  # 
  # Sets the value of the `inbound_peak` attribute.
  # 
  def inbound_peak=(value)
    @inbound_peak = value
  end
  
  # 
  # Returns the value of the `max_iops` attribute.
  # 
  def max_iops
    return @max_iops
  end
  
  # 
  # Sets the value of the `max_iops` attribute.
  # 
  def max_iops=(value)
    @max_iops = value
  end
  
  # 
  # Returns the value of the `max_read_iops` attribute.
  # 
  def max_read_iops
    return @max_read_iops
  end
  
  # 
  # Sets the value of the `max_read_iops` attribute.
  # 
  def max_read_iops=(value)
    @max_read_iops = value
  end
  
  # 
  # Returns the value of the `max_read_throughput` attribute.
  # 
  def max_read_throughput
    return @max_read_throughput
  end
  
  # 
  # Sets the value of the `max_read_throughput` attribute.
  # 
  def max_read_throughput=(value)
    @max_read_throughput = value
  end
  
  # 
  # Returns the value of the `max_throughput` attribute.
  # 
  def max_throughput
    return @max_throughput
  end
  
  # 
  # Sets the value of the `max_throughput` attribute.
  # 
  def max_throughput=(value)
    @max_throughput = value
  end
  
  # 
  # Returns the value of the `max_write_iops` attribute.
  # 
  def max_write_iops
    return @max_write_iops
  end
  
  # 
  # Sets the value of the `max_write_iops` attribute.
  # 
  def max_write_iops=(value)
    @max_write_iops = value
  end
  
  # 
  # Returns the value of the `max_write_throughput` attribute.
  # 
  def max_write_throughput
    return @max_write_throughput
  end
  
  # 
  # Sets the value of the `max_write_throughput` attribute.
  # 
  def max_write_throughput=(value)
    @max_write_throughput = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `outbound_average` attribute.
  # 
  def outbound_average
    return @outbound_average
  end
  
  # 
  # Sets the value of the `outbound_average` attribute.
  # 
  def outbound_average=(value)
    @outbound_average = value
  end
  
  # 
  # Returns the value of the `outbound_average_linkshare` attribute.
  # 
  def outbound_average_linkshare
    return @outbound_average_linkshare
  end
  
  # 
  # Sets the value of the `outbound_average_linkshare` attribute.
  # 
  def outbound_average_linkshare=(value)
    @outbound_average_linkshare = value
  end
  
  # 
  # Returns the value of the `outbound_average_realtime` attribute.
  # 
  def outbound_average_realtime
    return @outbound_average_realtime
  end
  
  # 
  # Sets the value of the `outbound_average_realtime` attribute.
  # 
  def outbound_average_realtime=(value)
    @outbound_average_realtime = value
  end
  
  # 
  # Returns the value of the `outbound_average_upperlimit` attribute.
  # 
  def outbound_average_upperlimit
    return @outbound_average_upperlimit
  end
  
  # 
  # Sets the value of the `outbound_average_upperlimit` attribute.
  # 
  def outbound_average_upperlimit=(value)
    @outbound_average_upperlimit = value
  end
  
  # 
  # Returns the value of the `outbound_burst` attribute.
  # 
  def outbound_burst
    return @outbound_burst
  end
  
  # 
  # Sets the value of the `outbound_burst` attribute.
  # 
  def outbound_burst=(value)
    @outbound_burst = value
  end
  
  # 
  # Returns the value of the `outbound_peak` attribute.
  # 
  def outbound_peak
    return @outbound_peak
  end
  
  # 
  # Sets the value of the `outbound_peak` attribute.
  # 
  def outbound_peak=(value)
    @outbound_peak = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {Qos} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :cpu_limit The value of attribute `cpu_limit`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :inbound_average The value of attribute `inbound_average`.
  # 
  # @option opts :inbound_burst The value of attribute `inbound_burst`.
  # 
  # @option opts :inbound_peak The value of attribute `inbound_peak`.
  # 
  # @option opts :max_iops The value of attribute `max_iops`.
  # 
  # @option opts :max_read_iops The value of attribute `max_read_iops`.
  # 
  # @option opts :max_read_throughput The value of attribute `max_read_throughput`.
  # 
  # @option opts :max_throughput The value of attribute `max_throughput`.
  # 
  # @option opts :max_write_iops The value of attribute `max_write_iops`.
  # 
  # @option opts :max_write_throughput The value of attribute `max_write_throughput`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :outbound_average The value of attribute `outbound_average`.
  # 
  # @option opts :outbound_average_linkshare The value of attribute `outbound_average_linkshare`.
  # 
  # @option opts :outbound_average_realtime The value of attribute `outbound_average_realtime`.
  # 
  # @option opts :outbound_average_upperlimit The value of attribute `outbound_average_upperlimit`.
  # 
  # @option opts :outbound_burst The value of attribute `outbound_burst`.
  # 
  # @option opts :outbound_peak The value of attribute `outbound_peak`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.cpu_limit = opts[:cpu_limit]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.id = opts[:id]
    self.inbound_average = opts[:inbound_average]
    self.inbound_burst = opts[:inbound_burst]
    self.inbound_peak = opts[:inbound_peak]
    self.max_iops = opts[:max_iops]
    self.max_read_iops = opts[:max_read_iops]
    self.max_read_throughput = opts[:max_read_throughput]
    self.max_throughput = opts[:max_throughput]
    self.max_write_iops = opts[:max_write_iops]
    self.max_write_throughput = opts[:max_write_throughput]
    self.name = opts[:name]
    self.outbound_average = opts[:outbound_average]
    self.outbound_average_linkshare = opts[:outbound_average_linkshare]
    self.outbound_average_realtime = opts[:outbound_average_realtime]
    self.outbound_average_upperlimit = opts[:outbound_average_upperlimit]
    self.outbound_burst = opts[:outbound_burst]
    self.outbound_peak = opts[:outbound_peak]
    self.type = opts[:type]
  end
  
end

class Quota < Identified
  
  # 
  # Returns the value of the `cluster_hard_limit_pct` attribute.
  # 
  def cluster_hard_limit_pct
    return @cluster_hard_limit_pct
  end
  
  # 
  # Sets the value of the `cluster_hard_limit_pct` attribute.
  # 
  def cluster_hard_limit_pct=(value)
    @cluster_hard_limit_pct = value
  end
  
  # 
  # Returns the value of the `cluster_soft_limit_pct` attribute.
  # 
  def cluster_soft_limit_pct
    return @cluster_soft_limit_pct
  end
  
  # 
  # Sets the value of the `cluster_soft_limit_pct` attribute.
  # 
  def cluster_soft_limit_pct=(value)
    @cluster_soft_limit_pct = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disks` attribute.
  # 
  def disks
    return @disks
  end
  
  # 
  # Sets the value of the `disks` attribute.
  # 
  def disks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Disk.new(value)
        end
      end
    end
    @disks = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `quota_cluster_limits` attribute.
  # 
  def quota_cluster_limits
    return @quota_cluster_limits
  end
  
  # 
  # Sets the value of the `quota_cluster_limits` attribute.
  # 
  def quota_cluster_limits=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = QuotaClusterLimit.new(value)
        end
      end
    end
    @quota_cluster_limits = list
  end
  
  # 
  # Returns the value of the `quota_storage_limits` attribute.
  # 
  def quota_storage_limits
    return @quota_storage_limits
  end
  
  # 
  # Sets the value of the `quota_storage_limits` attribute.
  # 
  def quota_storage_limits=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = QuotaStorageLimit.new(value)
        end
      end
    end
    @quota_storage_limits = list
  end
  
  # 
  # Returns the value of the `storage_hard_limit_pct` attribute.
  # 
  def storage_hard_limit_pct
    return @storage_hard_limit_pct
  end
  
  # 
  # Sets the value of the `storage_hard_limit_pct` attribute.
  # 
  def storage_hard_limit_pct=(value)
    @storage_hard_limit_pct = value
  end
  
  # 
  # Returns the value of the `storage_soft_limit_pct` attribute.
  # 
  def storage_soft_limit_pct
    return @storage_soft_limit_pct
  end
  
  # 
  # Sets the value of the `storage_soft_limit_pct` attribute.
  # 
  def storage_soft_limit_pct=(value)
    @storage_soft_limit_pct = value
  end
  
  # 
  # Returns the value of the `users` attribute.
  # 
  def users
    return @users
  end
  
  # 
  # Sets the value of the `users` attribute.
  # 
  def users=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = User.new(value)
        end
      end
    end
    @users = list
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {Quota} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :cluster_hard_limit_pct The value of attribute `cluster_hard_limit_pct`.
  # 
  # @option opts :cluster_soft_limit_pct The value of attribute `cluster_soft_limit_pct`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :disks The values of attribute `disks`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Array<Hash>] :quota_cluster_limits The values of attribute `quota_cluster_limits`.
  # 
  # @option opts [Array<Hash>] :quota_storage_limits The values of attribute `quota_storage_limits`.
  # 
  # @option opts :storage_hard_limit_pct The value of attribute `storage_hard_limit_pct`.
  # 
  # @option opts :storage_soft_limit_pct The value of attribute `storage_soft_limit_pct`.
  # 
  # @option opts [Array<Hash>] :users The values of attribute `users`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster_hard_limit_pct = opts[:cluster_hard_limit_pct]
    self.cluster_soft_limit_pct = opts[:cluster_soft_limit_pct]
    self.comment = opts[:comment]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.disks = opts[:disks]
    self.id = opts[:id]
    self.name = opts[:name]
    self.permissions = opts[:permissions]
    self.quota_cluster_limits = opts[:quota_cluster_limits]
    self.quota_storage_limits = opts[:quota_storage_limits]
    self.storage_hard_limit_pct = opts[:storage_hard_limit_pct]
    self.storage_soft_limit_pct = opts[:storage_soft_limit_pct]
    self.users = opts[:users]
    self.vms = opts[:vms]
  end
  
end

class QuotaClusterLimit < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `memory_limit` attribute.
  # 
  def memory_limit
    return @memory_limit
  end
  
  # 
  # Sets the value of the `memory_limit` attribute.
  # 
  def memory_limit=(value)
    @memory_limit = value
  end
  
  # 
  # Returns the value of the `memory_usage` attribute.
  # 
  def memory_usage
    return @memory_usage
  end
  
  # 
  # Sets the value of the `memory_usage` attribute.
  # 
  def memory_usage=(value)
    @memory_usage = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `quota` attribute.
  # 
  def quota
    return @quota
  end
  
  # 
  # Sets the value of the `quota` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Quota} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def quota=(value)
    if value.is_a?(Hash)
      value = Quota.new(value)
    end
    @quota = value
  end
  
  # 
  # Returns the value of the `vcpu_limit` attribute.
  # 
  def vcpu_limit
    return @vcpu_limit
  end
  
  # 
  # Sets the value of the `vcpu_limit` attribute.
  # 
  def vcpu_limit=(value)
    @vcpu_limit = value
  end
  
  # 
  # Returns the value of the `vcpu_usage` attribute.
  # 
  def vcpu_usage
    return @vcpu_usage
  end
  
  # 
  # Sets the value of the `vcpu_usage` attribute.
  # 
  def vcpu_usage=(value)
    @vcpu_usage = value
  end
  
  # 
  # Creates a new instance of the {QuotaClusterLimit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :memory_limit The value of attribute `memory_limit`.
  # 
  # @option opts :memory_usage The value of attribute `memory_usage`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :quota The value of attribute `quota`.
  # 
  # @option opts :vcpu_limit The value of attribute `vcpu_limit`.
  # 
  # @option opts :vcpu_usage The value of attribute `vcpu_usage`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.memory_limit = opts[:memory_limit]
    self.memory_usage = opts[:memory_usage]
    self.name = opts[:name]
    self.quota = opts[:quota]
    self.vcpu_limit = opts[:vcpu_limit]
    self.vcpu_usage = opts[:vcpu_usage]
  end
  
end

class QuotaStorageLimit < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `limit` attribute.
  # 
  def limit
    return @limit
  end
  
  # 
  # Sets the value of the `limit` attribute.
  # 
  def limit=(value)
    @limit = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `quota` attribute.
  # 
  def quota
    return @quota
  end
  
  # 
  # Sets the value of the `quota` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Quota} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def quota=(value)
    if value.is_a?(Hash)
      value = Quota.new(value)
    end
    @quota = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `usage` attribute.
  # 
  def usage
    return @usage
  end
  
  # 
  # Sets the value of the `usage` attribute.
  # 
  def usage=(value)
    @usage = value
  end
  
  # 
  # Creates a new instance of the {QuotaStorageLimit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :limit The value of attribute `limit`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :quota The value of attribute `quota`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts :usage The value of attribute `usage`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.limit = opts[:limit]
    self.name = opts[:name]
    self.quota = opts[:quota]
    self.storage_domain = opts[:storage_domain]
    self.usage = opts[:usage]
  end
  
end

class Range < Struct
  
  # 
  # Returns the value of the `from` attribute.
  # 
  def from
    return @from
  end
  
  # 
  # Sets the value of the `from` attribute.
  # 
  def from=(value)
    @from = value
  end
  
  # 
  # Returns the value of the `to` attribute.
  # 
  def to
    return @to
  end
  
  # 
  # Sets the value of the `to` attribute.
  # 
  def to=(value)
    @to = value
  end
  
  # 
  # Creates a new instance of the {Range} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :from The value of attribute `from`.
  # 
  # @option opts :to The value of attribute `to`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.from = opts[:from]
    self.to = opts[:to]
  end
  
end

class Rate < Struct
  
  # 
  # Returns the value of the `bytes` attribute.
  # 
  def bytes
    return @bytes
  end
  
  # 
  # Sets the value of the `bytes` attribute.
  # 
  def bytes=(value)
    @bytes = value
  end
  
  # 
  # Returns the value of the `period` attribute.
  # 
  def period
    return @period
  end
  
  # 
  # Sets the value of the `period` attribute.
  # 
  def period=(value)
    @period = value
  end
  
  # 
  # Creates a new instance of the {Rate} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :bytes The value of attribute `bytes`.
  # 
  # @option opts :period The value of attribute `period`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.bytes = opts[:bytes]
    self.period = opts[:period]
  end
  
end

class ReportedConfiguration < Struct
  
  # 
  # Returns the value of the `actual_value` attribute.
  # 
  def actual_value
    return @actual_value
  end
  
  # 
  # Sets the value of the `actual_value` attribute.
  # 
  def actual_value=(value)
    @actual_value = value
  end
  
  # 
  # Returns the value of the `expected_value` attribute.
  # 
  def expected_value
    return @expected_value
  end
  
  # 
  # Sets the value of the `expected_value` attribute.
  # 
  def expected_value=(value)
    @expected_value = value
  end
  
  # 
  # Returns the value of the `in_sync` attribute.
  # 
  def in_sync
    return @in_sync
  end
  
  # 
  # Sets the value of the `in_sync` attribute.
  # 
  def in_sync=(value)
    @in_sync = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {ReportedConfiguration} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :actual_value The value of attribute `actual_value`.
  # 
  # @option opts :expected_value The value of attribute `expected_value`.
  # 
  # @option opts :in_sync The value of attribute `in_sync`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.actual_value = opts[:actual_value]
    self.expected_value = opts[:expected_value]
    self.in_sync = opts[:in_sync]
    self.name = opts[:name]
  end
  
end

class ReportedDevice < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ips` attribute.
  # 
  def ips
    return @ips
  end
  
  # 
  # Sets the value of the `ips` attribute.
  # 
  def ips=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Ip.new(value)
        end
      end
    end
    @ips = list
  end
  
  # 
  # Returns the value of the `mac` attribute.
  # 
  def mac
    return @mac
  end
  
  # 
  # Sets the value of the `mac` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Mac} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def mac=(value)
    if value.is_a?(Hash)
      value = Mac.new(value)
    end
    @mac = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {ReportedDevice} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Array<Hash>] :ips The values of attribute `ips`.
  # 
  # @option opts [Hash] :mac The value of attribute `mac`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.ips = opts[:ips]
    self.mac = opts[:mac]
    self.name = opts[:name]
    self.type = opts[:type]
    self.vm = opts[:vm]
  end
  
end

class RngDevice < Struct
  
  # 
  # Returns the value of the `rate` attribute.
  # 
  def rate
    return @rate
  end
  
  # 
  # Sets the value of the `rate` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Rate} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rate=(value)
    if value.is_a?(Hash)
      value = Rate.new(value)
    end
    @rate = value
  end
  
  # 
  # Returns the value of the `source` attribute.
  # 
  def source
    return @source
  end
  
  # 
  # Sets the value of the `source` attribute.
  # 
  def source=(value)
    @source = value
  end
  
  # 
  # Creates a new instance of the {RngDevice} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :rate The value of attribute `rate`.
  # 
  # @option opts :source The value of attribute `source`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.rate = opts[:rate]
    self.source = opts[:source]
  end
  
end

class Role < Identified
  
  # 
  # Returns the value of the `administrative` attribute.
  # 
  def administrative
    return @administrative
  end
  
  # 
  # Sets the value of the `administrative` attribute.
  # 
  def administrative=(value)
    @administrative = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `mutable` attribute.
  # 
  def mutable
    return @mutable
  end
  
  # 
  # Sets the value of the `mutable` attribute.
  # 
  def mutable=(value)
    @mutable = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `permits` attribute.
  # 
  def permits
    return @permits
  end
  
  # 
  # Sets the value of the `permits` attribute.
  # 
  def permits=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permit.new(value)
        end
      end
    end
    @permits = list
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Creates a new instance of the {Role} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :administrative The value of attribute `administrative`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :mutable The value of attribute `mutable`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :permits The values of attribute `permits`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.administrative = opts[:administrative]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.mutable = opts[:mutable]
    self.name = opts[:name]
    self.permits = opts[:permits]
    self.user = opts[:user]
  end
  
end

class SchedulingPolicy < Identified
  
  # 
  # Returns the value of the `balances` attribute.
  # 
  def balances
    return @balances
  end
  
  # 
  # Sets the value of the `balances` attribute.
  # 
  def balances=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Balance.new(value)
        end
      end
    end
    @balances = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `default_policy` attribute.
  # 
  def default_policy
    return @default_policy
  end
  
  # 
  # Sets the value of the `default_policy` attribute.
  # 
  def default_policy=(value)
    @default_policy = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `filters` attribute.
  # 
  def filters
    return @filters
  end
  
  # 
  # Sets the value of the `filters` attribute.
  # 
  def filters=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Filter.new(value)
        end
      end
    end
    @filters = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `locked` attribute.
  # 
  def locked
    return @locked
  end
  
  # 
  # Sets the value of the `locked` attribute.
  # 
  def locked=(value)
    @locked = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `weight` attribute.
  # 
  def weight
    return @weight
  end
  
  # 
  # Sets the value of the `weight` attribute.
  # 
  def weight=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Weight.new(value)
        end
      end
    end
    @weight = list
  end
  
  # 
  # Creates a new instance of the {SchedulingPolicy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :balances The values of attribute `balances`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :default_policy The value of attribute `default_policy`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :filters The values of attribute `filters`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :locked The value of attribute `locked`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts [Array<Hash>] :weight The values of attribute `weight`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.balances = opts[:balances]
    self.comment = opts[:comment]
    self.default_policy = opts[:default_policy]
    self.description = opts[:description]
    self.filters = opts[:filters]
    self.id = opts[:id]
    self.locked = opts[:locked]
    self.name = opts[:name]
    self.properties = opts[:properties]
    self.weight = opts[:weight]
  end
  
end

class SchedulingPolicyUnit < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `internal` attribute.
  # 
  def internal
    return @internal
  end
  
  # 
  # Sets the value of the `internal` attribute.
  # 
  def internal=(value)
    @internal = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {SchedulingPolicyUnit} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :internal The value of attribute `internal`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.enabled = opts[:enabled]
    self.id = opts[:id]
    self.internal = opts[:internal]
    self.name = opts[:name]
    self.properties = opts[:properties]
    self.type = opts[:type]
  end
  
end

class SeLinux < Struct
  
  # 
  # Returns the value of the `mode` attribute.
  # 
  def mode
    return @mode
  end
  
  # 
  # Sets the value of the `mode` attribute.
  # 
  def mode=(value)
    @mode = value
  end
  
  # 
  # Creates a new instance of the {SeLinux} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :mode The value of attribute `mode`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.mode = opts[:mode]
  end
  
end

class SerialNumber < Struct
  
  # 
  # Returns the value of the `policy` attribute.
  # 
  def policy
    return @policy
  end
  
  # 
  # Sets the value of the `policy` attribute.
  # 
  def policy=(value)
    @policy = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {SerialNumber} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :policy The value of attribute `policy`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.policy = opts[:policy]
    self.value = opts[:value]
  end
  
end

class Session < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `console_user` attribute.
  # 
  def console_user
    return @console_user
  end
  
  # 
  # Sets the value of the `console_user` attribute.
  # 
  def console_user=(value)
    @console_user = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ip` attribute.
  # 
  def ip
    return @ip
  end
  
  # 
  # Sets the value of the `ip` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ip} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ip=(value)
    if value.is_a?(Hash)
      value = Ip.new(value)
    end
    @ip = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `protocol` attribute.
  # 
  def protocol
    return @protocol
  end
  
  # 
  # Sets the value of the `protocol` attribute.
  # 
  def protocol=(value)
    @protocol = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {Session} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :console_user The value of attribute `console_user`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :ip The value of attribute `ip`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :protocol The value of attribute `protocol`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.console_user = opts[:console_user]
    self.description = opts[:description]
    self.id = opts[:id]
    self.ip = opts[:ip]
    self.name = opts[:name]
    self.protocol = opts[:protocol]
    self.user = opts[:user]
    self.vm = opts[:vm]
  end
  
end

class SkipIfConnectivityBroken < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `threshold` attribute.
  # 
  def threshold
    return @threshold
  end
  
  # 
  # Sets the value of the `threshold` attribute.
  # 
  def threshold=(value)
    @threshold = value
  end
  
  # 
  # Creates a new instance of the {SkipIfConnectivityBroken} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts :threshold The value of attribute `threshold`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
    self.threshold = opts[:threshold]
  end
  
end

class SkipIfSdActive < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Creates a new instance of the {SkipIfSdActive} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
  end
  
end

class SpecialObjects < Struct
  
  # 
  # Returns the value of the `blank_template` attribute.
  # 
  def blank_template
    return @blank_template
  end
  
  # 
  # Sets the value of the `blank_template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def blank_template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @blank_template = value
  end
  
  # 
  # Returns the value of the `root_tag` attribute.
  # 
  def root_tag
    return @root_tag
  end
  
  # 
  # Sets the value of the `root_tag` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Tag} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def root_tag=(value)
    if value.is_a?(Hash)
      value = Tag.new(value)
    end
    @root_tag = value
  end
  
  # 
  # Creates a new instance of the {SpecialObjects} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :blank_template The value of attribute `blank_template`.
  # 
  # @option opts [Hash] :root_tag The value of attribute `root_tag`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.blank_template = opts[:blank_template]
    self.root_tag = opts[:root_tag]
  end
  
end

class Spm < Struct
  
  # 
  # Returns the value of the `priority` attribute.
  # 
  def priority
    return @priority
  end
  
  # 
  # Sets the value of the `priority` attribute.
  # 
  def priority=(value)
    @priority = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Creates a new instance of the {Spm} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :priority The value of attribute `priority`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.priority = opts[:priority]
    self.status = opts[:status]
  end
  
end

class Ssh < Identified
  
  # 
  # Returns the value of the `authentication_method` attribute.
  # 
  def authentication_method
    return @authentication_method
  end
  
  # 
  # Sets the value of the `authentication_method` attribute.
  # 
  def authentication_method=(value)
    @authentication_method = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `fingerprint` attribute.
  # 
  def fingerprint
    return @fingerprint
  end
  
  # 
  # Sets the value of the `fingerprint` attribute.
  # 
  def fingerprint=(value)
    @fingerprint = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Creates a new instance of the {Ssh} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :authentication_method The value of attribute `authentication_method`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :fingerprint The value of attribute `fingerprint`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authentication_method = opts[:authentication_method]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.fingerprint = opts[:fingerprint]
    self.id = opts[:id]
    self.name = opts[:name]
    self.port = opts[:port]
    self.user = opts[:user]
  end
  
end

class SshPublicKey < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `content` attribute.
  # 
  def content
    return @content
  end
  
  # 
  # Sets the value of the `content` attribute.
  # 
  def content=(value)
    @content = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Creates a new instance of the {SshPublicKey} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :content The value of attribute `content`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.content = opts[:content]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.user = opts[:user]
  end
  
end

class Sso < Struct
  
  # 
  # Returns the value of the `methods` attribute.
  # 
  def methods
    return @methods
  end
  
  # 
  # Sets the value of the `methods` attribute.
  # 
  def methods=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Method.new(value)
        end
      end
    end
    @methods = list
  end
  
  # 
  # Creates a new instance of the {Sso} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :methods The values of attribute `methods`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.methods = opts[:methods]
  end
  
end

class Statistic < Identified
  
  # 
  # Returns the value of the `brick` attribute.
  # 
  def brick
    return @brick
  end
  
  # 
  # Sets the value of the `brick` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GlusterBrick} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def brick=(value)
    if value.is_a?(Hash)
      value = GlusterBrick.new(value)
    end
    @brick = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk` attribute.
  # 
  def disk
    return @disk
  end
  
  # 
  # Sets the value of the `disk` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Disk} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk=(value)
    if value.is_a?(Hash)
      value = Disk.new(value)
    end
    @disk = value
  end
  
  # 
  # Returns the value of the `gluster_volume` attribute.
  # 
  def gluster_volume
    return @gluster_volume
  end
  
  # 
  # Sets the value of the `gluster_volume` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GlusterVolume} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def gluster_volume=(value)
    if value.is_a?(Hash)
      value = GlusterVolume.new(value)
    end
    @gluster_volume = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `host_nic` attribute.
  # 
  def host_nic
    return @host_nic
  end
  
  # 
  # Sets the value of the `host_nic` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNic} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host_nic=(value)
    if value.is_a?(Hash)
      value = HostNic.new(value)
    end
    @host_nic = value
  end
  
  # 
  # Returns the value of the `host_numa_node` attribute.
  # 
  def host_numa_node
    return @host_numa_node
  end
  
  # 
  # Sets the value of the `host_numa_node` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::NumaNode} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host_numa_node=(value)
    if value.is_a?(Hash)
      value = NumaNode.new(value)
    end
    @host_numa_node = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `kind` attribute.
  # 
  def kind
    return @kind
  end
  
  # 
  # Sets the value of the `kind` attribute.
  # 
  def kind=(value)
    @kind = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `nic` attribute.
  # 
  def nic
    return @nic
  end
  
  # 
  # Sets the value of the `nic` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Nic} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def nic=(value)
    if value.is_a?(Hash)
      value = Nic.new(value)
    end
    @nic = value
  end
  
  # 
  # Returns the value of the `step` attribute.
  # 
  def step
    return @step
  end
  
  # 
  # Sets the value of the `step` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Step} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def step=(value)
    if value.is_a?(Hash)
      value = Step.new(value)
    end
    @step = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `unit` attribute.
  # 
  def unit
    return @unit
  end
  
  # 
  # Sets the value of the `unit` attribute.
  # 
  def unit=(value)
    @unit = value
  end
  
  # 
  # Returns the value of the `values` attribute.
  # 
  def values
    return @values
  end
  
  # 
  # Sets the value of the `values` attribute.
  # 
  def values=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Value.new(value)
        end
      end
    end
    @values = list
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {Statistic} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :brick The value of attribute `brick`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :disk The value of attribute `disk`.
  # 
  # @option opts [Hash] :gluster_volume The value of attribute `gluster_volume`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts [Hash] :host_nic The value of attribute `host_nic`.
  # 
  # @option opts [Hash] :host_numa_node The value of attribute `host_numa_node`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :kind The value of attribute `kind`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :nic The value of attribute `nic`.
  # 
  # @option opts [Hash] :step The value of attribute `step`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :unit The value of attribute `unit`.
  # 
  # @option opts [Array<Hash>] :values The values of attribute `values`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.brick = opts[:brick]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.disk = opts[:disk]
    self.gluster_volume = opts[:gluster_volume]
    self.host = opts[:host]
    self.host_nic = opts[:host_nic]
    self.host_numa_node = opts[:host_numa_node]
    self.id = opts[:id]
    self.kind = opts[:kind]
    self.name = opts[:name]
    self.nic = opts[:nic]
    self.step = opts[:step]
    self.type = opts[:type]
    self.unit = opts[:unit]
    self.values = opts[:values]
    self.vm = opts[:vm]
  end
  
end

class Step < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `end_time` attribute.
  # 
  def end_time
    return @end_time
  end
  
  # 
  # Sets the value of the `end_time` attribute.
  # 
  def end_time=(value)
    @end_time = value
  end
  
  # 
  # Returns the value of the `external` attribute.
  # 
  def external
    return @external
  end
  
  # 
  # Sets the value of the `external` attribute.
  # 
  def external=(value)
    @external = value
  end
  
  # 
  # Returns the value of the `external_type` attribute.
  # 
  def external_type
    return @external_type
  end
  
  # 
  # Sets the value of the `external_type` attribute.
  # 
  def external_type=(value)
    @external_type = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `job` attribute.
  # 
  def job
    return @job
  end
  
  # 
  # Sets the value of the `job` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Job} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def job=(value)
    if value.is_a?(Hash)
      value = Job.new(value)
    end
    @job = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `number` attribute.
  # 
  def number
    return @number
  end
  
  # 
  # Sets the value of the `number` attribute.
  # 
  def number=(value)
    @number = value
  end
  
  # 
  # Returns the value of the `parent_step` attribute.
  # 
  def parent_step
    return @parent_step
  end
  
  # 
  # Sets the value of the `parent_step` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Step} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def parent_step=(value)
    if value.is_a?(Hash)
      value = Step.new(value)
    end
    @parent_step = value
  end
  
  # 
  # Returns the value of the `start_time` attribute.
  # 
  def start_time
    return @start_time
  end
  
  # 
  # Sets the value of the `start_time` attribute.
  # 
  def start_time=(value)
    @start_time = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {Step} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :end_time The value of attribute `end_time`.
  # 
  # @option opts :external The value of attribute `external`.
  # 
  # @option opts :external_type The value of attribute `external_type`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :job The value of attribute `job`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :number The value of attribute `number`.
  # 
  # @option opts [Hash] :parent_step The value of attribute `parent_step`.
  # 
  # @option opts :start_time The value of attribute `start_time`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.end_time = opts[:end_time]
    self.external = opts[:external]
    self.external_type = opts[:external_type]
    self.id = opts[:id]
    self.job = opts[:job]
    self.name = opts[:name]
    self.number = opts[:number]
    self.parent_step = opts[:parent_step]
    self.start_time = opts[:start_time]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.type = opts[:type]
  end
  
end

class StorageConnection < Identified
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `mount_options` attribute.
  # 
  def mount_options
    return @mount_options
  end
  
  # 
  # Sets the value of the `mount_options` attribute.
  # 
  def mount_options=(value)
    @mount_options = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `nfs_retrans` attribute.
  # 
  def nfs_retrans
    return @nfs_retrans
  end
  
  # 
  # Sets the value of the `nfs_retrans` attribute.
  # 
  def nfs_retrans=(value)
    @nfs_retrans = value
  end
  
  # 
  # Returns the value of the `nfs_timeo` attribute.
  # 
  def nfs_timeo
    return @nfs_timeo
  end
  
  # 
  # Sets the value of the `nfs_timeo` attribute.
  # 
  def nfs_timeo=(value)
    @nfs_timeo = value
  end
  
  # 
  # Returns the value of the `nfs_version` attribute.
  # 
  def nfs_version
    return @nfs_version
  end
  
  # 
  # Sets the value of the `nfs_version` attribute.
  # 
  def nfs_version=(value)
    @nfs_version = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `path` attribute.
  # 
  def path
    return @path
  end
  
  # 
  # Sets the value of the `path` attribute.
  # 
  def path=(value)
    @path = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `portal` attribute.
  # 
  def portal
    return @portal
  end
  
  # 
  # Sets the value of the `portal` attribute.
  # 
  def portal=(value)
    @portal = value
  end
  
  # 
  # Returns the value of the `target` attribute.
  # 
  def target
    return @target
  end
  
  # 
  # Sets the value of the `target` attribute.
  # 
  def target=(value)
    @target = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Returns the value of the `vfs_type` attribute.
  # 
  def vfs_type
    return @vfs_type
  end
  
  # 
  # Sets the value of the `vfs_type` attribute.
  # 
  def vfs_type=(value)
    @vfs_type = value
  end
  
  # 
  # Creates a new instance of the {StorageConnection} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :mount_options The value of attribute `mount_options`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :nfs_retrans The value of attribute `nfs_retrans`.
  # 
  # @option opts :nfs_timeo The value of attribute `nfs_timeo`.
  # 
  # @option opts :nfs_version The value of attribute `nfs_version`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :path The value of attribute `path`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :portal The value of attribute `portal`.
  # 
  # @option opts :target The value of attribute `target`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # @option opts :vfs_type The value of attribute `vfs_type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.mount_options = opts[:mount_options]
    self.name = opts[:name]
    self.nfs_retrans = opts[:nfs_retrans]
    self.nfs_timeo = opts[:nfs_timeo]
    self.nfs_version = opts[:nfs_version]
    self.password = opts[:password]
    self.path = opts[:path]
    self.port = opts[:port]
    self.portal = opts[:portal]
    self.target = opts[:target]
    self.type = opts[:type]
    self.username = opts[:username]
    self.vfs_type = opts[:vfs_type]
  end
  
end

class StorageConnectionExtension < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `target` attribute.
  # 
  def target
    return @target
  end
  
  # 
  # Sets the value of the `target` attribute.
  # 
  def target=(value)
    @target = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {StorageConnectionExtension} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :target The value of attribute `target`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.name = opts[:name]
    self.password = opts[:password]
    self.target = opts[:target]
    self.username = opts[:username]
  end
  
end

class StorageDomain < Identified
  
  # 
  # Returns the value of the `available` attribute.
  # 
  def available
    return @available
  end
  
  # 
  # Sets the value of the `available` attribute.
  # 
  def available=(value)
    @available = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `committed` attribute.
  # 
  def committed
    return @committed
  end
  
  # 
  # Sets the value of the `committed` attribute.
  # 
  def committed=(value)
    @committed = value
  end
  
  # 
  # Returns the value of the `critical_space_action_blocker` attribute.
  # 
  def critical_space_action_blocker
    return @critical_space_action_blocker
  end
  
  # 
  # Sets the value of the `critical_space_action_blocker` attribute.
  # 
  def critical_space_action_blocker=(value)
    @critical_space_action_blocker = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `data_centers` attribute.
  # 
  def data_centers
    return @data_centers
  end
  
  # 
  # Sets the value of the `data_centers` attribute.
  # 
  def data_centers=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DataCenter.new(value)
        end
      end
    end
    @data_centers = list
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk_profiles` attribute.
  # 
  def disk_profiles
    return @disk_profiles
  end
  
  # 
  # Sets the value of the `disk_profiles` attribute.
  # 
  def disk_profiles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DiskProfile.new(value)
        end
      end
    end
    @disk_profiles = list
  end
  
  # 
  # Returns the value of the `disk_snapshots` attribute.
  # 
  def disk_snapshots
    return @disk_snapshots
  end
  
  # 
  # Sets the value of the `disk_snapshots` attribute.
  # 
  def disk_snapshots=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DiskSnapshot.new(value)
        end
      end
    end
    @disk_snapshots = list
  end
  
  # 
  # Returns the value of the `disks` attribute.
  # 
  def disks
    return @disks
  end
  
  # 
  # Sets the value of the `disks` attribute.
  # 
  def disks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Disk.new(value)
        end
      end
    end
    @disks = list
  end
  
  # 
  # Returns the value of the `external_status` attribute.
  # 
  def external_status
    return @external_status
  end
  
  # 
  # Sets the value of the `external_status` attribute.
  # 
  def external_status=(value)
    @external_status = value
  end
  
  # 
  # Returns the value of the `files` attribute.
  # 
  def files
    return @files
  end
  
  # 
  # Sets the value of the `files` attribute.
  # 
  def files=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = File.new(value)
        end
      end
    end
    @files = list
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `images` attribute.
  # 
  def images
    return @images
  end
  
  # 
  # Sets the value of the `images` attribute.
  # 
  def images=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Image.new(value)
        end
      end
    end
    @images = list
  end
  
  # 
  # Returns the value of the `import` attribute.
  # 
  def import
    return @import
  end
  
  # 
  # Sets the value of the `import` attribute.
  # 
  def import=(value)
    @import = value
  end
  
  # 
  # Returns the value of the `master` attribute.
  # 
  def master
    return @master
  end
  
  # 
  # Sets the value of the `master` attribute.
  # 
  def master=(value)
    @master = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage` attribute.
  # 
  def storage
    return @storage
  end
  
  # 
  # Sets the value of the `storage` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostStorage} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage=(value)
    if value.is_a?(Hash)
      value = HostStorage.new(value)
    end
    @storage = value
  end
  
  # 
  # Returns the value of the `storage_connections` attribute.
  # 
  def storage_connections
    return @storage_connections
  end
  
  # 
  # Sets the value of the `storage_connections` attribute.
  # 
  def storage_connections=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageConnection.new(value)
        end
      end
    end
    @storage_connections = list
  end
  
  # 
  # Returns the value of the `storage_format` attribute.
  # 
  def storage_format
    return @storage_format
  end
  
  # 
  # Sets the value of the `storage_format` attribute.
  # 
  def storage_format=(value)
    @storage_format = value
  end
  
  # 
  # Returns the value of the `templates` attribute.
  # 
  def templates
    return @templates
  end
  
  # 
  # Sets the value of the `templates` attribute.
  # 
  def templates=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Template.new(value)
        end
      end
    end
    @templates = list
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `used` attribute.
  # 
  def used
    return @used
  end
  
  # 
  # Sets the value of the `used` attribute.
  # 
  def used=(value)
    @used = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Returns the value of the `warning_low_space_indicator` attribute.
  # 
  def warning_low_space_indicator
    return @warning_low_space_indicator
  end
  
  # 
  # Sets the value of the `warning_low_space_indicator` attribute.
  # 
  def warning_low_space_indicator=(value)
    @warning_low_space_indicator = value
  end
  
  # 
  # Returns the value of the `wipe_after_delete` attribute.
  # 
  def wipe_after_delete
    return @wipe_after_delete
  end
  
  # 
  # Sets the value of the `wipe_after_delete` attribute.
  # 
  def wipe_after_delete=(value)
    @wipe_after_delete = value
  end
  
  # 
  # Creates a new instance of the {StorageDomain} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :available The value of attribute `available`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :committed The value of attribute `committed`.
  # 
  # @option opts :critical_space_action_blocker The value of attribute `critical_space_action_blocker`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts [Array<Hash>] :data_centers The values of attribute `data_centers`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :disk_profiles The values of attribute `disk_profiles`.
  # 
  # @option opts [Array<Hash>] :disk_snapshots The values of attribute `disk_snapshots`.
  # 
  # @option opts [Array<Hash>] :disks The values of attribute `disks`.
  # 
  # @option opts :external_status The value of attribute `external_status`.
  # 
  # @option opts [Array<Hash>] :files The values of attribute `files`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Array<Hash>] :images The values of attribute `images`.
  # 
  # @option opts :import The value of attribute `import`.
  # 
  # @option opts :master The value of attribute `master`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :storage The value of attribute `storage`.
  # 
  # @option opts [Array<Hash>] :storage_connections The values of attribute `storage_connections`.
  # 
  # @option opts :storage_format The value of attribute `storage_format`.
  # 
  # @option opts [Array<Hash>] :templates The values of attribute `templates`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :used The value of attribute `used`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # @option opts :warning_low_space_indicator The value of attribute `warning_low_space_indicator`.
  # 
  # @option opts :wipe_after_delete The value of attribute `wipe_after_delete`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.available = opts[:available]
    self.comment = opts[:comment]
    self.committed = opts[:committed]
    self.critical_space_action_blocker = opts[:critical_space_action_blocker]
    self.data_center = opts[:data_center]
    self.data_centers = opts[:data_centers]
    self.description = opts[:description]
    self.disk_profiles = opts[:disk_profiles]
    self.disk_snapshots = opts[:disk_snapshots]
    self.disks = opts[:disks]
    self.external_status = opts[:external_status]
    self.files = opts[:files]
    self.host = opts[:host]
    self.id = opts[:id]
    self.images = opts[:images]
    self.import = opts[:import]
    self.master = opts[:master]
    self.name = opts[:name]
    self.permissions = opts[:permissions]
    self.status = opts[:status]
    self.storage = opts[:storage]
    self.storage_connections = opts[:storage_connections]
    self.storage_format = opts[:storage_format]
    self.templates = opts[:templates]
    self.type = opts[:type]
    self.used = opts[:used]
    self.vms = opts[:vms]
    self.warning_low_space_indicator = opts[:warning_low_space_indicator]
    self.wipe_after_delete = opts[:wipe_after_delete]
  end
  
end

class Tag < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `group` attribute.
  # 
  def group
    return @group
  end
  
  # 
  # Sets the value of the `group` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Group} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def group=(value)
    if value.is_a?(Hash)
      value = Group.new(value)
    end
    @group = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `parent` attribute.
  # 
  def parent
    return @parent
  end
  
  # 
  # Sets the value of the `parent` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Tag} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def parent=(value)
    if value.is_a?(Hash)
      value = Tag.new(value)
    end
    @parent = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {Tag} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :group The value of attribute `group`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :parent The value of attribute `parent`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.group = opts[:group]
    self.host = opts[:host]
    self.id = opts[:id]
    self.name = opts[:name]
    self.parent = opts[:parent]
    self.template = opts[:template]
    self.user = opts[:user]
    self.vm = opts[:vm]
  end
  
end

class TemplateVersion < Struct
  
  # 
  # Returns the value of the `base_template` attribute.
  # 
  def base_template
    return @base_template
  end
  
  # 
  # Sets the value of the `base_template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def base_template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @base_template = value
  end
  
  # 
  # Returns the value of the `version_name` attribute.
  # 
  def version_name
    return @version_name
  end
  
  # 
  # Sets the value of the `version_name` attribute.
  # 
  def version_name=(value)
    @version_name = value
  end
  
  # 
  # Returns the value of the `version_number` attribute.
  # 
  def version_number
    return @version_number
  end
  
  # 
  # Sets the value of the `version_number` attribute.
  # 
  def version_number=(value)
    @version_number = value
  end
  
  # 
  # Creates a new instance of the {TemplateVersion} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :base_template The value of attribute `base_template`.
  # 
  # @option opts :version_name The value of attribute `version_name`.
  # 
  # @option opts :version_number The value of attribute `version_number`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.base_template = opts[:base_template]
    self.version_name = opts[:version_name]
    self.version_number = opts[:version_number]
  end
  
end

class Ticket < Struct
  
  # 
  # Returns the value of the `expiry` attribute.
  # 
  def expiry
    return @expiry
  end
  
  # 
  # Sets the value of the `expiry` attribute.
  # 
  def expiry=(value)
    @expiry = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {Ticket} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :expiry The value of attribute `expiry`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.expiry = opts[:expiry]
    self.value = opts[:value]
  end
  
end

class TimeZone < Struct
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `utc_offset` attribute.
  # 
  def utc_offset
    return @utc_offset
  end
  
  # 
  # Sets the value of the `utc_offset` attribute.
  # 
  def utc_offset=(value)
    @utc_offset = value
  end
  
  # 
  # Creates a new instance of the {TimeZone} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :utc_offset The value of attribute `utc_offset`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.name = opts[:name]
    self.utc_offset = opts[:utc_offset]
  end
  
end

class TransparentHugePages < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Creates a new instance of the {TransparentHugePages} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
  end
  
end

class UnmanagedNetwork < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `host_nic` attribute.
  # 
  def host_nic
    return @host_nic
  end
  
  # 
  # Sets the value of the `host_nic` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNic} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host_nic=(value)
    if value.is_a?(Hash)
      value = HostNic.new(value)
    end
    @host_nic = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {UnmanagedNetwork} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts [Hash] :host_nic The value of attribute `host_nic`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.host_nic = opts[:host_nic]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class Usb < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {Usb} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
    self.type = opts[:type]
  end
  
end

class User < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `department` attribute.
  # 
  def department
    return @department
  end
  
  # 
  # Sets the value of the `department` attribute.
  # 
  def department=(value)
    @department = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `domain_entry_id` attribute.
  # 
  def domain_entry_id
    return @domain_entry_id
  end
  
  # 
  # Sets the value of the `domain_entry_id` attribute.
  # 
  def domain_entry_id=(value)
    @domain_entry_id = value
  end
  
  # 
  # Returns the value of the `email` attribute.
  # 
  def email
    return @email
  end
  
  # 
  # Sets the value of the `email` attribute.
  # 
  def email=(value)
    @email = value
  end
  
  # 
  # Returns the value of the `groups` attribute.
  # 
  def groups
    return @groups
  end
  
  # 
  # Sets the value of the `groups` attribute.
  # 
  def groups=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Group.new(value)
        end
      end
    end
    @groups = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `last_name` attribute.
  # 
  def last_name
    return @last_name
  end
  
  # 
  # Sets the value of the `last_name` attribute.
  # 
  def last_name=(value)
    @last_name = value
  end
  
  # 
  # Returns the value of the `logged_in` attribute.
  # 
  def logged_in
    return @logged_in
  end
  
  # 
  # Sets the value of the `logged_in` attribute.
  # 
  def logged_in=(value)
    @logged_in = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `namespace` attribute.
  # 
  def namespace
    return @namespace
  end
  
  # 
  # Sets the value of the `namespace` attribute.
  # 
  def namespace=(value)
    @namespace = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `principal` attribute.
  # 
  def principal
    return @principal
  end
  
  # 
  # Sets the value of the `principal` attribute.
  # 
  def principal=(value)
    @principal = value
  end
  
  # 
  # Returns the value of the `roles` attribute.
  # 
  def roles
    return @roles
  end
  
  # 
  # Sets the value of the `roles` attribute.
  # 
  def roles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Role.new(value)
        end
      end
    end
    @roles = list
  end
  
  # 
  # Returns the value of the `ssh_public_keys` attribute.
  # 
  def ssh_public_keys
    return @ssh_public_keys
  end
  
  # 
  # Sets the value of the `ssh_public_keys` attribute.
  # 
  def ssh_public_keys=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = SshPublicKey.new(value)
        end
      end
    end
    @ssh_public_keys = list
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Returns the value of the `user_name` attribute.
  # 
  def user_name
    return @user_name
  end
  
  # 
  # Sets the value of the `user_name` attribute.
  # 
  def user_name=(value)
    @user_name = value
  end
  
  # 
  # Creates a new instance of the {User} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :department The value of attribute `department`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts :domain_entry_id The value of attribute `domain_entry_id`.
  # 
  # @option opts :email The value of attribute `email`.
  # 
  # @option opts [Array<Hash>] :groups The values of attribute `groups`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :last_name The value of attribute `last_name`.
  # 
  # @option opts :logged_in The value of attribute `logged_in`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :namespace The value of attribute `namespace`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :principal The value of attribute `principal`.
  # 
  # @option opts [Array<Hash>] :roles The values of attribute `roles`.
  # 
  # @option opts [Array<Hash>] :ssh_public_keys The values of attribute `ssh_public_keys`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # @option opts :user_name The value of attribute `user_name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.department = opts[:department]
    self.description = opts[:description]
    self.domain = opts[:domain]
    self.domain_entry_id = opts[:domain_entry_id]
    self.email = opts[:email]
    self.groups = opts[:groups]
    self.id = opts[:id]
    self.last_name = opts[:last_name]
    self.logged_in = opts[:logged_in]
    self.name = opts[:name]
    self.namespace = opts[:namespace]
    self.password = opts[:password]
    self.permissions = opts[:permissions]
    self.principal = opts[:principal]
    self.roles = opts[:roles]
    self.ssh_public_keys = opts[:ssh_public_keys]
    self.tags = opts[:tags]
    self.user_name = opts[:user_name]
  end
  
end

class Value < Struct
  
  # 
  # Returns the value of the `datum` attribute.
  # 
  def datum
    return @datum
  end
  
  # 
  # Sets the value of the `datum` attribute.
  # 
  def datum=(value)
    @datum = value
  end
  
  # 
  # Returns the value of the `detail` attribute.
  # 
  def detail
    return @detail
  end
  
  # 
  # Sets the value of the `detail` attribute.
  # 
  def detail=(value)
    @detail = value
  end
  
  # 
  # Creates a new instance of the {Value} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :datum The value of attribute `datum`.
  # 
  # @option opts :detail The value of attribute `detail`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.datum = opts[:datum]
    self.detail = opts[:detail]
  end
  
end

class VcpuPin < Struct
  
  # 
  # Returns the value of the `cpu_set` attribute.
  # 
  def cpu_set
    return @cpu_set
  end
  
  # 
  # Sets the value of the `cpu_set` attribute.
  # 
  def cpu_set=(value)
    @cpu_set = value
  end
  
  # 
  # Returns the value of the `vcpu` attribute.
  # 
  def vcpu
    return @vcpu
  end
  
  # 
  # Sets the value of the `vcpu` attribute.
  # 
  def vcpu=(value)
    @vcpu = value
  end
  
  # 
  # Creates a new instance of the {VcpuPin} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :cpu_set The value of attribute `cpu_set`.
  # 
  # @option opts :vcpu The value of attribute `vcpu`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cpu_set = opts[:cpu_set]
    self.vcpu = opts[:vcpu]
  end
  
end

class Vendor < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {Vendor} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class Version < Identified
  
  # 
  # Returns the value of the `build` attribute.
  # 
  def build
    return @build
  end
  
  # 
  # Sets the value of the `build` attribute.
  # 
  def build=(value)
    @build = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `full_version` attribute.
  # 
  def full_version
    return @full_version
  end
  
  # 
  # Sets the value of the `full_version` attribute.
  # 
  def full_version=(value)
    @full_version = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `major` attribute.
  # 
  def major
    return @major
  end
  
  # 
  # Sets the value of the `major` attribute.
  # 
  def major=(value)
    @major = value
  end
  
  # 
  # Returns the value of the `minor` attribute.
  # 
  def minor
    return @minor
  end
  
  # 
  # Sets the value of the `minor` attribute.
  # 
  def minor=(value)
    @minor = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `revision` attribute.
  # 
  def revision
    return @revision
  end
  
  # 
  # Sets the value of the `revision` attribute.
  # 
  def revision=(value)
    @revision = value
  end
  
  # 
  # Creates a new instance of the {Version} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :build The value of attribute `build`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :full_version The value of attribute `full_version`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :major The value of attribute `major`.
  # 
  # @option opts :minor The value of attribute `minor`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :revision The value of attribute `revision`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.build = opts[:build]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.full_version = opts[:full_version]
    self.id = opts[:id]
    self.major = opts[:major]
    self.minor = opts[:minor]
    self.name = opts[:name]
    self.revision = opts[:revision]
  end
  
end

class VirtioScsi < Struct
  
  # 
  # Returns the value of the `enabled` attribute.
  # 
  def enabled
    return @enabled
  end
  
  # 
  # Sets the value of the `enabled` attribute.
  # 
  def enabled=(value)
    @enabled = value
  end
  
  # 
  # Creates a new instance of the {VirtioScsi} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :enabled The value of attribute `enabled`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.enabled = opts[:enabled]
  end
  
end

class VirtualNumaNode < NumaNode
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `index` attribute.
  # 
  def index
    return @index
  end
  
  # 
  # Sets the value of the `index` attribute.
  # 
  def index=(value)
    @index = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `node_distance` attribute.
  # 
  def node_distance
    return @node_distance
  end
  
  # 
  # Sets the value of the `node_distance` attribute.
  # 
  def node_distance=(value)
    @node_distance = value
  end
  
  # 
  # Returns the value of the `numa_node_pins` attribute.
  # 
  def numa_node_pins
    return @numa_node_pins
  end
  
  # 
  # Sets the value of the `numa_node_pins` attribute.
  # 
  def numa_node_pins=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NumaNodePin.new(value)
        end
      end
    end
    @numa_node_pins = list
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {VirtualNumaNode} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :index The value of attribute `index`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :node_distance The value of attribute `node_distance`.
  # 
  # @option opts [Array<Hash>] :numa_node_pins The values of attribute `numa_node_pins`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.cpu = opts[:cpu]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.index = opts[:index]
    self.memory = opts[:memory]
    self.name = opts[:name]
    self.node_distance = opts[:node_distance]
    self.numa_node_pins = opts[:numa_node_pins]
    self.statistics = opts[:statistics]
    self.vm = opts[:vm]
  end
  
end

class Vlan < Struct
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Creates a new instance of the {Vlan} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.id = opts[:id]
  end
  
end

class VmBase < Identified
  
  # 
  # Returns the value of the `bios` attribute.
  # 
  def bios
    return @bios
  end
  
  # 
  # Sets the value of the `bios` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Bios} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bios=(value)
    if value.is_a?(Hash)
      value = Bios.new(value)
    end
    @bios = value
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `console` attribute.
  # 
  def console
    return @console
  end
  
  # 
  # Sets the value of the `console` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Console} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def console=(value)
    if value.is_a?(Hash)
      value = Console.new(value)
    end
    @console = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `cpu_profile` attribute.
  # 
  def cpu_profile
    return @cpu_profile
  end
  
  # 
  # Sets the value of the `cpu_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu_profile=(value)
    if value.is_a?(Hash)
      value = CpuProfile.new(value)
    end
    @cpu_profile = value
  end
  
  # 
  # Returns the value of the `cpu_shares` attribute.
  # 
  def cpu_shares
    return @cpu_shares
  end
  
  # 
  # Sets the value of the `cpu_shares` attribute.
  # 
  def cpu_shares=(value)
    @cpu_shares = value
  end
  
  # 
  # Returns the value of the `creation_time` attribute.
  # 
  def creation_time
    return @creation_time
  end
  
  # 
  # Sets the value of the `creation_time` attribute.
  # 
  def creation_time=(value)
    @creation_time = value
  end
  
  # 
  # Returns the value of the `custom_compatibility_version` attribute.
  # 
  def custom_compatibility_version
    return @custom_compatibility_version
  end
  
  # 
  # Sets the value of the `custom_compatibility_version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def custom_compatibility_version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @custom_compatibility_version = value
  end
  
  # 
  # Returns the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model
    return @custom_cpu_model
  end
  
  # 
  # Sets the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model=(value)
    @custom_cpu_model = value
  end
  
  # 
  # Returns the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine
    return @custom_emulated_machine
  end
  
  # 
  # Sets the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine=(value)
    @custom_emulated_machine = value
  end
  
  # 
  # Returns the value of the `custom_properties` attribute.
  # 
  def custom_properties
    return @custom_properties
  end
  
  # 
  # Sets the value of the `custom_properties` attribute.
  # 
  def custom_properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CustomProperty.new(value)
        end
      end
    end
    @custom_properties = list
  end
  
  # 
  # Returns the value of the `delete_protected` attribute.
  # 
  def delete_protected
    return @delete_protected
  end
  
  # 
  # Sets the value of the `delete_protected` attribute.
  # 
  def delete_protected=(value)
    @delete_protected = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `high_availability` attribute.
  # 
  def high_availability
    return @high_availability
  end
  
  # 
  # Sets the value of the `high_availability` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HighAvailability} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def high_availability=(value)
    if value.is_a?(Hash)
      value = HighAvailability.new(value)
    end
    @high_availability = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `initialization` attribute.
  # 
  def initialization
    return @initialization
  end
  
  # 
  # Sets the value of the `initialization` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Initialization} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def initialization=(value)
    if value.is_a?(Hash)
      value = Initialization.new(value)
    end
    @initialization = value
  end
  
  # 
  # Returns the value of the `io` attribute.
  # 
  def io
    return @io
  end
  
  # 
  # Sets the value of the `io` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Io} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def io=(value)
    if value.is_a?(Hash)
      value = Io.new(value)
    end
    @io = value
  end
  
  # 
  # Returns the value of the `large_icon` attribute.
  # 
  def large_icon
    return @large_icon
  end
  
  # 
  # Sets the value of the `large_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def large_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @large_icon = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `memory_policy` attribute.
  # 
  def memory_policy
    return @memory_policy
  end
  
  # 
  # Sets the value of the `memory_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def memory_policy=(value)
    if value.is_a?(Hash)
      value = MemoryPolicy.new(value)
    end
    @memory_policy = value
  end
  
  # 
  # Returns the value of the `migration` attribute.
  # 
  def migration
    return @migration
  end
  
  # 
  # Sets the value of the `migration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationOptions} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def migration=(value)
    if value.is_a?(Hash)
      value = MigrationOptions.new(value)
    end
    @migration = value
  end
  
  # 
  # Returns the value of the `migration_downtime` attribute.
  # 
  def migration_downtime
    return @migration_downtime
  end
  
  # 
  # Sets the value of the `migration_downtime` attribute.
  # 
  def migration_downtime=(value)
    @migration_downtime = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `origin` attribute.
  # 
  def origin
    return @origin
  end
  
  # 
  # Sets the value of the `origin` attribute.
  # 
  def origin=(value)
    @origin = value
  end
  
  # 
  # Returns the value of the `os` attribute.
  # 
  def os
    return @os
  end
  
  # 
  # Sets the value of the `os` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def os=(value)
    if value.is_a?(Hash)
      value = OperatingSystem.new(value)
    end
    @os = value
  end
  
  # 
  # Returns the value of the `rng_device` attribute.
  # 
  def rng_device
    return @rng_device
  end
  
  # 
  # Sets the value of the `rng_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::RngDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rng_device=(value)
    if value.is_a?(Hash)
      value = RngDevice.new(value)
    end
    @rng_device = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SerialNumber} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def serial_number=(value)
    if value.is_a?(Hash)
      value = SerialNumber.new(value)
    end
    @serial_number = value
  end
  
  # 
  # Returns the value of the `small_icon` attribute.
  # 
  def small_icon
    return @small_icon
  end
  
  # 
  # Sets the value of the `small_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def small_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @small_icon = value
  end
  
  # 
  # Returns the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled
    return @soundcard_enabled
  end
  
  # 
  # Sets the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled=(value)
    @soundcard_enabled = value
  end
  
  # 
  # Returns the value of the `sso` attribute.
  # 
  def sso
    return @sso
  end
  
  # 
  # Sets the value of the `sso` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Sso} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def sso=(value)
    if value.is_a?(Hash)
      value = Sso.new(value)
    end
    @sso = value
  end
  
  # 
  # Returns the value of the `start_paused` attribute.
  # 
  def start_paused
    return @start_paused
  end
  
  # 
  # Sets the value of the `start_paused` attribute.
  # 
  def start_paused=(value)
    @start_paused = value
  end
  
  # 
  # Returns the value of the `stateless` attribute.
  # 
  def stateless
    return @stateless
  end
  
  # 
  # Sets the value of the `stateless` attribute.
  # 
  def stateless=(value)
    @stateless = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `time_zone` attribute.
  # 
  def time_zone
    return @time_zone
  end
  
  # 
  # Sets the value of the `time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @time_zone = value
  end
  
  # 
  # Returns the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration
    return @tunnel_migration
  end
  
  # 
  # Sets the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration=(value)
    @tunnel_migration = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `usb` attribute.
  # 
  def usb
    return @usb
  end
  
  # 
  # Sets the value of the `usb` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Usb} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def usb=(value)
    if value.is_a?(Hash)
      value = Usb.new(value)
    end
    @usb = value
  end
  
  # 
  # Returns the value of the `virtio_scsi` attribute.
  # 
  def virtio_scsi
    return @virtio_scsi
  end
  
  # 
  # Sets the value of the `virtio_scsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VirtioScsi} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtio_scsi=(value)
    if value.is_a?(Hash)
      value = VirtioScsi.new(value)
    end
    @virtio_scsi = value
  end
  
  # 
  # Creates a new instance of the {VmBase} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :bios The value of attribute `bios`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :console The value of attribute `console`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts [Hash] :cpu_profile The value of attribute `cpu_profile`.
  # 
  # @option opts :cpu_shares The value of attribute `cpu_shares`.
  # 
  # @option opts :creation_time The value of attribute `creation_time`.
  # 
  # @option opts [Hash] :custom_compatibility_version The value of attribute `custom_compatibility_version`.
  # 
  # @option opts :custom_cpu_model The value of attribute `custom_cpu_model`.
  # 
  # @option opts :custom_emulated_machine The value of attribute `custom_emulated_machine`.
  # 
  # @option opts [Array<Hash>] :custom_properties The values of attribute `custom_properties`.
  # 
  # @option opts :delete_protected The value of attribute `delete_protected`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts [Hash] :high_availability The value of attribute `high_availability`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :initialization The value of attribute `initialization`.
  # 
  # @option opts [Hash] :io The value of attribute `io`.
  # 
  # @option opts [Hash] :large_icon The value of attribute `large_icon`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts [Hash] :memory_policy The value of attribute `memory_policy`.
  # 
  # @option opts [Hash] :migration The value of attribute `migration`.
  # 
  # @option opts :migration_downtime The value of attribute `migration_downtime`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :origin The value of attribute `origin`.
  # 
  # @option opts [Hash] :os The value of attribute `os`.
  # 
  # @option opts [Hash] :rng_device The value of attribute `rng_device`.
  # 
  # @option opts [Hash] :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Hash] :small_icon The value of attribute `small_icon`.
  # 
  # @option opts :soundcard_enabled The value of attribute `soundcard_enabled`.
  # 
  # @option opts [Hash] :sso The value of attribute `sso`.
  # 
  # @option opts :start_paused The value of attribute `start_paused`.
  # 
  # @option opts :stateless The value of attribute `stateless`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Hash] :time_zone The value of attribute `time_zone`.
  # 
  # @option opts :tunnel_migration The value of attribute `tunnel_migration`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :usb The value of attribute `usb`.
  # 
  # @option opts [Hash] :virtio_scsi The value of attribute `virtio_scsi`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.bios = opts[:bios]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.console = opts[:console]
    self.cpu = opts[:cpu]
    self.cpu_profile = opts[:cpu_profile]
    self.cpu_shares = opts[:cpu_shares]
    self.creation_time = opts[:creation_time]
    self.custom_compatibility_version = opts[:custom_compatibility_version]
    self.custom_cpu_model = opts[:custom_cpu_model]
    self.custom_emulated_machine = opts[:custom_emulated_machine]
    self.custom_properties = opts[:custom_properties]
    self.delete_protected = opts[:delete_protected]
    self.description = opts[:description]
    self.display = opts[:display]
    self.domain = opts[:domain]
    self.high_availability = opts[:high_availability]
    self.id = opts[:id]
    self.initialization = opts[:initialization]
    self.io = opts[:io]
    self.large_icon = opts[:large_icon]
    self.memory = opts[:memory]
    self.memory_policy = opts[:memory_policy]
    self.migration = opts[:migration]
    self.migration_downtime = opts[:migration_downtime]
    self.name = opts[:name]
    self.origin = opts[:origin]
    self.os = opts[:os]
    self.rng_device = opts[:rng_device]
    self.serial_number = opts[:serial_number]
    self.small_icon = opts[:small_icon]
    self.soundcard_enabled = opts[:soundcard_enabled]
    self.sso = opts[:sso]
    self.start_paused = opts[:start_paused]
    self.stateless = opts[:stateless]
    self.storage_domain = opts[:storage_domain]
    self.time_zone = opts[:time_zone]
    self.tunnel_migration = opts[:tunnel_migration]
    self.type = opts[:type]
    self.usb = opts[:usb]
    self.virtio_scsi = opts[:virtio_scsi]
  end
  
end

class VmPlacementPolicy < Struct
  
  # 
  # Returns the value of the `affinity` attribute.
  # 
  def affinity
    return @affinity
  end
  
  # 
  # Sets the value of the `affinity` attribute.
  # 
  def affinity=(value)
    @affinity = value
  end
  
  # 
  # Returns the value of the `hosts` attribute.
  # 
  def hosts
    return @hosts
  end
  
  # 
  # Sets the value of the `hosts` attribute.
  # 
  def hosts=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Host.new(value)
        end
      end
    end
    @hosts = list
  end
  
  # 
  # Creates a new instance of the {VmPlacementPolicy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :affinity The value of attribute `affinity`.
  # 
  # @option opts [Array<Hash>] :hosts The values of attribute `hosts`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.affinity = opts[:affinity]
    self.hosts = opts[:hosts]
  end
  
end

class VmPool < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `max_user_vms` attribute.
  # 
  def max_user_vms
    return @max_user_vms
  end
  
  # 
  # Sets the value of the `max_user_vms` attribute.
  # 
  def max_user_vms=(value)
    @max_user_vms = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `prestarted_vms` attribute.
  # 
  def prestarted_vms
    return @prestarted_vms
  end
  
  # 
  # Sets the value of the `prestarted_vms` attribute.
  # 
  def prestarted_vms=(value)
    @prestarted_vms = value
  end
  
  # 
  # Returns the value of the `rng_device` attribute.
  # 
  def rng_device
    return @rng_device
  end
  
  # 
  # Sets the value of the `rng_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::RngDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rng_device=(value)
    if value.is_a?(Hash)
      value = RngDevice.new(value)
    end
    @rng_device = value
  end
  
  # 
  # Returns the value of the `size` attribute.
  # 
  def size
    return @size
  end
  
  # 
  # Sets the value of the `size` attribute.
  # 
  def size=(value)
    @size = value
  end
  
  # 
  # Returns the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled
    return @soundcard_enabled
  end
  
  # 
  # Sets the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled=(value)
    @soundcard_enabled = value
  end
  
  # 
  # Returns the value of the `stateful` attribute.
  # 
  def stateful
    return @stateful
  end
  
  # 
  # Sets the value of the `stateful` attribute.
  # 
  def stateful=(value)
    @stateful = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `use_latest_template_version` attribute.
  # 
  def use_latest_template_version
    return @use_latest_template_version
  end
  
  # 
  # Sets the value of the `use_latest_template_version` attribute.
  # 
  def use_latest_template_version=(value)
    @use_latest_template_version = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {VmPool} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :max_user_vms The value of attribute `max_user_vms`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :prestarted_vms The value of attribute `prestarted_vms`.
  # 
  # @option opts [Hash] :rng_device The value of attribute `rng_device`.
  # 
  # @option opts :size The value of attribute `size`.
  # 
  # @option opts :soundcard_enabled The value of attribute `soundcard_enabled`.
  # 
  # @option opts :stateful The value of attribute `stateful`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :use_latest_template_version The value of attribute `use_latest_template_version`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.display = opts[:display]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.max_user_vms = opts[:max_user_vms]
    self.name = opts[:name]
    self.permissions = opts[:permissions]
    self.prestarted_vms = opts[:prestarted_vms]
    self.rng_device = opts[:rng_device]
    self.size = opts[:size]
    self.soundcard_enabled = opts[:soundcard_enabled]
    self.stateful = opts[:stateful]
    self.template = opts[:template]
    self.type = opts[:type]
    self.use_latest_template_version = opts[:use_latest_template_version]
    self.vm = opts[:vm]
  end
  
end

class VmSummary < Struct
  
  # 
  # Returns the value of the `active` attribute.
  # 
  def active
    return @active
  end
  
  # 
  # Sets the value of the `active` attribute.
  # 
  def active=(value)
    @active = value
  end
  
  # 
  # Returns the value of the `migrating` attribute.
  # 
  def migrating
    return @migrating
  end
  
  # 
  # Sets the value of the `migrating` attribute.
  # 
  def migrating=(value)
    @migrating = value
  end
  
  # 
  # Returns the value of the `total` attribute.
  # 
  def total
    return @total
  end
  
  # 
  # Sets the value of the `total` attribute.
  # 
  def total=(value)
    @total = value
  end
  
  # 
  # Creates a new instance of the {VmSummary} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active The value of attribute `active`.
  # 
  # @option opts :migrating The value of attribute `migrating`.
  # 
  # @option opts :total The value of attribute `total`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active = opts[:active]
    self.migrating = opts[:migrating]
    self.total = opts[:total]
  end
  
end

class VnicPassThrough < Struct
  
  # 
  # Returns the value of the `mode` attribute.
  # 
  def mode
    return @mode
  end
  
  # 
  # Sets the value of the `mode` attribute.
  # 
  def mode=(value)
    @mode = value
  end
  
  # 
  # Creates a new instance of the {VnicPassThrough} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :mode The value of attribute `mode`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.mode = opts[:mode]
  end
  
end

class VnicProfile < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `custom_properties` attribute.
  # 
  def custom_properties
    return @custom_properties
  end
  
  # 
  # Sets the value of the `custom_properties` attribute.
  # 
  def custom_properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CustomProperty.new(value)
        end
      end
    end
    @custom_properties = list
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network` attribute.
  # 
  def network
    return @network
  end
  
  # 
  # Sets the value of the `network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Network} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network=(value)
    if value.is_a?(Hash)
      value = Network.new(value)
    end
    @network = value
  end
  
  # 
  # Returns the value of the `network_filter` attribute.
  # 
  def network_filter
    return @network_filter
  end
  
  # 
  # Sets the value of the `network_filter` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::NetworkFilter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network_filter=(value)
    if value.is_a?(Hash)
      value = NetworkFilter.new(value)
    end
    @network_filter = value
  end
  
  # 
  # Returns the value of the `pass_through` attribute.
  # 
  def pass_through
    return @pass_through
  end
  
  # 
  # Sets the value of the `pass_through` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VnicPassThrough} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def pass_through=(value)
    if value.is_a?(Hash)
      value = VnicPassThrough.new(value)
    end
    @pass_through = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `port_mirroring` attribute.
  # 
  def port_mirroring
    return @port_mirroring
  end
  
  # 
  # Sets the value of the `port_mirroring` attribute.
  # 
  def port_mirroring=(value)
    @port_mirroring = value
  end
  
  # 
  # Returns the value of the `qos` attribute.
  # 
  def qos
    return @qos
  end
  
  # 
  # Sets the value of the `qos` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Qos} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def qos=(value)
    if value.is_a?(Hash)
      value = Qos.new(value)
    end
    @qos = value
  end
  
  # 
  # Creates a new instance of the {VnicProfile} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Array<Hash>] :custom_properties The values of attribute `custom_properties`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :network The value of attribute `network`.
  # 
  # @option opts [Hash] :network_filter The value of attribute `network_filter`.
  # 
  # @option opts [Hash] :pass_through The value of attribute `pass_through`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :port_mirroring The value of attribute `port_mirroring`.
  # 
  # @option opts [Hash] :qos The value of attribute `qos`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.custom_properties = opts[:custom_properties]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.network = opts[:network]
    self.network_filter = opts[:network_filter]
    self.pass_through = opts[:pass_through]
    self.permissions = opts[:permissions]
    self.port_mirroring = opts[:port_mirroring]
    self.qos = opts[:qos]
  end
  
end

class VolumeGroup < Struct
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `logical_units` attribute.
  # 
  def logical_units
    return @logical_units
  end
  
  # 
  # Sets the value of the `logical_units` attribute.
  # 
  def logical_units=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = LogicalUnit.new(value)
        end
      end
    end
    @logical_units = list
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {VolumeGroup} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Array<Hash>] :logical_units The values of attribute `logical_units`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.id = opts[:id]
    self.logical_units = opts[:logical_units]
    self.name = opts[:name]
  end
  
end

class Weight < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `factor` attribute.
  # 
  def factor
    return @factor
  end
  
  # 
  # Sets the value of the `factor` attribute.
  # 
  def factor=(value)
    @factor = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `scheduling_policy` attribute.
  # 
  def scheduling_policy
    return @scheduling_policy
  end
  
  # 
  # Sets the value of the `scheduling_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SchedulingPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def scheduling_policy=(value)
    if value.is_a?(Hash)
      value = SchedulingPolicy.new(value)
    end
    @scheduling_policy = value
  end
  
  # 
  # Returns the value of the `scheduling_policy_unit` attribute.
  # 
  def scheduling_policy_unit
    return @scheduling_policy_unit
  end
  
  # 
  # Sets the value of the `scheduling_policy_unit` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SchedulingPolicyUnit} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def scheduling_policy_unit=(value)
    if value.is_a?(Hash)
      value = SchedulingPolicyUnit.new(value)
    end
    @scheduling_policy_unit = value
  end
  
  # 
  # Creates a new instance of the {Weight} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :factor The value of attribute `factor`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :scheduling_policy The value of attribute `scheduling_policy`.
  # 
  # @option opts [Hash] :scheduling_policy_unit The value of attribute `scheduling_policy_unit`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.factor = opts[:factor]
    self.id = opts[:id]
    self.name = opts[:name]
    self.scheduling_policy = opts[:scheduling_policy]
    self.scheduling_policy_unit = opts[:scheduling_policy_unit]
  end
  
end

class Action < Identified
  
  # 
  # Returns the value of the `async` attribute.
  # 
  def async
    return @async
  end
  
  # 
  # Sets the value of the `async` attribute.
  # 
  def async=(value)
    @async = value
  end
  
  # 
  # Returns the value of the `bricks` attribute.
  # 
  def bricks
    return @bricks
  end
  
  # 
  # Sets the value of the `bricks` attribute.
  # 
  def bricks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterBrick.new(value)
        end
      end
    end
    @bricks = list
  end
  
  # 
  # Returns the value of the `certificates` attribute.
  # 
  def certificates
    return @certificates
  end
  
  # 
  # Sets the value of the `certificates` attribute.
  # 
  def certificates=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Certificate.new(value)
        end
      end
    end
    @certificates = list
  end
  
  # 
  # Returns the value of the `check_connectivity` attribute.
  # 
  def check_connectivity
    return @check_connectivity
  end
  
  # 
  # Sets the value of the `check_connectivity` attribute.
  # 
  def check_connectivity=(value)
    @check_connectivity = value
  end
  
  # 
  # Returns the value of the `clone` attribute.
  # 
  def clone
    return @clone
  end
  
  # 
  # Sets the value of the `clone` attribute.
  # 
  def clone=(value)
    @clone = value
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `collapse_snapshots` attribute.
  # 
  def collapse_snapshots
    return @collapse_snapshots
  end
  
  # 
  # Sets the value of the `collapse_snapshots` attribute.
  # 
  def collapse_snapshots=(value)
    @collapse_snapshots = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `connectivity_timeout` attribute.
  # 
  def connectivity_timeout
    return @connectivity_timeout
  end
  
  # 
  # Sets the value of the `connectivity_timeout` attribute.
  # 
  def connectivity_timeout=(value)
    @connectivity_timeout = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `deploy_hosted_engine` attribute.
  # 
  def deploy_hosted_engine
    return @deploy_hosted_engine
  end
  
  # 
  # Sets the value of the `deploy_hosted_engine` attribute.
  # 
  def deploy_hosted_engine=(value)
    @deploy_hosted_engine = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `details` attribute.
  # 
  def details
    return @details
  end
  
  # 
  # Sets the value of the `details` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GlusterVolumeProfileDetails} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def details=(value)
    if value.is_a?(Hash)
      value = GlusterVolumeProfileDetails.new(value)
    end
    @details = value
  end
  
  # 
  # Returns the value of the `discard_snapshots` attribute.
  # 
  def discard_snapshots
    return @discard_snapshots
  end
  
  # 
  # Sets the value of the `discard_snapshots` attribute.
  # 
  def discard_snapshots=(value)
    @discard_snapshots = value
  end
  
  # 
  # Returns the value of the `disk` attribute.
  # 
  def disk
    return @disk
  end
  
  # 
  # Sets the value of the `disk` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Disk} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk=(value)
    if value.is_a?(Hash)
      value = Disk.new(value)
    end
    @disk = value
  end
  
  # 
  # Returns the value of the `disks` attribute.
  # 
  def disks
    return @disks
  end
  
  # 
  # Sets the value of the `disks` attribute.
  # 
  def disks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Disk.new(value)
        end
      end
    end
    @disks = list
  end
  
  # 
  # Returns the value of the `exclusive` attribute.
  # 
  def exclusive
    return @exclusive
  end
  
  # 
  # Sets the value of the `exclusive` attribute.
  # 
  def exclusive=(value)
    @exclusive = value
  end
  
  # 
  # Returns the value of the `exclussive` attribute.
  # 
  def exclussive
    return @exclussive
  end
  
  # 
  # Sets the value of the `exclussive` attribute.
  # 
  def exclussive=(value)
    @exclussive = value
  end
  
  # 
  # Returns the value of the `fault` attribute.
  # 
  def fault
    return @fault
  end
  
  # 
  # Sets the value of the `fault` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Fault} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def fault=(value)
    if value.is_a?(Hash)
      value = Fault.new(value)
    end
    @fault = value
  end
  
  # 
  # Returns the value of the `fence_type` attribute.
  # 
  def fence_type
    return @fence_type
  end
  
  # 
  # Sets the value of the `fence_type` attribute.
  # 
  def fence_type=(value)
    @fence_type = value
  end
  
  # 
  # Returns the value of the `filter` attribute.
  # 
  def filter
    return @filter
  end
  
  # 
  # Sets the value of the `filter` attribute.
  # 
  def filter=(value)
    @filter = value
  end
  
  # 
  # Returns the value of the `fix_layout` attribute.
  # 
  def fix_layout
    return @fix_layout
  end
  
  # 
  # Sets the value of the `fix_layout` attribute.
  # 
  def fix_layout=(value)
    @fix_layout = value
  end
  
  # 
  # Returns the value of the `force` attribute.
  # 
  def force
    return @force
  end
  
  # 
  # Sets the value of the `force` attribute.
  # 
  def force=(value)
    @force = value
  end
  
  # 
  # Returns the value of the `grace_period` attribute.
  # 
  def grace_period
    return @grace_period
  end
  
  # 
  # Sets the value of the `grace_period` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GracePeriod} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def grace_period=(value)
    if value.is_a?(Hash)
      value = GracePeriod.new(value)
    end
    @grace_period = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `image` attribute.
  # 
  def image
    return @image
  end
  
  # 
  # Sets the value of the `image` attribute.
  # 
  def image=(value)
    @image = value
  end
  
  # 
  # Returns the value of the `import_as_template` attribute.
  # 
  def import_as_template
    return @import_as_template
  end
  
  # 
  # Sets the value of the `import_as_template` attribute.
  # 
  def import_as_template=(value)
    @import_as_template = value
  end
  
  # 
  # Returns the value of the `is_attached` attribute.
  # 
  def is_attached
    return @is_attached
  end
  
  # 
  # Sets the value of the `is_attached` attribute.
  # 
  def is_attached=(value)
    @is_attached = value
  end
  
  # 
  # Returns the value of the `iscsi` attribute.
  # 
  def iscsi
    return @iscsi
  end
  
  # 
  # Sets the value of the `iscsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::IscsiDetails} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def iscsi=(value)
    if value.is_a?(Hash)
      value = IscsiDetails.new(value)
    end
    @iscsi = value
  end
  
  # 
  # Returns the value of the `iscsi_targets` attribute.
  # 
  def iscsi_targets
    return @iscsi_targets
  end
  
  # 
  # Sets the value of the `iscsi_targets` attribute.
  # 
  def iscsi_targets=(list)
    @iscsi_targets = list
  end
  
  # 
  # Returns the value of the `job` attribute.
  # 
  def job
    return @job
  end
  
  # 
  # Sets the value of the `job` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Job} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def job=(value)
    if value.is_a?(Hash)
      value = Job.new(value)
    end
    @job = value
  end
  
  # 
  # Returns the value of the `logical_units` attribute.
  # 
  def logical_units
    return @logical_units
  end
  
  # 
  # Sets the value of the `logical_units` attribute.
  # 
  def logical_units=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = LogicalUnit.new(value)
        end
      end
    end
    @logical_units = list
  end
  
  # 
  # Returns the value of the `maintenance_enabled` attribute.
  # 
  def maintenance_enabled
    return @maintenance_enabled
  end
  
  # 
  # Sets the value of the `maintenance_enabled` attribute.
  # 
  def maintenance_enabled=(value)
    @maintenance_enabled = value
  end
  
  # 
  # Returns the value of the `modified_bonds` attribute.
  # 
  def modified_bonds
    return @modified_bonds
  end
  
  # 
  # Sets the value of the `modified_bonds` attribute.
  # 
  def modified_bonds=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = HostNic.new(value)
        end
      end
    end
    @modified_bonds = list
  end
  
  # 
  # Returns the value of the `modified_labels` attribute.
  # 
  def modified_labels
    return @modified_labels
  end
  
  # 
  # Sets the value of the `modified_labels` attribute.
  # 
  def modified_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkLabel.new(value)
        end
      end
    end
    @modified_labels = list
  end
  
  # 
  # Returns the value of the `modified_network_attachments` attribute.
  # 
  def modified_network_attachments
    return @modified_network_attachments
  end
  
  # 
  # Sets the value of the `modified_network_attachments` attribute.
  # 
  def modified_network_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkAttachment.new(value)
        end
      end
    end
    @modified_network_attachments = list
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `option` attribute.
  # 
  def option
    return @option
  end
  
  # 
  # Sets the value of the `option` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Option} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def option=(value)
    if value.is_a?(Hash)
      value = Option.new(value)
    end
    @option = value
  end
  
  # 
  # Returns the value of the `pause` attribute.
  # 
  def pause
    return @pause
  end
  
  # 
  # Sets the value of the `pause` attribute.
  # 
  def pause=(value)
    @pause = value
  end
  
  # 
  # Returns the value of the `power_management` attribute.
  # 
  def power_management
    return @power_management
  end
  
  # 
  # Sets the value of the `power_management` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::PowerManagement} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def power_management=(value)
    if value.is_a?(Hash)
      value = PowerManagement.new(value)
    end
    @power_management = value
  end
  
  # 
  # Returns the value of the `proxy_ticket` attribute.
  # 
  def proxy_ticket
    return @proxy_ticket
  end
  
  # 
  # Sets the value of the `proxy_ticket` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ProxyTicket} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def proxy_ticket=(value)
    if value.is_a?(Hash)
      value = ProxyTicket.new(value)
    end
    @proxy_ticket = value
  end
  
  # 
  # Returns the value of the `reason` attribute.
  # 
  def reason
    return @reason
  end
  
  # 
  # Sets the value of the `reason` attribute.
  # 
  def reason=(value)
    @reason = value
  end
  
  # 
  # Returns the value of the `removed_bonds` attribute.
  # 
  def removed_bonds
    return @removed_bonds
  end
  
  # 
  # Sets the value of the `removed_bonds` attribute.
  # 
  def removed_bonds=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = HostNic.new(value)
        end
      end
    end
    @removed_bonds = list
  end
  
  # 
  # Returns the value of the `removed_labels` attribute.
  # 
  def removed_labels
    return @removed_labels
  end
  
  # 
  # Sets the value of the `removed_labels` attribute.
  # 
  def removed_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkLabel.new(value)
        end
      end
    end
    @removed_labels = list
  end
  
  # 
  # Returns the value of the `removed_network_attachments` attribute.
  # 
  def removed_network_attachments
    return @removed_network_attachments
  end
  
  # 
  # Sets the value of the `removed_network_attachments` attribute.
  # 
  def removed_network_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkAttachment.new(value)
        end
      end
    end
    @removed_network_attachments = list
  end
  
  # 
  # Returns the value of the `resolution_type` attribute.
  # 
  def resolution_type
    return @resolution_type
  end
  
  # 
  # Sets the value of the `resolution_type` attribute.
  # 
  def resolution_type=(value)
    @resolution_type = value
  end
  
  # 
  # Returns the value of the `restore_memory` attribute.
  # 
  def restore_memory
    return @restore_memory
  end
  
  # 
  # Sets the value of the `restore_memory` attribute.
  # 
  def restore_memory=(value)
    @restore_memory = value
  end
  
  # 
  # Returns the value of the `root_password` attribute.
  # 
  def root_password
    return @root_password
  end
  
  # 
  # Sets the value of the `root_password` attribute.
  # 
  def root_password=(value)
    @root_password = value
  end
  
  # 
  # Returns the value of the `snapshot` attribute.
  # 
  def snapshot
    return @snapshot
  end
  
  # 
  # Sets the value of the `snapshot` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Snapshot} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def snapshot=(value)
    if value.is_a?(Hash)
      value = Snapshot.new(value)
    end
    @snapshot = value
  end
  
  # 
  # Returns the value of the `ssh` attribute.
  # 
  def ssh
    return @ssh
  end
  
  # 
  # Sets the value of the `ssh` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ssh} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ssh=(value)
    if value.is_a?(Hash)
      value = Ssh.new(value)
    end
    @ssh = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `stop_gluster_service` attribute.
  # 
  def stop_gluster_service
    return @stop_gluster_service
  end
  
  # 
  # Sets the value of the `stop_gluster_service` attribute.
  # 
  def stop_gluster_service=(value)
    @stop_gluster_service = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `storage_domains` attribute.
  # 
  def storage_domains
    return @storage_domains
  end
  
  # 
  # Sets the value of the `storage_domains` attribute.
  # 
  def storage_domains=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageDomain.new(value)
        end
      end
    end
    @storage_domains = list
  end
  
  # 
  # Returns the value of the `succeeded` attribute.
  # 
  def succeeded
    return @succeeded
  end
  
  # 
  # Sets the value of the `succeeded` attribute.
  # 
  def succeeded=(value)
    @succeeded = value
  end
  
  # 
  # Returns the value of the `synchronized_network_attachments` attribute.
  # 
  def synchronized_network_attachments
    return @synchronized_network_attachments
  end
  
  # 
  # Sets the value of the `synchronized_network_attachments` attribute.
  # 
  def synchronized_network_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkAttachment.new(value)
        end
      end
    end
    @synchronized_network_attachments = list
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `ticket` attribute.
  # 
  def ticket
    return @ticket
  end
  
  # 
  # Sets the value of the `ticket` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ticket} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ticket=(value)
    if value.is_a?(Hash)
      value = Ticket.new(value)
    end
    @ticket = value
  end
  
  # 
  # Returns the value of the `undeploy_hosted_engine` attribute.
  # 
  def undeploy_hosted_engine
    return @undeploy_hosted_engine
  end
  
  # 
  # Sets the value of the `undeploy_hosted_engine` attribute.
  # 
  def undeploy_hosted_engine=(value)
    @undeploy_hosted_engine = value
  end
  
  # 
  # Returns the value of the `use_cloud_init` attribute.
  # 
  def use_cloud_init
    return @use_cloud_init
  end
  
  # 
  # Sets the value of the `use_cloud_init` attribute.
  # 
  def use_cloud_init=(value)
    @use_cloud_init = value
  end
  
  # 
  # Returns the value of the `use_sysprep` attribute.
  # 
  def use_sysprep
    return @use_sysprep
  end
  
  # 
  # Sets the value of the `use_sysprep` attribute.
  # 
  def use_sysprep=(value)
    @use_sysprep = value
  end
  
  # 
  # Returns the value of the `virtual_functions_configuration` attribute.
  # 
  def virtual_functions_configuration
    return @virtual_functions_configuration
  end
  
  # 
  # Sets the value of the `virtual_functions_configuration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNicVirtualFunctionsConfiguration} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtual_functions_configuration=(value)
    if value.is_a?(Hash)
      value = HostNicVirtualFunctionsConfiguration.new(value)
    end
    @virtual_functions_configuration = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {Action} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :async The value of attribute `async`.
  # 
  # @option opts [Array<Hash>] :bricks The values of attribute `bricks`.
  # 
  # @option opts [Array<Hash>] :certificates The values of attribute `certificates`.
  # 
  # @option opts :check_connectivity The value of attribute `check_connectivity`.
  # 
  # @option opts :clone The value of attribute `clone`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :collapse_snapshots The value of attribute `collapse_snapshots`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :connectivity_timeout The value of attribute `connectivity_timeout`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :deploy_hosted_engine The value of attribute `deploy_hosted_engine`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :details The value of attribute `details`.
  # 
  # @option opts :discard_snapshots The value of attribute `discard_snapshots`.
  # 
  # @option opts [Hash] :disk The value of attribute `disk`.
  # 
  # @option opts [Array<Hash>] :disks The values of attribute `disks`.
  # 
  # @option opts :exclusive The value of attribute `exclusive`.
  # 
  # @option opts :exclussive The value of attribute `exclussive`.
  # 
  # @option opts [Hash] :fault The value of attribute `fault`.
  # 
  # @option opts :fence_type The value of attribute `fence_type`.
  # 
  # @option opts :filter The value of attribute `filter`.
  # 
  # @option opts :fix_layout The value of attribute `fix_layout`.
  # 
  # @option opts :force The value of attribute `force`.
  # 
  # @option opts [Hash] :grace_period The value of attribute `grace_period`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :image The value of attribute `image`.
  # 
  # @option opts :import_as_template The value of attribute `import_as_template`.
  # 
  # @option opts :is_attached The value of attribute `is_attached`.
  # 
  # @option opts [Hash] :iscsi The value of attribute `iscsi`.
  # 
  # @option opts [Array<Hash>] :iscsi_targets The values of attribute `iscsi_targets`.
  # 
  # @option opts [Hash] :job The value of attribute `job`.
  # 
  # @option opts [Array<Hash>] :logical_units The values of attribute `logical_units`.
  # 
  # @option opts :maintenance_enabled The value of attribute `maintenance_enabled`.
  # 
  # @option opts [Array<Hash>] :modified_bonds The values of attribute `modified_bonds`.
  # 
  # @option opts [Array<Hash>] :modified_labels The values of attribute `modified_labels`.
  # 
  # @option opts [Array<Hash>] :modified_network_attachments The values of attribute `modified_network_attachments`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :option The value of attribute `option`.
  # 
  # @option opts :pause The value of attribute `pause`.
  # 
  # @option opts [Hash] :power_management The value of attribute `power_management`.
  # 
  # @option opts [Hash] :proxy_ticket The value of attribute `proxy_ticket`.
  # 
  # @option opts :reason The value of attribute `reason`.
  # 
  # @option opts [Array<Hash>] :removed_bonds The values of attribute `removed_bonds`.
  # 
  # @option opts [Array<Hash>] :removed_labels The values of attribute `removed_labels`.
  # 
  # @option opts [Array<Hash>] :removed_network_attachments The values of attribute `removed_network_attachments`.
  # 
  # @option opts :resolution_type The value of attribute `resolution_type`.
  # 
  # @option opts :restore_memory The value of attribute `restore_memory`.
  # 
  # @option opts :root_password The value of attribute `root_password`.
  # 
  # @option opts [Hash] :snapshot The value of attribute `snapshot`.
  # 
  # @option opts [Hash] :ssh The value of attribute `ssh`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :stop_gluster_service The value of attribute `stop_gluster_service`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :storage_domains The values of attribute `storage_domains`.
  # 
  # @option opts :succeeded The value of attribute `succeeded`.
  # 
  # @option opts [Array<Hash>] :synchronized_network_attachments The values of attribute `synchronized_network_attachments`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :ticket The value of attribute `ticket`.
  # 
  # @option opts :undeploy_hosted_engine The value of attribute `undeploy_hosted_engine`.
  # 
  # @option opts :use_cloud_init The value of attribute `use_cloud_init`.
  # 
  # @option opts :use_sysprep The value of attribute `use_sysprep`.
  # 
  # @option opts [Hash] :virtual_functions_configuration The value of attribute `virtual_functions_configuration`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.async = opts[:async]
    self.bricks = opts[:bricks]
    self.certificates = opts[:certificates]
    self.check_connectivity = opts[:check_connectivity]
    self.clone = opts[:clone]
    self.cluster = opts[:cluster]
    self.collapse_snapshots = opts[:collapse_snapshots]
    self.comment = opts[:comment]
    self.connectivity_timeout = opts[:connectivity_timeout]
    self.data_center = opts[:data_center]
    self.deploy_hosted_engine = opts[:deploy_hosted_engine]
    self.description = opts[:description]
    self.details = opts[:details]
    self.discard_snapshots = opts[:discard_snapshots]
    self.disk = opts[:disk]
    self.disks = opts[:disks]
    self.exclusive = opts[:exclusive]
    self.exclussive = opts[:exclussive]
    self.fault = opts[:fault]
    self.fence_type = opts[:fence_type]
    self.filter = opts[:filter]
    self.fix_layout = opts[:fix_layout]
    self.force = opts[:force]
    self.grace_period = opts[:grace_period]
    self.host = opts[:host]
    self.id = opts[:id]
    self.image = opts[:image]
    self.import_as_template = opts[:import_as_template]
    self.is_attached = opts[:is_attached]
    self.iscsi = opts[:iscsi]
    self.iscsi_targets = opts[:iscsi_targets]
    self.job = opts[:job]
    self.logical_units = opts[:logical_units]
    self.maintenance_enabled = opts[:maintenance_enabled]
    self.modified_bonds = opts[:modified_bonds]
    self.modified_labels = opts[:modified_labels]
    self.modified_network_attachments = opts[:modified_network_attachments]
    self.name = opts[:name]
    self.option = opts[:option]
    self.pause = opts[:pause]
    self.power_management = opts[:power_management]
    self.proxy_ticket = opts[:proxy_ticket]
    self.reason = opts[:reason]
    self.removed_bonds = opts[:removed_bonds]
    self.removed_labels = opts[:removed_labels]
    self.removed_network_attachments = opts[:removed_network_attachments]
    self.resolution_type = opts[:resolution_type]
    self.restore_memory = opts[:restore_memory]
    self.root_password = opts[:root_password]
    self.snapshot = opts[:snapshot]
    self.ssh = opts[:ssh]
    self.status = opts[:status]
    self.stop_gluster_service = opts[:stop_gluster_service]
    self.storage_domain = opts[:storage_domain]
    self.storage_domains = opts[:storage_domains]
    self.succeeded = opts[:succeeded]
    self.synchronized_network_attachments = opts[:synchronized_network_attachments]
    self.template = opts[:template]
    self.ticket = opts[:ticket]
    self.undeploy_hosted_engine = opts[:undeploy_hosted_engine]
    self.use_cloud_init = opts[:use_cloud_init]
    self.use_sysprep = opts[:use_sysprep]
    self.virtual_functions_configuration = opts[:virtual_functions_configuration]
    self.vm = opts[:vm]
  end
  
end

class AffinityGroup < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `enforcing` attribute.
  # 
  def enforcing
    return @enforcing
  end
  
  # 
  # Sets the value of the `enforcing` attribute.
  # 
  def enforcing=(value)
    @enforcing = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `positive` attribute.
  # 
  def positive
    return @positive
  end
  
  # 
  # Sets the value of the `positive` attribute.
  # 
  def positive=(value)
    @positive = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {AffinityGroup} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :enforcing The value of attribute `enforcing`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :positive The value of attribute `positive`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.enforcing = opts[:enforcing]
    self.id = opts[:id]
    self.name = opts[:name]
    self.positive = opts[:positive]
    self.vms = opts[:vms]
  end
  
end

class AffinityLabel < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `hosts` attribute.
  # 
  def hosts
    return @hosts
  end
  
  # 
  # Sets the value of the `hosts` attribute.
  # 
  def hosts=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Host.new(value)
        end
      end
    end
    @hosts = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `read_only` attribute.
  # 
  def read_only
    return @read_only
  end
  
  # 
  # Sets the value of the `read_only` attribute.
  # 
  def read_only=(value)
    @read_only = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {AffinityLabel} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :hosts The values of attribute `hosts`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :read_only The value of attribute `read_only`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.hosts = opts[:hosts]
    self.id = opts[:id]
    self.name = opts[:name]
    self.read_only = opts[:read_only]
    self.vms = opts[:vms]
  end
  
end

class Agent < Identified
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `concurrent` attribute.
  # 
  def concurrent
    return @concurrent
  end
  
  # 
  # Sets the value of the `concurrent` attribute.
  # 
  def concurrent=(value)
    @concurrent = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `encrypt_options` attribute.
  # 
  def encrypt_options
    return @encrypt_options
  end
  
  # 
  # Sets the value of the `encrypt_options` attribute.
  # 
  def encrypt_options=(value)
    @encrypt_options = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `options` attribute.
  # 
  def options
    return @options
  end
  
  # 
  # Sets the value of the `options` attribute.
  # 
  def options=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Option.new(value)
        end
      end
    end
    @options = list
  end
  
  # 
  # Returns the value of the `order` attribute.
  # 
  def order
    return @order
  end
  
  # 
  # Sets the value of the `order` attribute.
  # 
  def order=(value)
    @order = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {Agent} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :concurrent The value of attribute `concurrent`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :encrypt_options The value of attribute `encrypt_options`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :options The values of attribute `options`.
  # 
  # @option opts :order The value of attribute `order`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.comment = opts[:comment]
    self.concurrent = opts[:concurrent]
    self.description = opts[:description]
    self.encrypt_options = opts[:encrypt_options]
    self.host = opts[:host]
    self.id = opts[:id]
    self.name = opts[:name]
    self.options = opts[:options]
    self.order = opts[:order]
    self.password = opts[:password]
    self.port = opts[:port]
    self.type = opts[:type]
    self.username = opts[:username]
  end
  
end

class Application < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {Application} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.vm = opts[:vm]
  end
  
end

class AuthorizedKey < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `key` attribute.
  # 
  def key
    return @key
  end
  
  # 
  # Sets the value of the `key` attribute.
  # 
  def key=(value)
    @key = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Creates a new instance of the {AuthorizedKey} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :key The value of attribute `key`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.key = opts[:key]
    self.name = opts[:name]
    self.user = opts[:user]
  end
  
end

class Balance < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `scheduling_policy` attribute.
  # 
  def scheduling_policy
    return @scheduling_policy
  end
  
  # 
  # Sets the value of the `scheduling_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SchedulingPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def scheduling_policy=(value)
    if value.is_a?(Hash)
      value = SchedulingPolicy.new(value)
    end
    @scheduling_policy = value
  end
  
  # 
  # Returns the value of the `scheduling_policy_unit` attribute.
  # 
  def scheduling_policy_unit
    return @scheduling_policy_unit
  end
  
  # 
  # Sets the value of the `scheduling_policy_unit` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SchedulingPolicyUnit} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def scheduling_policy_unit=(value)
    if value.is_a?(Hash)
      value = SchedulingPolicyUnit.new(value)
    end
    @scheduling_policy_unit = value
  end
  
  # 
  # Creates a new instance of the {Balance} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :scheduling_policy The value of attribute `scheduling_policy`.
  # 
  # @option opts [Hash] :scheduling_policy_unit The value of attribute `scheduling_policy_unit`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.scheduling_policy = opts[:scheduling_policy]
    self.scheduling_policy_unit = opts[:scheduling_policy_unit]
  end
  
end

class Bookmark < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `value` attribute.
  # 
  def value
    return @value
  end
  
  # 
  # Sets the value of the `value` attribute.
  # 
  def value=(value)
    @value = value
  end
  
  # 
  # Creates a new instance of the {Bookmark} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :value The value of attribute `value`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.value = opts[:value]
  end
  
end

class BrickProfileDetail < EntityProfileDetail
  
  # 
  # Returns the value of the `brick` attribute.
  # 
  def brick
    return @brick
  end
  
  # 
  # Sets the value of the `brick` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GlusterBrick} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def brick=(value)
    if value.is_a?(Hash)
      value = GlusterBrick.new(value)
    end
    @brick = value
  end
  
  # 
  # Returns the value of the `profile_details` attribute.
  # 
  def profile_details
    return @profile_details
  end
  
  # 
  # Sets the value of the `profile_details` attribute.
  # 
  def profile_details=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ProfileDetail.new(value)
        end
      end
    end
    @profile_details = list
  end
  
  # 
  # Creates a new instance of the {BrickProfileDetail} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :brick The value of attribute `brick`.
  # 
  # @option opts [Array<Hash>] :profile_details The values of attribute `profile_details`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.brick = opts[:brick]
    self.profile_details = opts[:profile_details]
  end
  
end

class Certificate < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `content` attribute.
  # 
  def content
    return @content
  end
  
  # 
  # Sets the value of the `content` attribute.
  # 
  def content=(value)
    @content = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `organization` attribute.
  # 
  def organization
    return @organization
  end
  
  # 
  # Sets the value of the `organization` attribute.
  # 
  def organization=(value)
    @organization = value
  end
  
  # 
  # Returns the value of the `subject` attribute.
  # 
  def subject
    return @subject
  end
  
  # 
  # Sets the value of the `subject` attribute.
  # 
  def subject=(value)
    @subject = value
  end
  
  # 
  # Creates a new instance of the {Certificate} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :content The value of attribute `content`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :organization The value of attribute `organization`.
  # 
  # @option opts :subject The value of attribute `subject`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.content = opts[:content]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.organization = opts[:organization]
    self.subject = opts[:subject]
  end
  
end

class Cluster < Identified
  
  # 
  # Returns the value of the `affinity_groups` attribute.
  # 
  def affinity_groups
    return @affinity_groups
  end
  
  # 
  # Sets the value of the `affinity_groups` attribute.
  # 
  def affinity_groups=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = AffinityGroup.new(value)
        end
      end
    end
    @affinity_groups = list
  end
  
  # 
  # Returns the value of the `ballooning_enabled` attribute.
  # 
  def ballooning_enabled
    return @ballooning_enabled
  end
  
  # 
  # Sets the value of the `ballooning_enabled` attribute.
  # 
  def ballooning_enabled=(value)
    @ballooning_enabled = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `cpu_profiles` attribute.
  # 
  def cpu_profiles
    return @cpu_profiles
  end
  
  # 
  # Sets the value of the `cpu_profiles` attribute.
  # 
  def cpu_profiles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CpuProfile.new(value)
        end
      end
    end
    @cpu_profiles = list
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `error_handling` attribute.
  # 
  def error_handling
    return @error_handling
  end
  
  # 
  # Sets the value of the `error_handling` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ErrorHandling} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def error_handling=(value)
    if value.is_a?(Hash)
      value = ErrorHandling.new(value)
    end
    @error_handling = value
  end
  
  # 
  # Returns the value of the `fencing_policy` attribute.
  # 
  def fencing_policy
    return @fencing_policy
  end
  
  # 
  # Sets the value of the `fencing_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::FencingPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def fencing_policy=(value)
    if value.is_a?(Hash)
      value = FencingPolicy.new(value)
    end
    @fencing_policy = value
  end
  
  # 
  # Returns the value of the `gluster_hooks` attribute.
  # 
  def gluster_hooks
    return @gluster_hooks
  end
  
  # 
  # Sets the value of the `gluster_hooks` attribute.
  # 
  def gluster_hooks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterHook.new(value)
        end
      end
    end
    @gluster_hooks = list
  end
  
  # 
  # Returns the value of the `gluster_service` attribute.
  # 
  def gluster_service
    return @gluster_service
  end
  
  # 
  # Sets the value of the `gluster_service` attribute.
  # 
  def gluster_service=(value)
    @gluster_service = value
  end
  
  # 
  # Returns the value of the `gluster_volumes` attribute.
  # 
  def gluster_volumes
    return @gluster_volumes
  end
  
  # 
  # Sets the value of the `gluster_volumes` attribute.
  # 
  def gluster_volumes=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterVolume.new(value)
        end
      end
    end
    @gluster_volumes = list
  end
  
  # 
  # Returns the value of the `ha_reservation` attribute.
  # 
  def ha_reservation
    return @ha_reservation
  end
  
  # 
  # Sets the value of the `ha_reservation` attribute.
  # 
  def ha_reservation=(value)
    @ha_reservation = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ksm` attribute.
  # 
  def ksm
    return @ksm
  end
  
  # 
  # Sets the value of the `ksm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ksm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ksm=(value)
    if value.is_a?(Hash)
      value = Ksm.new(value)
    end
    @ksm = value
  end
  
  # 
  # Returns the value of the `mac_pool` attribute.
  # 
  def mac_pool
    return @mac_pool
  end
  
  # 
  # Sets the value of the `mac_pool` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MacPool} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def mac_pool=(value)
    if value.is_a?(Hash)
      value = MacPool.new(value)
    end
    @mac_pool = value
  end
  
  # 
  # Returns the value of the `maintenance_reason_required` attribute.
  # 
  def maintenance_reason_required
    return @maintenance_reason_required
  end
  
  # 
  # Sets the value of the `maintenance_reason_required` attribute.
  # 
  def maintenance_reason_required=(value)
    @maintenance_reason_required = value
  end
  
  # 
  # Returns the value of the `management_network` attribute.
  # 
  def management_network
    return @management_network
  end
  
  # 
  # Sets the value of the `management_network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Network} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def management_network=(value)
    if value.is_a?(Hash)
      value = Network.new(value)
    end
    @management_network = value
  end
  
  # 
  # Returns the value of the `memory_policy` attribute.
  # 
  def memory_policy
    return @memory_policy
  end
  
  # 
  # Sets the value of the `memory_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def memory_policy=(value)
    if value.is_a?(Hash)
      value = MemoryPolicy.new(value)
    end
    @memory_policy = value
  end
  
  # 
  # Returns the value of the `migration` attribute.
  # 
  def migration
    return @migration
  end
  
  # 
  # Sets the value of the `migration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationOptions} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def migration=(value)
    if value.is_a?(Hash)
      value = MigrationOptions.new(value)
    end
    @migration = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network_filters` attribute.
  # 
  def network_filters
    return @network_filters
  end
  
  # 
  # Sets the value of the `network_filters` attribute.
  # 
  def network_filters=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkFilter.new(value)
        end
      end
    end
    @network_filters = list
  end
  
  # 
  # Returns the value of the `networks` attribute.
  # 
  def networks
    return @networks
  end
  
  # 
  # Sets the value of the `networks` attribute.
  # 
  def networks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Network.new(value)
        end
      end
    end
    @networks = list
  end
  
  # 
  # Returns the value of the `optional_reason` attribute.
  # 
  def optional_reason
    return @optional_reason
  end
  
  # 
  # Sets the value of the `optional_reason` attribute.
  # 
  def optional_reason=(value)
    @optional_reason = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `required_rng_sources` attribute.
  # 
  def required_rng_sources
    return @required_rng_sources
  end
  
  # 
  # Sets the value of the `required_rng_sources` attribute.
  # 
  def required_rng_sources=(list)
    @required_rng_sources = list
  end
  
  # 
  # Returns the value of the `scheduling_policy` attribute.
  # 
  def scheduling_policy
    return @scheduling_policy
  end
  
  # 
  # Sets the value of the `scheduling_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SchedulingPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def scheduling_policy=(value)
    if value.is_a?(Hash)
      value = SchedulingPolicy.new(value)
    end
    @scheduling_policy = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SerialNumber} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def serial_number=(value)
    if value.is_a?(Hash)
      value = SerialNumber.new(value)
    end
    @serial_number = value
  end
  
  # 
  # Returns the value of the `supported_versions` attribute.
  # 
  def supported_versions
    return @supported_versions
  end
  
  # 
  # Sets the value of the `supported_versions` attribute.
  # 
  def supported_versions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Version.new(value)
        end
      end
    end
    @supported_versions = list
  end
  
  # 
  # Returns the value of the `switch_type` attribute.
  # 
  def switch_type
    return @switch_type
  end
  
  # 
  # Sets the value of the `switch_type` attribute.
  # 
  def switch_type=(value)
    @switch_type = value
  end
  
  # 
  # Returns the value of the `threads_as_cores` attribute.
  # 
  def threads_as_cores
    return @threads_as_cores
  end
  
  # 
  # Sets the value of the `threads_as_cores` attribute.
  # 
  def threads_as_cores=(value)
    @threads_as_cores = value
  end
  
  # 
  # Returns the value of the `trusted_service` attribute.
  # 
  def trusted_service
    return @trusted_service
  end
  
  # 
  # Sets the value of the `trusted_service` attribute.
  # 
  def trusted_service=(value)
    @trusted_service = value
  end
  
  # 
  # Returns the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration
    return @tunnel_migration
  end
  
  # 
  # Sets the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration=(value)
    @tunnel_migration = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Returns the value of the `virt_service` attribute.
  # 
  def virt_service
    return @virt_service
  end
  
  # 
  # Sets the value of the `virt_service` attribute.
  # 
  def virt_service=(value)
    @virt_service = value
  end
  
  # 
  # Creates a new instance of the {Cluster} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :affinity_groups The values of attribute `affinity_groups`.
  # 
  # @option opts :ballooning_enabled The value of attribute `ballooning_enabled`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts [Array<Hash>] :cpu_profiles The values of attribute `cpu_profiles`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :error_handling The value of attribute `error_handling`.
  # 
  # @option opts [Hash] :fencing_policy The value of attribute `fencing_policy`.
  # 
  # @option opts [Array<Hash>] :gluster_hooks The values of attribute `gluster_hooks`.
  # 
  # @option opts :gluster_service The value of attribute `gluster_service`.
  # 
  # @option opts [Array<Hash>] :gluster_volumes The values of attribute `gluster_volumes`.
  # 
  # @option opts :ha_reservation The value of attribute `ha_reservation`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :ksm The value of attribute `ksm`.
  # 
  # @option opts [Hash] :mac_pool The value of attribute `mac_pool`.
  # 
  # @option opts :maintenance_reason_required The value of attribute `maintenance_reason_required`.
  # 
  # @option opts [Hash] :management_network The value of attribute `management_network`.
  # 
  # @option opts [Hash] :memory_policy The value of attribute `memory_policy`.
  # 
  # @option opts [Hash] :migration The value of attribute `migration`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :network_filters The values of attribute `network_filters`.
  # 
  # @option opts [Array<Hash>] :networks The values of attribute `networks`.
  # 
  # @option opts :optional_reason The value of attribute `optional_reason`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Array<Hash>] :required_rng_sources The values of attribute `required_rng_sources`.
  # 
  # @option opts [Hash] :scheduling_policy The value of attribute `scheduling_policy`.
  # 
  # @option opts [Hash] :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Array<Hash>] :supported_versions The values of attribute `supported_versions`.
  # 
  # @option opts :switch_type The value of attribute `switch_type`.
  # 
  # @option opts :threads_as_cores The value of attribute `threads_as_cores`.
  # 
  # @option opts :trusted_service The value of attribute `trusted_service`.
  # 
  # @option opts :tunnel_migration The value of attribute `tunnel_migration`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # @option opts :virt_service The value of attribute `virt_service`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.affinity_groups = opts[:affinity_groups]
    self.ballooning_enabled = opts[:ballooning_enabled]
    self.comment = opts[:comment]
    self.cpu = opts[:cpu]
    self.cpu_profiles = opts[:cpu_profiles]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.display = opts[:display]
    self.error_handling = opts[:error_handling]
    self.fencing_policy = opts[:fencing_policy]
    self.gluster_hooks = opts[:gluster_hooks]
    self.gluster_service = opts[:gluster_service]
    self.gluster_volumes = opts[:gluster_volumes]
    self.ha_reservation = opts[:ha_reservation]
    self.id = opts[:id]
    self.ksm = opts[:ksm]
    self.mac_pool = opts[:mac_pool]
    self.maintenance_reason_required = opts[:maintenance_reason_required]
    self.management_network = opts[:management_network]
    self.memory_policy = opts[:memory_policy]
    self.migration = opts[:migration]
    self.name = opts[:name]
    self.network_filters = opts[:network_filters]
    self.networks = opts[:networks]
    self.optional_reason = opts[:optional_reason]
    self.permissions = opts[:permissions]
    self.required_rng_sources = opts[:required_rng_sources]
    self.scheduling_policy = opts[:scheduling_policy]
    self.serial_number = opts[:serial_number]
    self.supported_versions = opts[:supported_versions]
    self.switch_type = opts[:switch_type]
    self.threads_as_cores = opts[:threads_as_cores]
    self.trusted_service = opts[:trusted_service]
    self.tunnel_migration = opts[:tunnel_migration]
    self.version = opts[:version]
    self.virt_service = opts[:virt_service]
  end
  
end

class ClusterLevel < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `cpu_types` attribute.
  # 
  def cpu_types
    return @cpu_types
  end
  
  # 
  # Sets the value of the `cpu_types` attribute.
  # 
  def cpu_types=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CpuType.new(value)
        end
      end
    end
    @cpu_types = list
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {ClusterLevel} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Array<Hash>] :cpu_types The values of attribute `cpu_types`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.cpu_types = opts[:cpu_types]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class CpuProfile < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `qos` attribute.
  # 
  def qos
    return @qos
  end
  
  # 
  # Sets the value of the `qos` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Qos} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def qos=(value)
    if value.is_a?(Hash)
      value = Qos.new(value)
    end
    @qos = value
  end
  
  # 
  # Creates a new instance of the {CpuProfile} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Hash] :qos The value of attribute `qos`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.permissions = opts[:permissions]
    self.qos = opts[:qos]
  end
  
end

class DataCenter < Identified
  
  # 
  # Returns the value of the `clusters` attribute.
  # 
  def clusters
    return @clusters
  end
  
  # 
  # Sets the value of the `clusters` attribute.
  # 
  def clusters=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Cluster.new(value)
        end
      end
    end
    @clusters = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `iscsi_bonds` attribute.
  # 
  def iscsi_bonds
    return @iscsi_bonds
  end
  
  # 
  # Sets the value of the `iscsi_bonds` attribute.
  # 
  def iscsi_bonds=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = IscsiBond.new(value)
        end
      end
    end
    @iscsi_bonds = list
  end
  
  # 
  # Returns the value of the `local` attribute.
  # 
  def local
    return @local
  end
  
  # 
  # Sets the value of the `local` attribute.
  # 
  def local=(value)
    @local = value
  end
  
  # 
  # Returns the value of the `mac_pool` attribute.
  # 
  def mac_pool
    return @mac_pool
  end
  
  # 
  # Sets the value of the `mac_pool` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MacPool} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def mac_pool=(value)
    if value.is_a?(Hash)
      value = MacPool.new(value)
    end
    @mac_pool = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `networks` attribute.
  # 
  def networks
    return @networks
  end
  
  # 
  # Sets the value of the `networks` attribute.
  # 
  def networks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Network.new(value)
        end
      end
    end
    @networks = list
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `qoss` attribute.
  # 
  def qoss
    return @qoss
  end
  
  # 
  # Sets the value of the `qoss` attribute.
  # 
  def qoss=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Qos.new(value)
        end
      end
    end
    @qoss = list
  end
  
  # 
  # Returns the value of the `quota_mode` attribute.
  # 
  def quota_mode
    return @quota_mode
  end
  
  # 
  # Sets the value of the `quota_mode` attribute.
  # 
  def quota_mode=(value)
    @quota_mode = value
  end
  
  # 
  # Returns the value of the `quotas` attribute.
  # 
  def quotas
    return @quotas
  end
  
  # 
  # Sets the value of the `quotas` attribute.
  # 
  def quotas=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Quota.new(value)
        end
      end
    end
    @quotas = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domains` attribute.
  # 
  def storage_domains
    return @storage_domains
  end
  
  # 
  # Sets the value of the `storage_domains` attribute.
  # 
  def storage_domains=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageDomain.new(value)
        end
      end
    end
    @storage_domains = list
  end
  
  # 
  # Returns the value of the `storage_format` attribute.
  # 
  def storage_format
    return @storage_format
  end
  
  # 
  # Sets the value of the `storage_format` attribute.
  # 
  def storage_format=(value)
    @storage_format = value
  end
  
  # 
  # Returns the value of the `supported_versions` attribute.
  # 
  def supported_versions
    return @supported_versions
  end
  
  # 
  # Sets the value of the `supported_versions` attribute.
  # 
  def supported_versions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Version.new(value)
        end
      end
    end
    @supported_versions = list
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {DataCenter} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :clusters The values of attribute `clusters`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Array<Hash>] :iscsi_bonds The values of attribute `iscsi_bonds`.
  # 
  # @option opts :local The value of attribute `local`.
  # 
  # @option opts [Hash] :mac_pool The value of attribute `mac_pool`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :networks The values of attribute `networks`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Array<Hash>] :qoss The values of attribute `qoss`.
  # 
  # @option opts :quota_mode The value of attribute `quota_mode`.
  # 
  # @option opts [Array<Hash>] :quotas The values of attribute `quotas`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Array<Hash>] :storage_domains The values of attribute `storage_domains`.
  # 
  # @option opts :storage_format The value of attribute `storage_format`.
  # 
  # @option opts [Array<Hash>] :supported_versions The values of attribute `supported_versions`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.clusters = opts[:clusters]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.iscsi_bonds = opts[:iscsi_bonds]
    self.local = opts[:local]
    self.mac_pool = opts[:mac_pool]
    self.name = opts[:name]
    self.networks = opts[:networks]
    self.permissions = opts[:permissions]
    self.qoss = opts[:qoss]
    self.quota_mode = opts[:quota_mode]
    self.quotas = opts[:quotas]
    self.status = opts[:status]
    self.storage_domains = opts[:storage_domains]
    self.storage_format = opts[:storage_format]
    self.supported_versions = opts[:supported_versions]
    self.version = opts[:version]
  end
  
end

class Device < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {Device} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.name = opts[:name]
    self.template = opts[:template]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
  end
  
end

class Disk < Device
  
  # 
  # Returns the value of the `active` attribute.
  # 
  def active
    return @active
  end
  
  # 
  # Sets the value of the `active` attribute.
  # 
  def active=(value)
    @active = value
  end
  
  # 
  # Returns the value of the `actual_size` attribute.
  # 
  def actual_size
    return @actual_size
  end
  
  # 
  # Sets the value of the `actual_size` attribute.
  # 
  def actual_size=(value)
    @actual_size = value
  end
  
  # 
  # Returns the value of the `alias_` attribute.
  # 
  def alias_
    return @alias_
  end
  
  # 
  # Sets the value of the `alias_` attribute.
  # 
  def alias_=(value)
    @alias_ = value
  end
  
  # 
  # Returns the value of the `bootable` attribute.
  # 
  def bootable
    return @bootable
  end
  
  # 
  # Sets the value of the `bootable` attribute.
  # 
  def bootable=(value)
    @bootable = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk_profile` attribute.
  # 
  def disk_profile
    return @disk_profile
  end
  
  # 
  # Sets the value of the `disk_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DiskProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk_profile=(value)
    if value.is_a?(Hash)
      value = DiskProfile.new(value)
    end
    @disk_profile = value
  end
  
  # 
  # Returns the value of the `format` attribute.
  # 
  def format
    return @format
  end
  
  # 
  # Sets the value of the `format` attribute.
  # 
  def format=(value)
    @format = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `image_id` attribute.
  # 
  def image_id
    return @image_id
  end
  
  # 
  # Sets the value of the `image_id` attribute.
  # 
  def image_id=(value)
    @image_id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `interface` attribute.
  # 
  def interface
    return @interface
  end
  
  # 
  # Sets the value of the `interface` attribute.
  # 
  def interface=(value)
    @interface = value
  end
  
  # 
  # Returns the value of the `logical_name` attribute.
  # 
  def logical_name
    return @logical_name
  end
  
  # 
  # Sets the value of the `logical_name` attribute.
  # 
  def logical_name=(value)
    @logical_name = value
  end
  
  # 
  # Returns the value of the `lun_storage` attribute.
  # 
  def lun_storage
    return @lun_storage
  end
  
  # 
  # Sets the value of the `lun_storage` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostStorage} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def lun_storage=(value)
    if value.is_a?(Hash)
      value = HostStorage.new(value)
    end
    @lun_storage = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_volume_type` attribute.
  # 
  def openstack_volume_type
    return @openstack_volume_type
  end
  
  # 
  # Sets the value of the `openstack_volume_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackVolumeType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_volume_type=(value)
    if value.is_a?(Hash)
      value = OpenStackVolumeType.new(value)
    end
    @openstack_volume_type = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `propagate_errors` attribute.
  # 
  def propagate_errors
    return @propagate_errors
  end
  
  # 
  # Sets the value of the `propagate_errors` attribute.
  # 
  def propagate_errors=(value)
    @propagate_errors = value
  end
  
  # 
  # Returns the value of the `provisioned_size` attribute.
  # 
  def provisioned_size
    return @provisioned_size
  end
  
  # 
  # Sets the value of the `provisioned_size` attribute.
  # 
  def provisioned_size=(value)
    @provisioned_size = value
  end
  
  # 
  # Returns the value of the `quota` attribute.
  # 
  def quota
    return @quota
  end
  
  # 
  # Sets the value of the `quota` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Quota} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def quota=(value)
    if value.is_a?(Hash)
      value = Quota.new(value)
    end
    @quota = value
  end
  
  # 
  # Returns the value of the `read_only` attribute.
  # 
  def read_only
    return @read_only
  end
  
  # 
  # Sets the value of the `read_only` attribute.
  # 
  def read_only=(value)
    @read_only = value
  end
  
  # 
  # Returns the value of the `sgio` attribute.
  # 
  def sgio
    return @sgio
  end
  
  # 
  # Sets the value of the `sgio` attribute.
  # 
  def sgio=(value)
    @sgio = value
  end
  
  # 
  # Returns the value of the `shareable` attribute.
  # 
  def shareable
    return @shareable
  end
  
  # 
  # Sets the value of the `shareable` attribute.
  # 
  def shareable=(value)
    @shareable = value
  end
  
  # 
  # Returns the value of the `snapshot` attribute.
  # 
  def snapshot
    return @snapshot
  end
  
  # 
  # Sets the value of the `snapshot` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Snapshot} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def snapshot=(value)
    if value.is_a?(Hash)
      value = Snapshot.new(value)
    end
    @snapshot = value
  end
  
  # 
  # Returns the value of the `sparse` attribute.
  # 
  def sparse
    return @sparse
  end
  
  # 
  # Sets the value of the `sparse` attribute.
  # 
  def sparse=(value)
    @sparse = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `storage_domains` attribute.
  # 
  def storage_domains
    return @storage_domains
  end
  
  # 
  # Sets the value of the `storage_domains` attribute.
  # 
  def storage_domains=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageDomain.new(value)
        end
      end
    end
    @storage_domains = list
  end
  
  # 
  # Returns the value of the `storage_type` attribute.
  # 
  def storage_type
    return @storage_type
  end
  
  # 
  # Sets the value of the `storage_type` attribute.
  # 
  def storage_type=(value)
    @storage_type = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `uses_scsi_reservation` attribute.
  # 
  def uses_scsi_reservation
    return @uses_scsi_reservation
  end
  
  # 
  # Sets the value of the `uses_scsi_reservation` attribute.
  # 
  def uses_scsi_reservation=(value)
    @uses_scsi_reservation = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Returns the value of the `wipe_after_delete` attribute.
  # 
  def wipe_after_delete
    return @wipe_after_delete
  end
  
  # 
  # Sets the value of the `wipe_after_delete` attribute.
  # 
  def wipe_after_delete=(value)
    @wipe_after_delete = value
  end
  
  # 
  # Creates a new instance of the {Disk} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active The value of attribute `active`.
  # 
  # @option opts :actual_size The value of attribute `actual_size`.
  # 
  # @option opts :alias_ The value of attribute `alias_`.
  # 
  # @option opts :bootable The value of attribute `bootable`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :disk_profile The value of attribute `disk_profile`.
  # 
  # @option opts :format The value of attribute `format`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :image_id The value of attribute `image_id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :interface The value of attribute `interface`.
  # 
  # @option opts :logical_name The value of attribute `logical_name`.
  # 
  # @option opts [Hash] :lun_storage The value of attribute `lun_storage`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_volume_type The value of attribute `openstack_volume_type`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :propagate_errors The value of attribute `propagate_errors`.
  # 
  # @option opts :provisioned_size The value of attribute `provisioned_size`.
  # 
  # @option opts [Hash] :quota The value of attribute `quota`.
  # 
  # @option opts :read_only The value of attribute `read_only`.
  # 
  # @option opts :sgio The value of attribute `sgio`.
  # 
  # @option opts :shareable The value of attribute `shareable`.
  # 
  # @option opts [Hash] :snapshot The value of attribute `snapshot`.
  # 
  # @option opts :sparse The value of attribute `sparse`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :storage_domains The values of attribute `storage_domains`.
  # 
  # @option opts :storage_type The value of attribute `storage_type`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts :uses_scsi_reservation The value of attribute `uses_scsi_reservation`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # @option opts :wipe_after_delete The value of attribute `wipe_after_delete`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active = opts[:active]
    self.actual_size = opts[:actual_size]
    self.alias_ = opts[:alias_]
    self.bootable = opts[:bootable]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.disk_profile = opts[:disk_profile]
    self.format = opts[:format]
    self.id = opts[:id]
    self.image_id = opts[:image_id]
    self.instance_type = opts[:instance_type]
    self.interface = opts[:interface]
    self.logical_name = opts[:logical_name]
    self.lun_storage = opts[:lun_storage]
    self.name = opts[:name]
    self.openstack_volume_type = opts[:openstack_volume_type]
    self.permissions = opts[:permissions]
    self.propagate_errors = opts[:propagate_errors]
    self.provisioned_size = opts[:provisioned_size]
    self.quota = opts[:quota]
    self.read_only = opts[:read_only]
    self.sgio = opts[:sgio]
    self.shareable = opts[:shareable]
    self.snapshot = opts[:snapshot]
    self.sparse = opts[:sparse]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.storage_domain = opts[:storage_domain]
    self.storage_domains = opts[:storage_domains]
    self.storage_type = opts[:storage_type]
    self.template = opts[:template]
    self.uses_scsi_reservation = opts[:uses_scsi_reservation]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
    self.wipe_after_delete = opts[:wipe_after_delete]
  end
  
end

class DiskAttachment < Identified
  
  # 
  # Returns the value of the `active` attribute.
  # 
  def active
    return @active
  end
  
  # 
  # Sets the value of the `active` attribute.
  # 
  def active=(value)
    @active = value
  end
  
  # 
  # Returns the value of the `bootable` attribute.
  # 
  def bootable
    return @bootable
  end
  
  # 
  # Sets the value of the `bootable` attribute.
  # 
  def bootable=(value)
    @bootable = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk` attribute.
  # 
  def disk
    return @disk
  end
  
  # 
  # Sets the value of the `disk` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Disk} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk=(value)
    if value.is_a?(Hash)
      value = Disk.new(value)
    end
    @disk = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `interface` attribute.
  # 
  def interface
    return @interface
  end
  
  # 
  # Sets the value of the `interface` attribute.
  # 
  def interface=(value)
    @interface = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `pass_discard` attribute.
  # 
  def pass_discard
    return @pass_discard
  end
  
  # 
  # Sets the value of the `pass_discard` attribute.
  # 
  def pass_discard=(value)
    @pass_discard = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {DiskAttachment} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active The value of attribute `active`.
  # 
  # @option opts :bootable The value of attribute `bootable`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :disk The value of attribute `disk`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :interface The value of attribute `interface`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :pass_discard The value of attribute `pass_discard`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active = opts[:active]
    self.bootable = opts[:bootable]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.disk = opts[:disk]
    self.id = opts[:id]
    self.interface = opts[:interface]
    self.name = opts[:name]
    self.pass_discard = opts[:pass_discard]
    self.template = opts[:template]
    self.vm = opts[:vm]
  end
  
end

class DiskProfile < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `qos` attribute.
  # 
  def qos
    return @qos
  end
  
  # 
  # Sets the value of the `qos` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Qos} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def qos=(value)
    if value.is_a?(Hash)
      value = Qos.new(value)
    end
    @qos = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Creates a new instance of the {DiskProfile} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Hash] :qos The value of attribute `qos`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.permissions = opts[:permissions]
    self.qos = opts[:qos]
    self.storage_domain = opts[:storage_domain]
  end
  
end

class DiskSnapshot < Disk
  
  # 
  # Returns the value of the `active` attribute.
  # 
  def active
    return @active
  end
  
  # 
  # Sets the value of the `active` attribute.
  # 
  def active=(value)
    @active = value
  end
  
  # 
  # Returns the value of the `actual_size` attribute.
  # 
  def actual_size
    return @actual_size
  end
  
  # 
  # Sets the value of the `actual_size` attribute.
  # 
  def actual_size=(value)
    @actual_size = value
  end
  
  # 
  # Returns the value of the `alias_` attribute.
  # 
  def alias_
    return @alias_
  end
  
  # 
  # Sets the value of the `alias_` attribute.
  # 
  def alias_=(value)
    @alias_ = value
  end
  
  # 
  # Returns the value of the `bootable` attribute.
  # 
  def bootable
    return @bootable
  end
  
  # 
  # Sets the value of the `bootable` attribute.
  # 
  def bootable=(value)
    @bootable = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk` attribute.
  # 
  def disk
    return @disk
  end
  
  # 
  # Sets the value of the `disk` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Disk} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk=(value)
    if value.is_a?(Hash)
      value = Disk.new(value)
    end
    @disk = value
  end
  
  # 
  # Returns the value of the `disk_profile` attribute.
  # 
  def disk_profile
    return @disk_profile
  end
  
  # 
  # Sets the value of the `disk_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DiskProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def disk_profile=(value)
    if value.is_a?(Hash)
      value = DiskProfile.new(value)
    end
    @disk_profile = value
  end
  
  # 
  # Returns the value of the `format` attribute.
  # 
  def format
    return @format
  end
  
  # 
  # Sets the value of the `format` attribute.
  # 
  def format=(value)
    @format = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `image_id` attribute.
  # 
  def image_id
    return @image_id
  end
  
  # 
  # Sets the value of the `image_id` attribute.
  # 
  def image_id=(value)
    @image_id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `interface` attribute.
  # 
  def interface
    return @interface
  end
  
  # 
  # Sets the value of the `interface` attribute.
  # 
  def interface=(value)
    @interface = value
  end
  
  # 
  # Returns the value of the `logical_name` attribute.
  # 
  def logical_name
    return @logical_name
  end
  
  # 
  # Sets the value of the `logical_name` attribute.
  # 
  def logical_name=(value)
    @logical_name = value
  end
  
  # 
  # Returns the value of the `lun_storage` attribute.
  # 
  def lun_storage
    return @lun_storage
  end
  
  # 
  # Sets the value of the `lun_storage` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostStorage} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def lun_storage=(value)
    if value.is_a?(Hash)
      value = HostStorage.new(value)
    end
    @lun_storage = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `openstack_volume_type` attribute.
  # 
  def openstack_volume_type
    return @openstack_volume_type
  end
  
  # 
  # Sets the value of the `openstack_volume_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OpenStackVolumeType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def openstack_volume_type=(value)
    if value.is_a?(Hash)
      value = OpenStackVolumeType.new(value)
    end
    @openstack_volume_type = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `propagate_errors` attribute.
  # 
  def propagate_errors
    return @propagate_errors
  end
  
  # 
  # Sets the value of the `propagate_errors` attribute.
  # 
  def propagate_errors=(value)
    @propagate_errors = value
  end
  
  # 
  # Returns the value of the `provisioned_size` attribute.
  # 
  def provisioned_size
    return @provisioned_size
  end
  
  # 
  # Sets the value of the `provisioned_size` attribute.
  # 
  def provisioned_size=(value)
    @provisioned_size = value
  end
  
  # 
  # Returns the value of the `quota` attribute.
  # 
  def quota
    return @quota
  end
  
  # 
  # Sets the value of the `quota` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Quota} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def quota=(value)
    if value.is_a?(Hash)
      value = Quota.new(value)
    end
    @quota = value
  end
  
  # 
  # Returns the value of the `read_only` attribute.
  # 
  def read_only
    return @read_only
  end
  
  # 
  # Sets the value of the `read_only` attribute.
  # 
  def read_only=(value)
    @read_only = value
  end
  
  # 
  # Returns the value of the `sgio` attribute.
  # 
  def sgio
    return @sgio
  end
  
  # 
  # Sets the value of the `sgio` attribute.
  # 
  def sgio=(value)
    @sgio = value
  end
  
  # 
  # Returns the value of the `shareable` attribute.
  # 
  def shareable
    return @shareable
  end
  
  # 
  # Sets the value of the `shareable` attribute.
  # 
  def shareable=(value)
    @shareable = value
  end
  
  # 
  # Returns the value of the `snapshot` attribute.
  # 
  def snapshot
    return @snapshot
  end
  
  # 
  # Sets the value of the `snapshot` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Snapshot} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def snapshot=(value)
    if value.is_a?(Hash)
      value = Snapshot.new(value)
    end
    @snapshot = value
  end
  
  # 
  # Returns the value of the `sparse` attribute.
  # 
  def sparse
    return @sparse
  end
  
  # 
  # Sets the value of the `sparse` attribute.
  # 
  def sparse=(value)
    @sparse = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `storage_domains` attribute.
  # 
  def storage_domains
    return @storage_domains
  end
  
  # 
  # Sets the value of the `storage_domains` attribute.
  # 
  def storage_domains=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageDomain.new(value)
        end
      end
    end
    @storage_domains = list
  end
  
  # 
  # Returns the value of the `storage_type` attribute.
  # 
  def storage_type
    return @storage_type
  end
  
  # 
  # Sets the value of the `storage_type` attribute.
  # 
  def storage_type=(value)
    @storage_type = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `uses_scsi_reservation` attribute.
  # 
  def uses_scsi_reservation
    return @uses_scsi_reservation
  end
  
  # 
  # Sets the value of the `uses_scsi_reservation` attribute.
  # 
  def uses_scsi_reservation=(value)
    @uses_scsi_reservation = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Returns the value of the `wipe_after_delete` attribute.
  # 
  def wipe_after_delete
    return @wipe_after_delete
  end
  
  # 
  # Sets the value of the `wipe_after_delete` attribute.
  # 
  def wipe_after_delete=(value)
    @wipe_after_delete = value
  end
  
  # 
  # Creates a new instance of the {DiskSnapshot} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :active The value of attribute `active`.
  # 
  # @option opts :actual_size The value of attribute `actual_size`.
  # 
  # @option opts :alias_ The value of attribute `alias_`.
  # 
  # @option opts :bootable The value of attribute `bootable`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :disk The value of attribute `disk`.
  # 
  # @option opts [Hash] :disk_profile The value of attribute `disk_profile`.
  # 
  # @option opts :format The value of attribute `format`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :image_id The value of attribute `image_id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :interface The value of attribute `interface`.
  # 
  # @option opts :logical_name The value of attribute `logical_name`.
  # 
  # @option opts [Hash] :lun_storage The value of attribute `lun_storage`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :openstack_volume_type The value of attribute `openstack_volume_type`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :propagate_errors The value of attribute `propagate_errors`.
  # 
  # @option opts :provisioned_size The value of attribute `provisioned_size`.
  # 
  # @option opts [Hash] :quota The value of attribute `quota`.
  # 
  # @option opts :read_only The value of attribute `read_only`.
  # 
  # @option opts :sgio The value of attribute `sgio`.
  # 
  # @option opts :shareable The value of attribute `shareable`.
  # 
  # @option opts [Hash] :snapshot The value of attribute `snapshot`.
  # 
  # @option opts :sparse The value of attribute `sparse`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :storage_domains The values of attribute `storage_domains`.
  # 
  # @option opts :storage_type The value of attribute `storage_type`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts :uses_scsi_reservation The value of attribute `uses_scsi_reservation`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # @option opts :wipe_after_delete The value of attribute `wipe_after_delete`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.active = opts[:active]
    self.actual_size = opts[:actual_size]
    self.alias_ = opts[:alias_]
    self.bootable = opts[:bootable]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.disk = opts[:disk]
    self.disk_profile = opts[:disk_profile]
    self.format = opts[:format]
    self.id = opts[:id]
    self.image_id = opts[:image_id]
    self.instance_type = opts[:instance_type]
    self.interface = opts[:interface]
    self.logical_name = opts[:logical_name]
    self.lun_storage = opts[:lun_storage]
    self.name = opts[:name]
    self.openstack_volume_type = opts[:openstack_volume_type]
    self.permissions = opts[:permissions]
    self.propagate_errors = opts[:propagate_errors]
    self.provisioned_size = opts[:provisioned_size]
    self.quota = opts[:quota]
    self.read_only = opts[:read_only]
    self.sgio = opts[:sgio]
    self.shareable = opts[:shareable]
    self.snapshot = opts[:snapshot]
    self.sparse = opts[:sparse]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.storage_domain = opts[:storage_domain]
    self.storage_domains = opts[:storage_domains]
    self.storage_type = opts[:storage_type]
    self.template = opts[:template]
    self.uses_scsi_reservation = opts[:uses_scsi_reservation]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
    self.wipe_after_delete = opts[:wipe_after_delete]
  end
  
end

class Domain < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `groups` attribute.
  # 
  def groups
    return @groups
  end
  
  # 
  # Sets the value of the `groups` attribute.
  # 
  def groups=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Group.new(value)
        end
      end
    end
    @groups = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Returns the value of the `users` attribute.
  # 
  def users
    return @users
  end
  
  # 
  # Sets the value of the `users` attribute.
  # 
  def users=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = User.new(value)
        end
      end
    end
    @users = list
  end
  
  # 
  # Creates a new instance of the {Domain} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :groups The values of attribute `groups`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # @option opts [Array<Hash>] :users The values of attribute `users`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.groups = opts[:groups]
    self.id = opts[:id]
    self.name = opts[:name]
    self.user = opts[:user]
    self.users = opts[:users]
  end
  
end

class Event < Identified
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `code` attribute.
  # 
  def code
    return @code
  end
  
  # 
  # Sets the value of the `code` attribute.
  # 
  def code=(value)
    @code = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `correlation_id` attribute.
  # 
  def correlation_id
    return @correlation_id
  end
  
  # 
  # Sets the value of the `correlation_id` attribute.
  # 
  def correlation_id=(value)
    @correlation_id = value
  end
  
  # 
  # Returns the value of the `custom_data` attribute.
  # 
  def custom_data
    return @custom_data
  end
  
  # 
  # Sets the value of the `custom_data` attribute.
  # 
  def custom_data=(value)
    @custom_data = value
  end
  
  # 
  # Returns the value of the `custom_id` attribute.
  # 
  def custom_id
    return @custom_id
  end
  
  # 
  # Sets the value of the `custom_id` attribute.
  # 
  def custom_id=(value)
    @custom_id = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `flood_rate` attribute.
  # 
  def flood_rate
    return @flood_rate
  end
  
  # 
  # Sets the value of the `flood_rate` attribute.
  # 
  def flood_rate=(value)
    @flood_rate = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `origin` attribute.
  # 
  def origin
    return @origin
  end
  
  # 
  # Sets the value of the `origin` attribute.
  # 
  def origin=(value)
    @origin = value
  end
  
  # 
  # Returns the value of the `severity` attribute.
  # 
  def severity
    return @severity
  end
  
  # 
  # Sets the value of the `severity` attribute.
  # 
  def severity=(value)
    @severity = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `time` attribute.
  # 
  def time
    return @time
  end
  
  # 
  # Sets the value of the `time` attribute.
  # 
  def time=(value)
    @time = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::User} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def user=(value)
    if value.is_a?(Hash)
      value = User.new(value)
    end
    @user = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {Event} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :code The value of attribute `code`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :correlation_id The value of attribute `correlation_id`.
  # 
  # @option opts :custom_data The value of attribute `custom_data`.
  # 
  # @option opts :custom_id The value of attribute `custom_id`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :flood_rate The value of attribute `flood_rate`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :origin The value of attribute `origin`.
  # 
  # @option opts :severity The value of attribute `severity`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts :time The value of attribute `time`.
  # 
  # @option opts [Hash] :user The value of attribute `user`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.cluster = opts[:cluster]
    self.code = opts[:code]
    self.comment = opts[:comment]
    self.correlation_id = opts[:correlation_id]
    self.custom_data = opts[:custom_data]
    self.custom_id = opts[:custom_id]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.flood_rate = opts[:flood_rate]
    self.host = opts[:host]
    self.id = opts[:id]
    self.name = opts[:name]
    self.origin = opts[:origin]
    self.severity = opts[:severity]
    self.storage_domain = opts[:storage_domain]
    self.template = opts[:template]
    self.time = opts[:time]
    self.user = opts[:user]
    self.vm = opts[:vm]
  end
  
end

class ExternalComputeResource < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `provider` attribute.
  # 
  def provider
    return @provider
  end
  
  # 
  # Sets the value of the `provider` attribute.
  # 
  def provider=(value)
    @provider = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `user` attribute.
  # 
  def user
    return @user
  end
  
  # 
  # Sets the value of the `user` attribute.
  # 
  def user=(value)
    @user = value
  end
  
  # 
  # Creates a new instance of the {ExternalComputeResource} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :provider The value of attribute `provider`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :user The value of attribute `user`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.external_host_provider = opts[:external_host_provider]
    self.id = opts[:id]
    self.name = opts[:name]
    self.provider = opts[:provider]
    self.url = opts[:url]
    self.user = opts[:user]
  end
  
end

class ExternalDiscoveredHost < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ip` attribute.
  # 
  def ip
    return @ip
  end
  
  # 
  # Sets the value of the `ip` attribute.
  # 
  def ip=(value)
    @ip = value
  end
  
  # 
  # Returns the value of the `last_report` attribute.
  # 
  def last_report
    return @last_report
  end
  
  # 
  # Sets the value of the `last_report` attribute.
  # 
  def last_report=(value)
    @last_report = value
  end
  
  # 
  # Returns the value of the `mac` attribute.
  # 
  def mac
    return @mac
  end
  
  # 
  # Sets the value of the `mac` attribute.
  # 
  def mac=(value)
    @mac = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `subnet_name` attribute.
  # 
  def subnet_name
    return @subnet_name
  end
  
  # 
  # Sets the value of the `subnet_name` attribute.
  # 
  def subnet_name=(value)
    @subnet_name = value
  end
  
  # 
  # Creates a new instance of the {ExternalDiscoveredHost} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :ip The value of attribute `ip`.
  # 
  # @option opts :last_report The value of attribute `last_report`.
  # 
  # @option opts :mac The value of attribute `mac`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :subnet_name The value of attribute `subnet_name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.external_host_provider = opts[:external_host_provider]
    self.id = opts[:id]
    self.ip = opts[:ip]
    self.last_report = opts[:last_report]
    self.mac = opts[:mac]
    self.name = opts[:name]
    self.subnet_name = opts[:subnet_name]
  end
  
end

class ExternalHost < Identified
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {ExternalHost} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.external_host_provider = opts[:external_host_provider]
    self.id = opts[:id]
    self.name = opts[:name]
  end
  
end

class ExternalHostGroup < Identified
  
  # 
  # Returns the value of the `architecture_name` attribute.
  # 
  def architecture_name
    return @architecture_name
  end
  
  # 
  # Sets the value of the `architecture_name` attribute.
  # 
  def architecture_name=(value)
    @architecture_name = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `domain_name` attribute.
  # 
  def domain_name
    return @domain_name
  end
  
  # 
  # Sets the value of the `domain_name` attribute.
  # 
  def domain_name=(value)
    @domain_name = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `operating_system_name` attribute.
  # 
  def operating_system_name
    return @operating_system_name
  end
  
  # 
  # Sets the value of the `operating_system_name` attribute.
  # 
  def operating_system_name=(value)
    @operating_system_name = value
  end
  
  # 
  # Returns the value of the `subnet_name` attribute.
  # 
  def subnet_name
    return @subnet_name
  end
  
  # 
  # Sets the value of the `subnet_name` attribute.
  # 
  def subnet_name=(value)
    @subnet_name = value
  end
  
  # 
  # Creates a new instance of the {ExternalHostGroup} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :architecture_name The value of attribute `architecture_name`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :domain_name The value of attribute `domain_name`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :operating_system_name The value of attribute `operating_system_name`.
  # 
  # @option opts :subnet_name The value of attribute `subnet_name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.architecture_name = opts[:architecture_name]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.domain_name = opts[:domain_name]
    self.external_host_provider = opts[:external_host_provider]
    self.id = opts[:id]
    self.name = opts[:name]
    self.operating_system_name = opts[:operating_system_name]
    self.subnet_name = opts[:subnet_name]
  end
  
end

class ExternalProvider < Identified
  
  # 
  # Returns the value of the `authentication_url` attribute.
  # 
  def authentication_url
    return @authentication_url
  end
  
  # 
  # Sets the value of the `authentication_url` attribute.
  # 
  def authentication_url=(value)
    @authentication_url = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `requires_authentication` attribute.
  # 
  def requires_authentication
    return @requires_authentication
  end
  
  # 
  # Sets the value of the `requires_authentication` attribute.
  # 
  def requires_authentication=(value)
    @requires_authentication = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {ExternalProvider} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :authentication_url The value of attribute `authentication_url`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :requires_authentication The value of attribute `requires_authentication`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authentication_url = opts[:authentication_url]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.password = opts[:password]
    self.properties = opts[:properties]
    self.requires_authentication = opts[:requires_authentication]
    self.url = opts[:url]
    self.username = opts[:username]
  end
  
end

class File < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `content` attribute.
  # 
  def content
    return @content
  end
  
  # 
  # Sets the value of the `content` attribute.
  # 
  def content=(value)
    @content = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {File} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :content The value of attribute `content`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.content = opts[:content]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.storage_domain = opts[:storage_domain]
    self.type = opts[:type]
  end
  
end

class Filter < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `position` attribute.
  # 
  def position
    return @position
  end
  
  # 
  # Sets the value of the `position` attribute.
  # 
  def position=(value)
    @position = value
  end
  
  # 
  # Returns the value of the `scheduling_policy_unit` attribute.
  # 
  def scheduling_policy_unit
    return @scheduling_policy_unit
  end
  
  # 
  # Sets the value of the `scheduling_policy_unit` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SchedulingPolicyUnit} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def scheduling_policy_unit=(value)
    if value.is_a?(Hash)
      value = SchedulingPolicyUnit.new(value)
    end
    @scheduling_policy_unit = value
  end
  
  # 
  # Creates a new instance of the {Filter} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :position The value of attribute `position`.
  # 
  # @option opts [Hash] :scheduling_policy_unit The value of attribute `scheduling_policy_unit`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.position = opts[:position]
    self.scheduling_policy_unit = opts[:scheduling_policy_unit]
  end
  
end

class Floppy < Device
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `file` attribute.
  # 
  def file
    return @file
  end
  
  # 
  # Sets the value of the `file` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::File} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def file=(value)
    if value.is_a?(Hash)
      value = File.new(value)
    end
    @file = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {Floppy} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :file The value of attribute `file`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.file = opts[:file]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.name = opts[:name]
    self.template = opts[:template]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
  end
  
end

class GlusterBrickAdvancedDetails < Device
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `device` attribute.
  # 
  def device
    return @device
  end
  
  # 
  # Sets the value of the `device` attribute.
  # 
  def device=(value)
    @device = value
  end
  
  # 
  # Returns the value of the `fs_name` attribute.
  # 
  def fs_name
    return @fs_name
  end
  
  # 
  # Sets the value of the `fs_name` attribute.
  # 
  def fs_name=(value)
    @fs_name = value
  end
  
  # 
  # Returns the value of the `gluster_clients` attribute.
  # 
  def gluster_clients
    return @gluster_clients
  end
  
  # 
  # Sets the value of the `gluster_clients` attribute.
  # 
  def gluster_clients=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterClient.new(value)
        end
      end
    end
    @gluster_clients = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `memory_pools` attribute.
  # 
  def memory_pools
    return @memory_pools
  end
  
  # 
  # Sets the value of the `memory_pools` attribute.
  # 
  def memory_pools=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterMemoryPool.new(value)
        end
      end
    end
    @memory_pools = list
  end
  
  # 
  # Returns the value of the `mnt_options` attribute.
  # 
  def mnt_options
    return @mnt_options
  end
  
  # 
  # Sets the value of the `mnt_options` attribute.
  # 
  def mnt_options=(value)
    @mnt_options = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `pid` attribute.
  # 
  def pid
    return @pid
  end
  
  # 
  # Sets the value of the `pid` attribute.
  # 
  def pid=(value)
    @pid = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {GlusterBrickAdvancedDetails} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :device The value of attribute `device`.
  # 
  # @option opts :fs_name The value of attribute `fs_name`.
  # 
  # @option opts [Array<Hash>] :gluster_clients The values of attribute `gluster_clients`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts [Array<Hash>] :memory_pools The values of attribute `memory_pools`.
  # 
  # @option opts :mnt_options The value of attribute `mnt_options`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :pid The value of attribute `pid`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.device = opts[:device]
    self.fs_name = opts[:fs_name]
    self.gluster_clients = opts[:gluster_clients]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.memory_pools = opts[:memory_pools]
    self.mnt_options = opts[:mnt_options]
    self.name = opts[:name]
    self.pid = opts[:pid]
    self.port = opts[:port]
    self.template = opts[:template]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
  end
  
end

class GlusterHook < Identified
  
  # 
  # Returns the value of the `checksum` attribute.
  # 
  def checksum
    return @checksum
  end
  
  # 
  # Sets the value of the `checksum` attribute.
  # 
  def checksum=(value)
    @checksum = value
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `conflict_status` attribute.
  # 
  def conflict_status
    return @conflict_status
  end
  
  # 
  # Sets the value of the `conflict_status` attribute.
  # 
  def conflict_status=(value)
    @conflict_status = value
  end
  
  # 
  # Returns the value of the `conflicts` attribute.
  # 
  def conflicts
    return @conflicts
  end
  
  # 
  # Sets the value of the `conflicts` attribute.
  # 
  def conflicts=(value)
    @conflicts = value
  end
  
  # 
  # Returns the value of the `content` attribute.
  # 
  def content
    return @content
  end
  
  # 
  # Sets the value of the `content` attribute.
  # 
  def content=(value)
    @content = value
  end
  
  # 
  # Returns the value of the `content_type` attribute.
  # 
  def content_type
    return @content_type
  end
  
  # 
  # Sets the value of the `content_type` attribute.
  # 
  def content_type=(value)
    @content_type = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `gluster_command` attribute.
  # 
  def gluster_command
    return @gluster_command
  end
  
  # 
  # Sets the value of the `gluster_command` attribute.
  # 
  def gluster_command=(value)
    @gluster_command = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `server_hooks` attribute.
  # 
  def server_hooks
    return @server_hooks
  end
  
  # 
  # Sets the value of the `server_hooks` attribute.
  # 
  def server_hooks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterServerHook.new(value)
        end
      end
    end
    @server_hooks = list
  end
  
  # 
  # Returns the value of the `stage` attribute.
  # 
  def stage
    return @stage
  end
  
  # 
  # Sets the value of the `stage` attribute.
  # 
  def stage=(value)
    @stage = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Creates a new instance of the {GlusterHook} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :checksum The value of attribute `checksum`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :conflict_status The value of attribute `conflict_status`.
  # 
  # @option opts :conflicts The value of attribute `conflicts`.
  # 
  # @option opts :content The value of attribute `content`.
  # 
  # @option opts :content_type The value of attribute `content_type`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :gluster_command The value of attribute `gluster_command`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :server_hooks The values of attribute `server_hooks`.
  # 
  # @option opts :stage The value of attribute `stage`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.checksum = opts[:checksum]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.conflict_status = opts[:conflict_status]
    self.conflicts = opts[:conflicts]
    self.content = opts[:content]
    self.content_type = opts[:content_type]
    self.description = opts[:description]
    self.gluster_command = opts[:gluster_command]
    self.id = opts[:id]
    self.name = opts[:name]
    self.server_hooks = opts[:server_hooks]
    self.stage = opts[:stage]
    self.status = opts[:status]
  end
  
end

class GlusterMemoryPool < Identified
  
  # 
  # Returns the value of the `alloc_count` attribute.
  # 
  def alloc_count
    return @alloc_count
  end
  
  # 
  # Sets the value of the `alloc_count` attribute.
  # 
  def alloc_count=(value)
    @alloc_count = value
  end
  
  # 
  # Returns the value of the `cold_count` attribute.
  # 
  def cold_count
    return @cold_count
  end
  
  # 
  # Sets the value of the `cold_count` attribute.
  # 
  def cold_count=(value)
    @cold_count = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `hot_count` attribute.
  # 
  def hot_count
    return @hot_count
  end
  
  # 
  # Sets the value of the `hot_count` attribute.
  # 
  def hot_count=(value)
    @hot_count = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `max_alloc` attribute.
  # 
  def max_alloc
    return @max_alloc
  end
  
  # 
  # Sets the value of the `max_alloc` attribute.
  # 
  def max_alloc=(value)
    @max_alloc = value
  end
  
  # 
  # Returns the value of the `max_stdalloc` attribute.
  # 
  def max_stdalloc
    return @max_stdalloc
  end
  
  # 
  # Sets the value of the `max_stdalloc` attribute.
  # 
  def max_stdalloc=(value)
    @max_stdalloc = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `padded_size` attribute.
  # 
  def padded_size
    return @padded_size
  end
  
  # 
  # Sets the value of the `padded_size` attribute.
  # 
  def padded_size=(value)
    @padded_size = value
  end
  
  # 
  # Returns the value of the `pool_misses` attribute.
  # 
  def pool_misses
    return @pool_misses
  end
  
  # 
  # Sets the value of the `pool_misses` attribute.
  # 
  def pool_misses=(value)
    @pool_misses = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Creates a new instance of the {GlusterMemoryPool} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :alloc_count The value of attribute `alloc_count`.
  # 
  # @option opts :cold_count The value of attribute `cold_count`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :hot_count The value of attribute `hot_count`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :max_alloc The value of attribute `max_alloc`.
  # 
  # @option opts :max_stdalloc The value of attribute `max_stdalloc`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :padded_size The value of attribute `padded_size`.
  # 
  # @option opts :pool_misses The value of attribute `pool_misses`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.alloc_count = opts[:alloc_count]
    self.cold_count = opts[:cold_count]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.hot_count = opts[:hot_count]
    self.id = opts[:id]
    self.max_alloc = opts[:max_alloc]
    self.max_stdalloc = opts[:max_stdalloc]
    self.name = opts[:name]
    self.padded_size = opts[:padded_size]
    self.pool_misses = opts[:pool_misses]
    self.type = opts[:type]
  end
  
end

class GlusterServerHook < Identified
  
  # 
  # Returns the value of the `checksum` attribute.
  # 
  def checksum
    return @checksum
  end
  
  # 
  # Sets the value of the `checksum` attribute.
  # 
  def checksum=(value)
    @checksum = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `content_type` attribute.
  # 
  def content_type
    return @content_type
  end
  
  # 
  # Sets the value of the `content_type` attribute.
  # 
  def content_type=(value)
    @content_type = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Creates a new instance of the {GlusterServerHook} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :checksum The value of attribute `checksum`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :content_type The value of attribute `content_type`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.checksum = opts[:checksum]
    self.comment = opts[:comment]
    self.content_type = opts[:content_type]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.name = opts[:name]
    self.status = opts[:status]
  end
  
end

class GlusterVolume < Identified
  
  # 
  # Returns the value of the `bricks` attribute.
  # 
  def bricks
    return @bricks
  end
  
  # 
  # Sets the value of the `bricks` attribute.
  # 
  def bricks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterBrick.new(value)
        end
      end
    end
    @bricks = list
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disperse_count` attribute.
  # 
  def disperse_count
    return @disperse_count
  end
  
  # 
  # Sets the value of the `disperse_count` attribute.
  # 
  def disperse_count=(value)
    @disperse_count = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `options` attribute.
  # 
  def options
    return @options
  end
  
  # 
  # Sets the value of the `options` attribute.
  # 
  def options=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Option.new(value)
        end
      end
    end
    @options = list
  end
  
  # 
  # Returns the value of the `redundancy_count` attribute.
  # 
  def redundancy_count
    return @redundancy_count
  end
  
  # 
  # Sets the value of the `redundancy_count` attribute.
  # 
  def redundancy_count=(value)
    @redundancy_count = value
  end
  
  # 
  # Returns the value of the `replica_count` attribute.
  # 
  def replica_count
    return @replica_count
  end
  
  # 
  # Sets the value of the `replica_count` attribute.
  # 
  def replica_count=(value)
    @replica_count = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `stripe_count` attribute.
  # 
  def stripe_count
    return @stripe_count
  end
  
  # 
  # Sets the value of the `stripe_count` attribute.
  # 
  def stripe_count=(value)
    @stripe_count = value
  end
  
  # 
  # Returns the value of the `transport_types` attribute.
  # 
  def transport_types
    return @transport_types
  end
  
  # 
  # Sets the value of the `transport_types` attribute.
  # 
  def transport_types=(list)
    @transport_types = list
  end
  
  # 
  # Returns the value of the `volume_type` attribute.
  # 
  def volume_type
    return @volume_type
  end
  
  # 
  # Sets the value of the `volume_type` attribute.
  # 
  def volume_type=(value)
    @volume_type = value
  end
  
  # 
  # Creates a new instance of the {GlusterVolume} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :bricks The values of attribute `bricks`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :disperse_count The value of attribute `disperse_count`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :options The values of attribute `options`.
  # 
  # @option opts :redundancy_count The value of attribute `redundancy_count`.
  # 
  # @option opts :replica_count The value of attribute `replica_count`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :stripe_count The value of attribute `stripe_count`.
  # 
  # @option opts [Array<Hash>] :transport_types The values of attribute `transport_types`.
  # 
  # @option opts :volume_type The value of attribute `volume_type`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.bricks = opts[:bricks]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.disperse_count = opts[:disperse_count]
    self.id = opts[:id]
    self.name = opts[:name]
    self.options = opts[:options]
    self.redundancy_count = opts[:redundancy_count]
    self.replica_count = opts[:replica_count]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.stripe_count = opts[:stripe_count]
    self.transport_types = opts[:transport_types]
    self.volume_type = opts[:volume_type]
  end
  
end

class GlusterVolumeProfileDetails < Identified
  
  # 
  # Returns the value of the `brick_profile_details` attribute.
  # 
  def brick_profile_details
    return @brick_profile_details
  end
  
  # 
  # Sets the value of the `brick_profile_details` attribute.
  # 
  def brick_profile_details=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = BrickProfileDetail.new(value)
        end
      end
    end
    @brick_profile_details = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `nfs_profile_details` attribute.
  # 
  def nfs_profile_details
    return @nfs_profile_details
  end
  
  # 
  # Sets the value of the `nfs_profile_details` attribute.
  # 
  def nfs_profile_details=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NfsProfileDetail.new(value)
        end
      end
    end
    @nfs_profile_details = list
  end
  
  # 
  # Creates a new instance of the {GlusterVolumeProfileDetails} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :brick_profile_details The values of attribute `brick_profile_details`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :nfs_profile_details The values of attribute `nfs_profile_details`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.brick_profile_details = opts[:brick_profile_details]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.nfs_profile_details = opts[:nfs_profile_details]
  end
  
end

class GraphicsConsole < Identified
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `protocol` attribute.
  # 
  def protocol
    return @protocol
  end
  
  # 
  # Sets the value of the `protocol` attribute.
  # 
  def protocol=(value)
    @protocol = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `tls_port` attribute.
  # 
  def tls_port
    return @tls_port
  end
  
  # 
  # Sets the value of the `tls_port` attribute.
  # 
  def tls_port=(value)
    @tls_port = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {GraphicsConsole} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :protocol The value of attribute `protocol`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts :tls_port The value of attribute `tls_port`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.name = opts[:name]
    self.port = opts[:port]
    self.protocol = opts[:protocol]
    self.template = opts[:template]
    self.tls_port = opts[:tls_port]
    self.vm = opts[:vm]
  end
  
end

class Group < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `domain_entry_id` attribute.
  # 
  def domain_entry_id
    return @domain_entry_id
  end
  
  # 
  # Sets the value of the `domain_entry_id` attribute.
  # 
  def domain_entry_id=(value)
    @domain_entry_id = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `namespace` attribute.
  # 
  def namespace
    return @namespace
  end
  
  # 
  # Sets the value of the `namespace` attribute.
  # 
  def namespace=(value)
    @namespace = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `roles` attribute.
  # 
  def roles
    return @roles
  end
  
  # 
  # Sets the value of the `roles` attribute.
  # 
  def roles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Role.new(value)
        end
      end
    end
    @roles = list
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Creates a new instance of the {Group} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts :domain_entry_id The value of attribute `domain_entry_id`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :namespace The value of attribute `namespace`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Array<Hash>] :roles The values of attribute `roles`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.domain = opts[:domain]
    self.domain_entry_id = opts[:domain_entry_id]
    self.id = opts[:id]
    self.name = opts[:name]
    self.namespace = opts[:namespace]
    self.permissions = opts[:permissions]
    self.roles = opts[:roles]
    self.tags = opts[:tags]
  end
  
end

class Hook < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `event_name` attribute.
  # 
  def event_name
    return @event_name
  end
  
  # 
  # Sets the value of the `event_name` attribute.
  # 
  def event_name=(value)
    @event_name = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `md5` attribute.
  # 
  def md5
    return @md5
  end
  
  # 
  # Sets the value of the `md5` attribute.
  # 
  def md5=(value)
    @md5 = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {Hook} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :event_name The value of attribute `event_name`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :md5 The value of attribute `md5`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.event_name = opts[:event_name]
    self.host = opts[:host]
    self.id = opts[:id]
    self.md5 = opts[:md5]
    self.name = opts[:name]
  end
  
end

class Host < Identified
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `affinity_labels` attribute.
  # 
  def affinity_labels
    return @affinity_labels
  end
  
  # 
  # Sets the value of the `affinity_labels` attribute.
  # 
  def affinity_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = AffinityLabel.new(value)
        end
      end
    end
    @affinity_labels = list
  end
  
  # 
  # Returns the value of the `agents` attribute.
  # 
  def agents
    return @agents
  end
  
  # 
  # Sets the value of the `agents` attribute.
  # 
  def agents=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Agent.new(value)
        end
      end
    end
    @agents = list
  end
  
  # 
  # Returns the value of the `auto_numa_status` attribute.
  # 
  def auto_numa_status
    return @auto_numa_status
  end
  
  # 
  # Sets the value of the `auto_numa_status` attribute.
  # 
  def auto_numa_status=(value)
    @auto_numa_status = value
  end
  
  # 
  # Returns the value of the `certificate` attribute.
  # 
  def certificate
    return @certificate
  end
  
  # 
  # Sets the value of the `certificate` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Certificate} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def certificate=(value)
    if value.is_a?(Hash)
      value = Certificate.new(value)
    end
    @certificate = value
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `device_passthrough` attribute.
  # 
  def device_passthrough
    return @device_passthrough
  end
  
  # 
  # Sets the value of the `device_passthrough` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostDevicePassthrough} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def device_passthrough=(value)
    if value.is_a?(Hash)
      value = HostDevicePassthrough.new(value)
    end
    @device_passthrough = value
  end
  
  # 
  # Returns the value of the `devices` attribute.
  # 
  def devices
    return @devices
  end
  
  # 
  # Sets the value of the `devices` attribute.
  # 
  def devices=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Device.new(value)
        end
      end
    end
    @devices = list
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `external_status` attribute.
  # 
  def external_status
    return @external_status
  end
  
  # 
  # Sets the value of the `external_status` attribute.
  # 
  def external_status=(value)
    @external_status = value
  end
  
  # 
  # Returns the value of the `hardware_information` attribute.
  # 
  def hardware_information
    return @hardware_information
  end
  
  # 
  # Sets the value of the `hardware_information` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HardwareInformation} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def hardware_information=(value)
    if value.is_a?(Hash)
      value = HardwareInformation.new(value)
    end
    @hardware_information = value
  end
  
  # 
  # Returns the value of the `hooks` attribute.
  # 
  def hooks
    return @hooks
  end
  
  # 
  # Sets the value of the `hooks` attribute.
  # 
  def hooks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Hook.new(value)
        end
      end
    end
    @hooks = list
  end
  
  # 
  # Returns the value of the `hosted_engine` attribute.
  # 
  def hosted_engine
    return @hosted_engine
  end
  
  # 
  # Sets the value of the `hosted_engine` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostedEngine} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def hosted_engine=(value)
    if value.is_a?(Hash)
      value = HostedEngine.new(value)
    end
    @hosted_engine = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `iscsi` attribute.
  # 
  def iscsi
    return @iscsi
  end
  
  # 
  # Sets the value of the `iscsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::IscsiDetails} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def iscsi=(value)
    if value.is_a?(Hash)
      value = IscsiDetails.new(value)
    end
    @iscsi = value
  end
  
  # 
  # Returns the value of the `katello_errata` attribute.
  # 
  def katello_errata
    return @katello_errata
  end
  
  # 
  # Sets the value of the `katello_errata` attribute.
  # 
  def katello_errata=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = KatelloErratum.new(value)
        end
      end
    end
    @katello_errata = list
  end
  
  # 
  # Returns the value of the `kdump_status` attribute.
  # 
  def kdump_status
    return @kdump_status
  end
  
  # 
  # Sets the value of the `kdump_status` attribute.
  # 
  def kdump_status=(value)
    @kdump_status = value
  end
  
  # 
  # Returns the value of the `ksm` attribute.
  # 
  def ksm
    return @ksm
  end
  
  # 
  # Sets the value of the `ksm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ksm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ksm=(value)
    if value.is_a?(Hash)
      value = Ksm.new(value)
    end
    @ksm = value
  end
  
  # 
  # Returns the value of the `libvirt_version` attribute.
  # 
  def libvirt_version
    return @libvirt_version
  end
  
  # 
  # Sets the value of the `libvirt_version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def libvirt_version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @libvirt_version = value
  end
  
  # 
  # Returns the value of the `max_scheduling_memory` attribute.
  # 
  def max_scheduling_memory
    return @max_scheduling_memory
  end
  
  # 
  # Sets the value of the `max_scheduling_memory` attribute.
  # 
  def max_scheduling_memory=(value)
    @max_scheduling_memory = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network_attachments` attribute.
  # 
  def network_attachments
    return @network_attachments
  end
  
  # 
  # Sets the value of the `network_attachments` attribute.
  # 
  def network_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkAttachment.new(value)
        end
      end
    end
    @network_attachments = list
  end
  
  # 
  # Returns the value of the `nics` attribute.
  # 
  def nics
    return @nics
  end
  
  # 
  # Sets the value of the `nics` attribute.
  # 
  def nics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Nic.new(value)
        end
      end
    end
    @nics = list
  end
  
  # 
  # Returns the value of the `numa_nodes` attribute.
  # 
  def numa_nodes
    return @numa_nodes
  end
  
  # 
  # Sets the value of the `numa_nodes` attribute.
  # 
  def numa_nodes=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NumaNode.new(value)
        end
      end
    end
    @numa_nodes = list
  end
  
  # 
  # Returns the value of the `numa_supported` attribute.
  # 
  def numa_supported
    return @numa_supported
  end
  
  # 
  # Sets the value of the `numa_supported` attribute.
  # 
  def numa_supported=(value)
    @numa_supported = value
  end
  
  # 
  # Returns the value of the `os` attribute.
  # 
  def os
    return @os
  end
  
  # 
  # Sets the value of the `os` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def os=(value)
    if value.is_a?(Hash)
      value = OperatingSystem.new(value)
    end
    @os = value
  end
  
  # 
  # Returns the value of the `override_iptables` attribute.
  # 
  def override_iptables
    return @override_iptables
  end
  
  # 
  # Sets the value of the `override_iptables` attribute.
  # 
  def override_iptables=(value)
    @override_iptables = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `power_management` attribute.
  # 
  def power_management
    return @power_management
  end
  
  # 
  # Sets the value of the `power_management` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::PowerManagement} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def power_management=(value)
    if value.is_a?(Hash)
      value = PowerManagement.new(value)
    end
    @power_management = value
  end
  
  # 
  # Returns the value of the `protocol` attribute.
  # 
  def protocol
    return @protocol
  end
  
  # 
  # Sets the value of the `protocol` attribute.
  # 
  def protocol=(value)
    @protocol = value
  end
  
  # 
  # Returns the value of the `root_password` attribute.
  # 
  def root_password
    return @root_password
  end
  
  # 
  # Sets the value of the `root_password` attribute.
  # 
  def root_password=(value)
    @root_password = value
  end
  
  # 
  # Returns the value of the `se_linux` attribute.
  # 
  def se_linux
    return @se_linux
  end
  
  # 
  # Sets the value of the `se_linux` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SeLinux} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def se_linux=(value)
    if value.is_a?(Hash)
      value = SeLinux.new(value)
    end
    @se_linux = value
  end
  
  # 
  # Returns the value of the `spm` attribute.
  # 
  def spm
    return @spm
  end
  
  # 
  # Sets the value of the `spm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Spm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def spm=(value)
    if value.is_a?(Hash)
      value = Spm.new(value)
    end
    @spm = value
  end
  
  # 
  # Returns the value of the `ssh` attribute.
  # 
  def ssh
    return @ssh
  end
  
  # 
  # Sets the value of the `ssh` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ssh} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ssh=(value)
    if value.is_a?(Hash)
      value = Ssh.new(value)
    end
    @ssh = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `status_detail` attribute.
  # 
  def status_detail
    return @status_detail
  end
  
  # 
  # Sets the value of the `status_detail` attribute.
  # 
  def status_detail=(value)
    @status_detail = value
  end
  
  # 
  # Returns the value of the `storage_connection_extensions` attribute.
  # 
  def storage_connection_extensions
    return @storage_connection_extensions
  end
  
  # 
  # Sets the value of the `storage_connection_extensions` attribute.
  # 
  def storage_connection_extensions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = StorageConnectionExtension.new(value)
        end
      end
    end
    @storage_connection_extensions = list
  end
  
  # 
  # Returns the value of the `storages` attribute.
  # 
  def storages
    return @storages
  end
  
  # 
  # Sets the value of the `storages` attribute.
  # 
  def storages=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = HostStorage.new(value)
        end
      end
    end
    @storages = list
  end
  
  # 
  # Returns the value of the `summary` attribute.
  # 
  def summary
    return @summary
  end
  
  # 
  # Sets the value of the `summary` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VmSummary} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def summary=(value)
    if value.is_a?(Hash)
      value = VmSummary.new(value)
    end
    @summary = value
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Returns the value of the `transparent_huge_pages` attribute.
  # 
  def transparent_huge_pages
    return @transparent_huge_pages
  end
  
  # 
  # Sets the value of the `transparent_huge_pages` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TransparentHugePages} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def transparent_huge_pages=(value)
    if value.is_a?(Hash)
      value = TransparentHugePages.new(value)
    end
    @transparent_huge_pages = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `unmanaged_networks` attribute.
  # 
  def unmanaged_networks
    return @unmanaged_networks
  end
  
  # 
  # Sets the value of the `unmanaged_networks` attribute.
  # 
  def unmanaged_networks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = UnmanagedNetwork.new(value)
        end
      end
    end
    @unmanaged_networks = list
  end
  
  # 
  # Returns the value of the `update_available` attribute.
  # 
  def update_available
    return @update_available
  end
  
  # 
  # Sets the value of the `update_available` attribute.
  # 
  def update_available=(value)
    @update_available = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @version = value
  end
  
  # 
  # Creates a new instance of the {Host} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts [Array<Hash>] :affinity_labels The values of attribute `affinity_labels`.
  # 
  # @option opts [Array<Hash>] :agents The values of attribute `agents`.
  # 
  # @option opts :auto_numa_status The value of attribute `auto_numa_status`.
  # 
  # @option opts [Hash] :certificate The value of attribute `certificate`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :device_passthrough The value of attribute `device_passthrough`.
  # 
  # @option opts [Array<Hash>] :devices The values of attribute `devices`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts :external_status The value of attribute `external_status`.
  # 
  # @option opts [Hash] :hardware_information The value of attribute `hardware_information`.
  # 
  # @option opts [Array<Hash>] :hooks The values of attribute `hooks`.
  # 
  # @option opts [Hash] :hosted_engine The value of attribute `hosted_engine`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :iscsi The value of attribute `iscsi`.
  # 
  # @option opts [Array<Hash>] :katello_errata The values of attribute `katello_errata`.
  # 
  # @option opts :kdump_status The value of attribute `kdump_status`.
  # 
  # @option opts [Hash] :ksm The value of attribute `ksm`.
  # 
  # @option opts [Hash] :libvirt_version The value of attribute `libvirt_version`.
  # 
  # @option opts :max_scheduling_memory The value of attribute `max_scheduling_memory`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :network_attachments The values of attribute `network_attachments`.
  # 
  # @option opts [Array<Hash>] :nics The values of attribute `nics`.
  # 
  # @option opts [Array<Hash>] :numa_nodes The values of attribute `numa_nodes`.
  # 
  # @option opts :numa_supported The value of attribute `numa_supported`.
  # 
  # @option opts [Hash] :os The value of attribute `os`.
  # 
  # @option opts :override_iptables The value of attribute `override_iptables`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts [Hash] :power_management The value of attribute `power_management`.
  # 
  # @option opts :protocol The value of attribute `protocol`.
  # 
  # @option opts :root_password The value of attribute `root_password`.
  # 
  # @option opts [Hash] :se_linux The value of attribute `se_linux`.
  # 
  # @option opts [Hash] :spm The value of attribute `spm`.
  # 
  # @option opts [Hash] :ssh The value of attribute `ssh`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :status_detail The value of attribute `status_detail`.
  # 
  # @option opts [Array<Hash>] :storage_connection_extensions The values of attribute `storage_connection_extensions`.
  # 
  # @option opts [Array<Hash>] :storages The values of attribute `storages`.
  # 
  # @option opts [Hash] :summary The value of attribute `summary`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # @option opts [Hash] :transparent_huge_pages The value of attribute `transparent_huge_pages`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Array<Hash>] :unmanaged_networks The values of attribute `unmanaged_networks`.
  # 
  # @option opts :update_available The value of attribute `update_available`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.affinity_labels = opts[:affinity_labels]
    self.agents = opts[:agents]
    self.auto_numa_status = opts[:auto_numa_status]
    self.certificate = opts[:certificate]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.cpu = opts[:cpu]
    self.description = opts[:description]
    self.device_passthrough = opts[:device_passthrough]
    self.devices = opts[:devices]
    self.display = opts[:display]
    self.external_host_provider = opts[:external_host_provider]
    self.external_status = opts[:external_status]
    self.hardware_information = opts[:hardware_information]
    self.hooks = opts[:hooks]
    self.hosted_engine = opts[:hosted_engine]
    self.id = opts[:id]
    self.iscsi = opts[:iscsi]
    self.katello_errata = opts[:katello_errata]
    self.kdump_status = opts[:kdump_status]
    self.ksm = opts[:ksm]
    self.libvirt_version = opts[:libvirt_version]
    self.max_scheduling_memory = opts[:max_scheduling_memory]
    self.memory = opts[:memory]
    self.name = opts[:name]
    self.network_attachments = opts[:network_attachments]
    self.nics = opts[:nics]
    self.numa_nodes = opts[:numa_nodes]
    self.numa_supported = opts[:numa_supported]
    self.os = opts[:os]
    self.override_iptables = opts[:override_iptables]
    self.permissions = opts[:permissions]
    self.port = opts[:port]
    self.power_management = opts[:power_management]
    self.protocol = opts[:protocol]
    self.root_password = opts[:root_password]
    self.se_linux = opts[:se_linux]
    self.spm = opts[:spm]
    self.ssh = opts[:ssh]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.status_detail = opts[:status_detail]
    self.storage_connection_extensions = opts[:storage_connection_extensions]
    self.storages = opts[:storages]
    self.summary = opts[:summary]
    self.tags = opts[:tags]
    self.transparent_huge_pages = opts[:transparent_huge_pages]
    self.type = opts[:type]
    self.unmanaged_networks = opts[:unmanaged_networks]
    self.update_available = opts[:update_available]
    self.version = opts[:version]
  end
  
end

class HostDevice < Identified
  
  # 
  # Returns the value of the `capability` attribute.
  # 
  def capability
    return @capability
  end
  
  # 
  # Sets the value of the `capability` attribute.
  # 
  def capability=(value)
    @capability = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `iommu_group` attribute.
  # 
  def iommu_group
    return @iommu_group
  end
  
  # 
  # Sets the value of the `iommu_group` attribute.
  # 
  def iommu_group=(value)
    @iommu_group = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `parent_device` attribute.
  # 
  def parent_device
    return @parent_device
  end
  
  # 
  # Sets the value of the `parent_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def parent_device=(value)
    if value.is_a?(Hash)
      value = HostDevice.new(value)
    end
    @parent_device = value
  end
  
  # 
  # Returns the value of the `physical_function` attribute.
  # 
  def physical_function
    return @physical_function
  end
  
  # 
  # Sets the value of the `physical_function` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def physical_function=(value)
    if value.is_a?(Hash)
      value = HostDevice.new(value)
    end
    @physical_function = value
  end
  
  # 
  # Returns the value of the `placeholder` attribute.
  # 
  def placeholder
    return @placeholder
  end
  
  # 
  # Sets the value of the `placeholder` attribute.
  # 
  def placeholder=(value)
    @placeholder = value
  end
  
  # 
  # Returns the value of the `product` attribute.
  # 
  def product
    return @product
  end
  
  # 
  # Sets the value of the `product` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Product} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def product=(value)
    if value.is_a?(Hash)
      value = Product.new(value)
    end
    @product = value
  end
  
  # 
  # Returns the value of the `vendor` attribute.
  # 
  def vendor
    return @vendor
  end
  
  # 
  # Sets the value of the `vendor` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vendor} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vendor=(value)
    if value.is_a?(Hash)
      value = Vendor.new(value)
    end
    @vendor = value
  end
  
  # 
  # Returns the value of the `virtual_functions` attribute.
  # 
  def virtual_functions
    return @virtual_functions
  end
  
  # 
  # Sets the value of the `virtual_functions` attribute.
  # 
  def virtual_functions=(value)
    @virtual_functions = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Creates a new instance of the {HostDevice} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :capability The value of attribute `capability`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :iommu_group The value of attribute `iommu_group`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :parent_device The value of attribute `parent_device`.
  # 
  # @option opts [Hash] :physical_function The value of attribute `physical_function`.
  # 
  # @option opts :placeholder The value of attribute `placeholder`.
  # 
  # @option opts [Hash] :product The value of attribute `product`.
  # 
  # @option opts [Hash] :vendor The value of attribute `vendor`.
  # 
  # @option opts :virtual_functions The value of attribute `virtual_functions`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.capability = opts[:capability]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.iommu_group = opts[:iommu_group]
    self.name = opts[:name]
    self.parent_device = opts[:parent_device]
    self.physical_function = opts[:physical_function]
    self.placeholder = opts[:placeholder]
    self.product = opts[:product]
    self.vendor = opts[:vendor]
    self.virtual_functions = opts[:virtual_functions]
    self.vm = opts[:vm]
  end
  
end

class HostNic < Identified
  
  # 
  # Returns the value of the `base_interface` attribute.
  # 
  def base_interface
    return @base_interface
  end
  
  # 
  # Sets the value of the `base_interface` attribute.
  # 
  def base_interface=(value)
    @base_interface = value
  end
  
  # 
  # Returns the value of the `bonding` attribute.
  # 
  def bonding
    return @bonding
  end
  
  # 
  # Sets the value of the `bonding` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Bonding} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bonding=(value)
    if value.is_a?(Hash)
      value = Bonding.new(value)
    end
    @bonding = value
  end
  
  # 
  # Returns the value of the `boot_protocol` attribute.
  # 
  def boot_protocol
    return @boot_protocol
  end
  
  # 
  # Sets the value of the `boot_protocol` attribute.
  # 
  def boot_protocol=(value)
    @boot_protocol = value
  end
  
  # 
  # Returns the value of the `bridged` attribute.
  # 
  def bridged
    return @bridged
  end
  
  # 
  # Sets the value of the `bridged` attribute.
  # 
  def bridged=(value)
    @bridged = value
  end
  
  # 
  # Returns the value of the `check_connectivity` attribute.
  # 
  def check_connectivity
    return @check_connectivity
  end
  
  # 
  # Sets the value of the `check_connectivity` attribute.
  # 
  def check_connectivity=(value)
    @check_connectivity = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `custom_configuration` attribute.
  # 
  def custom_configuration
    return @custom_configuration
  end
  
  # 
  # Sets the value of the `custom_configuration` attribute.
  # 
  def custom_configuration=(value)
    @custom_configuration = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `ip` attribute.
  # 
  def ip
    return @ip
  end
  
  # 
  # Sets the value of the `ip` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ip} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ip=(value)
    if value.is_a?(Hash)
      value = Ip.new(value)
    end
    @ip = value
  end
  
  # 
  # Returns the value of the `ipv6` attribute.
  # 
  def ipv6
    return @ipv6
  end
  
  # 
  # Sets the value of the `ipv6` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Ip} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def ipv6=(value)
    if value.is_a?(Hash)
      value = Ip.new(value)
    end
    @ipv6 = value
  end
  
  # 
  # Returns the value of the `ipv6_boot_protocol` attribute.
  # 
  def ipv6_boot_protocol
    return @ipv6_boot_protocol
  end
  
  # 
  # Sets the value of the `ipv6_boot_protocol` attribute.
  # 
  def ipv6_boot_protocol=(value)
    @ipv6_boot_protocol = value
  end
  
  # 
  # Returns the value of the `mac` attribute.
  # 
  def mac
    return @mac
  end
  
  # 
  # Sets the value of the `mac` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Mac} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def mac=(value)
    if value.is_a?(Hash)
      value = Mac.new(value)
    end
    @mac = value
  end
  
  # 
  # Returns the value of the `mtu` attribute.
  # 
  def mtu
    return @mtu
  end
  
  # 
  # Sets the value of the `mtu` attribute.
  # 
  def mtu=(value)
    @mtu = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network` attribute.
  # 
  def network
    return @network
  end
  
  # 
  # Sets the value of the `network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Network} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network=(value)
    if value.is_a?(Hash)
      value = Network.new(value)
    end
    @network = value
  end
  
  # 
  # Returns the value of the `network_labels` attribute.
  # 
  def network_labels
    return @network_labels
  end
  
  # 
  # Sets the value of the `network_labels` attribute.
  # 
  def network_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkLabel.new(value)
        end
      end
    end
    @network_labels = list
  end
  
  # 
  # Returns the value of the `override_configuration` attribute.
  # 
  def override_configuration
    return @override_configuration
  end
  
  # 
  # Sets the value of the `override_configuration` attribute.
  # 
  def override_configuration=(value)
    @override_configuration = value
  end
  
  # 
  # Returns the value of the `physical_function` attribute.
  # 
  def physical_function
    return @physical_function
  end
  
  # 
  # Sets the value of the `physical_function` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNic} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def physical_function=(value)
    if value.is_a?(Hash)
      value = HostNic.new(value)
    end
    @physical_function = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `qos` attribute.
  # 
  def qos
    return @qos
  end
  
  # 
  # Sets the value of the `qos` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Qos} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def qos=(value)
    if value.is_a?(Hash)
      value = Qos.new(value)
    end
    @qos = value
  end
  
  # 
  # Returns the value of the `speed` attribute.
  # 
  def speed
    return @speed
  end
  
  # 
  # Sets the value of the `speed` attribute.
  # 
  def speed=(value)
    @speed = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `virtual_functions_configuration` attribute.
  # 
  def virtual_functions_configuration
    return @virtual_functions_configuration
  end
  
  # 
  # Sets the value of the `virtual_functions_configuration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HostNicVirtualFunctionsConfiguration} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtual_functions_configuration=(value)
    if value.is_a?(Hash)
      value = HostNicVirtualFunctionsConfiguration.new(value)
    end
    @virtual_functions_configuration = value
  end
  
  # 
  # Returns the value of the `vlan` attribute.
  # 
  def vlan
    return @vlan
  end
  
  # 
  # Sets the value of the `vlan` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vlan} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vlan=(value)
    if value.is_a?(Hash)
      value = Vlan.new(value)
    end
    @vlan = value
  end
  
  # 
  # Creates a new instance of the {HostNic} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :base_interface The value of attribute `base_interface`.
  # 
  # @option opts [Hash] :bonding The value of attribute `bonding`.
  # 
  # @option opts :boot_protocol The value of attribute `boot_protocol`.
  # 
  # @option opts :bridged The value of attribute `bridged`.
  # 
  # @option opts :check_connectivity The value of attribute `check_connectivity`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :custom_configuration The value of attribute `custom_configuration`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :ip The value of attribute `ip`.
  # 
  # @option opts [Hash] :ipv6 The value of attribute `ipv6`.
  # 
  # @option opts :ipv6_boot_protocol The value of attribute `ipv6_boot_protocol`.
  # 
  # @option opts [Hash] :mac The value of attribute `mac`.
  # 
  # @option opts :mtu The value of attribute `mtu`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :network The value of attribute `network`.
  # 
  # @option opts [Array<Hash>] :network_labels The values of attribute `network_labels`.
  # 
  # @option opts :override_configuration The value of attribute `override_configuration`.
  # 
  # @option opts [Hash] :physical_function The value of attribute `physical_function`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts [Hash] :qos The value of attribute `qos`.
  # 
  # @option opts :speed The value of attribute `speed`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :virtual_functions_configuration The value of attribute `virtual_functions_configuration`.
  # 
  # @option opts [Hash] :vlan The value of attribute `vlan`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.base_interface = opts[:base_interface]
    self.bonding = opts[:bonding]
    self.boot_protocol = opts[:boot_protocol]
    self.bridged = opts[:bridged]
    self.check_connectivity = opts[:check_connectivity]
    self.comment = opts[:comment]
    self.custom_configuration = opts[:custom_configuration]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.ip = opts[:ip]
    self.ipv6 = opts[:ipv6]
    self.ipv6_boot_protocol = opts[:ipv6_boot_protocol]
    self.mac = opts[:mac]
    self.mtu = opts[:mtu]
    self.name = opts[:name]
    self.network = opts[:network]
    self.network_labels = opts[:network_labels]
    self.override_configuration = opts[:override_configuration]
    self.physical_function = opts[:physical_function]
    self.properties = opts[:properties]
    self.qos = opts[:qos]
    self.speed = opts[:speed]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.virtual_functions_configuration = opts[:virtual_functions_configuration]
    self.vlan = opts[:vlan]
  end
  
end

class HostStorage < Identified
  
  # 
  # Returns the value of the `address` attribute.
  # 
  def address
    return @address
  end
  
  # 
  # Sets the value of the `address` attribute.
  # 
  def address=(value)
    @address = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `logical_units` attribute.
  # 
  def logical_units
    return @logical_units
  end
  
  # 
  # Sets the value of the `logical_units` attribute.
  # 
  def logical_units=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = LogicalUnit.new(value)
        end
      end
    end
    @logical_units = list
  end
  
  # 
  # Returns the value of the `mount_options` attribute.
  # 
  def mount_options
    return @mount_options
  end
  
  # 
  # Sets the value of the `mount_options` attribute.
  # 
  def mount_options=(value)
    @mount_options = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `nfs_retrans` attribute.
  # 
  def nfs_retrans
    return @nfs_retrans
  end
  
  # 
  # Sets the value of the `nfs_retrans` attribute.
  # 
  def nfs_retrans=(value)
    @nfs_retrans = value
  end
  
  # 
  # Returns the value of the `nfs_timeo` attribute.
  # 
  def nfs_timeo
    return @nfs_timeo
  end
  
  # 
  # Sets the value of the `nfs_timeo` attribute.
  # 
  def nfs_timeo=(value)
    @nfs_timeo = value
  end
  
  # 
  # Returns the value of the `nfs_version` attribute.
  # 
  def nfs_version
    return @nfs_version
  end
  
  # 
  # Sets the value of the `nfs_version` attribute.
  # 
  def nfs_version=(value)
    @nfs_version = value
  end
  
  # 
  # Returns the value of the `override_luns` attribute.
  # 
  def override_luns
    return @override_luns
  end
  
  # 
  # Sets the value of the `override_luns` attribute.
  # 
  def override_luns=(value)
    @override_luns = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `path` attribute.
  # 
  def path
    return @path
  end
  
  # 
  # Sets the value of the `path` attribute.
  # 
  def path=(value)
    @path = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `portal` attribute.
  # 
  def portal
    return @portal
  end
  
  # 
  # Sets the value of the `portal` attribute.
  # 
  def portal=(value)
    @portal = value
  end
  
  # 
  # Returns the value of the `target` attribute.
  # 
  def target
    return @target
  end
  
  # 
  # Sets the value of the `target` attribute.
  # 
  def target=(value)
    @target = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Returns the value of the `vfs_type` attribute.
  # 
  def vfs_type
    return @vfs_type
  end
  
  # 
  # Sets the value of the `vfs_type` attribute.
  # 
  def vfs_type=(value)
    @vfs_type = value
  end
  
  # 
  # Returns the value of the `volume_group` attribute.
  # 
  def volume_group
    return @volume_group
  end
  
  # 
  # Sets the value of the `volume_group` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VolumeGroup} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def volume_group=(value)
    if value.is_a?(Hash)
      value = VolumeGroup.new(value)
    end
    @volume_group = value
  end
  
  # 
  # Creates a new instance of the {HostStorage} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :address The value of attribute `address`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Array<Hash>] :logical_units The values of attribute `logical_units`.
  # 
  # @option opts :mount_options The value of attribute `mount_options`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :nfs_retrans The value of attribute `nfs_retrans`.
  # 
  # @option opts :nfs_timeo The value of attribute `nfs_timeo`.
  # 
  # @option opts :nfs_version The value of attribute `nfs_version`.
  # 
  # @option opts :override_luns The value of attribute `override_luns`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :path The value of attribute `path`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :portal The value of attribute `portal`.
  # 
  # @option opts :target The value of attribute `target`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # @option opts :vfs_type The value of attribute `vfs_type`.
  # 
  # @option opts [Hash] :volume_group The value of attribute `volume_group`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.address = opts[:address]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.host = opts[:host]
    self.id = opts[:id]
    self.logical_units = opts[:logical_units]
    self.mount_options = opts[:mount_options]
    self.name = opts[:name]
    self.nfs_retrans = opts[:nfs_retrans]
    self.nfs_timeo = opts[:nfs_timeo]
    self.nfs_version = opts[:nfs_version]
    self.override_luns = opts[:override_luns]
    self.password = opts[:password]
    self.path = opts[:path]
    self.port = opts[:port]
    self.portal = opts[:portal]
    self.target = opts[:target]
    self.type = opts[:type]
    self.username = opts[:username]
    self.vfs_type = opts[:vfs_type]
    self.volume_group = opts[:volume_group]
  end
  
end

class Icon < Identified
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `data` attribute.
  # 
  def data
    return @data
  end
  
  # 
  # Sets the value of the `data` attribute.
  # 
  def data=(value)
    @data = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `media_type` attribute.
  # 
  def media_type
    return @media_type
  end
  
  # 
  # Sets the value of the `media_type` attribute.
  # 
  def media_type=(value)
    @media_type = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Creates a new instance of the {Icon} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :data The value of attribute `data`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :media_type The value of attribute `media_type`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.data = opts[:data]
    self.description = opts[:description]
    self.id = opts[:id]
    self.media_type = opts[:media_type]
    self.name = opts[:name]
  end
  
end

class Nic < Device
  
  # 
  # Returns the value of the `boot_protocol` attribute.
  # 
  def boot_protocol
    return @boot_protocol
  end
  
  # 
  # Sets the value of the `boot_protocol` attribute.
  # 
  def boot_protocol=(value)
    @boot_protocol = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `interface` attribute.
  # 
  def interface
    return @interface
  end
  
  # 
  # Sets the value of the `interface` attribute.
  # 
  def interface=(value)
    @interface = value
  end
  
  # 
  # Returns the value of the `linked` attribute.
  # 
  def linked
    return @linked
  end
  
  # 
  # Sets the value of the `linked` attribute.
  # 
  def linked=(value)
    @linked = value
  end
  
  # 
  # Returns the value of the `mac` attribute.
  # 
  def mac
    return @mac
  end
  
  # 
  # Sets the value of the `mac` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Mac} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def mac=(value)
    if value.is_a?(Hash)
      value = Mac.new(value)
    end
    @mac = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `network` attribute.
  # 
  def network
    return @network
  end
  
  # 
  # Sets the value of the `network` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Network} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def network=(value)
    if value.is_a?(Hash)
      value = Network.new(value)
    end
    @network = value
  end
  
  # 
  # Returns the value of the `network_attachments` attribute.
  # 
  def network_attachments
    return @network_attachments
  end
  
  # 
  # Sets the value of the `network_attachments` attribute.
  # 
  def network_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkAttachment.new(value)
        end
      end
    end
    @network_attachments = list
  end
  
  # 
  # Returns the value of the `network_labels` attribute.
  # 
  def network_labels
    return @network_labels
  end
  
  # 
  # Sets the value of the `network_labels` attribute.
  # 
  def network_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkLabel.new(value)
        end
      end
    end
    @network_labels = list
  end
  
  # 
  # Returns the value of the `on_boot` attribute.
  # 
  def on_boot
    return @on_boot
  end
  
  # 
  # Sets the value of the `on_boot` attribute.
  # 
  def on_boot=(value)
    @on_boot = value
  end
  
  # 
  # Returns the value of the `plugged` attribute.
  # 
  def plugged
    return @plugged
  end
  
  # 
  # Sets the value of the `plugged` attribute.
  # 
  def plugged=(value)
    @plugged = value
  end
  
  # 
  # Returns the value of the `reported_devices` attribute.
  # 
  def reported_devices
    return @reported_devices
  end
  
  # 
  # Sets the value of the `reported_devices` attribute.
  # 
  def reported_devices=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ReportedDevice.new(value)
        end
      end
    end
    @reported_devices = list
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `virtual_function_allowed_labels` attribute.
  # 
  def virtual_function_allowed_labels
    return @virtual_function_allowed_labels
  end
  
  # 
  # Sets the value of the `virtual_function_allowed_labels` attribute.
  # 
  def virtual_function_allowed_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NetworkLabel.new(value)
        end
      end
    end
    @virtual_function_allowed_labels = list
  end
  
  # 
  # Returns the value of the `virtual_function_allowed_networks` attribute.
  # 
  def virtual_function_allowed_networks
    return @virtual_function_allowed_networks
  end
  
  # 
  # Sets the value of the `virtual_function_allowed_networks` attribute.
  # 
  def virtual_function_allowed_networks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Network.new(value)
        end
      end
    end
    @virtual_function_allowed_networks = list
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Returns the value of the `vnic_profile` attribute.
  # 
  def vnic_profile
    return @vnic_profile
  end
  
  # 
  # Sets the value of the `vnic_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VnicProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vnic_profile=(value)
    if value.is_a?(Hash)
      value = VnicProfile.new(value)
    end
    @vnic_profile = value
  end
  
  # 
  # Creates a new instance of the {Nic} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :boot_protocol The value of attribute `boot_protocol`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :interface The value of attribute `interface`.
  # 
  # @option opts :linked The value of attribute `linked`.
  # 
  # @option opts [Hash] :mac The value of attribute `mac`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :network The value of attribute `network`.
  # 
  # @option opts [Array<Hash>] :network_attachments The values of attribute `network_attachments`.
  # 
  # @option opts [Array<Hash>] :network_labels The values of attribute `network_labels`.
  # 
  # @option opts :on_boot The value of attribute `on_boot`.
  # 
  # @option opts :plugged The value of attribute `plugged`.
  # 
  # @option opts [Array<Hash>] :reported_devices The values of attribute `reported_devices`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Array<Hash>] :virtual_function_allowed_labels The values of attribute `virtual_function_allowed_labels`.
  # 
  # @option opts [Array<Hash>] :virtual_function_allowed_networks The values of attribute `virtual_function_allowed_networks`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # @option opts [Hash] :vnic_profile The value of attribute `vnic_profile`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.boot_protocol = opts[:boot_protocol]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.interface = opts[:interface]
    self.linked = opts[:linked]
    self.mac = opts[:mac]
    self.name = opts[:name]
    self.network = opts[:network]
    self.network_attachments = opts[:network_attachments]
    self.network_labels = opts[:network_labels]
    self.on_boot = opts[:on_boot]
    self.plugged = opts[:plugged]
    self.reported_devices = opts[:reported_devices]
    self.statistics = opts[:statistics]
    self.template = opts[:template]
    self.virtual_function_allowed_labels = opts[:virtual_function_allowed_labels]
    self.virtual_function_allowed_networks = opts[:virtual_function_allowed_networks]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
    self.vnic_profile = opts[:vnic_profile]
  end
  
end

class OpenStackProvider < ExternalProvider
  
  # 
  # Returns the value of the `authentication_url` attribute.
  # 
  def authentication_url
    return @authentication_url
  end
  
  # 
  # Sets the value of the `authentication_url` attribute.
  # 
  def authentication_url=(value)
    @authentication_url = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `requires_authentication` attribute.
  # 
  def requires_authentication
    return @requires_authentication
  end
  
  # 
  # Sets the value of the `requires_authentication` attribute.
  # 
  def requires_authentication=(value)
    @requires_authentication = value
  end
  
  # 
  # Returns the value of the `tenant_name` attribute.
  # 
  def tenant_name
    return @tenant_name
  end
  
  # 
  # Sets the value of the `tenant_name` attribute.
  # 
  def tenant_name=(value)
    @tenant_name = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {OpenStackProvider} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :authentication_url The value of attribute `authentication_url`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :requires_authentication The value of attribute `requires_authentication`.
  # 
  # @option opts :tenant_name The value of attribute `tenant_name`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authentication_url = opts[:authentication_url]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.password = opts[:password]
    self.properties = opts[:properties]
    self.requires_authentication = opts[:requires_authentication]
    self.tenant_name = opts[:tenant_name]
    self.url = opts[:url]
    self.username = opts[:username]
  end
  
end

class OpenStackVolumeProvider < OpenStackProvider
  
  # 
  # Returns the value of the `authentication_keys` attribute.
  # 
  def authentication_keys
    return @authentication_keys
  end
  
  # 
  # Sets the value of the `authentication_keys` attribute.
  # 
  def authentication_keys=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = OpenstackVolumeAuthenticationKey.new(value)
        end
      end
    end
    @authentication_keys = list
  end
  
  # 
  # Returns the value of the `authentication_url` attribute.
  # 
  def authentication_url
    return @authentication_url
  end
  
  # 
  # Sets the value of the `authentication_url` attribute.
  # 
  def authentication_url=(value)
    @authentication_url = value
  end
  
  # 
  # Returns the value of the `certificates` attribute.
  # 
  def certificates
    return @certificates
  end
  
  # 
  # Sets the value of the `certificates` attribute.
  # 
  def certificates=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Certificate.new(value)
        end
      end
    end
    @certificates = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `data_center` attribute.
  # 
  def data_center
    return @data_center
  end
  
  # 
  # Sets the value of the `data_center` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::DataCenter} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def data_center=(value)
    if value.is_a?(Hash)
      value = DataCenter.new(value)
    end
    @data_center = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `requires_authentication` attribute.
  # 
  def requires_authentication
    return @requires_authentication
  end
  
  # 
  # Sets the value of the `requires_authentication` attribute.
  # 
  def requires_authentication=(value)
    @requires_authentication = value
  end
  
  # 
  # Returns the value of the `tenant_name` attribute.
  # 
  def tenant_name
    return @tenant_name
  end
  
  # 
  # Sets the value of the `tenant_name` attribute.
  # 
  def tenant_name=(value)
    @tenant_name = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Returns the value of the `volume_types` attribute.
  # 
  def volume_types
    return @volume_types
  end
  
  # 
  # Sets the value of the `volume_types` attribute.
  # 
  def volume_types=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = OpenStackVolumeType.new(value)
        end
      end
    end
    @volume_types = list
  end
  
  # 
  # Creates a new instance of the {OpenStackVolumeProvider} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :authentication_keys The values of attribute `authentication_keys`.
  # 
  # @option opts :authentication_url The value of attribute `authentication_url`.
  # 
  # @option opts [Array<Hash>] :certificates The values of attribute `certificates`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :data_center The value of attribute `data_center`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :requires_authentication The value of attribute `requires_authentication`.
  # 
  # @option opts :tenant_name The value of attribute `tenant_name`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # @option opts [Array<Hash>] :volume_types The values of attribute `volume_types`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authentication_keys = opts[:authentication_keys]
    self.authentication_url = opts[:authentication_url]
    self.certificates = opts[:certificates]
    self.comment = opts[:comment]
    self.data_center = opts[:data_center]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.password = opts[:password]
    self.properties = opts[:properties]
    self.requires_authentication = opts[:requires_authentication]
    self.tenant_name = opts[:tenant_name]
    self.url = opts[:url]
    self.username = opts[:username]
    self.volume_types = opts[:volume_types]
  end
  
end

class Template < VmBase
  
  # 
  # Returns the value of the `bios` attribute.
  # 
  def bios
    return @bios
  end
  
  # 
  # Sets the value of the `bios` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Bios} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bios=(value)
    if value.is_a?(Hash)
      value = Bios.new(value)
    end
    @bios = value
  end
  
  # 
  # Returns the value of the `cdroms` attribute.
  # 
  def cdroms
    return @cdroms
  end
  
  # 
  # Sets the value of the `cdroms` attribute.
  # 
  def cdroms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Cdrom.new(value)
        end
      end
    end
    @cdroms = list
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `console` attribute.
  # 
  def console
    return @console
  end
  
  # 
  # Sets the value of the `console` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Console} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def console=(value)
    if value.is_a?(Hash)
      value = Console.new(value)
    end
    @console = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `cpu_profile` attribute.
  # 
  def cpu_profile
    return @cpu_profile
  end
  
  # 
  # Sets the value of the `cpu_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu_profile=(value)
    if value.is_a?(Hash)
      value = CpuProfile.new(value)
    end
    @cpu_profile = value
  end
  
  # 
  # Returns the value of the `cpu_shares` attribute.
  # 
  def cpu_shares
    return @cpu_shares
  end
  
  # 
  # Sets the value of the `cpu_shares` attribute.
  # 
  def cpu_shares=(value)
    @cpu_shares = value
  end
  
  # 
  # Returns the value of the `creation_time` attribute.
  # 
  def creation_time
    return @creation_time
  end
  
  # 
  # Sets the value of the `creation_time` attribute.
  # 
  def creation_time=(value)
    @creation_time = value
  end
  
  # 
  # Returns the value of the `custom_compatibility_version` attribute.
  # 
  def custom_compatibility_version
    return @custom_compatibility_version
  end
  
  # 
  # Sets the value of the `custom_compatibility_version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def custom_compatibility_version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @custom_compatibility_version = value
  end
  
  # 
  # Returns the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model
    return @custom_cpu_model
  end
  
  # 
  # Sets the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model=(value)
    @custom_cpu_model = value
  end
  
  # 
  # Returns the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine
    return @custom_emulated_machine
  end
  
  # 
  # Sets the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine=(value)
    @custom_emulated_machine = value
  end
  
  # 
  # Returns the value of the `custom_properties` attribute.
  # 
  def custom_properties
    return @custom_properties
  end
  
  # 
  # Sets the value of the `custom_properties` attribute.
  # 
  def custom_properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CustomProperty.new(value)
        end
      end
    end
    @custom_properties = list
  end
  
  # 
  # Returns the value of the `delete_protected` attribute.
  # 
  def delete_protected
    return @delete_protected
  end
  
  # 
  # Sets the value of the `delete_protected` attribute.
  # 
  def delete_protected=(value)
    @delete_protected = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk_attachments` attribute.
  # 
  def disk_attachments
    return @disk_attachments
  end
  
  # 
  # Sets the value of the `disk_attachments` attribute.
  # 
  def disk_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DiskAttachment.new(value)
        end
      end
    end
    @disk_attachments = list
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles
    return @graphics_consoles
  end
  
  # 
  # Sets the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GraphicsConsole.new(value)
        end
      end
    end
    @graphics_consoles = list
  end
  
  # 
  # Returns the value of the `high_availability` attribute.
  # 
  def high_availability
    return @high_availability
  end
  
  # 
  # Sets the value of the `high_availability` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HighAvailability} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def high_availability=(value)
    if value.is_a?(Hash)
      value = HighAvailability.new(value)
    end
    @high_availability = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `initialization` attribute.
  # 
  def initialization
    return @initialization
  end
  
  # 
  # Sets the value of the `initialization` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Initialization} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def initialization=(value)
    if value.is_a?(Hash)
      value = Initialization.new(value)
    end
    @initialization = value
  end
  
  # 
  # Returns the value of the `io` attribute.
  # 
  def io
    return @io
  end
  
  # 
  # Sets the value of the `io` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Io} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def io=(value)
    if value.is_a?(Hash)
      value = Io.new(value)
    end
    @io = value
  end
  
  # 
  # Returns the value of the `large_icon` attribute.
  # 
  def large_icon
    return @large_icon
  end
  
  # 
  # Sets the value of the `large_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def large_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @large_icon = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `memory_policy` attribute.
  # 
  def memory_policy
    return @memory_policy
  end
  
  # 
  # Sets the value of the `memory_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def memory_policy=(value)
    if value.is_a?(Hash)
      value = MemoryPolicy.new(value)
    end
    @memory_policy = value
  end
  
  # 
  # Returns the value of the `migration` attribute.
  # 
  def migration
    return @migration
  end
  
  # 
  # Sets the value of the `migration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationOptions} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def migration=(value)
    if value.is_a?(Hash)
      value = MigrationOptions.new(value)
    end
    @migration = value
  end
  
  # 
  # Returns the value of the `migration_downtime` attribute.
  # 
  def migration_downtime
    return @migration_downtime
  end
  
  # 
  # Sets the value of the `migration_downtime` attribute.
  # 
  def migration_downtime=(value)
    @migration_downtime = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `nics` attribute.
  # 
  def nics
    return @nics
  end
  
  # 
  # Sets the value of the `nics` attribute.
  # 
  def nics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Nic.new(value)
        end
      end
    end
    @nics = list
  end
  
  # 
  # Returns the value of the `origin` attribute.
  # 
  def origin
    return @origin
  end
  
  # 
  # Sets the value of the `origin` attribute.
  # 
  def origin=(value)
    @origin = value
  end
  
  # 
  # Returns the value of the `os` attribute.
  # 
  def os
    return @os
  end
  
  # 
  # Sets the value of the `os` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def os=(value)
    if value.is_a?(Hash)
      value = OperatingSystem.new(value)
    end
    @os = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `rng_device` attribute.
  # 
  def rng_device
    return @rng_device
  end
  
  # 
  # Sets the value of the `rng_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::RngDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rng_device=(value)
    if value.is_a?(Hash)
      value = RngDevice.new(value)
    end
    @rng_device = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SerialNumber} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def serial_number=(value)
    if value.is_a?(Hash)
      value = SerialNumber.new(value)
    end
    @serial_number = value
  end
  
  # 
  # Returns the value of the `small_icon` attribute.
  # 
  def small_icon
    return @small_icon
  end
  
  # 
  # Sets the value of the `small_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def small_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @small_icon = value
  end
  
  # 
  # Returns the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled
    return @soundcard_enabled
  end
  
  # 
  # Sets the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled=(value)
    @soundcard_enabled = value
  end
  
  # 
  # Returns the value of the `sso` attribute.
  # 
  def sso
    return @sso
  end
  
  # 
  # Sets the value of the `sso` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Sso} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def sso=(value)
    if value.is_a?(Hash)
      value = Sso.new(value)
    end
    @sso = value
  end
  
  # 
  # Returns the value of the `start_paused` attribute.
  # 
  def start_paused
    return @start_paused
  end
  
  # 
  # Sets the value of the `start_paused` attribute.
  # 
  def start_paused=(value)
    @start_paused = value
  end
  
  # 
  # Returns the value of the `stateless` attribute.
  # 
  def stateless
    return @stateless
  end
  
  # 
  # Sets the value of the `stateless` attribute.
  # 
  def stateless=(value)
    @stateless = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Returns the value of the `time_zone` attribute.
  # 
  def time_zone
    return @time_zone
  end
  
  # 
  # Sets the value of the `time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @time_zone = value
  end
  
  # 
  # Returns the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration
    return @tunnel_migration
  end
  
  # 
  # Sets the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration=(value)
    @tunnel_migration = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `usb` attribute.
  # 
  def usb
    return @usb
  end
  
  # 
  # Sets the value of the `usb` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Usb} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def usb=(value)
    if value.is_a?(Hash)
      value = Usb.new(value)
    end
    @usb = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TemplateVersion} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = TemplateVersion.new(value)
    end
    @version = value
  end
  
  # 
  # Returns the value of the `virtio_scsi` attribute.
  # 
  def virtio_scsi
    return @virtio_scsi
  end
  
  # 
  # Sets the value of the `virtio_scsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VirtioScsi} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtio_scsi=(value)
    if value.is_a?(Hash)
      value = VirtioScsi.new(value)
    end
    @virtio_scsi = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `watchdogs` attribute.
  # 
  def watchdogs
    return @watchdogs
  end
  
  # 
  # Sets the value of the `watchdogs` attribute.
  # 
  def watchdogs=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Watchdog.new(value)
        end
      end
    end
    @watchdogs = list
  end
  
  # 
  # Creates a new instance of the {Template} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :bios The value of attribute `bios`.
  # 
  # @option opts [Array<Hash>] :cdroms The values of attribute `cdroms`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :console The value of attribute `console`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts [Hash] :cpu_profile The value of attribute `cpu_profile`.
  # 
  # @option opts :cpu_shares The value of attribute `cpu_shares`.
  # 
  # @option opts :creation_time The value of attribute `creation_time`.
  # 
  # @option opts [Hash] :custom_compatibility_version The value of attribute `custom_compatibility_version`.
  # 
  # @option opts :custom_cpu_model The value of attribute `custom_cpu_model`.
  # 
  # @option opts :custom_emulated_machine The value of attribute `custom_emulated_machine`.
  # 
  # @option opts [Array<Hash>] :custom_properties The values of attribute `custom_properties`.
  # 
  # @option opts :delete_protected The value of attribute `delete_protected`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :disk_attachments The values of attribute `disk_attachments`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts [Array<Hash>] :graphics_consoles The values of attribute `graphics_consoles`.
  # 
  # @option opts [Hash] :high_availability The value of attribute `high_availability`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :initialization The value of attribute `initialization`.
  # 
  # @option opts [Hash] :io The value of attribute `io`.
  # 
  # @option opts [Hash] :large_icon The value of attribute `large_icon`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts [Hash] :memory_policy The value of attribute `memory_policy`.
  # 
  # @option opts [Hash] :migration The value of attribute `migration`.
  # 
  # @option opts :migration_downtime The value of attribute `migration_downtime`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :nics The values of attribute `nics`.
  # 
  # @option opts :origin The value of attribute `origin`.
  # 
  # @option opts [Hash] :os The value of attribute `os`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Hash] :rng_device The value of attribute `rng_device`.
  # 
  # @option opts [Hash] :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Hash] :small_icon The value of attribute `small_icon`.
  # 
  # @option opts :soundcard_enabled The value of attribute `soundcard_enabled`.
  # 
  # @option opts [Hash] :sso The value of attribute `sso`.
  # 
  # @option opts :start_paused The value of attribute `start_paused`.
  # 
  # @option opts :stateless The value of attribute `stateless`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # @option opts [Hash] :time_zone The value of attribute `time_zone`.
  # 
  # @option opts :tunnel_migration The value of attribute `tunnel_migration`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :usb The value of attribute `usb`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # @option opts [Hash] :virtio_scsi The value of attribute `virtio_scsi`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :watchdogs The values of attribute `watchdogs`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.bios = opts[:bios]
    self.cdroms = opts[:cdroms]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.console = opts[:console]
    self.cpu = opts[:cpu]
    self.cpu_profile = opts[:cpu_profile]
    self.cpu_shares = opts[:cpu_shares]
    self.creation_time = opts[:creation_time]
    self.custom_compatibility_version = opts[:custom_compatibility_version]
    self.custom_cpu_model = opts[:custom_cpu_model]
    self.custom_emulated_machine = opts[:custom_emulated_machine]
    self.custom_properties = opts[:custom_properties]
    self.delete_protected = opts[:delete_protected]
    self.description = opts[:description]
    self.disk_attachments = opts[:disk_attachments]
    self.display = opts[:display]
    self.domain = opts[:domain]
    self.graphics_consoles = opts[:graphics_consoles]
    self.high_availability = opts[:high_availability]
    self.id = opts[:id]
    self.initialization = opts[:initialization]
    self.io = opts[:io]
    self.large_icon = opts[:large_icon]
    self.memory = opts[:memory]
    self.memory_policy = opts[:memory_policy]
    self.migration = opts[:migration]
    self.migration_downtime = opts[:migration_downtime]
    self.name = opts[:name]
    self.nics = opts[:nics]
    self.origin = opts[:origin]
    self.os = opts[:os]
    self.permissions = opts[:permissions]
    self.rng_device = opts[:rng_device]
    self.serial_number = opts[:serial_number]
    self.small_icon = opts[:small_icon]
    self.soundcard_enabled = opts[:soundcard_enabled]
    self.sso = opts[:sso]
    self.start_paused = opts[:start_paused]
    self.stateless = opts[:stateless]
    self.status = opts[:status]
    self.storage_domain = opts[:storage_domain]
    self.tags = opts[:tags]
    self.time_zone = opts[:time_zone]
    self.tunnel_migration = opts[:tunnel_migration]
    self.type = opts[:type]
    self.usb = opts[:usb]
    self.version = opts[:version]
    self.virtio_scsi = opts[:virtio_scsi]
    self.vm = opts[:vm]
    self.watchdogs = opts[:watchdogs]
  end
  
end

class Vm < VmBase
  
  # 
  # Returns the value of the `affinity_labels` attribute.
  # 
  def affinity_labels
    return @affinity_labels
  end
  
  # 
  # Sets the value of the `affinity_labels` attribute.
  # 
  def affinity_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = AffinityLabel.new(value)
        end
      end
    end
    @affinity_labels = list
  end
  
  # 
  # Returns the value of the `applications` attribute.
  # 
  def applications
    return @applications
  end
  
  # 
  # Sets the value of the `applications` attribute.
  # 
  def applications=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Application.new(value)
        end
      end
    end
    @applications = list
  end
  
  # 
  # Returns the value of the `bios` attribute.
  # 
  def bios
    return @bios
  end
  
  # 
  # Sets the value of the `bios` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Bios} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bios=(value)
    if value.is_a?(Hash)
      value = Bios.new(value)
    end
    @bios = value
  end
  
  # 
  # Returns the value of the `cdroms` attribute.
  # 
  def cdroms
    return @cdroms
  end
  
  # 
  # Sets the value of the `cdroms` attribute.
  # 
  def cdroms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Cdrom.new(value)
        end
      end
    end
    @cdroms = list
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `console` attribute.
  # 
  def console
    return @console
  end
  
  # 
  # Sets the value of the `console` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Console} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def console=(value)
    if value.is_a?(Hash)
      value = Console.new(value)
    end
    @console = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `cpu_profile` attribute.
  # 
  def cpu_profile
    return @cpu_profile
  end
  
  # 
  # Sets the value of the `cpu_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu_profile=(value)
    if value.is_a?(Hash)
      value = CpuProfile.new(value)
    end
    @cpu_profile = value
  end
  
  # 
  # Returns the value of the `cpu_shares` attribute.
  # 
  def cpu_shares
    return @cpu_shares
  end
  
  # 
  # Sets the value of the `cpu_shares` attribute.
  # 
  def cpu_shares=(value)
    @cpu_shares = value
  end
  
  # 
  # Returns the value of the `creation_time` attribute.
  # 
  def creation_time
    return @creation_time
  end
  
  # 
  # Sets the value of the `creation_time` attribute.
  # 
  def creation_time=(value)
    @creation_time = value
  end
  
  # 
  # Returns the value of the `custom_compatibility_version` attribute.
  # 
  def custom_compatibility_version
    return @custom_compatibility_version
  end
  
  # 
  # Sets the value of the `custom_compatibility_version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def custom_compatibility_version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @custom_compatibility_version = value
  end
  
  # 
  # Returns the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model
    return @custom_cpu_model
  end
  
  # 
  # Sets the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model=(value)
    @custom_cpu_model = value
  end
  
  # 
  # Returns the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine
    return @custom_emulated_machine
  end
  
  # 
  # Sets the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine=(value)
    @custom_emulated_machine = value
  end
  
  # 
  # Returns the value of the `custom_properties` attribute.
  # 
  def custom_properties
    return @custom_properties
  end
  
  # 
  # Sets the value of the `custom_properties` attribute.
  # 
  def custom_properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CustomProperty.new(value)
        end
      end
    end
    @custom_properties = list
  end
  
  # 
  # Returns the value of the `delete_protected` attribute.
  # 
  def delete_protected
    return @delete_protected
  end
  
  # 
  # Sets the value of the `delete_protected` attribute.
  # 
  def delete_protected=(value)
    @delete_protected = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk_attachments` attribute.
  # 
  def disk_attachments
    return @disk_attachments
  end
  
  # 
  # Sets the value of the `disk_attachments` attribute.
  # 
  def disk_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DiskAttachment.new(value)
        end
      end
    end
    @disk_attachments = list
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `floppies` attribute.
  # 
  def floppies
    return @floppies
  end
  
  # 
  # Sets the value of the `floppies` attribute.
  # 
  def floppies=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Floppy.new(value)
        end
      end
    end
    @floppies = list
  end
  
  # 
  # Returns the value of the `fqdn` attribute.
  # 
  def fqdn
    return @fqdn
  end
  
  # 
  # Sets the value of the `fqdn` attribute.
  # 
  def fqdn=(value)
    @fqdn = value
  end
  
  # 
  # Returns the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles
    return @graphics_consoles
  end
  
  # 
  # Sets the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GraphicsConsole.new(value)
        end
      end
    end
    @graphics_consoles = list
  end
  
  # 
  # Returns the value of the `guest_operating_system` attribute.
  # 
  def guest_operating_system
    return @guest_operating_system
  end
  
  # 
  # Sets the value of the `guest_operating_system` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GuestOperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def guest_operating_system=(value)
    if value.is_a?(Hash)
      value = GuestOperatingSystem.new(value)
    end
    @guest_operating_system = value
  end
  
  # 
  # Returns the value of the `guest_time_zone` attribute.
  # 
  def guest_time_zone
    return @guest_time_zone
  end
  
  # 
  # Sets the value of the `guest_time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def guest_time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @guest_time_zone = value
  end
  
  # 
  # Returns the value of the `high_availability` attribute.
  # 
  def high_availability
    return @high_availability
  end
  
  # 
  # Sets the value of the `high_availability` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HighAvailability} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def high_availability=(value)
    if value.is_a?(Hash)
      value = HighAvailability.new(value)
    end
    @high_availability = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `host_devices` attribute.
  # 
  def host_devices
    return @host_devices
  end
  
  # 
  # Sets the value of the `host_devices` attribute.
  # 
  def host_devices=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = HostDevice.new(value)
        end
      end
    end
    @host_devices = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `initialization` attribute.
  # 
  def initialization
    return @initialization
  end
  
  # 
  # Sets the value of the `initialization` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Initialization} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def initialization=(value)
    if value.is_a?(Hash)
      value = Initialization.new(value)
    end
    @initialization = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `io` attribute.
  # 
  def io
    return @io
  end
  
  # 
  # Sets the value of the `io` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Io} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def io=(value)
    if value.is_a?(Hash)
      value = Io.new(value)
    end
    @io = value
  end
  
  # 
  # Returns the value of the `katello_errata` attribute.
  # 
  def katello_errata
    return @katello_errata
  end
  
  # 
  # Sets the value of the `katello_errata` attribute.
  # 
  def katello_errata=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = KatelloErratum.new(value)
        end
      end
    end
    @katello_errata = list
  end
  
  # 
  # Returns the value of the `large_icon` attribute.
  # 
  def large_icon
    return @large_icon
  end
  
  # 
  # Sets the value of the `large_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def large_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @large_icon = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `memory_policy` attribute.
  # 
  def memory_policy
    return @memory_policy
  end
  
  # 
  # Sets the value of the `memory_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def memory_policy=(value)
    if value.is_a?(Hash)
      value = MemoryPolicy.new(value)
    end
    @memory_policy = value
  end
  
  # 
  # Returns the value of the `migration` attribute.
  # 
  def migration
    return @migration
  end
  
  # 
  # Sets the value of the `migration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationOptions} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def migration=(value)
    if value.is_a?(Hash)
      value = MigrationOptions.new(value)
    end
    @migration = value
  end
  
  # 
  # Returns the value of the `migration_downtime` attribute.
  # 
  def migration_downtime
    return @migration_downtime
  end
  
  # 
  # Sets the value of the `migration_downtime` attribute.
  # 
  def migration_downtime=(value)
    @migration_downtime = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `next_run_configuration_exists` attribute.
  # 
  def next_run_configuration_exists
    return @next_run_configuration_exists
  end
  
  # 
  # Sets the value of the `next_run_configuration_exists` attribute.
  # 
  def next_run_configuration_exists=(value)
    @next_run_configuration_exists = value
  end
  
  # 
  # Returns the value of the `nics` attribute.
  # 
  def nics
    return @nics
  end
  
  # 
  # Sets the value of the `nics` attribute.
  # 
  def nics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Nic.new(value)
        end
      end
    end
    @nics = list
  end
  
  # 
  # Returns the value of the `numa_nodes` attribute.
  # 
  def numa_nodes
    return @numa_nodes
  end
  
  # 
  # Sets the value of the `numa_nodes` attribute.
  # 
  def numa_nodes=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NumaNode.new(value)
        end
      end
    end
    @numa_nodes = list
  end
  
  # 
  # Returns the value of the `numa_tune_mode` attribute.
  # 
  def numa_tune_mode
    return @numa_tune_mode
  end
  
  # 
  # Sets the value of the `numa_tune_mode` attribute.
  # 
  def numa_tune_mode=(value)
    @numa_tune_mode = value
  end
  
  # 
  # Returns the value of the `origin` attribute.
  # 
  def origin
    return @origin
  end
  
  # 
  # Sets the value of the `origin` attribute.
  # 
  def origin=(value)
    @origin = value
  end
  
  # 
  # Returns the value of the `os` attribute.
  # 
  def os
    return @os
  end
  
  # 
  # Sets the value of the `os` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def os=(value)
    if value.is_a?(Hash)
      value = OperatingSystem.new(value)
    end
    @os = value
  end
  
  # 
  # Returns the value of the `payloads` attribute.
  # 
  def payloads
    return @payloads
  end
  
  # 
  # Sets the value of the `payloads` attribute.
  # 
  def payloads=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Payload.new(value)
        end
      end
    end
    @payloads = list
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `placement_policy` attribute.
  # 
  def placement_policy
    return @placement_policy
  end
  
  # 
  # Sets the value of the `placement_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VmPlacementPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def placement_policy=(value)
    if value.is_a?(Hash)
      value = VmPlacementPolicy.new(value)
    end
    @placement_policy = value
  end
  
  # 
  # Returns the value of the `quota` attribute.
  # 
  def quota
    return @quota
  end
  
  # 
  # Sets the value of the `quota` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Quota} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def quota=(value)
    if value.is_a?(Hash)
      value = Quota.new(value)
    end
    @quota = value
  end
  
  # 
  # Returns the value of the `reported_devices` attribute.
  # 
  def reported_devices
    return @reported_devices
  end
  
  # 
  # Sets the value of the `reported_devices` attribute.
  # 
  def reported_devices=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ReportedDevice.new(value)
        end
      end
    end
    @reported_devices = list
  end
  
  # 
  # Returns the value of the `rng_device` attribute.
  # 
  def rng_device
    return @rng_device
  end
  
  # 
  # Sets the value of the `rng_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::RngDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rng_device=(value)
    if value.is_a?(Hash)
      value = RngDevice.new(value)
    end
    @rng_device = value
  end
  
  # 
  # Returns the value of the `run_once` attribute.
  # 
  def run_once
    return @run_once
  end
  
  # 
  # Sets the value of the `run_once` attribute.
  # 
  def run_once=(value)
    @run_once = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SerialNumber} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def serial_number=(value)
    if value.is_a?(Hash)
      value = SerialNumber.new(value)
    end
    @serial_number = value
  end
  
  # 
  # Returns the value of the `sessions` attribute.
  # 
  def sessions
    return @sessions
  end
  
  # 
  # Sets the value of the `sessions` attribute.
  # 
  def sessions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Session.new(value)
        end
      end
    end
    @sessions = list
  end
  
  # 
  # Returns the value of the `small_icon` attribute.
  # 
  def small_icon
    return @small_icon
  end
  
  # 
  # Sets the value of the `small_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def small_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @small_icon = value
  end
  
  # 
  # Returns the value of the `snapshots` attribute.
  # 
  def snapshots
    return @snapshots
  end
  
  # 
  # Sets the value of the `snapshots` attribute.
  # 
  def snapshots=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Snapshot.new(value)
        end
      end
    end
    @snapshots = list
  end
  
  # 
  # Returns the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled
    return @soundcard_enabled
  end
  
  # 
  # Sets the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled=(value)
    @soundcard_enabled = value
  end
  
  # 
  # Returns the value of the `sso` attribute.
  # 
  def sso
    return @sso
  end
  
  # 
  # Sets the value of the `sso` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Sso} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def sso=(value)
    if value.is_a?(Hash)
      value = Sso.new(value)
    end
    @sso = value
  end
  
  # 
  # Returns the value of the `start_paused` attribute.
  # 
  def start_paused
    return @start_paused
  end
  
  # 
  # Sets the value of the `start_paused` attribute.
  # 
  def start_paused=(value)
    @start_paused = value
  end
  
  # 
  # Returns the value of the `start_time` attribute.
  # 
  def start_time
    return @start_time
  end
  
  # 
  # Sets the value of the `start_time` attribute.
  # 
  def start_time=(value)
    @start_time = value
  end
  
  # 
  # Returns the value of the `stateless` attribute.
  # 
  def stateless
    return @stateless
  end
  
  # 
  # Sets the value of the `stateless` attribute.
  # 
  def stateless=(value)
    @stateless = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `status_detail` attribute.
  # 
  def status_detail
    return @status_detail
  end
  
  # 
  # Sets the value of the `status_detail` attribute.
  # 
  def status_detail=(value)
    @status_detail = value
  end
  
  # 
  # Returns the value of the `stop_reason` attribute.
  # 
  def stop_reason
    return @stop_reason
  end
  
  # 
  # Sets the value of the `stop_reason` attribute.
  # 
  def stop_reason=(value)
    @stop_reason = value
  end
  
  # 
  # Returns the value of the `stop_time` attribute.
  # 
  def stop_time
    return @stop_time
  end
  
  # 
  # Sets the value of the `stop_time` attribute.
  # 
  def stop_time=(value)
    @stop_time = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `time_zone` attribute.
  # 
  def time_zone
    return @time_zone
  end
  
  # 
  # Sets the value of the `time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @time_zone = value
  end
  
  # 
  # Returns the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration
    return @tunnel_migration
  end
  
  # 
  # Sets the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration=(value)
    @tunnel_migration = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `usb` attribute.
  # 
  def usb
    return @usb
  end
  
  # 
  # Sets the value of the `usb` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Usb} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def usb=(value)
    if value.is_a?(Hash)
      value = Usb.new(value)
    end
    @usb = value
  end
  
  # 
  # Returns the value of the `use_latest_template_version` attribute.
  # 
  def use_latest_template_version
    return @use_latest_template_version
  end
  
  # 
  # Sets the value of the `use_latest_template_version` attribute.
  # 
  def use_latest_template_version=(value)
    @use_latest_template_version = value
  end
  
  # 
  # Returns the value of the `virtio_scsi` attribute.
  # 
  def virtio_scsi
    return @virtio_scsi
  end
  
  # 
  # Sets the value of the `virtio_scsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VirtioScsi} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtio_scsi=(value)
    if value.is_a?(Hash)
      value = VirtioScsi.new(value)
    end
    @virtio_scsi = value
  end
  
  # 
  # Returns the value of the `vm_pool` attribute.
  # 
  def vm_pool
    return @vm_pool
  end
  
  # 
  # Sets the value of the `vm_pool` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VmPool} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm_pool=(value)
    if value.is_a?(Hash)
      value = VmPool.new(value)
    end
    @vm_pool = value
  end
  
  # 
  # Returns the value of the `watchdogs` attribute.
  # 
  def watchdogs
    return @watchdogs
  end
  
  # 
  # Sets the value of the `watchdogs` attribute.
  # 
  def watchdogs=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Watchdog.new(value)
        end
      end
    end
    @watchdogs = list
  end
  
  # 
  # Creates a new instance of the {Vm} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :affinity_labels The values of attribute `affinity_labels`.
  # 
  # @option opts [Array<Hash>] :applications The values of attribute `applications`.
  # 
  # @option opts [Hash] :bios The value of attribute `bios`.
  # 
  # @option opts [Array<Hash>] :cdroms The values of attribute `cdroms`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :console The value of attribute `console`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts [Hash] :cpu_profile The value of attribute `cpu_profile`.
  # 
  # @option opts :cpu_shares The value of attribute `cpu_shares`.
  # 
  # @option opts :creation_time The value of attribute `creation_time`.
  # 
  # @option opts [Hash] :custom_compatibility_version The value of attribute `custom_compatibility_version`.
  # 
  # @option opts :custom_cpu_model The value of attribute `custom_cpu_model`.
  # 
  # @option opts :custom_emulated_machine The value of attribute `custom_emulated_machine`.
  # 
  # @option opts [Array<Hash>] :custom_properties The values of attribute `custom_properties`.
  # 
  # @option opts :delete_protected The value of attribute `delete_protected`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :disk_attachments The values of attribute `disk_attachments`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts [Array<Hash>] :floppies The values of attribute `floppies`.
  # 
  # @option opts :fqdn The value of attribute `fqdn`.
  # 
  # @option opts [Array<Hash>] :graphics_consoles The values of attribute `graphics_consoles`.
  # 
  # @option opts [Hash] :guest_operating_system The value of attribute `guest_operating_system`.
  # 
  # @option opts [Hash] :guest_time_zone The value of attribute `guest_time_zone`.
  # 
  # @option opts [Hash] :high_availability The value of attribute `high_availability`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts [Array<Hash>] :host_devices The values of attribute `host_devices`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :initialization The value of attribute `initialization`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts [Hash] :io The value of attribute `io`.
  # 
  # @option opts [Array<Hash>] :katello_errata The values of attribute `katello_errata`.
  # 
  # @option opts [Hash] :large_icon The value of attribute `large_icon`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts [Hash] :memory_policy The value of attribute `memory_policy`.
  # 
  # @option opts [Hash] :migration The value of attribute `migration`.
  # 
  # @option opts :migration_downtime The value of attribute `migration_downtime`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :next_run_configuration_exists The value of attribute `next_run_configuration_exists`.
  # 
  # @option opts [Array<Hash>] :nics The values of attribute `nics`.
  # 
  # @option opts [Array<Hash>] :numa_nodes The values of attribute `numa_nodes`.
  # 
  # @option opts :numa_tune_mode The value of attribute `numa_tune_mode`.
  # 
  # @option opts :origin The value of attribute `origin`.
  # 
  # @option opts [Hash] :os The value of attribute `os`.
  # 
  # @option opts [Array<Hash>] :payloads The values of attribute `payloads`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Hash] :placement_policy The value of attribute `placement_policy`.
  # 
  # @option opts [Hash] :quota The value of attribute `quota`.
  # 
  # @option opts [Array<Hash>] :reported_devices The values of attribute `reported_devices`.
  # 
  # @option opts [Hash] :rng_device The value of attribute `rng_device`.
  # 
  # @option opts :run_once The value of attribute `run_once`.
  # 
  # @option opts [Hash] :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Array<Hash>] :sessions The values of attribute `sessions`.
  # 
  # @option opts [Hash] :small_icon The value of attribute `small_icon`.
  # 
  # @option opts [Array<Hash>] :snapshots The values of attribute `snapshots`.
  # 
  # @option opts :soundcard_enabled The value of attribute `soundcard_enabled`.
  # 
  # @option opts [Hash] :sso The value of attribute `sso`.
  # 
  # @option opts :start_paused The value of attribute `start_paused`.
  # 
  # @option opts :start_time The value of attribute `start_time`.
  # 
  # @option opts :stateless The value of attribute `stateless`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :status_detail The value of attribute `status_detail`.
  # 
  # @option opts :stop_reason The value of attribute `stop_reason`.
  # 
  # @option opts :stop_time The value of attribute `stop_time`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :time_zone The value of attribute `time_zone`.
  # 
  # @option opts :tunnel_migration The value of attribute `tunnel_migration`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :usb The value of attribute `usb`.
  # 
  # @option opts :use_latest_template_version The value of attribute `use_latest_template_version`.
  # 
  # @option opts [Hash] :virtio_scsi The value of attribute `virtio_scsi`.
  # 
  # @option opts [Hash] :vm_pool The value of attribute `vm_pool`.
  # 
  # @option opts [Array<Hash>] :watchdogs The values of attribute `watchdogs`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.affinity_labels = opts[:affinity_labels]
    self.applications = opts[:applications]
    self.bios = opts[:bios]
    self.cdroms = opts[:cdroms]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.console = opts[:console]
    self.cpu = opts[:cpu]
    self.cpu_profile = opts[:cpu_profile]
    self.cpu_shares = opts[:cpu_shares]
    self.creation_time = opts[:creation_time]
    self.custom_compatibility_version = opts[:custom_compatibility_version]
    self.custom_cpu_model = opts[:custom_cpu_model]
    self.custom_emulated_machine = opts[:custom_emulated_machine]
    self.custom_properties = opts[:custom_properties]
    self.delete_protected = opts[:delete_protected]
    self.description = opts[:description]
    self.disk_attachments = opts[:disk_attachments]
    self.display = opts[:display]
    self.domain = opts[:domain]
    self.external_host_provider = opts[:external_host_provider]
    self.floppies = opts[:floppies]
    self.fqdn = opts[:fqdn]
    self.graphics_consoles = opts[:graphics_consoles]
    self.guest_operating_system = opts[:guest_operating_system]
    self.guest_time_zone = opts[:guest_time_zone]
    self.high_availability = opts[:high_availability]
    self.host = opts[:host]
    self.host_devices = opts[:host_devices]
    self.id = opts[:id]
    self.initialization = opts[:initialization]
    self.instance_type = opts[:instance_type]
    self.io = opts[:io]
    self.katello_errata = opts[:katello_errata]
    self.large_icon = opts[:large_icon]
    self.memory = opts[:memory]
    self.memory_policy = opts[:memory_policy]
    self.migration = opts[:migration]
    self.migration_downtime = opts[:migration_downtime]
    self.name = opts[:name]
    self.next_run_configuration_exists = opts[:next_run_configuration_exists]
    self.nics = opts[:nics]
    self.numa_nodes = opts[:numa_nodes]
    self.numa_tune_mode = opts[:numa_tune_mode]
    self.origin = opts[:origin]
    self.os = opts[:os]
    self.payloads = opts[:payloads]
    self.permissions = opts[:permissions]
    self.placement_policy = opts[:placement_policy]
    self.quota = opts[:quota]
    self.reported_devices = opts[:reported_devices]
    self.rng_device = opts[:rng_device]
    self.run_once = opts[:run_once]
    self.serial_number = opts[:serial_number]
    self.sessions = opts[:sessions]
    self.small_icon = opts[:small_icon]
    self.snapshots = opts[:snapshots]
    self.soundcard_enabled = opts[:soundcard_enabled]
    self.sso = opts[:sso]
    self.start_paused = opts[:start_paused]
    self.start_time = opts[:start_time]
    self.stateless = opts[:stateless]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.status_detail = opts[:status_detail]
    self.stop_reason = opts[:stop_reason]
    self.stop_time = opts[:stop_time]
    self.storage_domain = opts[:storage_domain]
    self.tags = opts[:tags]
    self.template = opts[:template]
    self.time_zone = opts[:time_zone]
    self.tunnel_migration = opts[:tunnel_migration]
    self.type = opts[:type]
    self.usb = opts[:usb]
    self.use_latest_template_version = opts[:use_latest_template_version]
    self.virtio_scsi = opts[:virtio_scsi]
    self.vm_pool = opts[:vm_pool]
    self.watchdogs = opts[:watchdogs]
  end
  
end

class Watchdog < Device
  
  # 
  # Returns the value of the `action` attribute.
  # 
  def action
    return @action
  end
  
  # 
  # Sets the value of the `action` attribute.
  # 
  def action=(value)
    @action = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `model` attribute.
  # 
  def model
    return @model
  end
  
  # 
  # Sets the value of the `model` attribute.
  # 
  def model=(value)
    @model = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {Watchdog} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :action The value of attribute `action`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :model The value of attribute `model`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.action = opts[:action]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.model = opts[:model]
    self.name = opts[:name]
    self.template = opts[:template]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
  end
  
end

class Cdrom < Device
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `file` attribute.
  # 
  def file
    return @file
  end
  
  # 
  # Sets the value of the `file` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::File} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def file=(value)
    if value.is_a?(Hash)
      value = File.new(value)
    end
    @file = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {Cdrom} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Hash] :file The value of attribute `file`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.file = opts[:file]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.name = opts[:name]
    self.template = opts[:template]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
  end
  
end

class ExternalHostProvider < ExternalProvider
  
  # 
  # Returns the value of the `authentication_url` attribute.
  # 
  def authentication_url
    return @authentication_url
  end
  
  # 
  # Sets the value of the `authentication_url` attribute.
  # 
  def authentication_url=(value)
    @authentication_url = value
  end
  
  # 
  # Returns the value of the `certificates` attribute.
  # 
  def certificates
    return @certificates
  end
  
  # 
  # Sets the value of the `certificates` attribute.
  # 
  def certificates=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Certificate.new(value)
        end
      end
    end
    @certificates = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `compute_resources` attribute.
  # 
  def compute_resources
    return @compute_resources
  end
  
  # 
  # Sets the value of the `compute_resources` attribute.
  # 
  def compute_resources=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ExternalComputeResource.new(value)
        end
      end
    end
    @compute_resources = list
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `discovered_hosts` attribute.
  # 
  def discovered_hosts
    return @discovered_hosts
  end
  
  # 
  # Sets the value of the `discovered_hosts` attribute.
  # 
  def discovered_hosts=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ExternalDiscoveredHost.new(value)
        end
      end
    end
    @discovered_hosts = list
  end
  
  # 
  # Returns the value of the `host_groups` attribute.
  # 
  def host_groups
    return @host_groups
  end
  
  # 
  # Sets the value of the `host_groups` attribute.
  # 
  def host_groups=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ExternalHostGroup.new(value)
        end
      end
    end
    @host_groups = list
  end
  
  # 
  # Returns the value of the `hosts` attribute.
  # 
  def hosts
    return @hosts
  end
  
  # 
  # Sets the value of the `hosts` attribute.
  # 
  def hosts=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Host.new(value)
        end
      end
    end
    @hosts = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `requires_authentication` attribute.
  # 
  def requires_authentication
    return @requires_authentication
  end
  
  # 
  # Sets the value of the `requires_authentication` attribute.
  # 
  def requires_authentication=(value)
    @requires_authentication = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {ExternalHostProvider} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :authentication_url The value of attribute `authentication_url`.
  # 
  # @option opts [Array<Hash>] :certificates The values of attribute `certificates`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Array<Hash>] :compute_resources The values of attribute `compute_resources`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :discovered_hosts The values of attribute `discovered_hosts`.
  # 
  # @option opts [Array<Hash>] :host_groups The values of attribute `host_groups`.
  # 
  # @option opts [Array<Hash>] :hosts The values of attribute `hosts`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :requires_authentication The value of attribute `requires_authentication`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authentication_url = opts[:authentication_url]
    self.certificates = opts[:certificates]
    self.comment = opts[:comment]
    self.compute_resources = opts[:compute_resources]
    self.description = opts[:description]
    self.discovered_hosts = opts[:discovered_hosts]
    self.host_groups = opts[:host_groups]
    self.hosts = opts[:hosts]
    self.id = opts[:id]
    self.name = opts[:name]
    self.password = opts[:password]
    self.properties = opts[:properties]
    self.requires_authentication = opts[:requires_authentication]
    self.url = opts[:url]
    self.username = opts[:username]
  end
  
end

class GlusterBrick < GlusterBrickAdvancedDetails
  
  # 
  # Returns the value of the `brick_dir` attribute.
  # 
  def brick_dir
    return @brick_dir
  end
  
  # 
  # Sets the value of the `brick_dir` attribute.
  # 
  def brick_dir=(value)
    @brick_dir = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `device` attribute.
  # 
  def device
    return @device
  end
  
  # 
  # Sets the value of the `device` attribute.
  # 
  def device=(value)
    @device = value
  end
  
  # 
  # Returns the value of the `fs_name` attribute.
  # 
  def fs_name
    return @fs_name
  end
  
  # 
  # Sets the value of the `fs_name` attribute.
  # 
  def fs_name=(value)
    @fs_name = value
  end
  
  # 
  # Returns the value of the `gluster_clients` attribute.
  # 
  def gluster_clients
    return @gluster_clients
  end
  
  # 
  # Sets the value of the `gluster_clients` attribute.
  # 
  def gluster_clients=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterClient.new(value)
        end
      end
    end
    @gluster_clients = list
  end
  
  # 
  # Returns the value of the `gluster_volume` attribute.
  # 
  def gluster_volume
    return @gluster_volume
  end
  
  # 
  # Sets the value of the `gluster_volume` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GlusterVolume} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def gluster_volume=(value)
    if value.is_a?(Hash)
      value = GlusterVolume.new(value)
    end
    @gluster_volume = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `memory_pools` attribute.
  # 
  def memory_pools
    return @memory_pools
  end
  
  # 
  # Sets the value of the `memory_pools` attribute.
  # 
  def memory_pools=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GlusterMemoryPool.new(value)
        end
      end
    end
    @memory_pools = list
  end
  
  # 
  # Returns the value of the `mnt_options` attribute.
  # 
  def mnt_options
    return @mnt_options
  end
  
  # 
  # Sets the value of the `mnt_options` attribute.
  # 
  def mnt_options=(value)
    @mnt_options = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `pid` attribute.
  # 
  def pid
    return @pid
  end
  
  # 
  # Sets the value of the `pid` attribute.
  # 
  def pid=(value)
    @pid = value
  end
  
  # 
  # Returns the value of the `port` attribute.
  # 
  def port
    return @port
  end
  
  # 
  # Sets the value of the `port` attribute.
  # 
  def port=(value)
    @port = value
  end
  
  # 
  # Returns the value of the `server_id` attribute.
  # 
  def server_id
    return @server_id
  end
  
  # 
  # Sets the value of the `server_id` attribute.
  # 
  def server_id=(value)
    @server_id = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vms` attribute.
  # 
  def vms
    return @vms
  end
  
  # 
  # Sets the value of the `vms` attribute.
  # 
  def vms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Vm.new(value)
        end
      end
    end
    @vms = list
  end
  
  # 
  # Creates a new instance of the {GlusterBrick} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :brick_dir The value of attribute `brick_dir`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :device The value of attribute `device`.
  # 
  # @option opts :fs_name The value of attribute `fs_name`.
  # 
  # @option opts [Array<Hash>] :gluster_clients The values of attribute `gluster_clients`.
  # 
  # @option opts [Hash] :gluster_volume The value of attribute `gluster_volume`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts [Array<Hash>] :memory_pools The values of attribute `memory_pools`.
  # 
  # @option opts :mnt_options The value of attribute `mnt_options`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :pid The value of attribute `pid`.
  # 
  # @option opts :port The value of attribute `port`.
  # 
  # @option opts :server_id The value of attribute `server_id`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :vms The values of attribute `vms`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.brick_dir = opts[:brick_dir]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.device = opts[:device]
    self.fs_name = opts[:fs_name]
    self.gluster_clients = opts[:gluster_clients]
    self.gluster_volume = opts[:gluster_volume]
    self.id = opts[:id]
    self.instance_type = opts[:instance_type]
    self.memory_pools = opts[:memory_pools]
    self.mnt_options = opts[:mnt_options]
    self.name = opts[:name]
    self.pid = opts[:pid]
    self.port = opts[:port]
    self.server_id = opts[:server_id]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.template = opts[:template]
    self.vm = opts[:vm]
    self.vms = opts[:vms]
  end
  
end

class InstanceType < Template
  
  # 
  # Returns the value of the `bios` attribute.
  # 
  def bios
    return @bios
  end
  
  # 
  # Sets the value of the `bios` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Bios} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bios=(value)
    if value.is_a?(Hash)
      value = Bios.new(value)
    end
    @bios = value
  end
  
  # 
  # Returns the value of the `cdroms` attribute.
  # 
  def cdroms
    return @cdroms
  end
  
  # 
  # Sets the value of the `cdroms` attribute.
  # 
  def cdroms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Cdrom.new(value)
        end
      end
    end
    @cdroms = list
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `console` attribute.
  # 
  def console
    return @console
  end
  
  # 
  # Sets the value of the `console` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Console} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def console=(value)
    if value.is_a?(Hash)
      value = Console.new(value)
    end
    @console = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `cpu_profile` attribute.
  # 
  def cpu_profile
    return @cpu_profile
  end
  
  # 
  # Sets the value of the `cpu_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu_profile=(value)
    if value.is_a?(Hash)
      value = CpuProfile.new(value)
    end
    @cpu_profile = value
  end
  
  # 
  # Returns the value of the `cpu_shares` attribute.
  # 
  def cpu_shares
    return @cpu_shares
  end
  
  # 
  # Sets the value of the `cpu_shares` attribute.
  # 
  def cpu_shares=(value)
    @cpu_shares = value
  end
  
  # 
  # Returns the value of the `creation_time` attribute.
  # 
  def creation_time
    return @creation_time
  end
  
  # 
  # Sets the value of the `creation_time` attribute.
  # 
  def creation_time=(value)
    @creation_time = value
  end
  
  # 
  # Returns the value of the `custom_compatibility_version` attribute.
  # 
  def custom_compatibility_version
    return @custom_compatibility_version
  end
  
  # 
  # Sets the value of the `custom_compatibility_version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def custom_compatibility_version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @custom_compatibility_version = value
  end
  
  # 
  # Returns the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model
    return @custom_cpu_model
  end
  
  # 
  # Sets the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model=(value)
    @custom_cpu_model = value
  end
  
  # 
  # Returns the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine
    return @custom_emulated_machine
  end
  
  # 
  # Sets the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine=(value)
    @custom_emulated_machine = value
  end
  
  # 
  # Returns the value of the `custom_properties` attribute.
  # 
  def custom_properties
    return @custom_properties
  end
  
  # 
  # Sets the value of the `custom_properties` attribute.
  # 
  def custom_properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CustomProperty.new(value)
        end
      end
    end
    @custom_properties = list
  end
  
  # 
  # Returns the value of the `delete_protected` attribute.
  # 
  def delete_protected
    return @delete_protected
  end
  
  # 
  # Sets the value of the `delete_protected` attribute.
  # 
  def delete_protected=(value)
    @delete_protected = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk_attachments` attribute.
  # 
  def disk_attachments
    return @disk_attachments
  end
  
  # 
  # Sets the value of the `disk_attachments` attribute.
  # 
  def disk_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DiskAttachment.new(value)
        end
      end
    end
    @disk_attachments = list
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles
    return @graphics_consoles
  end
  
  # 
  # Sets the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GraphicsConsole.new(value)
        end
      end
    end
    @graphics_consoles = list
  end
  
  # 
  # Returns the value of the `high_availability` attribute.
  # 
  def high_availability
    return @high_availability
  end
  
  # 
  # Sets the value of the `high_availability` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HighAvailability} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def high_availability=(value)
    if value.is_a?(Hash)
      value = HighAvailability.new(value)
    end
    @high_availability = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `initialization` attribute.
  # 
  def initialization
    return @initialization
  end
  
  # 
  # Sets the value of the `initialization` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Initialization} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def initialization=(value)
    if value.is_a?(Hash)
      value = Initialization.new(value)
    end
    @initialization = value
  end
  
  # 
  # Returns the value of the `io` attribute.
  # 
  def io
    return @io
  end
  
  # 
  # Sets the value of the `io` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Io} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def io=(value)
    if value.is_a?(Hash)
      value = Io.new(value)
    end
    @io = value
  end
  
  # 
  # Returns the value of the `large_icon` attribute.
  # 
  def large_icon
    return @large_icon
  end
  
  # 
  # Sets the value of the `large_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def large_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @large_icon = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `memory_policy` attribute.
  # 
  def memory_policy
    return @memory_policy
  end
  
  # 
  # Sets the value of the `memory_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def memory_policy=(value)
    if value.is_a?(Hash)
      value = MemoryPolicy.new(value)
    end
    @memory_policy = value
  end
  
  # 
  # Returns the value of the `migration` attribute.
  # 
  def migration
    return @migration
  end
  
  # 
  # Sets the value of the `migration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationOptions} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def migration=(value)
    if value.is_a?(Hash)
      value = MigrationOptions.new(value)
    end
    @migration = value
  end
  
  # 
  # Returns the value of the `migration_downtime` attribute.
  # 
  def migration_downtime
    return @migration_downtime
  end
  
  # 
  # Sets the value of the `migration_downtime` attribute.
  # 
  def migration_downtime=(value)
    @migration_downtime = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `nics` attribute.
  # 
  def nics
    return @nics
  end
  
  # 
  # Sets the value of the `nics` attribute.
  # 
  def nics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Nic.new(value)
        end
      end
    end
    @nics = list
  end
  
  # 
  # Returns the value of the `origin` attribute.
  # 
  def origin
    return @origin
  end
  
  # 
  # Sets the value of the `origin` attribute.
  # 
  def origin=(value)
    @origin = value
  end
  
  # 
  # Returns the value of the `os` attribute.
  # 
  def os
    return @os
  end
  
  # 
  # Sets the value of the `os` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def os=(value)
    if value.is_a?(Hash)
      value = OperatingSystem.new(value)
    end
    @os = value
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `rng_device` attribute.
  # 
  def rng_device
    return @rng_device
  end
  
  # 
  # Sets the value of the `rng_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::RngDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rng_device=(value)
    if value.is_a?(Hash)
      value = RngDevice.new(value)
    end
    @rng_device = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SerialNumber} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def serial_number=(value)
    if value.is_a?(Hash)
      value = SerialNumber.new(value)
    end
    @serial_number = value
  end
  
  # 
  # Returns the value of the `small_icon` attribute.
  # 
  def small_icon
    return @small_icon
  end
  
  # 
  # Sets the value of the `small_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def small_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @small_icon = value
  end
  
  # 
  # Returns the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled
    return @soundcard_enabled
  end
  
  # 
  # Sets the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled=(value)
    @soundcard_enabled = value
  end
  
  # 
  # Returns the value of the `sso` attribute.
  # 
  def sso
    return @sso
  end
  
  # 
  # Sets the value of the `sso` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Sso} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def sso=(value)
    if value.is_a?(Hash)
      value = Sso.new(value)
    end
    @sso = value
  end
  
  # 
  # Returns the value of the `start_paused` attribute.
  # 
  def start_paused
    return @start_paused
  end
  
  # 
  # Sets the value of the `start_paused` attribute.
  # 
  def start_paused=(value)
    @start_paused = value
  end
  
  # 
  # Returns the value of the `stateless` attribute.
  # 
  def stateless
    return @stateless
  end
  
  # 
  # Sets the value of the `stateless` attribute.
  # 
  def stateless=(value)
    @stateless = value
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Returns the value of the `time_zone` attribute.
  # 
  def time_zone
    return @time_zone
  end
  
  # 
  # Sets the value of the `time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @time_zone = value
  end
  
  # 
  # Returns the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration
    return @tunnel_migration
  end
  
  # 
  # Sets the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration=(value)
    @tunnel_migration = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `usb` attribute.
  # 
  def usb
    return @usb
  end
  
  # 
  # Sets the value of the `usb` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Usb} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def usb=(value)
    if value.is_a?(Hash)
      value = Usb.new(value)
    end
    @usb = value
  end
  
  # 
  # Returns the value of the `version` attribute.
  # 
  def version
    return @version
  end
  
  # 
  # Sets the value of the `version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TemplateVersion} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def version=(value)
    if value.is_a?(Hash)
      value = TemplateVersion.new(value)
    end
    @version = value
  end
  
  # 
  # Returns the value of the `virtio_scsi` attribute.
  # 
  def virtio_scsi
    return @virtio_scsi
  end
  
  # 
  # Sets the value of the `virtio_scsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VirtioScsi} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtio_scsi=(value)
    if value.is_a?(Hash)
      value = VirtioScsi.new(value)
    end
    @virtio_scsi = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `watchdogs` attribute.
  # 
  def watchdogs
    return @watchdogs
  end
  
  # 
  # Sets the value of the `watchdogs` attribute.
  # 
  def watchdogs=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Watchdog.new(value)
        end
      end
    end
    @watchdogs = list
  end
  
  # 
  # Creates a new instance of the {InstanceType} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :bios The value of attribute `bios`.
  # 
  # @option opts [Array<Hash>] :cdroms The values of attribute `cdroms`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :console The value of attribute `console`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts [Hash] :cpu_profile The value of attribute `cpu_profile`.
  # 
  # @option opts :cpu_shares The value of attribute `cpu_shares`.
  # 
  # @option opts :creation_time The value of attribute `creation_time`.
  # 
  # @option opts [Hash] :custom_compatibility_version The value of attribute `custom_compatibility_version`.
  # 
  # @option opts :custom_cpu_model The value of attribute `custom_cpu_model`.
  # 
  # @option opts :custom_emulated_machine The value of attribute `custom_emulated_machine`.
  # 
  # @option opts [Array<Hash>] :custom_properties The values of attribute `custom_properties`.
  # 
  # @option opts :delete_protected The value of attribute `delete_protected`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :disk_attachments The values of attribute `disk_attachments`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts [Array<Hash>] :graphics_consoles The values of attribute `graphics_consoles`.
  # 
  # @option opts [Hash] :high_availability The value of attribute `high_availability`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :initialization The value of attribute `initialization`.
  # 
  # @option opts [Hash] :io The value of attribute `io`.
  # 
  # @option opts [Hash] :large_icon The value of attribute `large_icon`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts [Hash] :memory_policy The value of attribute `memory_policy`.
  # 
  # @option opts [Hash] :migration The value of attribute `migration`.
  # 
  # @option opts :migration_downtime The value of attribute `migration_downtime`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :nics The values of attribute `nics`.
  # 
  # @option opts :origin The value of attribute `origin`.
  # 
  # @option opts [Hash] :os The value of attribute `os`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts [Hash] :rng_device The value of attribute `rng_device`.
  # 
  # @option opts [Hash] :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Hash] :small_icon The value of attribute `small_icon`.
  # 
  # @option opts :soundcard_enabled The value of attribute `soundcard_enabled`.
  # 
  # @option opts [Hash] :sso The value of attribute `sso`.
  # 
  # @option opts :start_paused The value of attribute `start_paused`.
  # 
  # @option opts :stateless The value of attribute `stateless`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # @option opts [Hash] :time_zone The value of attribute `time_zone`.
  # 
  # @option opts :tunnel_migration The value of attribute `tunnel_migration`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :usb The value of attribute `usb`.
  # 
  # @option opts [Hash] :version The value of attribute `version`.
  # 
  # @option opts [Hash] :virtio_scsi The value of attribute `virtio_scsi`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Array<Hash>] :watchdogs The values of attribute `watchdogs`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.bios = opts[:bios]
    self.cdroms = opts[:cdroms]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.console = opts[:console]
    self.cpu = opts[:cpu]
    self.cpu_profile = opts[:cpu_profile]
    self.cpu_shares = opts[:cpu_shares]
    self.creation_time = opts[:creation_time]
    self.custom_compatibility_version = opts[:custom_compatibility_version]
    self.custom_cpu_model = opts[:custom_cpu_model]
    self.custom_emulated_machine = opts[:custom_emulated_machine]
    self.custom_properties = opts[:custom_properties]
    self.delete_protected = opts[:delete_protected]
    self.description = opts[:description]
    self.disk_attachments = opts[:disk_attachments]
    self.display = opts[:display]
    self.domain = opts[:domain]
    self.graphics_consoles = opts[:graphics_consoles]
    self.high_availability = opts[:high_availability]
    self.id = opts[:id]
    self.initialization = opts[:initialization]
    self.io = opts[:io]
    self.large_icon = opts[:large_icon]
    self.memory = opts[:memory]
    self.memory_policy = opts[:memory_policy]
    self.migration = opts[:migration]
    self.migration_downtime = opts[:migration_downtime]
    self.name = opts[:name]
    self.nics = opts[:nics]
    self.origin = opts[:origin]
    self.os = opts[:os]
    self.permissions = opts[:permissions]
    self.rng_device = opts[:rng_device]
    self.serial_number = opts[:serial_number]
    self.small_icon = opts[:small_icon]
    self.soundcard_enabled = opts[:soundcard_enabled]
    self.sso = opts[:sso]
    self.start_paused = opts[:start_paused]
    self.stateless = opts[:stateless]
    self.status = opts[:status]
    self.storage_domain = opts[:storage_domain]
    self.tags = opts[:tags]
    self.time_zone = opts[:time_zone]
    self.tunnel_migration = opts[:tunnel_migration]
    self.type = opts[:type]
    self.usb = opts[:usb]
    self.version = opts[:version]
    self.virtio_scsi = opts[:virtio_scsi]
    self.vm = opts[:vm]
    self.watchdogs = opts[:watchdogs]
  end
  
end

class OpenStackImageProvider < OpenStackProvider
  
  # 
  # Returns the value of the `authentication_url` attribute.
  # 
  def authentication_url
    return @authentication_url
  end
  
  # 
  # Sets the value of the `authentication_url` attribute.
  # 
  def authentication_url=(value)
    @authentication_url = value
  end
  
  # 
  # Returns the value of the `certificates` attribute.
  # 
  def certificates
    return @certificates
  end
  
  # 
  # Sets the value of the `certificates` attribute.
  # 
  def certificates=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Certificate.new(value)
        end
      end
    end
    @certificates = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `images` attribute.
  # 
  def images
    return @images
  end
  
  # 
  # Sets the value of the `images` attribute.
  # 
  def images=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = OpenStackImage.new(value)
        end
      end
    end
    @images = list
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `requires_authentication` attribute.
  # 
  def requires_authentication
    return @requires_authentication
  end
  
  # 
  # Sets the value of the `requires_authentication` attribute.
  # 
  def requires_authentication=(value)
    @requires_authentication = value
  end
  
  # 
  # Returns the value of the `tenant_name` attribute.
  # 
  def tenant_name
    return @tenant_name
  end
  
  # 
  # Sets the value of the `tenant_name` attribute.
  # 
  def tenant_name=(value)
    @tenant_name = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {OpenStackImageProvider} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts :authentication_url The value of attribute `authentication_url`.
  # 
  # @option opts [Array<Hash>] :certificates The values of attribute `certificates`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Array<Hash>] :images The values of attribute `images`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :requires_authentication The value of attribute `requires_authentication`.
  # 
  # @option opts :tenant_name The value of attribute `tenant_name`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.authentication_url = opts[:authentication_url]
    self.certificates = opts[:certificates]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.images = opts[:images]
    self.name = opts[:name]
    self.password = opts[:password]
    self.properties = opts[:properties]
    self.requires_authentication = opts[:requires_authentication]
    self.tenant_name = opts[:tenant_name]
    self.url = opts[:url]
    self.username = opts[:username]
  end
  
end

class OpenStackNetworkProvider < OpenStackProvider
  
  # 
  # Returns the value of the `agent_configuration` attribute.
  # 
  def agent_configuration
    return @agent_configuration
  end
  
  # 
  # Sets the value of the `agent_configuration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::AgentConfiguration} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def agent_configuration=(value)
    if value.is_a?(Hash)
      value = AgentConfiguration.new(value)
    end
    @agent_configuration = value
  end
  
  # 
  # Returns the value of the `authentication_url` attribute.
  # 
  def authentication_url
    return @authentication_url
  end
  
  # 
  # Sets the value of the `authentication_url` attribute.
  # 
  def authentication_url=(value)
    @authentication_url = value
  end
  
  # 
  # Returns the value of the `certificates` attribute.
  # 
  def certificates
    return @certificates
  end
  
  # 
  # Sets the value of the `certificates` attribute.
  # 
  def certificates=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Certificate.new(value)
        end
      end
    end
    @certificates = list
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `networks` attribute.
  # 
  def networks
    return @networks
  end
  
  # 
  # Sets the value of the `networks` attribute.
  # 
  def networks=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = OpenStackNetwork.new(value)
        end
      end
    end
    @networks = list
  end
  
  # 
  # Returns the value of the `password` attribute.
  # 
  def password
    return @password
  end
  
  # 
  # Sets the value of the `password` attribute.
  # 
  def password=(value)
    @password = value
  end
  
  # 
  # Returns the value of the `plugin_type` attribute.
  # 
  def plugin_type
    return @plugin_type
  end
  
  # 
  # Sets the value of the `plugin_type` attribute.
  # 
  def plugin_type=(value)
    @plugin_type = value
  end
  
  # 
  # Returns the value of the `properties` attribute.
  # 
  def properties
    return @properties
  end
  
  # 
  # Sets the value of the `properties` attribute.
  # 
  def properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Property.new(value)
        end
      end
    end
    @properties = list
  end
  
  # 
  # Returns the value of the `read_only` attribute.
  # 
  def read_only
    return @read_only
  end
  
  # 
  # Sets the value of the `read_only` attribute.
  # 
  def read_only=(value)
    @read_only = value
  end
  
  # 
  # Returns the value of the `requires_authentication` attribute.
  # 
  def requires_authentication
    return @requires_authentication
  end
  
  # 
  # Sets the value of the `requires_authentication` attribute.
  # 
  def requires_authentication=(value)
    @requires_authentication = value
  end
  
  # 
  # Returns the value of the `subnets` attribute.
  # 
  def subnets
    return @subnets
  end
  
  # 
  # Sets the value of the `subnets` attribute.
  # 
  def subnets=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = OpenStackSubnet.new(value)
        end
      end
    end
    @subnets = list
  end
  
  # 
  # Returns the value of the `tenant_name` attribute.
  # 
  def tenant_name
    return @tenant_name
  end
  
  # 
  # Sets the value of the `tenant_name` attribute.
  # 
  def tenant_name=(value)
    @tenant_name = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `url` attribute.
  # 
  def url
    return @url
  end
  
  # 
  # Sets the value of the `url` attribute.
  # 
  def url=(value)
    @url = value
  end
  
  # 
  # Returns the value of the `username` attribute.
  # 
  def username
    return @username
  end
  
  # 
  # Sets the value of the `username` attribute.
  # 
  def username=(value)
    @username = value
  end
  
  # 
  # Creates a new instance of the {OpenStackNetworkProvider} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Hash] :agent_configuration The value of attribute `agent_configuration`.
  # 
  # @option opts :authentication_url The value of attribute `authentication_url`.
  # 
  # @option opts [Array<Hash>] :certificates The values of attribute `certificates`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts [Array<Hash>] :networks The values of attribute `networks`.
  # 
  # @option opts :password The value of attribute `password`.
  # 
  # @option opts :plugin_type The value of attribute `plugin_type`.
  # 
  # @option opts [Array<Hash>] :properties The values of attribute `properties`.
  # 
  # @option opts :read_only The value of attribute `read_only`.
  # 
  # @option opts :requires_authentication The value of attribute `requires_authentication`.
  # 
  # @option opts [Array<Hash>] :subnets The values of attribute `subnets`.
  # 
  # @option opts :tenant_name The value of attribute `tenant_name`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts :url The value of attribute `url`.
  # 
  # @option opts :username The value of attribute `username`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.agent_configuration = opts[:agent_configuration]
    self.authentication_url = opts[:authentication_url]
    self.certificates = opts[:certificates]
    self.comment = opts[:comment]
    self.description = opts[:description]
    self.id = opts[:id]
    self.name = opts[:name]
    self.networks = opts[:networks]
    self.password = opts[:password]
    self.plugin_type = opts[:plugin_type]
    self.properties = opts[:properties]
    self.read_only = opts[:read_only]
    self.requires_authentication = opts[:requires_authentication]
    self.subnets = opts[:subnets]
    self.tenant_name = opts[:tenant_name]
    self.type = opts[:type]
    self.url = opts[:url]
    self.username = opts[:username]
  end
  
end

class Snapshot < Vm
  
  # 
  # Returns the value of the `affinity_labels` attribute.
  # 
  def affinity_labels
    return @affinity_labels
  end
  
  # 
  # Sets the value of the `affinity_labels` attribute.
  # 
  def affinity_labels=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = AffinityLabel.new(value)
        end
      end
    end
    @affinity_labels = list
  end
  
  # 
  # Returns the value of the `applications` attribute.
  # 
  def applications
    return @applications
  end
  
  # 
  # Sets the value of the `applications` attribute.
  # 
  def applications=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Application.new(value)
        end
      end
    end
    @applications = list
  end
  
  # 
  # Returns the value of the `bios` attribute.
  # 
  def bios
    return @bios
  end
  
  # 
  # Sets the value of the `bios` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Bios} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def bios=(value)
    if value.is_a?(Hash)
      value = Bios.new(value)
    end
    @bios = value
  end
  
  # 
  # Returns the value of the `cdroms` attribute.
  # 
  def cdroms
    return @cdroms
  end
  
  # 
  # Sets the value of the `cdroms` attribute.
  # 
  def cdroms=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Cdrom.new(value)
        end
      end
    end
    @cdroms = list
  end
  
  # 
  # Returns the value of the `cluster` attribute.
  # 
  def cluster
    return @cluster
  end
  
  # 
  # Sets the value of the `cluster` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cluster} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cluster=(value)
    if value.is_a?(Hash)
      value = Cluster.new(value)
    end
    @cluster = value
  end
  
  # 
  # Returns the value of the `comment` attribute.
  # 
  def comment
    return @comment
  end
  
  # 
  # Sets the value of the `comment` attribute.
  # 
  def comment=(value)
    @comment = value
  end
  
  # 
  # Returns the value of the `console` attribute.
  # 
  def console
    return @console
  end
  
  # 
  # Sets the value of the `console` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Console} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def console=(value)
    if value.is_a?(Hash)
      value = Console.new(value)
    end
    @console = value
  end
  
  # 
  # Returns the value of the `cpu` attribute.
  # 
  def cpu
    return @cpu
  end
  
  # 
  # Sets the value of the `cpu` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Cpu} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu=(value)
    if value.is_a?(Hash)
      value = Cpu.new(value)
    end
    @cpu = value
  end
  
  # 
  # Returns the value of the `cpu_profile` attribute.
  # 
  def cpu_profile
    return @cpu_profile
  end
  
  # 
  # Sets the value of the `cpu_profile` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::CpuProfile} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def cpu_profile=(value)
    if value.is_a?(Hash)
      value = CpuProfile.new(value)
    end
    @cpu_profile = value
  end
  
  # 
  # Returns the value of the `cpu_shares` attribute.
  # 
  def cpu_shares
    return @cpu_shares
  end
  
  # 
  # Sets the value of the `cpu_shares` attribute.
  # 
  def cpu_shares=(value)
    @cpu_shares = value
  end
  
  # 
  # Returns the value of the `creation_time` attribute.
  # 
  def creation_time
    return @creation_time
  end
  
  # 
  # Sets the value of the `creation_time` attribute.
  # 
  def creation_time=(value)
    @creation_time = value
  end
  
  # 
  # Returns the value of the `custom_compatibility_version` attribute.
  # 
  def custom_compatibility_version
    return @custom_compatibility_version
  end
  
  # 
  # Sets the value of the `custom_compatibility_version` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Version} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def custom_compatibility_version=(value)
    if value.is_a?(Hash)
      value = Version.new(value)
    end
    @custom_compatibility_version = value
  end
  
  # 
  # Returns the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model
    return @custom_cpu_model
  end
  
  # 
  # Sets the value of the `custom_cpu_model` attribute.
  # 
  def custom_cpu_model=(value)
    @custom_cpu_model = value
  end
  
  # 
  # Returns the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine
    return @custom_emulated_machine
  end
  
  # 
  # Sets the value of the `custom_emulated_machine` attribute.
  # 
  def custom_emulated_machine=(value)
    @custom_emulated_machine = value
  end
  
  # 
  # Returns the value of the `custom_properties` attribute.
  # 
  def custom_properties
    return @custom_properties
  end
  
  # 
  # Sets the value of the `custom_properties` attribute.
  # 
  def custom_properties=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = CustomProperty.new(value)
        end
      end
    end
    @custom_properties = list
  end
  
  # 
  # Returns the value of the `date` attribute.
  # 
  def date
    return @date
  end
  
  # 
  # Sets the value of the `date` attribute.
  # 
  def date=(value)
    @date = value
  end
  
  # 
  # Returns the value of the `delete_protected` attribute.
  # 
  def delete_protected
    return @delete_protected
  end
  
  # 
  # Sets the value of the `delete_protected` attribute.
  # 
  def delete_protected=(value)
    @delete_protected = value
  end
  
  # 
  # Returns the value of the `description` attribute.
  # 
  def description
    return @description
  end
  
  # 
  # Sets the value of the `description` attribute.
  # 
  def description=(value)
    @description = value
  end
  
  # 
  # Returns the value of the `disk_attachments` attribute.
  # 
  def disk_attachments
    return @disk_attachments
  end
  
  # 
  # Sets the value of the `disk_attachments` attribute.
  # 
  def disk_attachments=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = DiskAttachment.new(value)
        end
      end
    end
    @disk_attachments = list
  end
  
  # 
  # Returns the value of the `display` attribute.
  # 
  def display
    return @display
  end
  
  # 
  # Sets the value of the `display` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Display} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def display=(value)
    if value.is_a?(Hash)
      value = Display.new(value)
    end
    @display = value
  end
  
  # 
  # Returns the value of the `domain` attribute.
  # 
  def domain
    return @domain
  end
  
  # 
  # Sets the value of the `domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Domain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def domain=(value)
    if value.is_a?(Hash)
      value = Domain.new(value)
    end
    @domain = value
  end
  
  # 
  # Returns the value of the `external_host_provider` attribute.
  # 
  def external_host_provider
    return @external_host_provider
  end
  
  # 
  # Sets the value of the `external_host_provider` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::ExternalHostProvider} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def external_host_provider=(value)
    if value.is_a?(Hash)
      value = ExternalHostProvider.new(value)
    end
    @external_host_provider = value
  end
  
  # 
  # Returns the value of the `floppies` attribute.
  # 
  def floppies
    return @floppies
  end
  
  # 
  # Sets the value of the `floppies` attribute.
  # 
  def floppies=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Floppy.new(value)
        end
      end
    end
    @floppies = list
  end
  
  # 
  # Returns the value of the `fqdn` attribute.
  # 
  def fqdn
    return @fqdn
  end
  
  # 
  # Sets the value of the `fqdn` attribute.
  # 
  def fqdn=(value)
    @fqdn = value
  end
  
  # 
  # Returns the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles
    return @graphics_consoles
  end
  
  # 
  # Sets the value of the `graphics_consoles` attribute.
  # 
  def graphics_consoles=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = GraphicsConsole.new(value)
        end
      end
    end
    @graphics_consoles = list
  end
  
  # 
  # Returns the value of the `guest_operating_system` attribute.
  # 
  def guest_operating_system
    return @guest_operating_system
  end
  
  # 
  # Sets the value of the `guest_operating_system` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::GuestOperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def guest_operating_system=(value)
    if value.is_a?(Hash)
      value = GuestOperatingSystem.new(value)
    end
    @guest_operating_system = value
  end
  
  # 
  # Returns the value of the `guest_time_zone` attribute.
  # 
  def guest_time_zone
    return @guest_time_zone
  end
  
  # 
  # Sets the value of the `guest_time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def guest_time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @guest_time_zone = value
  end
  
  # 
  # Returns the value of the `high_availability` attribute.
  # 
  def high_availability
    return @high_availability
  end
  
  # 
  # Sets the value of the `high_availability` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::HighAvailability} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def high_availability=(value)
    if value.is_a?(Hash)
      value = HighAvailability.new(value)
    end
    @high_availability = value
  end
  
  # 
  # Returns the value of the `host` attribute.
  # 
  def host
    return @host
  end
  
  # 
  # Sets the value of the `host` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Host} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def host=(value)
    if value.is_a?(Hash)
      value = Host.new(value)
    end
    @host = value
  end
  
  # 
  # Returns the value of the `host_devices` attribute.
  # 
  def host_devices
    return @host_devices
  end
  
  # 
  # Sets the value of the `host_devices` attribute.
  # 
  def host_devices=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = HostDevice.new(value)
        end
      end
    end
    @host_devices = list
  end
  
  # 
  # Returns the value of the `id` attribute.
  # 
  def id
    return @id
  end
  
  # 
  # Sets the value of the `id` attribute.
  # 
  def id=(value)
    @id = value
  end
  
  # 
  # Returns the value of the `initialization` attribute.
  # 
  def initialization
    return @initialization
  end
  
  # 
  # Sets the value of the `initialization` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Initialization} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def initialization=(value)
    if value.is_a?(Hash)
      value = Initialization.new(value)
    end
    @initialization = value
  end
  
  # 
  # Returns the value of the `instance_type` attribute.
  # 
  def instance_type
    return @instance_type
  end
  
  # 
  # Sets the value of the `instance_type` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::InstanceType} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def instance_type=(value)
    if value.is_a?(Hash)
      value = InstanceType.new(value)
    end
    @instance_type = value
  end
  
  # 
  # Returns the value of the `io` attribute.
  # 
  def io
    return @io
  end
  
  # 
  # Sets the value of the `io` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Io} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def io=(value)
    if value.is_a?(Hash)
      value = Io.new(value)
    end
    @io = value
  end
  
  # 
  # Returns the value of the `katello_errata` attribute.
  # 
  def katello_errata
    return @katello_errata
  end
  
  # 
  # Sets the value of the `katello_errata` attribute.
  # 
  def katello_errata=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = KatelloErratum.new(value)
        end
      end
    end
    @katello_errata = list
  end
  
  # 
  # Returns the value of the `large_icon` attribute.
  # 
  def large_icon
    return @large_icon
  end
  
  # 
  # Sets the value of the `large_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def large_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @large_icon = value
  end
  
  # 
  # Returns the value of the `memory` attribute.
  # 
  def memory
    return @memory
  end
  
  # 
  # Sets the value of the `memory` attribute.
  # 
  def memory=(value)
    @memory = value
  end
  
  # 
  # Returns the value of the `memory_policy` attribute.
  # 
  def memory_policy
    return @memory_policy
  end
  
  # 
  # Sets the value of the `memory_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MemoryPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def memory_policy=(value)
    if value.is_a?(Hash)
      value = MemoryPolicy.new(value)
    end
    @memory_policy = value
  end
  
  # 
  # Returns the value of the `migration` attribute.
  # 
  def migration
    return @migration
  end
  
  # 
  # Sets the value of the `migration` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::MigrationOptions} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def migration=(value)
    if value.is_a?(Hash)
      value = MigrationOptions.new(value)
    end
    @migration = value
  end
  
  # 
  # Returns the value of the `migration_downtime` attribute.
  # 
  def migration_downtime
    return @migration_downtime
  end
  
  # 
  # Sets the value of the `migration_downtime` attribute.
  # 
  def migration_downtime=(value)
    @migration_downtime = value
  end
  
  # 
  # Returns the value of the `name` attribute.
  # 
  def name
    return @name
  end
  
  # 
  # Sets the value of the `name` attribute.
  # 
  def name=(value)
    @name = value
  end
  
  # 
  # Returns the value of the `next_run_configuration_exists` attribute.
  # 
  def next_run_configuration_exists
    return @next_run_configuration_exists
  end
  
  # 
  # Sets the value of the `next_run_configuration_exists` attribute.
  # 
  def next_run_configuration_exists=(value)
    @next_run_configuration_exists = value
  end
  
  # 
  # Returns the value of the `nics` attribute.
  # 
  def nics
    return @nics
  end
  
  # 
  # Sets the value of the `nics` attribute.
  # 
  def nics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Nic.new(value)
        end
      end
    end
    @nics = list
  end
  
  # 
  # Returns the value of the `numa_nodes` attribute.
  # 
  def numa_nodes
    return @numa_nodes
  end
  
  # 
  # Sets the value of the `numa_nodes` attribute.
  # 
  def numa_nodes=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = NumaNode.new(value)
        end
      end
    end
    @numa_nodes = list
  end
  
  # 
  # Returns the value of the `numa_tune_mode` attribute.
  # 
  def numa_tune_mode
    return @numa_tune_mode
  end
  
  # 
  # Sets the value of the `numa_tune_mode` attribute.
  # 
  def numa_tune_mode=(value)
    @numa_tune_mode = value
  end
  
  # 
  # Returns the value of the `origin` attribute.
  # 
  def origin
    return @origin
  end
  
  # 
  # Sets the value of the `origin` attribute.
  # 
  def origin=(value)
    @origin = value
  end
  
  # 
  # Returns the value of the `os` attribute.
  # 
  def os
    return @os
  end
  
  # 
  # Sets the value of the `os` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::OperatingSystem} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def os=(value)
    if value.is_a?(Hash)
      value = OperatingSystem.new(value)
    end
    @os = value
  end
  
  # 
  # Returns the value of the `payloads` attribute.
  # 
  def payloads
    return @payloads
  end
  
  # 
  # Sets the value of the `payloads` attribute.
  # 
  def payloads=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Payload.new(value)
        end
      end
    end
    @payloads = list
  end
  
  # 
  # Returns the value of the `permissions` attribute.
  # 
  def permissions
    return @permissions
  end
  
  # 
  # Sets the value of the `permissions` attribute.
  # 
  def permissions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Permission.new(value)
        end
      end
    end
    @permissions = list
  end
  
  # 
  # Returns the value of the `persist_memorystate` attribute.
  # 
  def persist_memorystate
    return @persist_memorystate
  end
  
  # 
  # Sets the value of the `persist_memorystate` attribute.
  # 
  def persist_memorystate=(value)
    @persist_memorystate = value
  end
  
  # 
  # Returns the value of the `placement_policy` attribute.
  # 
  def placement_policy
    return @placement_policy
  end
  
  # 
  # Sets the value of the `placement_policy` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VmPlacementPolicy} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def placement_policy=(value)
    if value.is_a?(Hash)
      value = VmPlacementPolicy.new(value)
    end
    @placement_policy = value
  end
  
  # 
  # Returns the value of the `quota` attribute.
  # 
  def quota
    return @quota
  end
  
  # 
  # Sets the value of the `quota` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Quota} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def quota=(value)
    if value.is_a?(Hash)
      value = Quota.new(value)
    end
    @quota = value
  end
  
  # 
  # Returns the value of the `reported_devices` attribute.
  # 
  def reported_devices
    return @reported_devices
  end
  
  # 
  # Sets the value of the `reported_devices` attribute.
  # 
  def reported_devices=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = ReportedDevice.new(value)
        end
      end
    end
    @reported_devices = list
  end
  
  # 
  # Returns the value of the `rng_device` attribute.
  # 
  def rng_device
    return @rng_device
  end
  
  # 
  # Sets the value of the `rng_device` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::RngDevice} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def rng_device=(value)
    if value.is_a?(Hash)
      value = RngDevice.new(value)
    end
    @rng_device = value
  end
  
  # 
  # Returns the value of the `run_once` attribute.
  # 
  def run_once
    return @run_once
  end
  
  # 
  # Sets the value of the `run_once` attribute.
  # 
  def run_once=(value)
    @run_once = value
  end
  
  # 
  # Returns the value of the `serial_number` attribute.
  # 
  def serial_number
    return @serial_number
  end
  
  # 
  # Sets the value of the `serial_number` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::SerialNumber} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def serial_number=(value)
    if value.is_a?(Hash)
      value = SerialNumber.new(value)
    end
    @serial_number = value
  end
  
  # 
  # Returns the value of the `sessions` attribute.
  # 
  def sessions
    return @sessions
  end
  
  # 
  # Sets the value of the `sessions` attribute.
  # 
  def sessions=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Session.new(value)
        end
      end
    end
    @sessions = list
  end
  
  # 
  # Returns the value of the `small_icon` attribute.
  # 
  def small_icon
    return @small_icon
  end
  
  # 
  # Sets the value of the `small_icon` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Icon} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def small_icon=(value)
    if value.is_a?(Hash)
      value = Icon.new(value)
    end
    @small_icon = value
  end
  
  # 
  # Returns the value of the `snapshot_status` attribute.
  # 
  def snapshot_status
    return @snapshot_status
  end
  
  # 
  # Sets the value of the `snapshot_status` attribute.
  # 
  def snapshot_status=(value)
    @snapshot_status = value
  end
  
  # 
  # Returns the value of the `snapshot_type` attribute.
  # 
  def snapshot_type
    return @snapshot_type
  end
  
  # 
  # Sets the value of the `snapshot_type` attribute.
  # 
  def snapshot_type=(value)
    @snapshot_type = value
  end
  
  # 
  # Returns the value of the `snapshots` attribute.
  # 
  def snapshots
    return @snapshots
  end
  
  # 
  # Sets the value of the `snapshots` attribute.
  # 
  def snapshots=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Snapshot.new(value)
        end
      end
    end
    @snapshots = list
  end
  
  # 
  # Returns the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled
    return @soundcard_enabled
  end
  
  # 
  # Sets the value of the `soundcard_enabled` attribute.
  # 
  def soundcard_enabled=(value)
    @soundcard_enabled = value
  end
  
  # 
  # Returns the value of the `sso` attribute.
  # 
  def sso
    return @sso
  end
  
  # 
  # Sets the value of the `sso` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Sso} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def sso=(value)
    if value.is_a?(Hash)
      value = Sso.new(value)
    end
    @sso = value
  end
  
  # 
  # Returns the value of the `start_paused` attribute.
  # 
  def start_paused
    return @start_paused
  end
  
  # 
  # Sets the value of the `start_paused` attribute.
  # 
  def start_paused=(value)
    @start_paused = value
  end
  
  # 
  # Returns the value of the `start_time` attribute.
  # 
  def start_time
    return @start_time
  end
  
  # 
  # Sets the value of the `start_time` attribute.
  # 
  def start_time=(value)
    @start_time = value
  end
  
  # 
  # Returns the value of the `stateless` attribute.
  # 
  def stateless
    return @stateless
  end
  
  # 
  # Sets the value of the `stateless` attribute.
  # 
  def stateless=(value)
    @stateless = value
  end
  
  # 
  # Returns the value of the `statistics` attribute.
  # 
  def statistics
    return @statistics
  end
  
  # 
  # Sets the value of the `statistics` attribute.
  # 
  def statistics=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Statistic.new(value)
        end
      end
    end
    @statistics = list
  end
  
  # 
  # Returns the value of the `status` attribute.
  # 
  def status
    return @status
  end
  
  # 
  # Sets the value of the `status` attribute.
  # 
  def status=(value)
    @status = value
  end
  
  # 
  # Returns the value of the `status_detail` attribute.
  # 
  def status_detail
    return @status_detail
  end
  
  # 
  # Sets the value of the `status_detail` attribute.
  # 
  def status_detail=(value)
    @status_detail = value
  end
  
  # 
  # Returns the value of the `stop_reason` attribute.
  # 
  def stop_reason
    return @stop_reason
  end
  
  # 
  # Sets the value of the `stop_reason` attribute.
  # 
  def stop_reason=(value)
    @stop_reason = value
  end
  
  # 
  # Returns the value of the `stop_time` attribute.
  # 
  def stop_time
    return @stop_time
  end
  
  # 
  # Sets the value of the `stop_time` attribute.
  # 
  def stop_time=(value)
    @stop_time = value
  end
  
  # 
  # Returns the value of the `storage_domain` attribute.
  # 
  def storage_domain
    return @storage_domain
  end
  
  # 
  # Sets the value of the `storage_domain` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::StorageDomain} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def storage_domain=(value)
    if value.is_a?(Hash)
      value = StorageDomain.new(value)
    end
    @storage_domain = value
  end
  
  # 
  # Returns the value of the `tags` attribute.
  # 
  def tags
    return @tags
  end
  
  # 
  # Sets the value of the `tags` attribute.
  # 
  def tags=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Tag.new(value)
        end
      end
    end
    @tags = list
  end
  
  # 
  # Returns the value of the `template` attribute.
  # 
  def template
    return @template
  end
  
  # 
  # Sets the value of the `template` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Template} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def template=(value)
    if value.is_a?(Hash)
      value = Template.new(value)
    end
    @template = value
  end
  
  # 
  # Returns the value of the `time_zone` attribute.
  # 
  def time_zone
    return @time_zone
  end
  
  # 
  # Sets the value of the `time_zone` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::TimeZone} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def time_zone=(value)
    if value.is_a?(Hash)
      value = TimeZone.new(value)
    end
    @time_zone = value
  end
  
  # 
  # Returns the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration
    return @tunnel_migration
  end
  
  # 
  # Sets the value of the `tunnel_migration` attribute.
  # 
  def tunnel_migration=(value)
    @tunnel_migration = value
  end
  
  # 
  # Returns the value of the `type` attribute.
  # 
  def type
    return @type
  end
  
  # 
  # Sets the value of the `type` attribute.
  # 
  def type=(value)
    @type = value
  end
  
  # 
  # Returns the value of the `usb` attribute.
  # 
  def usb
    return @usb
  end
  
  # 
  # Sets the value of the `usb` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Usb} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def usb=(value)
    if value.is_a?(Hash)
      value = Usb.new(value)
    end
    @usb = value
  end
  
  # 
  # Returns the value of the `use_latest_template_version` attribute.
  # 
  def use_latest_template_version
    return @use_latest_template_version
  end
  
  # 
  # Sets the value of the `use_latest_template_version` attribute.
  # 
  def use_latest_template_version=(value)
    @use_latest_template_version = value
  end
  
  # 
  # Returns the value of the `virtio_scsi` attribute.
  # 
  def virtio_scsi
    return @virtio_scsi
  end
  
  # 
  # Sets the value of the `virtio_scsi` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VirtioScsi} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def virtio_scsi=(value)
    if value.is_a?(Hash)
      value = VirtioScsi.new(value)
    end
    @virtio_scsi = value
  end
  
  # 
  # Returns the value of the `vm` attribute.
  # 
  def vm
    return @vm
  end
  
  # 
  # Sets the value of the `vm` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::Vm} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm=(value)
    if value.is_a?(Hash)
      value = Vm.new(value)
    end
    @vm = value
  end
  
  # 
  # Returns the value of the `vm_pool` attribute.
  # 
  def vm_pool
    return @vm_pool
  end
  
  # 
  # Sets the value of the `vm_pool` attribute.
  # 
  # @param value [Hash]
  # 
  # The `value` parameter can be an instance of {ovirtsdk4::VmPool} or a hash.
  # If it is a hash then a new instance will be created passing the hash as the 
  # `opts` parameter to the constructor.
  # 
  def vm_pool=(value)
    if value.is_a?(Hash)
      value = VmPool.new(value)
    end
    @vm_pool = value
  end
  
  # 
  # Returns the value of the `watchdogs` attribute.
  # 
  def watchdogs
    return @watchdogs
  end
  
  # 
  # Sets the value of the `watchdogs` attribute.
  # 
  def watchdogs=(list)
    if list.class == Array
      list = List.new(list)
      list.each_with_index do |value, index|
        if value.is_a?(Hash)
          list[index] = Watchdog.new(value)
        end
      end
    end
    @watchdogs = list
  end
  
  # 
  # Creates a new instance of the {Snapshot} class.
  # 
  # @param opts [Hash] A hash containing the attributes of the object. The keys of the hash 
  #   should be symbols corresponding to the names of the attributes. The values of the hash 
  #   should be the values of the attributes.
  # 
  # @option opts [Array<Hash>] :affinity_labels The values of attribute `affinity_labels`.
  # 
  # @option opts [Array<Hash>] :applications The values of attribute `applications`.
  # 
  # @option opts [Hash] :bios The value of attribute `bios`.
  # 
  # @option opts [Array<Hash>] :cdroms The values of attribute `cdroms`.
  # 
  # @option opts [Hash] :cluster The value of attribute `cluster`.
  # 
  # @option opts :comment The value of attribute `comment`.
  # 
  # @option opts [Hash] :console The value of attribute `console`.
  # 
  # @option opts [Hash] :cpu The value of attribute `cpu`.
  # 
  # @option opts [Hash] :cpu_profile The value of attribute `cpu_profile`.
  # 
  # @option opts :cpu_shares The value of attribute `cpu_shares`.
  # 
  # @option opts :creation_time The value of attribute `creation_time`.
  # 
  # @option opts [Hash] :custom_compatibility_version The value of attribute `custom_compatibility_version`.
  # 
  # @option opts :custom_cpu_model The value of attribute `custom_cpu_model`.
  # 
  # @option opts :custom_emulated_machine The value of attribute `custom_emulated_machine`.
  # 
  # @option opts [Array<Hash>] :custom_properties The values of attribute `custom_properties`.
  # 
  # @option opts :date The value of attribute `date`.
  # 
  # @option opts :delete_protected The value of attribute `delete_protected`.
  # 
  # @option opts :description The value of attribute `description`.
  # 
  # @option opts [Array<Hash>] :disk_attachments The values of attribute `disk_attachments`.
  # 
  # @option opts [Hash] :display The value of attribute `display`.
  # 
  # @option opts [Hash] :domain The value of attribute `domain`.
  # 
  # @option opts [Hash] :external_host_provider The value of attribute `external_host_provider`.
  # 
  # @option opts [Array<Hash>] :floppies The values of attribute `floppies`.
  # 
  # @option opts :fqdn The value of attribute `fqdn`.
  # 
  # @option opts [Array<Hash>] :graphics_consoles The values of attribute `graphics_consoles`.
  # 
  # @option opts [Hash] :guest_operating_system The value of attribute `guest_operating_system`.
  # 
  # @option opts [Hash] :guest_time_zone The value of attribute `guest_time_zone`.
  # 
  # @option opts [Hash] :high_availability The value of attribute `high_availability`.
  # 
  # @option opts [Hash] :host The value of attribute `host`.
  # 
  # @option opts [Array<Hash>] :host_devices The values of attribute `host_devices`.
  # 
  # @option opts :id The value of attribute `id`.
  # 
  # @option opts [Hash] :initialization The value of attribute `initialization`.
  # 
  # @option opts [Hash] :instance_type The value of attribute `instance_type`.
  # 
  # @option opts [Hash] :io The value of attribute `io`.
  # 
  # @option opts [Array<Hash>] :katello_errata The values of attribute `katello_errata`.
  # 
  # @option opts [Hash] :large_icon The value of attribute `large_icon`.
  # 
  # @option opts :memory The value of attribute `memory`.
  # 
  # @option opts [Hash] :memory_policy The value of attribute `memory_policy`.
  # 
  # @option opts [Hash] :migration The value of attribute `migration`.
  # 
  # @option opts :migration_downtime The value of attribute `migration_downtime`.
  # 
  # @option opts :name The value of attribute `name`.
  # 
  # @option opts :next_run_configuration_exists The value of attribute `next_run_configuration_exists`.
  # 
  # @option opts [Array<Hash>] :nics The values of attribute `nics`.
  # 
  # @option opts [Array<Hash>] :numa_nodes The values of attribute `numa_nodes`.
  # 
  # @option opts :numa_tune_mode The value of attribute `numa_tune_mode`.
  # 
  # @option opts :origin The value of attribute `origin`.
  # 
  # @option opts [Hash] :os The value of attribute `os`.
  # 
  # @option opts [Array<Hash>] :payloads The values of attribute `payloads`.
  # 
  # @option opts [Array<Hash>] :permissions The values of attribute `permissions`.
  # 
  # @option opts :persist_memorystate The value of attribute `persist_memorystate`.
  # 
  # @option opts [Hash] :placement_policy The value of attribute `placement_policy`.
  # 
  # @option opts [Hash] :quota The value of attribute `quota`.
  # 
  # @option opts [Array<Hash>] :reported_devices The values of attribute `reported_devices`.
  # 
  # @option opts [Hash] :rng_device The value of attribute `rng_device`.
  # 
  # @option opts :run_once The value of attribute `run_once`.
  # 
  # @option opts [Hash] :serial_number The value of attribute `serial_number`.
  # 
  # @option opts [Array<Hash>] :sessions The values of attribute `sessions`.
  # 
  # @option opts [Hash] :small_icon The value of attribute `small_icon`.
  # 
  # @option opts :snapshot_status The value of attribute `snapshot_status`.
  # 
  # @option opts :snapshot_type The value of attribute `snapshot_type`.
  # 
  # @option opts [Array<Hash>] :snapshots The values of attribute `snapshots`.
  # 
  # @option opts :soundcard_enabled The value of attribute `soundcard_enabled`.
  # 
  # @option opts [Hash] :sso The value of attribute `sso`.
  # 
  # @option opts :start_paused The value of attribute `start_paused`.
  # 
  # @option opts :start_time The value of attribute `start_time`.
  # 
  # @option opts :stateless The value of attribute `stateless`.
  # 
  # @option opts [Array<Hash>] :statistics The values of attribute `statistics`.
  # 
  # @option opts :status The value of attribute `status`.
  # 
  # @option opts :status_detail The value of attribute `status_detail`.
  # 
  # @option opts :stop_reason The value of attribute `stop_reason`.
  # 
  # @option opts :stop_time The value of attribute `stop_time`.
  # 
  # @option opts [Hash] :storage_domain The value of attribute `storage_domain`.
  # 
  # @option opts [Array<Hash>] :tags The values of attribute `tags`.
  # 
  # @option opts [Hash] :template The value of attribute `template`.
  # 
  # @option opts [Hash] :time_zone The value of attribute `time_zone`.
  # 
  # @option opts :tunnel_migration The value of attribute `tunnel_migration`.
  # 
  # @option opts :type The value of attribute `type`.
  # 
  # @option opts [Hash] :usb The value of attribute `usb`.
  # 
  # @option opts :use_latest_template_version The value of attribute `use_latest_template_version`.
  # 
  # @option opts [Hash] :virtio_scsi The value of attribute `virtio_scsi`.
  # 
  # @option opts [Hash] :vm The value of attribute `vm`.
  # 
  # @option opts [Hash] :vm_pool The value of attribute `vm_pool`.
  # 
  # @option opts [Array<Hash>] :watchdogs The values of attribute `watchdogs`.
  # 
  # 
  def initialize(opts = {})
    super(opts)
    self.affinity_labels = opts[:affinity_labels]
    self.applications = opts[:applications]
    self.bios = opts[:bios]
    self.cdroms = opts[:cdroms]
    self.cluster = opts[:cluster]
    self.comment = opts[:comment]
    self.console = opts[:console]
    self.cpu = opts[:cpu]
    self.cpu_profile = opts[:cpu_profile]
    self.cpu_shares = opts[:cpu_shares]
    self.creation_time = opts[:creation_time]
    self.custom_compatibility_version = opts[:custom_compatibility_version]
    self.custom_cpu_model = opts[:custom_cpu_model]
    self.custom_emulated_machine = opts[:custom_emulated_machine]
    self.custom_properties = opts[:custom_properties]
    self.date = opts[:date]
    self.delete_protected = opts[:delete_protected]
    self.description = opts[:description]
    self.disk_attachments = opts[:disk_attachments]
    self.display = opts[:display]
    self.domain = opts[:domain]
    self.external_host_provider = opts[:external_host_provider]
    self.floppies = opts[:floppies]
    self.fqdn = opts[:fqdn]
    self.graphics_consoles = opts[:graphics_consoles]
    self.guest_operating_system = opts[:guest_operating_system]
    self.guest_time_zone = opts[:guest_time_zone]
    self.high_availability = opts[:high_availability]
    self.host = opts[:host]
    self.host_devices = opts[:host_devices]
    self.id = opts[:id]
    self.initialization = opts[:initialization]
    self.instance_type = opts[:instance_type]
    self.io = opts[:io]
    self.katello_errata = opts[:katello_errata]
    self.large_icon = opts[:large_icon]
    self.memory = opts[:memory]
    self.memory_policy = opts[:memory_policy]
    self.migration = opts[:migration]
    self.migration_downtime = opts[:migration_downtime]
    self.name = opts[:name]
    self.next_run_configuration_exists = opts[:next_run_configuration_exists]
    self.nics = opts[:nics]
    self.numa_nodes = opts[:numa_nodes]
    self.numa_tune_mode = opts[:numa_tune_mode]
    self.origin = opts[:origin]
    self.os = opts[:os]
    self.payloads = opts[:payloads]
    self.permissions = opts[:permissions]
    self.persist_memorystate = opts[:persist_memorystate]
    self.placement_policy = opts[:placement_policy]
    self.quota = opts[:quota]
    self.reported_devices = opts[:reported_devices]
    self.rng_device = opts[:rng_device]
    self.run_once = opts[:run_once]
    self.serial_number = opts[:serial_number]
    self.sessions = opts[:sessions]
    self.small_icon = opts[:small_icon]
    self.snapshot_status = opts[:snapshot_status]
    self.snapshot_type = opts[:snapshot_type]
    self.snapshots = opts[:snapshots]
    self.soundcard_enabled = opts[:soundcard_enabled]
    self.sso = opts[:sso]
    self.start_paused = opts[:start_paused]
    self.start_time = opts[:start_time]
    self.stateless = opts[:stateless]
    self.statistics = opts[:statistics]
    self.status = opts[:status]
    self.status_detail = opts[:status_detail]
    self.stop_reason = opts[:stop_reason]
    self.stop_time = opts[:stop_time]
    self.storage_domain = opts[:storage_domain]
    self.tags = opts[:tags]
    self.template = opts[:template]
    self.time_zone = opts[:time_zone]
    self.tunnel_migration = opts[:tunnel_migration]
    self.type = opts[:type]
    self.usb = opts[:usb]
    self.use_latest_template_version = opts[:use_latest_template_version]
    self.virtio_scsi = opts[:virtio_scsi]
    self.vm = opts[:vm]
    self.vm_pool = opts[:vm_pool]
    self.watchdogs = opts[:watchdogs]
  end
  
end

package AccessProtocol
CIFS = 'cifs'.freeze
GLUSTER = 'gluster'.freeze
NFS = 'nfs'.freeze
package Architecture
PPC64 = 'ppc64'.freeze
UNDEFINED = 'undefined'.freeze
X86_64 = 'x86_64'.freeze
package AuthenticationMethod
PASSWORD = 'password'.freeze
PUBLICKEY = 'publickey'.freeze
package AutoNumaStatus
DISABLE = 'disable'.freeze
ENABLE = 'enable'.freeze
UNKNOWN = 'unknown'.freeze
package BootDevice
CDROM = 'cdrom'.freeze
HD = 'hd'.freeze
NETWORK = 'network'.freeze
package BootProtocol
AUTOCONF = 'autoconf'.freeze
DHCP = 'dhcp'.freeze
NONE = 'none'.freeze
STATIC = 'static'.freeze
package ConfigurationType
OVF = 'ovf'.freeze
package CpuMode
CUSTOM = 'custom'.freeze
HOST_MODEL = 'host_model'.freeze
HOST_PASSTHROUGH = 'host_passthrough'.freeze
package CreationStatus
COMPLETE = 'complete'.freeze
FAILED = 'failed'.freeze
IN_PROGRESS = 'in_progress'.freeze
PENDING = 'pending'.freeze
package DataCenterStatus
CONTEND = 'contend'.freeze
MAINTENANCE = 'maintenance'.freeze
NOT_OPERATIONAL = 'not_operational'.freeze
PROBLEMATIC = 'problematic'.freeze
UNINITIALIZED = 'uninitialized'.freeze
UP = 'up'.freeze
package DiskFormat
COW = 'cow'.freeze
RAW = 'raw'.freeze
package DiskInterface
IDE = 'ide'.freeze
SPAPR_VSCSI = 'spapr_vscsi'.freeze
VIRTIO = 'virtio'.freeze
VIRTIO_SCSI = 'virtio_scsi'.freeze
package DiskStatus
ILLEGAL = 'illegal'.freeze
LOCKED = 'locked'.freeze
OK = 'ok'.freeze
package DiskStorageType
CINDER = 'cinder'.freeze
IMAGE = 'image'.freeze
LUN = 'lun'.freeze
package DiskType
DATA = 'data'.freeze
SYSTEM = 'system'.freeze
package DisplayType
SPICE = 'spice'.freeze
VNC = 'vnc'.freeze
package EntityExternalStatus
ERROR = 'error'.freeze
FAILURE = 'failure'.freeze
INFO = 'info'.freeze
OK = 'ok'.freeze
WARNING = 'warning'.freeze
package ExternalStatus
ERROR = 'error'.freeze
FAILURE = 'failure'.freeze
INFO = 'info'.freeze
OK = 'ok'.freeze
WARNING = 'warning'.freeze
package ExternalSystemType
GLUSTER = 'gluster'.freeze
VDSM = 'vdsm'.freeze
package FenceType
MANUAL = 'manual'.freeze
RESTART = 'restart'.freeze
START = 'start'.freeze
STATUS = 'status'.freeze
STOP = 'stop'.freeze
package GlusterBrickStatus
DOWN = 'down'.freeze
UNKNOWN = 'unknown'.freeze
UP = 'up'.freeze
package GlusterHookStatus
DISABLED = 'disabled'.freeze
ENABLED = 'enabled'.freeze
MISSING = 'missing'.freeze
package GlusterState
DOWN = 'down'.freeze
UNKNOWN = 'unknown'.freeze
UP = 'up'.freeze
package GlusterVolumeStatus
DOWN = 'down'.freeze
UNKNOWN = 'unknown'.freeze
UP = 'up'.freeze
package GlusterVolumeType
DISPERSE = 'disperse'.freeze
DISTRIBUTE = 'distribute'.freeze
DISTRIBUTED_DISPERSE = 'distributed_disperse'.freeze
DISTRIBUTED_REPLICATE = 'distributed_replicate'.freeze
DISTRIBUTED_STRIPE = 'distributed_stripe'.freeze
DISTRIBUTED_STRIPED_REPLICATE = 'distributed_striped_replicate'.freeze
REPLICATE = 'replicate'.freeze
STRIPE = 'stripe'.freeze
STRIPED_REPLICATE = 'striped_replicate'.freeze
package GraphicsType
SPICE = 'spice'.freeze
VNC = 'vnc'.freeze
package HookContentType
BINARY = 'binary'.freeze
TEXT = 'text'.freeze
package HookStage
POST = 'post'.freeze
PRE = 'pre'.freeze
package HookStatus
DISABLED = 'disabled'.freeze
ENABLED = 'enabled'.freeze
MISSING = 'missing'.freeze
package HostProtocol
STOMP = 'stomp'.freeze
XML = 'xml'.freeze
package HostStatus
CONNECTING = 'connecting'.freeze
DOWN = 'down'.freeze
ERROR = 'error'.freeze
INITIALIZING = 'initializing'.freeze
INSTALL_FAILED = 'install_failed'.freeze
INSTALLING = 'installing'.freeze
INSTALLING_OS = 'installing_os'.freeze
KDUMPING = 'kdumping'.freeze
MAINTENANCE = 'maintenance'.freeze
NON_OPERATIONAL = 'non_operational'.freeze
NON_RESPONSIVE = 'non_responsive'.freeze
PENDING_APPROVAL = 'pending_approval'.freeze
PREPARING_FOR_MAINTENANCE = 'preparing_for_maintenance'.freeze
REBOOT = 'reboot'.freeze
UNASSIGNED = 'unassigned'.freeze
UP = 'up'.freeze
package HostType
OVIRT_NODE = 'ovirt_node'.freeze
RHEL = 'rhel'.freeze
RHEV_H = 'rhev_h'.freeze
package InheritableBoolean
FALSE = 'false_'.freeze
INHERIT = 'inherit'.freeze
TRUE = 'true_'.freeze
package IpVersion
V4 = 'v4'.freeze
V6 = 'v6'.freeze
package JobStatus
ABORTED = 'aborted'.freeze
FAILED = 'failed'.freeze
FINISHED = 'finished'.freeze
STARTED = 'started'.freeze
UNKNOWN = 'unknown'.freeze
package KdumpStatus
DISABLED = 'disabled'.freeze
ENABLED = 'enabled'.freeze
UNKNOWN = 'unknown'.freeze
package LogSeverity
ALERT = 'alert'.freeze
ERROR = 'error'.freeze
NORMAL = 'normal'.freeze
WARNING = 'warning'.freeze
package LunStatus
FREE = 'free'.freeze
UNUSABLE = 'unusable'.freeze
USED = 'used'.freeze
package MessageBrokerType
QPID = 'qpid'.freeze
RABBIT_MQ = 'rabbit_mq'.freeze
package MigrateOnError
DO_NOT_MIGRATE = 'do_not_migrate'.freeze
MIGRATE = 'migrate'.freeze
MIGRATE_HIGHLY_AVAILABLE = 'migrate_highly_available'.freeze
package MigrationBandwidthAssignmentMethod
AUTO = 'auto'.freeze
CUSTOM = 'custom'.freeze
HYPERVISOR_DEFAULT = 'hypervisor_default'.freeze
package NetworkPluginType
OPEN_VSWITCH = 'open_vswitch'.freeze
package NetworkStatus
NON_OPERATIONAL = 'non_operational'.freeze
OPERATIONAL = 'operational'.freeze
package NetworkUsage
DISPLAY = 'display'.freeze
MANAGEMENT = 'management'.freeze
MIGRATION = 'migration'.freeze
VM = 'vm'.freeze
package NfsVersion
AUTO = 'auto'.freeze
V3 = 'v3'.freeze
V4 = 'v4'.freeze
V4_1 = 'v4_1'.freeze
package NicInterface
E1000 = 'e1000'.freeze
PCI_PASSTHROUGH = 'pci_passthrough'.freeze
RTL8139 = 'rtl8139'.freeze
RTL8139_VIRTIO = 'rtl8139_virtio'.freeze
SPAPR_VLAN = 'spapr_vlan'.freeze
VIRTIO = 'virtio'.freeze
package NicStatus
DOWN = 'down'.freeze
UP = 'up'.freeze
package NumaTuneMode
INTERLEAVE = 'interleave'.freeze
PREFERRED = 'preferred'.freeze
STRICT = 'strict'.freeze
package OpenStackNetworkProviderType
EXTERNAL = 'external'.freeze
NEUTRON = 'neutron'.freeze
package OpenstackVolumeAuthenticationKeyUsageType
CEPH = 'ceph'.freeze
package OsType
OTHER = 'other'.freeze
OTHER_LINUX = 'other_linux'.freeze
RHEL_3 = 'rhel_3'.freeze
RHEL_3X64 = 'rhel_3x64'.freeze
RHEL_4 = 'rhel_4'.freeze
RHEL_4X64 = 'rhel_4x64'.freeze
RHEL_5 = 'rhel_5'.freeze
RHEL_5X64 = 'rhel_5x64'.freeze
RHEL_6 = 'rhel_6'.freeze
RHEL_6X64 = 'rhel_6x64'.freeze
UNASSIGNED = 'unassigned'.freeze
WINDOWS_2003 = 'windows_2003'.freeze
WINDOWS_2003X64 = 'windows_2003x64'.freeze
WINDOWS_2008 = 'windows_2008'.freeze
WINDOWS_2008R2X64 = 'windows_2008r2x64'.freeze
WINDOWS_2008X64 = 'windows_2008x64'.freeze
WINDOWS_2012X64 = 'windows_2012x64'.freeze
WINDOWS_7 = 'windows_7'.freeze
WINDOWS_7X64 = 'windows_7x64'.freeze
WINDOWS_8 = 'windows_8'.freeze
WINDOWS_8X64 = 'windows_8x64'.freeze
WINDOWS_XP = 'windows_xp'.freeze
package PayloadEncoding
BASE64 = 'base64'.freeze
PLAINTEXT = 'plaintext'.freeze
package PmProxyType
CLUSTER = 'cluster'.freeze
DC = 'dc'.freeze
OTHER_DC = 'other_dc'.freeze
package PolicyUnitType
FILTER = 'filter'.freeze
LOAD_BALANCING = 'load_balancing'.freeze
WEIGHT = 'weight'.freeze
package PowerManagementStatus
OFF = 'off'.freeze
ON = 'on'.freeze
UNKNOWN = 'unknown'.freeze
package QosType
CPU = 'cpu'.freeze
HOSTNETWORK = 'hostnetwork'.freeze
NETWORK = 'network'.freeze
STORAGE = 'storage'.freeze
package QuotaModeType
AUDIT = 'audit'.freeze
DISABLED = 'disabled'.freeze
ENABLED = 'enabled'.freeze
package ReportedDeviceType
NETWORK = 'network'.freeze
package ResolutionType
ADD = 'add'.freeze
COPY = 'copy'.freeze
package RngSource
HWRNG = 'hwrng'.freeze
RANDOM = 'random'.freeze
package RoleType
ADMIN = 'admin'.freeze
USER = 'user'.freeze
package ScsiGenericIO
FILTERED = 'filtered'.freeze
UNFILTERED = 'unfiltered'.freeze
package SeLinuxMode
DISABLED = 'disabled'.freeze
ENFORCING = 'enforcing'.freeze
PERMISSIVE = 'permissive'.freeze
package SerialNumberPolicy
CUSTOM = 'custom'.freeze
HOST = 'host'.freeze
VM = 'vm'.freeze
package SnapshotStatus
IN_PREVIEW = 'in_preview'.freeze
LOCKED = 'locked'.freeze
OK = 'ok'.freeze
package SnapshotType
ACTIVE = 'active'.freeze
PREVIEW = 'preview'.freeze
REGULAR = 'regular'.freeze
STATELESS = 'stateless'.freeze
package SpmStatus
CONTENDING = 'contending'.freeze
NONE = 'none'.freeze
SPM = 'spm'.freeze
package SsoMethod
GUEST_AGENT = 'guest_agent'.freeze
package StatisticKind
COUNTER = 'counter'.freeze
GAUGE = 'gauge'.freeze
package StatisticUnit
BITS_PER_SECOND = 'bits_per_second'.freeze
BYTES = 'bytes'.freeze
BYTES_PER_SECOND = 'bytes_per_second'.freeze
COUNT_PER_SECOND = 'count_per_second'.freeze
NONE = 'none'.freeze
PERCENT = 'percent'.freeze
SECONDS = 'seconds'.freeze
package StepEnum
EXECUTING = 'executing'.freeze
FINALIZING = 'finalizing'.freeze
REBALANCING_VOLUME = 'rebalancing_volume'.freeze
REMOVING_BRICKS = 'removing_bricks'.freeze
UNKNOWN = 'unknown'.freeze
VALIDATING = 'validating'.freeze
package StepStatus
ABORTED = 'aborted'.freeze
FAILED = 'failed'.freeze
FINISHED = 'finished'.freeze
STARTED = 'started'.freeze
UNKNOWN = 'unknown'.freeze
package StorageDomainStatus
ACTIVATING = 'activating'.freeze
ACTIVE = 'active'.freeze
DETACHING = 'detaching'.freeze
INACTIVE = 'inactive'.freeze
LOCKED = 'locked'.freeze
MAINTENANCE = 'maintenance'.freeze
MIXED = 'mixed'.freeze
PREPARING_FOR_MAINTENANCE = 'preparing_for_maintenance'.freeze
UNATTACHED = 'unattached'.freeze
UNKNOWN = 'unknown'.freeze
package StorageDomainType
DATA = 'data'.freeze
EXPORT = 'export'.freeze
IMAGE = 'image'.freeze
ISO = 'iso'.freeze
VOLUME = 'volume'.freeze
package StorageFormat
V1 = 'v1'.freeze
V2 = 'v2'.freeze
V3 = 'v3'.freeze
package StorageType
CINDER = 'cinder'.freeze
FCP = 'fcp'.freeze
GLANCE = 'glance'.freeze
GLUSTERFS = 'glusterfs'.freeze
ISCSI = 'iscsi'.freeze
LOCALFS = 'localfs'.freeze
NFS = 'nfs'.freeze
POSIXFS = 'posixfs'.freeze
package SwitchType
LEGACY = 'legacy'.freeze
OVS = 'ovs'.freeze
package TemplateStatus
ILLEGAL = 'illegal'.freeze
LOCKED = 'locked'.freeze
OK = 'ok'.freeze
package TransportType
RDMA = 'rdma'.freeze
TCP = 'tcp'.freeze
package UsbType
LEGACY = 'legacy'.freeze
NATIVE = 'native'.freeze
package ValueType
DECIMAL = 'decimal'.freeze
INTEGER = 'integer'.freeze
STRING = 'string'.freeze
package VmAffinity
MIGRATABLE = 'migratable'.freeze
PINNED = 'pinned'.freeze
USER_MIGRATABLE = 'user_migratable'.freeze
package VmDeviceType
CDROM = 'cdrom'.freeze
FLOPPY = 'floppy'.freeze
package VmPoolType
AUTOMATIC = 'automatic'.freeze
MANUAL = 'manual'.freeze
package VmStatus
DOWN = 'down'.freeze
IMAGE_LOCKED = 'image_locked'.freeze
MIGRATING = 'migrating'.freeze
NOT_RESPONDING = 'not_responding'.freeze
PAUSED = 'paused'.freeze
POWERING_DOWN = 'powering_down'.freeze
POWERING_UP = 'powering_up'.freeze
REBOOT_IN_PROGRESS = 'reboot_in_progress'.freeze
RESTORING_STATE = 'restoring_state'.freeze
SAVING_STATE = 'saving_state'.freeze
SUSPENDED = 'suspended'.freeze
UNASSIGNED = 'unassigned'.freeze
UNKNOWN = 'unknown'.freeze
UP = 'up'.freeze
WAIT_FOR_LAUNCH = 'wait_for_launch'.freeze
package VmType
DESKTOP = 'desktop'.freeze
SERVER = 'server'.freeze
package VnicPassThroughMode
DISABLED = 'disabled'.freeze
ENABLED = 'enabled'.freeze
package WatchdogAction
DUMP = 'dump'.freeze
NONE = 'none'.freeze
PAUSE = 'pause'.freeze
POWEROFF = 'poweroff'.freeze
RESET = 'reset'.freeze
package WatchdogModel
I6300ESB = 'i6300esb'.freeze

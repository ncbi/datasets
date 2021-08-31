# V1reportsAssemblyDataReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** |  | [optional] 
**OrganismName** | Pointer to **string** |  | [optional] 
**Breed** | Pointer to **string** |  | [optional] 
**Cultivar** | Pointer to **string** |  | [optional] 
**Ecotype** | Pointer to **string** |  | [optional] 
**Isolate** | Pointer to **string** |  | [optional] 
**Sex** | Pointer to **string** |  | [optional] 
**Strain** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **int32** |  | [optional] 
**AssemblyInfo** | Pointer to [**V1reportsAssemblyInfo**](V1reportsAssemblyInfo.md) |  | [optional] 
**AssemblyStats** | Pointer to [**V1reportsAssemblyStats**](V1reportsAssemblyStats.md) |  | [optional] 
**OrganelleInfo** | Pointer to [**[]V1reportsOrganelleInfo**](V1reportsOrganelleInfo.md) |  | [optional] 
**AnnotationInfo** | Pointer to [**V1reportsAnnotationInfo**](V1reportsAnnotationInfo.md) |  | [optional] 
**WgsInfo** | Pointer to [**V1reportsWGSInfo**](V1reportsWGSInfo.md) |  | [optional] 

## Methods

### NewV1reportsAssemblyDataReport

`func NewV1reportsAssemblyDataReport() *V1reportsAssemblyDataReport`

NewV1reportsAssemblyDataReport instantiates a new V1reportsAssemblyDataReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsAssemblyDataReportWithDefaults

`func NewV1reportsAssemblyDataReportWithDefaults() *V1reportsAssemblyDataReport`

NewV1reportsAssemblyDataReportWithDefaults instantiates a new V1reportsAssemblyDataReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *V1reportsAssemblyDataReport) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1reportsAssemblyDataReport) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1reportsAssemblyDataReport) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1reportsAssemblyDataReport) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganismName

`func (o *V1reportsAssemblyDataReport) GetOrganismName() string`

GetOrganismName returns the OrganismName field if non-nil, zero value otherwise.

### GetOrganismNameOk

`func (o *V1reportsAssemblyDataReport) GetOrganismNameOk() (*string, bool)`

GetOrganismNameOk returns a tuple with the OrganismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismName

`func (o *V1reportsAssemblyDataReport) SetOrganismName(v string)`

SetOrganismName sets OrganismName field to given value.

### HasOrganismName

`func (o *V1reportsAssemblyDataReport) HasOrganismName() bool`

HasOrganismName returns a boolean if a field has been set.

### GetBreed

`func (o *V1reportsAssemblyDataReport) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *V1reportsAssemblyDataReport) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *V1reportsAssemblyDataReport) SetBreed(v string)`

SetBreed sets Breed field to given value.

### HasBreed

`func (o *V1reportsAssemblyDataReport) HasBreed() bool`

HasBreed returns a boolean if a field has been set.

### GetCultivar

`func (o *V1reportsAssemblyDataReport) GetCultivar() string`

GetCultivar returns the Cultivar field if non-nil, zero value otherwise.

### GetCultivarOk

`func (o *V1reportsAssemblyDataReport) GetCultivarOk() (*string, bool)`

GetCultivarOk returns a tuple with the Cultivar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultivar

`func (o *V1reportsAssemblyDataReport) SetCultivar(v string)`

SetCultivar sets Cultivar field to given value.

### HasCultivar

`func (o *V1reportsAssemblyDataReport) HasCultivar() bool`

HasCultivar returns a boolean if a field has been set.

### GetEcotype

`func (o *V1reportsAssemblyDataReport) GetEcotype() string`

GetEcotype returns the Ecotype field if non-nil, zero value otherwise.

### GetEcotypeOk

`func (o *V1reportsAssemblyDataReport) GetEcotypeOk() (*string, bool)`

GetEcotypeOk returns a tuple with the Ecotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcotype

`func (o *V1reportsAssemblyDataReport) SetEcotype(v string)`

SetEcotype sets Ecotype field to given value.

### HasEcotype

`func (o *V1reportsAssemblyDataReport) HasEcotype() bool`

HasEcotype returns a boolean if a field has been set.

### GetIsolate

`func (o *V1reportsAssemblyDataReport) GetIsolate() string`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V1reportsAssemblyDataReport) GetIsolateOk() (*string, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V1reportsAssemblyDataReport) SetIsolate(v string)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V1reportsAssemblyDataReport) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetSex

`func (o *V1reportsAssemblyDataReport) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *V1reportsAssemblyDataReport) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *V1reportsAssemblyDataReport) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *V1reportsAssemblyDataReport) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetStrain

`func (o *V1reportsAssemblyDataReport) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *V1reportsAssemblyDataReport) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *V1reportsAssemblyDataReport) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *V1reportsAssemblyDataReport) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetTaxId

`func (o *V1reportsAssemblyDataReport) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1reportsAssemblyDataReport) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1reportsAssemblyDataReport) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1reportsAssemblyDataReport) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetAssemblyInfo

`func (o *V1reportsAssemblyDataReport) GetAssemblyInfo() V1reportsAssemblyInfo`

GetAssemblyInfo returns the AssemblyInfo field if non-nil, zero value otherwise.

### GetAssemblyInfoOk

`func (o *V1reportsAssemblyDataReport) GetAssemblyInfoOk() (*V1reportsAssemblyInfo, bool)`

GetAssemblyInfoOk returns a tuple with the AssemblyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyInfo

`func (o *V1reportsAssemblyDataReport) SetAssemblyInfo(v V1reportsAssemblyInfo)`

SetAssemblyInfo sets AssemblyInfo field to given value.

### HasAssemblyInfo

`func (o *V1reportsAssemblyDataReport) HasAssemblyInfo() bool`

HasAssemblyInfo returns a boolean if a field has been set.

### GetAssemblyStats

`func (o *V1reportsAssemblyDataReport) GetAssemblyStats() V1reportsAssemblyStats`

GetAssemblyStats returns the AssemblyStats field if non-nil, zero value otherwise.

### GetAssemblyStatsOk

`func (o *V1reportsAssemblyDataReport) GetAssemblyStatsOk() (*V1reportsAssemblyStats, bool)`

GetAssemblyStatsOk returns a tuple with the AssemblyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyStats

`func (o *V1reportsAssemblyDataReport) SetAssemblyStats(v V1reportsAssemblyStats)`

SetAssemblyStats sets AssemblyStats field to given value.

### HasAssemblyStats

`func (o *V1reportsAssemblyDataReport) HasAssemblyStats() bool`

HasAssemblyStats returns a boolean if a field has been set.

### GetOrganelleInfo

`func (o *V1reportsAssemblyDataReport) GetOrganelleInfo() []V1reportsOrganelleInfo`

GetOrganelleInfo returns the OrganelleInfo field if non-nil, zero value otherwise.

### GetOrganelleInfoOk

`func (o *V1reportsAssemblyDataReport) GetOrganelleInfoOk() (*[]V1reportsOrganelleInfo, bool)`

GetOrganelleInfoOk returns a tuple with the OrganelleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganelleInfo

`func (o *V1reportsAssemblyDataReport) SetOrganelleInfo(v []V1reportsOrganelleInfo)`

SetOrganelleInfo sets OrganelleInfo field to given value.

### HasOrganelleInfo

`func (o *V1reportsAssemblyDataReport) HasOrganelleInfo() bool`

HasOrganelleInfo returns a boolean if a field has been set.

### GetAnnotationInfo

`func (o *V1reportsAssemblyDataReport) GetAnnotationInfo() V1reportsAnnotationInfo`

GetAnnotationInfo returns the AnnotationInfo field if non-nil, zero value otherwise.

### GetAnnotationInfoOk

`func (o *V1reportsAssemblyDataReport) GetAnnotationInfoOk() (*V1reportsAnnotationInfo, bool)`

GetAnnotationInfoOk returns a tuple with the AnnotationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationInfo

`func (o *V1reportsAssemblyDataReport) SetAnnotationInfo(v V1reportsAnnotationInfo)`

SetAnnotationInfo sets AnnotationInfo field to given value.

### HasAnnotationInfo

`func (o *V1reportsAssemblyDataReport) HasAnnotationInfo() bool`

HasAnnotationInfo returns a boolean if a field has been set.

### GetWgsInfo

`func (o *V1reportsAssemblyDataReport) GetWgsInfo() V1reportsWGSInfo`

GetWgsInfo returns the WgsInfo field if non-nil, zero value otherwise.

### GetWgsInfoOk

`func (o *V1reportsAssemblyDataReport) GetWgsInfoOk() (*V1reportsWGSInfo, bool)`

GetWgsInfoOk returns a tuple with the WgsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgsInfo

`func (o *V1reportsAssemblyDataReport) SetWgsInfo(v V1reportsWGSInfo)`

SetWgsInfo sets WgsInfo field to given value.

### HasWgsInfo

`func (o *V1reportsAssemblyDataReport) HasWgsInfo() bool`

HasWgsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



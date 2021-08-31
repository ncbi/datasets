# V1reportsMicroBiggeReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAcc** | Pointer to **string** |  | [optional] 
**Element** | Pointer to [**V1reportsElement**](V1reportsElement.md) |  | [optional] 
**Location** | Pointer to [**V1reportsSeqRangeSet**](V1reportsSeqRangeSet.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**Subclass** | Pointer to **string** |  | [optional] 
**AmrMethod** | Pointer to **string** |  | [optional] 
**IsPlus** | Pointer to **bool** |  | [optional] 
**ClosestReferenceSequenceComparison** | Pointer to [**V1reportsClosestReference**](V1reportsClosestReference.md) |  | [optional] 
**Taxonomy** | Pointer to [**V1reportsTaxonomy**](V1reportsTaxonomy.md) |  | [optional] 
**Biosample** | Pointer to [**V1reportsBiosample**](V1reportsBiosample.md) |  | [optional] 
**ReadToAssemblyCoverage** | Pointer to [**V1reportsReadToAssemblyCoverage**](V1reportsReadToAssemblyCoverage.md) |  | [optional] 
**AmrFinderPlus** | Pointer to [**V1reportsAmrFinderPlus**](V1reportsAmrFinderPlus.md) |  | [optional] 
**GenesOnContig** | Pointer to **[]string** |  | [optional] 
**GenesOnIsolate** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1reportsMicroBiggeReport

`func NewV1reportsMicroBiggeReport() *V1reportsMicroBiggeReport`

NewV1reportsMicroBiggeReport instantiates a new V1reportsMicroBiggeReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsMicroBiggeReportWithDefaults

`func NewV1reportsMicroBiggeReportWithDefaults() *V1reportsMicroBiggeReport`

NewV1reportsMicroBiggeReportWithDefaults instantiates a new V1reportsMicroBiggeReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAcc

`func (o *V1reportsMicroBiggeReport) GetTargetAcc() string`

GetTargetAcc returns the TargetAcc field if non-nil, zero value otherwise.

### GetTargetAccOk

`func (o *V1reportsMicroBiggeReport) GetTargetAccOk() (*string, bool)`

GetTargetAccOk returns a tuple with the TargetAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAcc

`func (o *V1reportsMicroBiggeReport) SetTargetAcc(v string)`

SetTargetAcc sets TargetAcc field to given value.

### HasTargetAcc

`func (o *V1reportsMicroBiggeReport) HasTargetAcc() bool`

HasTargetAcc returns a boolean if a field has been set.

### GetElement

`func (o *V1reportsMicroBiggeReport) GetElement() V1reportsElement`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *V1reportsMicroBiggeReport) GetElementOk() (*V1reportsElement, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *V1reportsMicroBiggeReport) SetElement(v V1reportsElement)`

SetElement sets Element field to given value.

### HasElement

`func (o *V1reportsMicroBiggeReport) HasElement() bool`

HasElement returns a boolean if a field has been set.

### GetLocation

`func (o *V1reportsMicroBiggeReport) GetLocation() V1reportsSeqRangeSet`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1reportsMicroBiggeReport) GetLocationOk() (*V1reportsSeqRangeSet, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1reportsMicroBiggeReport) SetLocation(v V1reportsSeqRangeSet)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1reportsMicroBiggeReport) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *V1reportsMicroBiggeReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1reportsMicroBiggeReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1reportsMicroBiggeReport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1reportsMicroBiggeReport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *V1reportsMicroBiggeReport) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *V1reportsMicroBiggeReport) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *V1reportsMicroBiggeReport) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *V1reportsMicroBiggeReport) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetClass

`func (o *V1reportsMicroBiggeReport) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *V1reportsMicroBiggeReport) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *V1reportsMicroBiggeReport) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *V1reportsMicroBiggeReport) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSubclass

`func (o *V1reportsMicroBiggeReport) GetSubclass() string`

GetSubclass returns the Subclass field if non-nil, zero value otherwise.

### GetSubclassOk

`func (o *V1reportsMicroBiggeReport) GetSubclassOk() (*string, bool)`

GetSubclassOk returns a tuple with the Subclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclass

`func (o *V1reportsMicroBiggeReport) SetSubclass(v string)`

SetSubclass sets Subclass field to given value.

### HasSubclass

`func (o *V1reportsMicroBiggeReport) HasSubclass() bool`

HasSubclass returns a boolean if a field has been set.

### GetAmrMethod

`func (o *V1reportsMicroBiggeReport) GetAmrMethod() string`

GetAmrMethod returns the AmrMethod field if non-nil, zero value otherwise.

### GetAmrMethodOk

`func (o *V1reportsMicroBiggeReport) GetAmrMethodOk() (*string, bool)`

GetAmrMethodOk returns a tuple with the AmrMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmrMethod

`func (o *V1reportsMicroBiggeReport) SetAmrMethod(v string)`

SetAmrMethod sets AmrMethod field to given value.

### HasAmrMethod

`func (o *V1reportsMicroBiggeReport) HasAmrMethod() bool`

HasAmrMethod returns a boolean if a field has been set.

### GetIsPlus

`func (o *V1reportsMicroBiggeReport) GetIsPlus() bool`

GetIsPlus returns the IsPlus field if non-nil, zero value otherwise.

### GetIsPlusOk

`func (o *V1reportsMicroBiggeReport) GetIsPlusOk() (*bool, bool)`

GetIsPlusOk returns a tuple with the IsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlus

`func (o *V1reportsMicroBiggeReport) SetIsPlus(v bool)`

SetIsPlus sets IsPlus field to given value.

### HasIsPlus

`func (o *V1reportsMicroBiggeReport) HasIsPlus() bool`

HasIsPlus returns a boolean if a field has been set.

### GetClosestReferenceSequenceComparison

`func (o *V1reportsMicroBiggeReport) GetClosestReferenceSequenceComparison() V1reportsClosestReference`

GetClosestReferenceSequenceComparison returns the ClosestReferenceSequenceComparison field if non-nil, zero value otherwise.

### GetClosestReferenceSequenceComparisonOk

`func (o *V1reportsMicroBiggeReport) GetClosestReferenceSequenceComparisonOk() (*V1reportsClosestReference, bool)`

GetClosestReferenceSequenceComparisonOk returns a tuple with the ClosestReferenceSequenceComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosestReferenceSequenceComparison

`func (o *V1reportsMicroBiggeReport) SetClosestReferenceSequenceComparison(v V1reportsClosestReference)`

SetClosestReferenceSequenceComparison sets ClosestReferenceSequenceComparison field to given value.

### HasClosestReferenceSequenceComparison

`func (o *V1reportsMicroBiggeReport) HasClosestReferenceSequenceComparison() bool`

HasClosestReferenceSequenceComparison returns a boolean if a field has been set.

### GetTaxonomy

`func (o *V1reportsMicroBiggeReport) GetTaxonomy() V1reportsTaxonomy`

GetTaxonomy returns the Taxonomy field if non-nil, zero value otherwise.

### GetTaxonomyOk

`func (o *V1reportsMicroBiggeReport) GetTaxonomyOk() (*V1reportsTaxonomy, bool)`

GetTaxonomyOk returns a tuple with the Taxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomy

`func (o *V1reportsMicroBiggeReport) SetTaxonomy(v V1reportsTaxonomy)`

SetTaxonomy sets Taxonomy field to given value.

### HasTaxonomy

`func (o *V1reportsMicroBiggeReport) HasTaxonomy() bool`

HasTaxonomy returns a boolean if a field has been set.

### GetBiosample

`func (o *V1reportsMicroBiggeReport) GetBiosample() V1reportsBiosample`

GetBiosample returns the Biosample field if non-nil, zero value otherwise.

### GetBiosampleOk

`func (o *V1reportsMicroBiggeReport) GetBiosampleOk() (*V1reportsBiosample, bool)`

GetBiosampleOk returns a tuple with the Biosample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosample

`func (o *V1reportsMicroBiggeReport) SetBiosample(v V1reportsBiosample)`

SetBiosample sets Biosample field to given value.

### HasBiosample

`func (o *V1reportsMicroBiggeReport) HasBiosample() bool`

HasBiosample returns a boolean if a field has been set.

### GetReadToAssemblyCoverage

`func (o *V1reportsMicroBiggeReport) GetReadToAssemblyCoverage() V1reportsReadToAssemblyCoverage`

GetReadToAssemblyCoverage returns the ReadToAssemblyCoverage field if non-nil, zero value otherwise.

### GetReadToAssemblyCoverageOk

`func (o *V1reportsMicroBiggeReport) GetReadToAssemblyCoverageOk() (*V1reportsReadToAssemblyCoverage, bool)`

GetReadToAssemblyCoverageOk returns a tuple with the ReadToAssemblyCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToAssemblyCoverage

`func (o *V1reportsMicroBiggeReport) SetReadToAssemblyCoverage(v V1reportsReadToAssemblyCoverage)`

SetReadToAssemblyCoverage sets ReadToAssemblyCoverage field to given value.

### HasReadToAssemblyCoverage

`func (o *V1reportsMicroBiggeReport) HasReadToAssemblyCoverage() bool`

HasReadToAssemblyCoverage returns a boolean if a field has been set.

### GetAmrFinderPlus

`func (o *V1reportsMicroBiggeReport) GetAmrFinderPlus() V1reportsAmrFinderPlus`

GetAmrFinderPlus returns the AmrFinderPlus field if non-nil, zero value otherwise.

### GetAmrFinderPlusOk

`func (o *V1reportsMicroBiggeReport) GetAmrFinderPlusOk() (*V1reportsAmrFinderPlus, bool)`

GetAmrFinderPlusOk returns a tuple with the AmrFinderPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmrFinderPlus

`func (o *V1reportsMicroBiggeReport) SetAmrFinderPlus(v V1reportsAmrFinderPlus)`

SetAmrFinderPlus sets AmrFinderPlus field to given value.

### HasAmrFinderPlus

`func (o *V1reportsMicroBiggeReport) HasAmrFinderPlus() bool`

HasAmrFinderPlus returns a boolean if a field has been set.

### GetGenesOnContig

`func (o *V1reportsMicroBiggeReport) GetGenesOnContig() []string`

GetGenesOnContig returns the GenesOnContig field if non-nil, zero value otherwise.

### GetGenesOnContigOk

`func (o *V1reportsMicroBiggeReport) GetGenesOnContigOk() (*[]string, bool)`

GetGenesOnContigOk returns a tuple with the GenesOnContig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesOnContig

`func (o *V1reportsMicroBiggeReport) SetGenesOnContig(v []string)`

SetGenesOnContig sets GenesOnContig field to given value.

### HasGenesOnContig

`func (o *V1reportsMicroBiggeReport) HasGenesOnContig() bool`

HasGenesOnContig returns a boolean if a field has been set.

### GetGenesOnIsolate

`func (o *V1reportsMicroBiggeReport) GetGenesOnIsolate() []string`

GetGenesOnIsolate returns the GenesOnIsolate field if non-nil, zero value otherwise.

### GetGenesOnIsolateOk

`func (o *V1reportsMicroBiggeReport) GetGenesOnIsolateOk() (*[]string, bool)`

GetGenesOnIsolateOk returns a tuple with the GenesOnIsolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesOnIsolate

`func (o *V1reportsMicroBiggeReport) SetGenesOnIsolate(v []string)`

SetGenesOnIsolate sets GenesOnIsolate field to given value.

### HasGenesOnIsolate

`func (o *V1reportsMicroBiggeReport) HasGenesOnIsolate() bool`

HasGenesOnIsolate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



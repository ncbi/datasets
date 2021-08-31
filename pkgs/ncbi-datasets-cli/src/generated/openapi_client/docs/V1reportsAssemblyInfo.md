# V1reportsAssemblyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyAccession** | Pointer to **string** |  | [optional] 
**PairedAssemblyAccession** | Pointer to **string** |  | [optional] 
**AssemblyLevel** | Pointer to **string** |  | [optional] 
**AssemblyName** | Pointer to **string** |  | [optional] 
**AssemblyType** | Pointer to **string** |  | [optional] 
**BioprojectLineage** | Pointer to [**[]V1reportsBioProjectLineage**](V1reportsBioProjectLineage.md) |  | [optional] 
**SubmissionDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GenbankAssmAccession** | Pointer to **string** |  | [optional] 
**Submitter** | Pointer to **string** |  | [optional] 
**RefseqCategory** | Pointer to **string** |  | [optional] 
**RefseqAssmAccession** | Pointer to **string** |  | [optional] 
**UcscAssmName** | Pointer to **string** |  | [optional] 
**LinkedAssembly** | Pointer to **string** |  | [optional] 
**SequencingTech** | Pointer to **string** |  | [optional] 
**BiosampleAccession** | Pointer to **string** |  | [optional] 
**BlastUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewV1reportsAssemblyInfo

`func NewV1reportsAssemblyInfo() *V1reportsAssemblyInfo`

NewV1reportsAssemblyInfo instantiates a new V1reportsAssemblyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsAssemblyInfoWithDefaults

`func NewV1reportsAssemblyInfoWithDefaults() *V1reportsAssemblyInfo`

NewV1reportsAssemblyInfoWithDefaults instantiates a new V1reportsAssemblyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyAccession

`func (o *V1reportsAssemblyInfo) GetAssemblyAccession() string`

GetAssemblyAccession returns the AssemblyAccession field if non-nil, zero value otherwise.

### GetAssemblyAccessionOk

`func (o *V1reportsAssemblyInfo) GetAssemblyAccessionOk() (*string, bool)`

GetAssemblyAccessionOk returns a tuple with the AssemblyAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyAccession

`func (o *V1reportsAssemblyInfo) SetAssemblyAccession(v string)`

SetAssemblyAccession sets AssemblyAccession field to given value.

### HasAssemblyAccession

`func (o *V1reportsAssemblyInfo) HasAssemblyAccession() bool`

HasAssemblyAccession returns a boolean if a field has been set.

### GetPairedAssemblyAccession

`func (o *V1reportsAssemblyInfo) GetPairedAssemblyAccession() string`

GetPairedAssemblyAccession returns the PairedAssemblyAccession field if non-nil, zero value otherwise.

### GetPairedAssemblyAccessionOk

`func (o *V1reportsAssemblyInfo) GetPairedAssemblyAccessionOk() (*string, bool)`

GetPairedAssemblyAccessionOk returns a tuple with the PairedAssemblyAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedAssemblyAccession

`func (o *V1reportsAssemblyInfo) SetPairedAssemblyAccession(v string)`

SetPairedAssemblyAccession sets PairedAssemblyAccession field to given value.

### HasPairedAssemblyAccession

`func (o *V1reportsAssemblyInfo) HasPairedAssemblyAccession() bool`

HasPairedAssemblyAccession returns a boolean if a field has been set.

### GetAssemblyLevel

`func (o *V1reportsAssemblyInfo) GetAssemblyLevel() string`

GetAssemblyLevel returns the AssemblyLevel field if non-nil, zero value otherwise.

### GetAssemblyLevelOk

`func (o *V1reportsAssemblyInfo) GetAssemblyLevelOk() (*string, bool)`

GetAssemblyLevelOk returns a tuple with the AssemblyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyLevel

`func (o *V1reportsAssemblyInfo) SetAssemblyLevel(v string)`

SetAssemblyLevel sets AssemblyLevel field to given value.

### HasAssemblyLevel

`func (o *V1reportsAssemblyInfo) HasAssemblyLevel() bool`

HasAssemblyLevel returns a boolean if a field has been set.

### GetAssemblyName

`func (o *V1reportsAssemblyInfo) GetAssemblyName() string`

GetAssemblyName returns the AssemblyName field if non-nil, zero value otherwise.

### GetAssemblyNameOk

`func (o *V1reportsAssemblyInfo) GetAssemblyNameOk() (*string, bool)`

GetAssemblyNameOk returns a tuple with the AssemblyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyName

`func (o *V1reportsAssemblyInfo) SetAssemblyName(v string)`

SetAssemblyName sets AssemblyName field to given value.

### HasAssemblyName

`func (o *V1reportsAssemblyInfo) HasAssemblyName() bool`

HasAssemblyName returns a boolean if a field has been set.

### GetAssemblyType

`func (o *V1reportsAssemblyInfo) GetAssemblyType() string`

GetAssemblyType returns the AssemblyType field if non-nil, zero value otherwise.

### GetAssemblyTypeOk

`func (o *V1reportsAssemblyInfo) GetAssemblyTypeOk() (*string, bool)`

GetAssemblyTypeOk returns a tuple with the AssemblyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyType

`func (o *V1reportsAssemblyInfo) SetAssemblyType(v string)`

SetAssemblyType sets AssemblyType field to given value.

### HasAssemblyType

`func (o *V1reportsAssemblyInfo) HasAssemblyType() bool`

HasAssemblyType returns a boolean if a field has been set.

### GetBioprojectLineage

`func (o *V1reportsAssemblyInfo) GetBioprojectLineage() []V1reportsBioProjectLineage`

GetBioprojectLineage returns the BioprojectLineage field if non-nil, zero value otherwise.

### GetBioprojectLineageOk

`func (o *V1reportsAssemblyInfo) GetBioprojectLineageOk() (*[]V1reportsBioProjectLineage, bool)`

GetBioprojectLineageOk returns a tuple with the BioprojectLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojectLineage

`func (o *V1reportsAssemblyInfo) SetBioprojectLineage(v []V1reportsBioProjectLineage)`

SetBioprojectLineage sets BioprojectLineage field to given value.

### HasBioprojectLineage

`func (o *V1reportsAssemblyInfo) HasBioprojectLineage() bool`

HasBioprojectLineage returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *V1reportsAssemblyInfo) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *V1reportsAssemblyInfo) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *V1reportsAssemblyInfo) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *V1reportsAssemblyInfo) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetDescription

`func (o *V1reportsAssemblyInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1reportsAssemblyInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1reportsAssemblyInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1reportsAssemblyInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenbankAssmAccession

`func (o *V1reportsAssemblyInfo) GetGenbankAssmAccession() string`

GetGenbankAssmAccession returns the GenbankAssmAccession field if non-nil, zero value otherwise.

### GetGenbankAssmAccessionOk

`func (o *V1reportsAssemblyInfo) GetGenbankAssmAccessionOk() (*string, bool)`

GetGenbankAssmAccessionOk returns a tuple with the GenbankAssmAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenbankAssmAccession

`func (o *V1reportsAssemblyInfo) SetGenbankAssmAccession(v string)`

SetGenbankAssmAccession sets GenbankAssmAccession field to given value.

### HasGenbankAssmAccession

`func (o *V1reportsAssemblyInfo) HasGenbankAssmAccession() bool`

HasGenbankAssmAccession returns a boolean if a field has been set.

### GetSubmitter

`func (o *V1reportsAssemblyInfo) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *V1reportsAssemblyInfo) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *V1reportsAssemblyInfo) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *V1reportsAssemblyInfo) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetRefseqCategory

`func (o *V1reportsAssemblyInfo) GetRefseqCategory() string`

GetRefseqCategory returns the RefseqCategory field if non-nil, zero value otherwise.

### GetRefseqCategoryOk

`func (o *V1reportsAssemblyInfo) GetRefseqCategoryOk() (*string, bool)`

GetRefseqCategoryOk returns a tuple with the RefseqCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqCategory

`func (o *V1reportsAssemblyInfo) SetRefseqCategory(v string)`

SetRefseqCategory sets RefseqCategory field to given value.

### HasRefseqCategory

`func (o *V1reportsAssemblyInfo) HasRefseqCategory() bool`

HasRefseqCategory returns a boolean if a field has been set.

### GetRefseqAssmAccession

`func (o *V1reportsAssemblyInfo) GetRefseqAssmAccession() string`

GetRefseqAssmAccession returns the RefseqAssmAccession field if non-nil, zero value otherwise.

### GetRefseqAssmAccessionOk

`func (o *V1reportsAssemblyInfo) GetRefseqAssmAccessionOk() (*string, bool)`

GetRefseqAssmAccessionOk returns a tuple with the RefseqAssmAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqAssmAccession

`func (o *V1reportsAssemblyInfo) SetRefseqAssmAccession(v string)`

SetRefseqAssmAccession sets RefseqAssmAccession field to given value.

### HasRefseqAssmAccession

`func (o *V1reportsAssemblyInfo) HasRefseqAssmAccession() bool`

HasRefseqAssmAccession returns a boolean if a field has been set.

### GetUcscAssmName

`func (o *V1reportsAssemblyInfo) GetUcscAssmName() string`

GetUcscAssmName returns the UcscAssmName field if non-nil, zero value otherwise.

### GetUcscAssmNameOk

`func (o *V1reportsAssemblyInfo) GetUcscAssmNameOk() (*string, bool)`

GetUcscAssmNameOk returns a tuple with the UcscAssmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcscAssmName

`func (o *V1reportsAssemblyInfo) SetUcscAssmName(v string)`

SetUcscAssmName sets UcscAssmName field to given value.

### HasUcscAssmName

`func (o *V1reportsAssemblyInfo) HasUcscAssmName() bool`

HasUcscAssmName returns a boolean if a field has been set.

### GetLinkedAssembly

`func (o *V1reportsAssemblyInfo) GetLinkedAssembly() string`

GetLinkedAssembly returns the LinkedAssembly field if non-nil, zero value otherwise.

### GetLinkedAssemblyOk

`func (o *V1reportsAssemblyInfo) GetLinkedAssemblyOk() (*string, bool)`

GetLinkedAssemblyOk returns a tuple with the LinkedAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAssembly

`func (o *V1reportsAssemblyInfo) SetLinkedAssembly(v string)`

SetLinkedAssembly sets LinkedAssembly field to given value.

### HasLinkedAssembly

`func (o *V1reportsAssemblyInfo) HasLinkedAssembly() bool`

HasLinkedAssembly returns a boolean if a field has been set.

### GetSequencingTech

`func (o *V1reportsAssemblyInfo) GetSequencingTech() string`

GetSequencingTech returns the SequencingTech field if non-nil, zero value otherwise.

### GetSequencingTechOk

`func (o *V1reportsAssemblyInfo) GetSequencingTechOk() (*string, bool)`

GetSequencingTechOk returns a tuple with the SequencingTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequencingTech

`func (o *V1reportsAssemblyInfo) SetSequencingTech(v string)`

SetSequencingTech sets SequencingTech field to given value.

### HasSequencingTech

`func (o *V1reportsAssemblyInfo) HasSequencingTech() bool`

HasSequencingTech returns a boolean if a field has been set.

### GetBiosampleAccession

`func (o *V1reportsAssemblyInfo) GetBiosampleAccession() string`

GetBiosampleAccession returns the BiosampleAccession field if non-nil, zero value otherwise.

### GetBiosampleAccessionOk

`func (o *V1reportsAssemblyInfo) GetBiosampleAccessionOk() (*string, bool)`

GetBiosampleAccessionOk returns a tuple with the BiosampleAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosampleAccession

`func (o *V1reportsAssemblyInfo) SetBiosampleAccession(v string)`

SetBiosampleAccession sets BiosampleAccession field to given value.

### HasBiosampleAccession

`func (o *V1reportsAssemblyInfo) HasBiosampleAccession() bool`

HasBiosampleAccession returns a boolean if a field has been set.

### GetBlastUrl

`func (o *V1reportsAssemblyInfo) GetBlastUrl() string`

GetBlastUrl returns the BlastUrl field if non-nil, zero value otherwise.

### GetBlastUrlOk

`func (o *V1reportsAssemblyInfo) GetBlastUrlOk() (*string, bool)`

GetBlastUrlOk returns a tuple with the BlastUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastUrl

`func (o *V1reportsAssemblyInfo) SetBlastUrl(v string)`

SetBlastUrl sets BlastUrl field to given value.

### HasBlastUrl

`func (o *V1reportsAssemblyInfo) HasBlastUrl() bool`

HasBlastUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



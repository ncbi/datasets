# NCBI Datasets jq cheatsheet for genome metadata
Try out jq commands on the web: [https://jqplay.org/](https://jqplay.org/)

## Installation
### Download datasets
[https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/)
### Download jq
[https://stedolan.github.io/jq/](https://stedolan.github.io/jq/)

## Quick commands:

### First generate a json file with metadata for all cow genomes
```
datasets summary genome taxon cow > cow_genomes.json
```

### Pretty-print the data (and only show the first 10 lines)
Note that the data is hierarchically structured: the busco information is nested within annotation_metadata, and annotation_metadata is nested within the assembly object
```
datasets summary genome taxon cow | jq . | head
{
  "assemblies": [
    {
      "assembly": {
        "annotation_metadata": {
          "busco": {
            "busco_lineage": "cetartiodactyla_odb10",
            "busco_ver": "4.0.2 ",
            "complete": 0.98672664,
            "duplicated": 0.005024372,
```

### Show the assembly count
```
jq '.total_count' cow_genomes.json
7
```

### Only show data for the first assembly in a set of multiple assemblies (and only show the first 10 lines)

`assemblies[0]` is used to specify the first assembly in the set, `assemblies[1]` refers to the second assembly, etc.
```
datasets summary genome taxon cow | jq '.assemblies[0]' | head 
{
  "assembly": {
    "annotation_metadata": {
      "busco": {
        "busco_lineage": "cetartiodactyla_odb10",
        "busco_ver": "4.0.2 ",
        "complete": 0.98672664,
        "duplicated": 0.005024372,
        "fragmented": 0.0045744283,
        "missing": 0.008698912,
```

### Show the BUSCO data for the first assembly in a set
```
datasets summary genome taxon cow | jq '.assemblies[0].assembly.annotation_metadata.busco'        
{
  "busco_lineage": "cetartiodactyla_odb10",
  "busco_ver": "4.0.2 ",
  "complete": 0.98672664,
  "duplicated": 0.005024372,
  "fragmented": 0.0045744283,
  "missing": 0.008698912,
  "single_copy": 0.98170227,
  "total_count": "13335"
}
```

### Show the gene counts for the first assembly in a set
```
datasets summary genome taxon cow | jq '.assemblies[0].assembly.annotation_metadata.stats.gene_counts' 
{
  "protein_coding": 21039,
  "total": 35143
}
```

### Show the assembly accession, submitter, and submission date for the first assembly in a set and format the output in a new JSON object with custom key names
```
datasets summary genome taxon cow | jq '.assemblies[0].assembly | {accession: .assembly_accession, submitter: .submitter, date: .submission_date}' 
{
  "accession": "GCF_002263795.1",
  "submitter": "USDA ARS",
  "date": "2018-04-11"
}
```

### Generate a table of 3 columns including assembly accession, submission date and submitter
```
datasets summary genome taxon cow | jq -r '.assemblies[].assembly | [.assembly_accession, .submission_date, .submitter] | @tsv' 
GCF_002263795.1	2018-04-11	USDA ARS
GCF_000003055.6	2014-11-25	Center for Bioinformatics and Computational Biology, University of Maryland
GCF_000003205.5	2011-11-02	Cattle Genome Sequencing International Consortium
GCF_000003205.7	2015-11-19	Cattle Genome Sequencing International Consortium
GCA_000003055.5	2014-11-25	Center for Bioinformatics and Computational Biology, University of Maryland
GCA_000003205.6	2015-11-19	Cattle Genome Sequencing International Consortium
GCA_002263795.2	2018-04-11	USDA ARS
```

### Show the assembly accession and the chromosome count for the first assembly in a set
**Note #1**: we will use jq length to count the number of chromosomes
**Note #2**: chromosome count includes all assembled chromosomes, the set of unplaced scaffolds counts as 1 chromosome, and each organelle genome counts as 1 chromosome, so in this example 29 autosomes + 1 X chromosome + 1 set of unplaced scaffolds + 1 mitochondrial genome = 32
```
datasets summary genome taxon cow | jq '.assemblies[0].assembly | {accession: .assembly_accession, chromosome_count: [.chromosomes[]] | length}'        
{
  "accession": "GCF_002263795.1",
  "chromosome_count": 32
}
```

### Available fields for genome assembly summary metadata


![](https://github.com/ncbi/datasets/blob/workshop-cshl-2021/training/cshl-2021/images/json8.png)


GENERAL FIELDS
```
.total_count
.assemblies[].assembly.biosample_accession
.assemblies[].assembly.blast_url
.assemblies[].assembly.gc_count
.assemblies[].assembly.contig_n50
.assemblies[].assembly.display_name
.assemblies[].assembly.estimated_size
.assemblies[].assembly.seq_length
.assemblies[].assembly.paired_assembly_accession
.assemblies[].assembly.submission_date
.assemblies[].assembly.submitter
```

ANNOTATION
```
.assemblies[].assembly.annotation_metadata.busco.busco_lineage
.assemblies[].assembly.annotation_metadata.busco.busco_ver
.assemblies[].assembly.annotation_metadata.busco.complete
.assemblies[].assembly.annotation_metadata.busco.duplicated
.assemblies[].assembly.annotation_metadata.busco.fragmented
.assemblies[].assembly.annotation_metadata.busco.missing
.assemblies[].assembly.annotation_metadata.busco.single_copy
.assemblies[].assembly.annotation_metadata.busco.total_count
.assemblies[].assembly.annotation_metadata.file[].estimated_size
.assemblies[].assembly.annotation_metadata.file[].type
.assemblies[].assembly.annotation_metadata.name
.assemblies[].assembly.annotation_metadata.release_date
.assemblies[].assembly.annotation_metadata.release_number
.assemblies[].assembly.annotation_metadata.report_url
.assemblies[].assembly.annotation_metadata.source
.assemblies[].assembly.annotation_metadata.stats.gene_counts.protein_coding
.assemblies[].assembly.annotation_metadata.stats.gene_counts.total
.assemblies[].assembly.assembly_accession
```

BIOPROJECT
```
.assemblies[].assembly.bioproject_lineages[].bioprojects[].accession
.assemblies[].assembly.bioproject_lineages[].bioprojects[].parent_accessions[]
.assemblies[].assembly.bioproject_lineages[].bioprojects[].title
```

CHROMOSOME
```
.assemblies[].assembly.chromosomes[].accession_version
.assemblies[].assembly.chromosomes[].gc_count
.assemblies[].assembly.chromosomes[].length
.assemblies[].assembly.chromosomes[].name
```

ORGANISM
```
.assemblies[].assembly.org.assembly_counts.node
.assemblies[].assembly.org.assembly_counts.subtree
.assemblies[].assembly.org.breed
.assemblies[].assembly.org.common_name
.assemblies[].assembly.org.ecotype
.assemblies[].assembly.org.isolate
.assemblies[].assembly.org.key
.assemblies[].assembly.org.merged_tax_ids[]
.assemblies[].assembly.org.parent_tax_id
.assemblies[].assembly.org.rank
.assemblies[].assembly.org.sci_name
.assemblies[].assembly.org.sex
.assemblies[].assembly.org.strain
.assemblies[].assembly.org.tax_id
.assemblies[].assembly.org.title
```

# NCBI Datasets

NCBI Datasets is a resource that lets you easily gather data from across NCBI databases. You can use it to find and download sequence, annotation, and metadata for genes and genomes using our command-line interface (CLI) tools or [NCBI Datasets](https://www.ncbi.nlm.nih.gov/datasets/) web interface.

NCBI Datasets tools are under active development. To submit feedback, please create a [GitHub issue](https://github.com/ncbi/datasets/issues/new/choose) or [contact NCBI](mailto:info@ncbi.nlm.nih.gov) directly with your questions, comments or feature requests.  

## Install the NCBI Datasets command-line tools

[![Anaconda.org badge](https://anaconda.org/conda-forge/ncbi-datasets-cli/badges/version.svg)](https://anaconda.org/conda-forge/ncbi-datasets-cli)
[![Platforms badge](https://anaconda.org/conda-forge/ncbi-datasets-cli/badges/platforms.svg)](https://anaconda.org/conda-forge/ncbi-datasets-cli)
[![Total downloads badge](https://anaconda.org/conda-forge/ncbi-datasets-cli/badges/downloads.svg)](https://anaconda.org/conda-forge/ncbi-datasets-cli)

Install the latest version (CLI v16.x) of the NCBI Datasets CLI tools, *datasets* and *dataformat*, using conda:

`conda install -c conda-forge ncbi-datasets-cli`

For other installation options, see our CLI tools [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) instructions. 

## Use the NCBI Datasets command-line tools

Use *datasets* to download biological sequence data across all domains of life from NCBI.

Use *dataformat* to convert metadata included as part of the data package from JSON Lines format to other formats.

### Examples:
Use *datasets* to download a genome data package for the human reference genome GRCh38:

`datasets download genome taxon human --reference --filename human-reference.zip`

Use *dataformat* to extract selected fields of metadata from the downloaded data package for the human reference genome, GRCh38:
```
dataformat tsv genome --package human-reference.zip --fields organism-name,assminfo-name,accession,assminfo-submitter
Organism name	Assembly Name	Assembly Accession	Assembly Submitter
Homo sapiens	GRCh38.p14	GCF_000001405.40	Genome Reference Consortium
```

The Datasets CLI schematic below also outlines the available commands for the *datasets* CLI. 
![Datasets CLI schematic](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/datasets_schema_taxonomy.png)

### Download large numbers of genomes

Download large numbers of genomes by first downloading a dehydrated zip archive and then accessing the data in three steps.

1. Download the dehydrated zip archive
1. Unzip the downloaded zip archive
1. Rehydrate to access the data


Try this example for the human reference genome:

1. Download the dehydrated zip archive:  
`datasets download genome accession GCF_000001405.40 --dehydrated --filename human_GRCh38_dataset.zip`

1. Unzip the downloaded zip archive:  
`unzip human_GRCh38_dataset.zip -d my_human_dataset`

1. Rehydrate to access the data:  
`datasets rehydrate --directory my_human_dataset/`

For more information, see [how to download large genome data packages](https://www.ncbi.nlm.nih.gov/datasets/docs/how-tos/genomes/large-download/).

### Use your API key with the NCBI Datasets command-line tools
NCBI Datasets API and command-line tool requests are rate-limited. By default, this rate limit is set at 5 requests per second (rps). By using your API key, you can increase this rate limit to 10 rps. For more information, see our documentation on [how to get an API key](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/api/api-keys/#get-your-api-key) and [how to use your API key.](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/api/api-keys/#use-your-api-key-with-the-ncbi-datasets-command-line-tools)

## NCBI Datasets data packages
NCBI Datasets provides sequence, annotation, metadata and other biological data as [NCBI Datasets Data Package zip archives](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/data-packages/).

We currently offer four types of data package: 
1. An [NCBI Datasets Gene Data Package](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/data-packages/gene-package/)
1. An [NCBI Datasets Genome Data Package](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/data-packages/genome/)
1. A specialized [NCBI Datasets Virus Data Package](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/data-packages/virus-genome/).
1. An [NCBI Datasets Taxonomy Data Package](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/data-packages/taxonomy/)

## NCBI Datasets data reports
NCBI Datasets data packages include data report files that contain metadata about the requested records. [Data report schemas](https://www.ncbi.nlm.nih.gov/datasets/docs/reference-docs/data-reports/) describe each type of data report, including available fields, with descriptions and examples.

## Citing NCBI Datasets
### Exploring and retrieving sequence and metadata for species across the tree of life with NCBI Datasets

O'Leary NA, Cox E, Holmes JB, Anderson WR, Falk R, Hem V, Tsuchiya MTN, Schuler GD, Zhang X, Torcivia J, Ketter A, Breen L, Cothran J, Bajwa H, Tinne J, Meric PA, Hlavina W, Schneider VA. [Exploring and retrieving sequence and metadata for species across the tree of life with NCBI Datasets.](https://www.nature.com/articles/s41597-024-03571-y) Sci Data. 2024 Jul 5;11(1):732. doi: 10.1038/s41597-024-03571-y. PMID: 38969627; PMCID: PMC11226681.

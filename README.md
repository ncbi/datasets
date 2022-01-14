# NCBI Datasets

NCBI Datasets is a new resource that lets you easily gather data from across NCBI databases.

Find and download sequence, annotation, and metadata for genes and genomes using our command-line tools or [web interface](https://www.ncbi.nlm.nih.gov/datasets/).

NCBI Datasets tools are under active development. Submit feedback by creating a GitHub issue or you may [contact NCBI](mailto:info@ncbi.nlm.nih.gov) directly with your questions, comments or feature requests.

## Install the Datasets command-line tools

[![Anaconda-Server Badge](https://anaconda.org/conda-forge/ncbi-datasets-cli/badges/installer/conda.svg)](https://anaconda.org/conda-forge/ncbi-datasets-cli)

Download and install the NCBI Datasets command-line tools, datasets and dataformat:

`conda install -c conda-forge ncbi-datasets-cli`

For other ways to install, see our command-line tool [quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/quickstarts/command-line-tools/).

## Use the Datasets command-line tools

Use **datasets** to download biological sequence data across all domains of life from NCBI.

Use **dataformat** to convert metadata included as part of the data package from JSON Lines format to other formats.

### Examples:
Use **datasets** to download a genome data package for the human reference genome, GRCh38:

`datasets download genome taxon human --reference --filename human-reference.zip`

Use **dataformat** to extract selected fields of metadata from the downloaded data package for the human reference genome, GRCh38:
```
dataformat tsv genome --package human-reference.zip --fields organism-name,assminfo-name,assminfo-accession,assminfo-submitter
Organism name	Assembly Name	Assembly Accession	Assembly Submitter
Homo sapiens	GRCh38.p13	GCF_000001405.39	Genome Reference Consortium
```

The schematic below outlines the available commands for the **datasets** command-line tool:
![Datasets CLI schematic](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/quickstarts/datasets_schema_github.png)

### Download large numbers of genomes

Download large numbers of genomes by first downloading a dehydrated zip archive and then getting the data in three steps.

1. Download the dehydrated zip archive
1. Unzip the downloaded zip archive
1. Rehydrate to get the data


Try this example for the human reference genome:

1. Download the dehydrated zip archive
`datasets download genome accession GCF_000001405.39 --dehydrated --filename human_GRCh38_dataset.zip`

2. Unzip the downloaded zip archive
`unzip human_GRCh38_dataset.zip -d my_human_dataset`

3. Rehydrate to get the data
`datasets rehydrate --directory my_human_dataset/`

For more information, see [how to download large genome data packages](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/).

## Datasets data packages
NCBI Datasets provides sequence, annotation, metadata and other biological data as an [NCBI Datasets Data Package zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/data-packages/).

We currently offer three types of data package: a [gene data package](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/data-packages/gene-package/), a [genome data package](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/data-packages/genome/), and a specialized [SARS-CoV-2 data package](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/data-packages/sars-cov-2-genome/).

## Datasets data reports
NCBI Datasets data packages include data report files that contain metadata about the requested records. [Data report schemas](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/reference-docs/data-reports/) describe each type of data report, including the available fields, with descriptions and examples.

## Jupyter notebooks
[Jupyter notebooks](https://github.com/ncbi/datasets/tree/master/examples/jupyter) provide a way to explore our command-line tool, python package and API.


## Python library
Programmers may be interested in trying the NCBI Datasets [python library](https://github.com/ncbi/datasets/tree/master/client_docs/python) to interact with the NCBI Datasets [REST API](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/reference-docs/rest-api/). Note that we are offering limited support for the python package at this time.



**NOTE:** We recommend avoiding the use of this package in an environment also containing the following packages:
```bash
google-cloud-bigquery
pandas-gbq
```

We have observed some dependency issues with these packages such that some environments may fail to interact properly with python BigQuery API endpoints.

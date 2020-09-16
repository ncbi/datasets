# Welcome to NCBI Datasets

NCBI Datasets is a new resource that lets you easily gather data from across NCBI databases.

NCBI Datasets uses an iterative software development approach based on user research and interviews, with the goal of helping users to find the data they want more quickly and easily, in file formats that can be used with common bioinformatics workflows.

Use this python package with our [RESTful API](https://www.ncbi.nlm.nih.gov/datasets/docs/datasets-api/) to get genome datasets including sequence, annotation and a data report containing detailed metadata, for organisms representing all domains of life.

We have recently added viral genome and protein datasets for SARS-CoV-2.
These datasets include sequence, annotation and a data report containing isolate, host, collection date and other metadata. Viral genome datasets are also available for other viruses in the coronavirus family.

We also offer a [command-line tool](https://www.ncbi.nlm.nih.gov/datasets/docs/command-line-start/), for access to genome data for all taxa, and a [web interface that is focused on eukaryotic genomes.](https://www.ncbi.nlm.nih.gov/datasets/)

Gene datasets are available through our API service, command-line tool, and our newest web tool, gene [data tables](https://www.ncbi.nlm.nih.gov/datasets/tables/genes/), for all genes available in NCBI Gene (currently >30k organisms).

The NCBI Datasets [github repository](https://github.com/ncbi/datasets/) contains documentation for the python library that interacts with the API and works with the downloaded data.  To familiarize yourself with the library functions, we suggest exploring the [Jupyter Notebooks](https://github.com/ncbi/datasets/tree/master/examples/jupyter) in this repository.

Gathering user feedback is central to this project, so we encourage you to [send us any questions, comments, or ideas](mailto:info@ncbi.nlm.nih.gov) that could help us to better serve your research needs.

**NOTE:** We recommend avoiding the use of this package in an environment also containing the following packages:
```bash
google-cloud-bigquery
pandas-gbq
```
We have observed some dependency issues with these packages such that some environments may fail to interact properly with python BigQuery API endpoints.


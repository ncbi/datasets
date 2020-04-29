# Welcome to NCBI Datasets

NCBI Datasets is a new resource that lets you easily gather data from across NCBI databases.

NCBI Datasets is also a work-in-progress that uses an iterative software development approach and extensive user research and interviews, with the goal of helping users to find the data they want more quickly and easily, in file formats most useful for common bioinformatics workflows.

In our first release, you can use this python package with our RESTful API to get data on genomes across all domains of life. We also offer a [command-line tool](https://www.ncbi.nlm.nih.gov/datasets/docs/command-line-start/), with access to the same broad taxonomic scope, and a [web interface that is focused on eukaryotic genomes.](https://www.ncbi.nlm.nih.gov/datasets/) At this time, gene data is available through our API service and command-line tool for a select group of organisms, including human, mouse, cattle, dog, rat, chimp and pig.

This repository contains documentation for the python library that can be used to interact with the API and work with the downloaded data. To faimiliarize with the library functions, we suggest exploring the [Jupyter Notebooks](examples/jupyter) in this repository. 

Gathering user feedback is central to this project, so we encourage you to [send us any questions, comments, or ideas](mailto:info@ncbi.nlm.nih.gov) that could help us to better serve your research needs.

**NOTE:** We recommend avoiding the use of this package in an environment also containing the following packages:
```bash
google-cloud-bigquery
pandas-gbq
```
We have observed some dependency issues with these packages such that some environments may fail to interact properly with python BigQuery API endpoints.
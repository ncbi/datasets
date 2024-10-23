# NCBI Datasets (ASM-NGS - 10/16/2024)


1. [NCBI Datasets: an overview](https://github.com/ncbi/datasets/new/master/training#1-ncbi-datasets-an-overview)
2. [Retrieving bacterial data and metadata using datasets](https://github.com/ncbi/datasets/new/master/training#retrieving-bacterial-data-and-metadata-using-datasets)
3. [Retrieving Virus information using NCBI Datasets](https://github.com/ncbi/datasets/new/master/training#3-retrieving-virus-information-using-ncbi-datasets)
4. [Important resources](https://github.com/ncbi/datasets/new/master/training#4-important-resources)



## 1. NCBI Datasets: an overview

[NCBI Datasets](https://www.ncbi.nlm.nih.gov/datasets/) is a resource that allows users to download data and metadata from [API](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/rest-api/), [web](https://www.ncbi.nlm.nih.gov/datasets/) and [command-line tool](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/). In this workshop, we will be focusing on the command-line tool, its structure and organization.

![getting_started](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/datasets_getting_started.png)

### 1A. Command-line tools: *datasets* and *dataformat*

While the web interface is helpful, there are times when it's more convenient to access genomes through a command-line environment. For example, let's say you are working on your institution's high-performance computing (HPC) system and you need to download dozens (or hundreds of genomes). Even if you're using the Datasets web interface, this would potentially be a two step process: 

1. Download the genome data package locally;
2. Transfer the files to the HPC system.

With the NCBI Datasets command-line interface (CLI), you can do this process in a single step. Our CLI allows users to access not only genomes, but also genes, ortholog sets and virus genomes. 

The program follows a hierarchy that makes it easier for users to select exact which options they would like to use. In addition to the program commands, additional flags are available for filtering the results. We will go over those during this tutorial.

![datasets-schema](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/datasets_schema_taxonomy.svg)

In this virtual machine, we have the necessary tools installed for you to explore NCBI Datasets without the need to configure anything. When you decide to use NCBI Datasets on your own machine or HPC system, you need to install it. More information on how to install NCBI Datasets can be found in [our documentation page](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/).

The NCBI Datasets CLI command structure is very intuitive. If you take a look at the diagram below, you will notice that the commands are built by choosing one option from each vertical rectangle. Let's start!

![datasets-command](https://github.com/user-attachments/assets/212db0a8-2e1e-44ec-ac60-d0660064dd22)

In addition to *datasets*, we also have *dataformat*, a companion tool to explore and convert metadata to TSV or Excel formats. We will cover the *dataformat* command syntax and use in the metadata section.

Bacteria is the taxonomic group with the largest number of available genomes - no surprises here. Currently there are 2.24 million bacterial genomes (on 10/09/2024), with *Salmonella* being the most represented taxon (>558,000 genome sequences).  

As expected, retrieving such a large number of genome sequences is a challenging task. In this workshop, we will show some strategies to potentially improve your download experience and data retrieval from NCBI using *datasets*. 


## Retrieving bacterial data and metadata using datasets

### 2A. Downloading a "small" genome data package
 
Small and large are very relative terms. When thinking about genome projects, that answer also can vary depending on the taxon in question. For vertebrates, genome sizes varies tremendously (more than 200 times), from around 400 million base pairs (Mb) in the fugu (*Tetraodon nigroviridis*) to almost 90 billion base pairs (Gb) in the South American Lungfish (*Lepidosiren paradoxa*). In Bacteria, genome sizes also vary significantly, from around 0.1 Mb in Candidatus *Nasuia deltocephalinicola* to 15 Mb in *Streptomyces prasinosporus*. 

<div style="text-align: center;">
<img src="https://tom-e-white.com/datavision/assets/img/05-genome-size.png" alt="genome-size" style="width: 700px;">
<figcaption>Image source: Tom E. White <a href="https://tom-e-white.com/datavision/05-genome-size.html">Datavision 2020</a></figcaption>
</div>
                                                       
<p>&nbsp;</p>

Back to the idea of small and big genome projects: if we consider genome size, a project that aims to evaluate 1000 human genomes will involve roughly 3,100 Gb, while the same concept (1,000 genomes) for an average bacterial genome (around 5 Mb) will entail only 5 Gb, which is 620 times less data. In addition to the storage space, we should also consider the network over which the download is happening: if your network is not very fast or reliable, you would benefit from using an option that allows for downloads to be resumed and that can handle slower speeds.  

For our small example, let's look into *Salmonella bongori* (88 genomes). Using *datasets*, you can download the genomes by NCBI TaxID, scientific name and common name (when available). In this case, we're using taxid.   


**Download S. bongori genomes by taxid:**

```
datasets download genome taxon 54736 --filename 54736.zip
```
**Unzip the package and inspect the folder structure and contents**

```
unzip 54736.zip -d sbongori
```
**Check folder contents using the tree command**

```
tree sbongori/
```

### 2B. Before we continue... What is a data package?

NCBI Datasets uses the concept of **data packages** for data retrieval. Data packages are compressed ZIP archives that contain both data and metadata. Each data endpoint (genome, gene, virus and taxonomy) have a set of files included by default plus a metadata report in [JSON-Lines](https://jsonlines.org/) format. You can read more about each type of data package and filetypes in our [Data Package reference page](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/reference-docs/data-packages/).


![data-package](https://github.com/user-attachments/assets/3ae28643-f5e6-4a28-852b-bdf5377b740d)

We always include at least one metadata file in all data packages. In the default genome data package, we have two reports:  

- `dataset_catalog.json`: lists all files included in the data package, grouped by accession. It also includes the uncompressed size of the entire package, filetypes and the API version.
- `assembly_data_report`: main metadata report. It has information on the assembly (stats, quality checks such as BUSCO, ANI and CheckM when applicable) biosample, annotation information (when available), owner/submitter, organism, and more.

If you are not familiar with JSON/JSON-Lines file formats, you can think of it as boxes inside boxes. The image below shows all the "boxes" in the genome assembly report based on the *dataformat* report schema.

<div style="text-align: center;">
<img src="https://github.com/user-attachments/assets/ff936e07-6c4d-47db-954b-97f31e4304a8" alt="dataformat-genome-schema">
</div>

### 2C. Let's take a look at the metadata report for one entry and look at quality scores

```
datasets summary genome taxon "salmonella bongori"  --reference | jq .
```

If we want to generate a table with only the CheckM quality scores for all samples available for that species, we could use *dataformat*. *dataformat* has specific schemas for each type of metadata reports, and can be used to convert JSON-Lines to TSV or excel formats.

![dataformat-choice](https://github.com/user-attachments/assets/c2696699-8b78-496d-8d83-1d1e87d34123)

Let's take a look at the help menu:

```
dataformat tsv genome --help
```

Now let's generate a TSV of all *Salmonella bongori* genomes with the following information:  
- Accession
- Best ANI match
- Scaffold N50
- CheckM completeness
- CheckM contamination

```
datasets summary genome taxon "salmonella bongori" --as-json-lines | \
dataformat tsv genome \
--fields accession,ani-best-ani-match-organism,assmstats-scaffold-n50,checkm-completeness,checkm-contamination | column -ts $'\t'

Assembly Accession  ANI Best ANI match Organism  Assembly Stats Scaffold N50  CheckM completeness  CheckM contamination
GCF_002035285.1     Salmonella bongori                                        99.61                1.29
GCF_002035475.1     Salmonella bongori                                        99.48                0.82
GCF_002119365.1     Salmonella bongori                                        99.5                 1.04
GCF_003448395.1     Salmonella bongori           112432                       99.33                0.87
GCF_003497195.1     Salmonella bongori           116904                       97.76                1.08
GCF_003521435.1     Salmonella bongori           103793                       96.43                0.78
GCF_003521525.1     Salmonella bongori           196552                       98.99                0.64
GCF_003522735.1     Salmonella bongori           358514                       98.59                0.64
GCF_006051015.1     Salmonella bongori           4476215                      98.5                 0.93
GCF_006051555.1     Salmonella bongori           351113                       99.61                0.88
GCF_006113225.1     Salmonella bongori           4487548                      99.53                0.71
GCF_007019345.1     Salmonella bongori           369720                       99.5                 2.34
GCF_016653555.1     Salmonella bongori           4459441                      99.61                0.64
GCF_019356415.1     Salmonella bongori           4440439                      99.27                0.77
GCF_900635485.1     Salmonella bongori           4466174                      42.81                0.58
...

```

You can save this output and look at the results in the program of your preference. To save the output, you would redirect it to a file, like this:

```
datasets summary genome taxon "salmonella bongori" --as-json-lines | \
dataformat tsv genome \
--fields accession,ani-best-ani-match-organism,assmstats-scaffold-n50,checkm-completeness,checkm-contamination > sbongori_stats.tsv
```

*datasets* also allows users to filter the data to be retrieved by using the available flags. Let's take a look at what's available and test some of them:

```
datasets summary genome taxon --help

Flags
      --tax-exact-match   Exclude sub-species when a species-level taxon is specified


Global Flags
      --annotated                 Limit to annotated genomes
      --api-key string            Specify an NCBI API key
      --as-json-lines             Output results in JSON Lines format
      --assembly-level string     Limit to genomes at one or more assembly levels (comma-separated):
                                    * chromosome
                                    * complete
                                    * contig
                                    * scaffold
                                     (default "[]")
      --assembly-source string    Limit to 'RefSeq' (GCF_) or 'GenBank' (GCA_) genomes (default "all")
      --assembly-version string   Limit to 'latest' assembly accession version or include 'all' (latest + previous versions)
      --debug                     Emit debugging info
      --exclude-atypical          Exclude atypical assemblies
      --exclude-multi-isolate     Exclude assemblies from multi-isolate projects
      --from-type                 Only return records with type material
      --help                      Print detailed help about a datasets command
      --limit string              Limit the number of genome summaries returned
                                    * all:      returns all matching genome summaries
                                    * a number: returns the specified number of matching genome summaries
                                       (default "all")
      --mag string                Limit to metagenome assembled genomes (only) or remove them from the results (exclude) (default "all")
      --reference                 Limit to reference genomes
      --released-after string     Limit to genomes released on or after a specified date (free format, ISO 8601 YYYY-MM-DD recommended)
      --released-before string    Limit to genomes released on or before a specified date (free format, ISO 8601 YYYY-MM-DD recommended)
      --report string             Choose the output type:
                                    * genome:   Retrieve the primary genome report
                                    * sequence: Retrieve the sequence report
                                    * ids_only: Retrieve only the genome identifiers
                                     (default "genome")
      --search strings            Limit results to genomes with specified text in the searchable fields:
                                  species and infraspecies, assembly name and submitter.
                                  To search multiple strings, use the flag multiple times.
      --version                   Print version of datasets

```

So, let's get the accession numbers of Salmonella samples released on or after 10/04/2024 and before or on 10/06/2024. And let's pipe the output to `jq` to make it easier to read it.

```
datasets summary genome taxon salmonella --released-after 2024-10-04 --released-before 2024-10-06 --report ids_only | jq .

```
---

**Now it's your turn**
Using the `summary` command, try another flag available for the genome subcommand. You could increase the data range, while filtering by annotated genomes only, for example.

---


### 2D. How about a big download?

Now let's say that you actually need to download all *Salmonella* genome sequences, which amounts to more than half million sequences. Or maybe something smaller, like *Salmonella enterica* subsp. *diarizonae*, which has around 800 genomes. You *can* download it directly, but that's not our recommendation. The *datasets* CLI has the option of downloading *dehydrated* genome packages.  

A dehydrated package doesn't include any data. It has a file (`fetch.txt`) that holds the location information of the requested data files. To retrieve those files, the option *rehydrate* is invoked on the CLI and the files are retrieved. 

**Advantages of rehydration:**

* It's faster than a regular download (it doesn't look like much in this case, but when you're dealing with 500x more sequences, that makes a difference)

|                                   | Salmonella enterica subsp. diarizonae (844 genomes, taxid 59204) |
|-----------------------------------|:-----------------------------------------------------:|
| Regular download                  | ~ 4 minutes                                           |
| Dehydrated download + rehydration | 7 seconds + 90 seconds                                |


* It allows users to download the files as gzip during rehydration, thus taking up less storage space;
* It can be resumed (instead of having to restart in case of failure);
* The `fetch.txt` file can be shared, so you can potentially share your data with other collaborators without having to put hundreds of gigabytes of data on a sharing platform;


**Download a dehydrated package with *Salmonella enterica* subsp. *diarizonae* (taxid 59204) genome and GFF3 files**

```
datasets download genome taxon 59204 --dehydrated --include genome,gff3 --filename 59204-dehydrated.zip
```

Let's unzip and explore the package contents, and compare it to the previous package we downloaded before.


```
unzip 59204-dehydrated.zip -d 59204

Archive:  59204-dehydrated.zip
  inflating: 59204/README.md         
  inflating: 59204/ncbi_dataset/data/assembly_data_report.jsonl  
  inflating: 59204/ncbi_dataset/fetch.txt  
  inflating: 59204/ncbi_dataset/data/dataset_catalog.json  
  inflating: 59204/md5sum.txt        
 
```

As we can see, instead of all the files in their respective folders, we have a file called `fetch.txt`. Let's take a quick look at it:

Now, let's take a look at the `fetch.txt` file inside the `ncbi_dataset` folder:

```
head 59204/ncbi_dataset/fetch.txt 

https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DcyN9M3tzQF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1kvLS9RLrzJgtGAEAIMhH_A	0	data/GCF_001276795.1/GCF_001276795.1_ASM127679v1_genomic.fna
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DcyN9M3tzQF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1ktPS9NLrzJgtGAEAIMLH-4	0	data/GCF_001276795.1/genomic.gff
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DcyN9O3MDcF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1kvLS9RLrzJgtGAEAIIvH-s	0	data/GCF_001276875.1/GCF_001276875.1_ASM127687v1_genomic.fna
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DcyN9O3MDcF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1ktPS9NLrzJgtGAEAIIZH-k	0	data/GCF_001276875.1/genomic.gff
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DczstA3NzUF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1kvLS9RLrzJgtGAEAIBWH-E	0	data/GCF_001628755.1/GCF_001628755.1_ASM162875v1_genomic.fna
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DczstA3NzUF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1ktPS9NLrzJgtGAEAIBAH98	0	data/GCF_001628755.1/genomic.gff
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DczstQ3NzUF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1kvLS9RLrzJgtGAEAIFYH-Y	0	data/GCF_001629755.1/GCF_001629755.1_ASM162975v1_genomic.fna
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DczstQ3NzUF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1ktPS9NLrzJgtGAEAIFCH-Q	0	data/GCF_001629755.1/genomic.gff
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DczstQ3NzcF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1kvLS9RLrzJgtGAEAINGH_A	0	data/GCF_001629775.1/GCF_001629775.1_ASM162977v1_genomic.fna
https://api.ncbi.nlm.nih.gov/datasets/fetch_h/R2V0UmVtb3RlRGF0YWZpbGU/eNqTyufKzUtOytROKymw0tdPT83Lz00t1k_MydF3d3bTNzAw1DczstQ3NzcF8eOBfCAXyNMzjHcM9oVwygzxycWDzcxM1ktPS9NLrzJgtGAEAIMwH-4	0	data/GCF_001629775.1/genomic.gff
```

You will notice that this file has three columns:

1. NCBI Datasets API path 
2. 0 column (CHECK THE MEANING)  
3. Final file name and location to download it  

Let's take a closer look at the third column (using the command `cut -f3`):

```
head ncbi_dataset/fetch.txt | cut -f3

data/GCF_001276795.1/GCF_001276795.1_ASM127679v1_genomic.fna
data/GCF_001276795.1/genomic.gff
data/GCF_001276875.1/GCF_001276875.1_ASM127687v1_genomic.fna
data/GCF_001276875.1/genomic.gff
data/GCF_001628755.1/GCF_001628755.1_ASM162875v1_genomic.fna
data/GCF_001628755.1/genomic.gff
data/GCF_001629755.1/GCF_001629755.1_ASM162975v1_genomic.fna
data/GCF_001629755.1/genomic.gff
data/GCF_001629775.1/GCF_001629775.1_ASM162977v1_genomic.fna
data/GCF_001629775.1/genomic.gff

```
Here we can see:  
- Where to download the file
- Files to be downloaded

During rehydration, users can filter which files they want to download by string matching. Let's take a look at the `rehydrate` subcommand help menu:

```
datasets rehydrate --help

Download data files for an unzipped, dehydrated genome data package. Data files specified in fetch.txt will be downloaded from NCBI. Read more about how rehydration can help with large genome downloads: https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/

Usage
  datasets rehydrate [flags] --directory <directory_name>

Flags
      --directory string   Specify the directory containing the unzipped dehydrated bag
      --gzip               rehydrate files to gzip format
      --list               List files that would be downloaded during rehydration
      --match string       Specify substring that matches files for rehydration
      --max-workers int    Limit the maximum number of concurrent download workers (allowed range is 1-30) (default 10)
      --no-progressbar     Hide progress bar


Global Flags
      --api-key string   Specify an NCBI API key
      --debug            Emit debugging info
      --help             Print detailed help about a datasets command
      --version          Print version of datasets
```
So, let's say that first we want to download only the GFF3 files. Here's how we would build this command: 

```
datasets		calls the datasets program
rehydrate		calls the rehydrate subcommand
--directory	specifies the directory where the folder ncbi_dataset is.
--match		matches the name of the files to be rehYDrated.
--list			shows which files WILL be downloaded
| head			pipe the output to the command head, which prints the first ten lines.

```
Let's put it all together now:  

```
datasets rehydrate --directory 59204 --match gff --list | head

data/GCF_001276795.1/genomic.gff
data/GCF_001276875.1/genomic.gff
data/GCF_001628755.1/genomic.gff
data/GCF_001629755.1/genomic.gff
data/GCF_001629775.1/genomic.gff
data/GCF_002794415.1/genomic.gff
data/GCF_003872565.1/genomic.gff
data/GCF_008692785.1/genomic.gff
data/GCF_016030095.1/genomic.gff
data/GCF_019222725.1/genomic.gff
```

---


**Try it yourself:**
 
-  what would happen if you ran the same command without the `--match` flag?   
-  How can you download only the genomic FASTA files and none of the GFF3?

---

In addition to the genome service (covered in the previous session), users can retrieve viral genome sequences and metadata using the Virus service from NCBI Datasets CLI. 

## 3. Retrieving Virus information using NCBI Datasets

### 3A. What's the difference between the Virus and Genome endpoints?

The data available through these endpoints originate from two distinct sources. Virus data is sourced from NCBI Virus, which employs both manual and automated curation processes to verify viral sequence data provided by the International Nucleotide Sequence Database Collaboration (INSDC) and standardized metadata. In contrast, the NCBI Datasets genome service, provides access to a subset of virus sequences from NCBI Virus that have been assembled and assigned an NCBI Assembly accession (GCA\_/GCF\_). 

**What does it mean for those working on Viruses?**  

- Virus data from the Virus endpoint has an additional layer of curation
- Data packages downloaded from Virus will likely be different (larger) from those downloaded from the Genome: 
   -  For SARS-CoV-2, there are currently (10/15/2024) 8,972,850 genomes available through the Virus service, and only 93 through the genome service. 

      <details>

        <summary><i>Try it yourself</i></summary>

      ---
      * Check the number of SARS-CoV-2 sequences in the Virus endpoint:
      ```
      datasets summary virus genome taxon sars2 --limit 0
      ```

      * Now check the number of sequences in the genome endpoint:
      ```
      datasets summary genome taxon sars2 --limit 0
      ```
      ---

      </details>


### 3B. Special virus cases: cache packages for SARS-CoV-2 and Influenza

For both SARS-CoV-2 and (Alpha)Influenza, NCBI Datasets CLI provides a cache package. A cache package is pre-packaged with all genomes available for those two taxa. 

An cache package is the equivalent to a grab-n-go sandwich versus a made-to-order one (regular packages). It's faster to download a cache package because it doesn't need to be assembled, but it also travels through faster download channels at NCBI.


### 3C. Retrieving genome information for Dengue virus

In this exercise, we will take a look at the genomes available for the Dengue virus. 

- Let's take a look at the options for the virus genome subcommand:

```
datasets download virus genome taxon --help

Download a virus genome data package by taxon (NCBI Taxonomy ID, scientific or common name for any virus at any tax rank). Virus genome data packages include genome, transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

The default virus genome data package includes the following files:
  * genomic.fna (genomic sequences)
  * data_report.jsonl (data report with virus genome metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)

Usage
  datasets download virus genome taxon  <taxon> [flags]

Sample Commands
  datasets download virus genome taxon sars-cov-2 --host dog --include protein
  datasets download virus genome taxon coronaviridae --host "manis javanica"

Flags
      --fast-zip-validation   Skip zip checksum validation after download


Global Flags
      --annotated                 Limit to annotated genomes
      --api-key string            Specify an NCBI API key
      --complete-only             Limit to complete genomes
      --debug                     Emit debugging info
      --filename string           Specify a custom file name for the downloaded data package (default "ncbi_dataset.zip")
      --geo-location string       Limit to genomes isolated from a specified geographic location (continent or country)
      --help                      Print detailed help about a datasets command
      --host string               Limit to virus genomes isolated from a specified host species
      --include string(,string)   Specify virus genome sequence types to download
                                    * genome:     genomic sequences
                                    * cds:        nucleotide coding sequences
                                    * protein:    amino acid sequences
                                    * annotation: annotation report
                                    * biosample:  biosample report
                                    * none:       no sequence data, only primary data report
                                       (default [genome])
      --lineage string            Limit results by Pango lineage (only SARS-CoV-2)
      --no-progressbar            Hide progress bar
      --refseq                    Limit to RefSeq genomes
      --released-after string     Limit to genomes released on or after a specified date (free format, ISO 8601 YYYY-MM-DD recommended)
      --updated-after string      Limit to genomes updated on or after a specified date (free format, ISO 8601 YYYY-MM-DD recommended)
      --usa-state string          Limit to genomes isolated from a specified U.S. state (two-letter abbreviation)
      --version                   Print version of datasets
```

#### 3C1. Downloading genomes

- Download all genomes

```
datasets download virus genome taxon 12637 --filename dengue-all.zip
```

- Download reference genomes (4 genomes, Dengue virus 1-4)

```
datasets download virus genome taxon 12637 --refseq --filename dengue-all-ref.zip
```

#### 3C2. Filtering download based on metadata information

- Look at the first record and all the fields with `jq`

```
datasets summary virus genome taxon 12637 --limit 1 | jq
```

- Use dataformat to pull metadata information. Let's look at unique entries for the geo-location field. 
   - `datasets` will retrieve the metadata;
   - `dataformat` will pull the information from the metadata field `geo-location`
   - `sort` will sort all the `geo-location` entries in alphabetical order
   - `uniq -c` will count the number of each unique entry

```
datasets summary virus genome taxon 12637 --as-json-lines | \
dataformat tsv virus-genome --fields geo-location | sort | uniq -c

```


- Let's look at all genomes filtered by geo-location (Brazil)

```
datasets summary virus genome taxon 12637 --geo-location Brazil | jq .total_count
3721
```

- We have a new field to filter by US state that's separate from the `--geo-location` flag. Let's take a look at how many genomes we have from Florida using the `summary` subcommand and `jq`:

```
datasets summary virus genome taxon 12637 --usa-state FL | jq .total_count
526
```

---

## 4. Important resources

- ASM-NGS Wiki (for the CLI tutorial): [https://github.com/ncbi/workshop-asm-ngs-2024/wiki#2-datasets](https://github.com/ncbi/workshop-asm-ngs-2024/wiki#2-datasets)


##### NCBI Datasets main resources

- NCBI Datasets homepage: [https://www.ncbi.nlm.nih.gov/datasets/](https://www.ncbi.nlm.nih.gov/datasets/)
- Github: [https://github.com/ncbi/datasets](https://github.com/ncbi/datasets)


##### Download and installation instructions (CLI)

- Instructions:  
 [https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/)
  
#### Tutorials, how-to guides and past workshops
 
- How-to guides (short, one-line CLI tasks):   
[https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/)
- Tutorials (multi-task, longer tutorials, mostly based on feedback or questions we get from users): [https://www.ncbi.nlm.nih.gov/datasets/docs/v2/tutorials/](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/tutorials/)
- Past training sessions and workshops (Jupyter notebooks used in previous *datasets* training events): [https://github.com/ncbi/datasets/tree/master/training](https://github.com/ncbi/datasets/tree/master/training)

##### How to get help

- Email the helpdesk: [info@ncbi.nlm.nih.gov](mailto:info@ncbi.nlm.nih.gov)
- Github: [https://github.com/ncbi/datasets](https://github.com/ncbi/datasets)
- Yellow feedback button on our pages 


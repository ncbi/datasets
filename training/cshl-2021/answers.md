# NCBI Datasets - CSHL (11/02/2021): answers

## Tutorial - I: Accessing genomes

### A little bit more about json files
A JSON (JavaScript Object Notation) file stores data structures and objects. In a very simplified (and non-technical) way, a JSON file is a box, that might contain other boxes with more boxes inside. In `datasets summary genome` our JSON "box" is organized like this:
![](https://hackmd.io/_uploads/HyQS80YUY.png)


### Exercises

Now we will practice what we learned about `datasets`. Take a look at the questions below and feel free to ask questions. Useful resources for this exercise are the `--help` from the command line and the [jq cheatsheet](). 



```bash
%%bash
# How many reference genomes in the family Formicidae? (hint --reference)
datasets summary genome taxon formicidae --reference | jq '.total_count'

```


```bash
%%bash
# How many reference genomes are annotated? (hint: --annotated)
datasets summary genome taxon formicidae --reference --annotated | jq '.total_count'

```


```bash
%%bash
# How many genomes have NCBI (RefSeq) annotations? (hint: --assembly-source)
datasets summary genome taxon formicidae --reference --annotated --assembly-source refseq | jq '.total_count'

```


```bash
%%bash
# For the question above, they might also find a different answer if they don't use the --reference flag. 
# Nuala will cover the genbank/reference/refseq in the next section
datasets summary genome taxon formicidae --annotated --assembly-source refseq | jq '.total_count'
```

### Bonus questions:


```bash
%%bash
## Take a look at the jq cheat sheet (link here) and try to build a jq query for the metadata
datasets summary genome taxon formicidae | jq -r '.assemblies[].assembly 
| [.bioproject_lineages[].bioprojects[].accession, .org.sci_name] | @tsv'

```


```bash
%%bash
# Now look at the summary metadata for your organism of interest 
# (if you don't have a favorite, go with bears, Ursidae - taxid: 9632)
datasets summary genome taxon bears

```


```bash
%%bash
# How many genomes? (if you used scientific or common name in the search above, try with taxid now)
datasets summary genome taxon 9632 | jq '.total_count'

```


```bash
%%bash
# Assembly level breakdown
datasets summary genome taxon 9632 | jq '.assemblies[].assembly.assembly_level' | sort | uniq -c

```


```bash
%%bash
# How many with contig N50 ?
datasets summary genome taxon 9632 | jq '.assemblies[].assembly | select(.contig_n50 > 100000) | (.assembly_accession)' | wc -l

```

## Tutorial - II: Accessing genes
### GENES


When choosing any of those three options, you will retrieve the gene information for the **reference** taxon. Like this:
![](https://hackmd.io/_uploads/rydJGkkDF.png)

Now let's take a look at a gene example:

### Back to ants
We will download the gene *orco* for the species *Acromyrmex echinatior*. We will use the gene-id 105147775 instead of the symbol.
The reason for it is that sometimes even when a known gene is characterized in a species, the gene symbol is not necessarily propagated.

Now we are going to take advantage of the fact that we are using a Jupyter Notebook and use the package `pandas` to look at the gene data table

### Exercises

1. Look for the summary data for a gene of interest (check the [etherpad](https://etherpad.wikimedia.org/p/CSHL_Datasets_Workshop_2021) for suggestions)
2. What is the gene location?
3. What is the gene range?
4. Now, download a list of gene symbols using the file genes.txt (provided). Save it as gene_list.zip
5. Unzip gene_list.zip and explore the folder structure
6. How many fasta files?


```bash
%%bash
# Summary data
datasets summary gene symbol p53

```


```bash
%%bash
# Gene location
datasets summary gene symbol p53 | jq '.genes[].gene .chromosomes[]'

```


```bash
%%bash
# Gene range
datasets summary gene symbol p53 | jq '.genes[].gene.genomic_ranges[].range[] |(.begin, .end)'

```


```bash
%%bash
# Download a list of genes and save the data package as gene_list.zip (--filename gene_list.zip)
datasets download gene symbol --inputfile genes.txt --filename gene_list.zip --no-progressbar

```


```bash
%%bash
# Explore the folder structure
unzip gene_list.zip -d gene_list
tree gene_list
```


```bash
%%bash
# How many genes were downloaded?
cat genes.txt

```

How many fasta files in the data package?



## Tutorial - III: Accessing orthologs

![](https://hackmd.io/_uploads/rkr2-J1DF.png)

## Tutorial - V: Downloading large datasets (dehydration/rehydration) and `dataformat`

Now you learned how to download genomes, genes and ortholog gene sets from NCBI with one command using `datasets`. Now we want to show you another feature of `datasets` that allows you to download what we call a `dehydrated` package. Let's download a dehydrated package and explore the files inside it.

## Exercise
* Download a dehydrated package for all *Mycobacterium tuberculosis* genomes that meet all of the following criteria (hint: use flags)
    1. submitted/released in 2021
    2. annotated
    3. assembly level of complete_genome
* use dataformat to view the sequencing technology used for each of these genomes
* use rehydrate to get the genome sequence for one genome generated using Oxford Nanopore


```bash
%%bash
datasets download genome -h
```


```bash
%%bash
# Download a dehydrated genome data package
datasets download genome taxon "Mycobacterium tuberculosis" \
--dehydrated --annotated \
--assembly-level complete_genome \
--released-since 01/01/2021 \
--filename mycobacterium-dry.zip --no-progressbar

```


```bash
%%bash
# Unzip the data package
unzip mycobacterium-dry.zip -d mycobacterium-dry


```


```bash
%%bash
# Use dataformat to generate a table that includes sequencing technology
dataformat tsv genome --fields assminfo-accession,assminfo-sequencing-tech,assminfo-submission-date \
--package mycobacterium-dry.zip | column -t -s$'\t'

```


```bash
%%bash
# Use rehydrate to get genome sequence generated using Oxford Nanopore
datasets rehydrate --directory mycobacterium-dry/ --list

```


```bash
%%bash
# Use rehydrate to get genome sequence generated using Oxford Nanopore
datasets rehydrate --directory mycobacterium-dry/ --match GCA_019670745.1/GCA_019670745.1_ASM1967074v1_genomic.fna

```

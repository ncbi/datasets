# This script depends on having R_USER_LIBS set properly, or the user has write permissions to install packages.
local({r <- getOption("repos")
       r["CRAN"] <- "http://cran.r-project.org"
       options(repos=r)
})
if (!require("pacman")) install.packages("pacman")
pacman::p_load(jsonlite, httr, caTools)

if (!require(ncbi.dataset)) {
    install.packages("https://ftp.ncbi.nlm.nih.gov/pub/datasets/r_client_lib/ncbi.datasets_LATEST.tar.gz", repos = NULL)
    library(ncbi.datasets)
}

print("Demo 1: Gene API, GeneMetadataById")
api.gene_instance <- GeneApi$new()
result_gene <- api.gene_instance$GeneMetadataById(
    '56,2',
    returned.content='COMPLETE',
    sort.schema.field='SORT_FIELD_GENE_ID'
)
prettify(result_gene$toJSONString())
for (gene_match in result_gene$genes)  {
    cat(gene_match$gene$gene_id, " - ", gene_match$gene$symbol, "\n")
}


print("Demo 2: Genome API, AssemblyDescriptorsByAccessions")
api.genome_instance <- GenomeApi$new()
result_genome <- api.genome_instance$AssemblyDescriptorsByAccessions('GCF_000001405.39,GCA_007922845.1')
prettify(result_genome$toJSONString())


print("Demo 3: Virus API, Sars2ProteinSummary")
api.virus_instance <- VirusApi$new()
result_genome <- api.virus_instance$Sars2ProteinSummary(
    'spike protein,E',
    released.since='2020-11-01T00:00:00.000Z',
    host='human',
    geo.location='French')
prettify(result_genome$toJSONString())

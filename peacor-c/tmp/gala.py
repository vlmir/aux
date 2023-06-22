#! /usr/bin/python
import sys
import os.path
import getopt
import networkx as nx
import pprint

pth2src = sys.argv[1]

def xpcr(pth2src, dlm='\t'):
    dat = {}
    with open(pth2src) as fp:
        for line in fp:
            line = line.rstrip()
            if not line:
                continue
            (lblA, lblB, pcr) = line.split(dlm)
            if lblA not in dat.keys():
                dat[lblA] = {}
            if lblB not in dat[lblA].keys():
                dat[lblA][lblB] = {}
            dat[lblA][lblB] = pcr

G = nx.readwrite.edgelist.read_weighted_edgelist(pth2src)
print(G.number_of_edges())
data = nx.adjacency_data(G)
#pprint.pprint(data)
wtk = 'weight'
#print(G.degree['gene1'])
#print(list(G.degree(['gene1'], wtk)))
#print(dict(G.degree(['gene1'], wtk)))
#print(tuple(G.degree(['gene1'], wtk)))
#print(set(G.degree(['gene1'], wtk)))

N = list(G.nodes)
D = dict(G.degree(N, wtk))
C = nx.clustering(G, N, wtk)
mxitr = 100
tol = 1.0e-6
nstart = None
E = nx.eigenvector_centrality(G, mxitr, tol, nstart, wtk)
for nd in N:
    print('%s\t%.3f\t%.3f\t%.3f' % (nd, D[nd], C[nd], E[nd]))
#    print(G.degree([nd], wtk)[nd])
#    print(D[nd]) # the same as above
#    #print(nx.clustering(G, [nd], wtk)[nd])
#    print(C[nd]) # the same as above
#    print(E[nd])

for nd, nbrs in G.adj.items():
    for nbr, eattr in nbrs.items():
        wt = eattr[wtk]
#        if wt < 0.0: print('(%s, %s, %.3f)' % (nd, nbr, wt))

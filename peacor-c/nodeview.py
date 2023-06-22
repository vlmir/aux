#! /usr/bin/python
from __future__ import print_function
import sys
import math
import getopt
import networkx as nx
import pprint

def eprint(*args, **kwargs):
    print(*args, file=sys.stderr, **kwargs)

def usage():
    print('Usage:')
    print(sys.argv[0] + ' [-cdehv] [-p <float> | -w <float>] -i <string>')
    print('Options:')
    opts = (
    '-c: compute clustering',
    '-d: compute degree',
    '-e: compute eigen vector centrality',
    '-h: print help',
    '-i: input adjacency matrix',
    '-p: percent of top nodes to print: (0,100) exclusive',
    '-v: print version',
    '-w: node weight cutoff for printing: (0,1) exclusive',
    )
    for opt in opts:
        print(opt)
    print('Examples:')
    print(sys.argv[0] + ' -h')
    print(sys.argv[0] + ' -cde -i adjacency.tsv')
    print(sys.argv[0] + ' -cde -p 50 -i adjacency.tsv')

def main():
    nixopts = 'cdehvp:w:i:'
    version = 0.1
    ## flags
    ierr = 0
    cc = 0
    dd = 0
    ee = 0
    ii = 0
    hh = 0
    vv = 0
    pp = -1
    ww = -1

    try:
        opts, args = getopt.getopt(sys.argv[1:], nixopts)
    except getopt.GetoptError as err:
        print('ERROR:', err)
        ierr = 1
    if not opts:
        ierr = 1
    for opt, val in opts:
        if opt in ('-v', '--version'):
            print('Version: ' + str(version))
            sys.exit(1)
        if opt in ("-h", "--help"):
            ierr = 1
        if opt == '-c': cc = 1
        if opt == '-d': dd = 1
        if opt == '-e': ee = 1
        if opt == '-p': pp = float(val)
        if opt == '-w': ww = float(val)
        if opt in ('-i', '--input'): ii = val
    if not ii: ierr = 1
    if ierr:
        usage()
        sys.exit(1)
    if (cc+dd+ee == 0):
        print('ERROR: at least one of "-cde" to be provided')
        sys.exit(1)
    if (pp == 0.0 or pp == 100.0):
        print('ERROR: "-p" argument must be BETWEEN "0" and "100"')
        sys.exit(1)
    if (ww == 0.0 or ww == 1.0):
        print('ERROR: "-w" argument must be BETWEEN "0" and "1"')
        sys.exit(1)
    if pp > 0 and ww > 0:
        print('ERROR: "-pw" are mutually exclusive')
        sys.exit(1)

    pth2dat = ii
    wtk = 'weight'
    ndk = 'node'
    rank = {
    wtk: {},
    ndk: {},
    }
    G = nx.readwrite.edgelist.read_weighted_edgelist(pth2dat)
    N = list(G.nodes)
    nn = len(N)

    ## computing node weights
    for nd, nbrs in G.adj.items():
        ndwt = 0.0
        cnt = 0
        for nbr, eattr in nbrs.items():
            wt = abs(eattr[wtk])
            cnt += 1
## geometrical mean of abs values, values too low for large data sets
#        ndwt = 1.0
#            ndwt *= wt
#        ndwt = ndwt ** 1.0/cnt
            ndwt += wt
        ndwt = ndwt/cnt
        if ndwt not in rank[wtk].keys():
            rank[wtk][ndwt] = {}
        rank[wtk][ndwt][nd] = 1
        if nd not in rank[ndk].keys():
            rank[ndk][nd] = {}
        rank[ndk][nd][ndwt] = 1

    allwts = sorted(rank[wtk].keys())
    eprint('! min node weght: ' + str(allwts[0]) + ' max node weght: ' + str(allwts[-1]))
    eprint('! all nodes: ' + str(nn))

    ## pruning
    if pp > 0:
        ppns = []
        noff = int(round(nn - nn*pp/100))
        xwts = []
        cnt = 0
        for wt in sorted(rank[wtk].keys()):
            if cnt >= noff: break
            nds = rank[wtk][wt].keys()
            xwts.append(wt)
            cnt += len(nds)
        for nd in N:
            ndmx = max(rank[ndk][nd].keys())
            if ndmx not in xwts: ppns.append(nd)
        N = ppns

    if ww > 0:
        wwns = []
        for nd in N:
            ndwts = rank[ndk][nd]
            ndmx = max(ndwts.keys())
            if ndmx > ww: wwns.append(nd)
        N = wwns
    nn = len(N)
    eprint('! top nodes: ' + str(nn))

    ## computing
    if dd:
        D = dict(G.degree(N, wtk))
    if cc:
        C = nx.clustering(G, N, wtk) # works only for non-negative weights
        #C = nx.clustering(G, N) # the same value 1.0 for all nodes if G is complete
    if ee:
        # networkx defaults
        mxitr = 100
        tol = 1.0e-6
        nstart = None
        E = nx.eigenvector_centrality(G, mxitr, tol, nstart, wtk)

    ## exporting
    for nd in N:
        ndwts = sorted(rank[ndk][nd].keys())
        if len(ndwts) != 1:
            eprint('ERROR: node: ' + nd + 'weights: ', ndwts)
            sys.exit(1)
        ndwt = ndwts[0]
        if dd: dnd = D[nd]
        else: dnd = 0.0
        if cc: cnd = C[nd]
        else: cnd = 0.0
        if ee: end = E[nd]
        else: end = 0.0
        #fp.write('%s\t%.6f\t%.6f\t%.6f\n' % (nd, dnd, cnd, end))
        print('%s\t%.6f\t%.6f\t%.6f\t%.6f' % (nd, ndwt, dnd, cnd, end))
    #fp.close()
if __name__ == "__main__":
    main()

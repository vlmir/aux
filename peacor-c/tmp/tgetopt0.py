#! /usr/bin/python
# from; https://stackabuse.com/command-line-arguments-in-python/
import sys
import getopt
import pprint

nixopts = "ho:v"  
gnuopts = ["help", "output=", "verbose"]

def usage():
        print(sys.argv[0] + ' [options]')
def main():
    try:
        opts, rest = getopt.getopt(sys.argv[1:], nixopts, gnuopts)
    except getopt.GetoptError as err:
        # print help information and exit:
#    print (str(err))
        print(err) # will print something like "option -a not recognized"
        usage()
        sys.exit(2)
    output = None
    verbose = False
    for opt, val in opts:
        if opt in ("-v", "--verbose"):
            print ("enabling verbose mode")
            verbose = True
        elif opt in ("-h", "--help"):
            print ("displaying help")
            usage()
            sys.exit()
        elif opt in ("-o", "--output"):
            print (("enabling special output mode (%s)") % (val))
            output = val
        else:
            assert False, "unhandled option"
    # ...

if __name__ == "__main__":
    main()

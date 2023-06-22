CC=clang
CFLAGS=-std=c11 # 2017-08: the latest standard
CFLAGS+=-W -Wall
LDLIBS=-lm # libm - basic maths
LDLIBS+=-lgsl
LDLIBS+=-lgslcblas
mode=1
ifeq ($(mode), 0) # production mode 
CFLAGS+=-O3 # max optimization
endif

ifeq ($(mode), 1) # development mode 
CFLAGS+=-O0 -g -ggdb
CFLAGS+=-pg
LDFLAGS+=-pg
endif

CFLAGS+=-MMD  # generate dependency .d files
CFLAGS+=-I../include # custom 
job=peacor
input=var2obs
srcs=$(job).c ../ext/util.c
objs=$(srcs:.c=.o)
deps=$(objs:.o=.d)

$(job) : $(objs)
-include $(deps)
# %: %.o # all targerts for all objects
# obj/%.o: %.c # to be tested
trgs=$(job) # single target in our case 

### command line parameters:
C=5 # 
R=3 # 

.PHONY: run
run :
	./$(job) -a -c $C -r $R -i ./data/$(input).tsv 2> ./$(input).err

.PHONY: clean
clean:
	$(RM) $(objs)
	$(RM) $(deps)

### everything below is relevant only in dev mode

# cscope.out
cscope.out: $(srcs)
	cscope -f $@ -I$(CURDIR)/inc -bq $^

# ctags
tags: $(srcs)
	ctags -f $@ -R $(CURDIR) $^

# '-pg' required for both compilation and linking
# usually: gprof peacor gmon.out > gprof.out
gmon.out: $(job)
gprof.out: gmon.out
	gprof $(CURDIR)/$(job) > $@
.PHONY: prof
prof: gprof.out

.PHONY: more
more: cscope.out tags prof

.PHONY: cleanall
cleanall:
	$(RM) $(trgs)
	$(RM) $(objs)
	$(RM) $(deps)
	$(RM) $(job).err
	$(RM) cscope.out cscope.out.in cscope.out.po
	$(RM) tags
	$(RM) gmon.out
	$(RM) gprof.out

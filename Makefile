test-slice:
	@SLICES ?= $(shell nproc)
	@for i in $(shell seq 1 $$(SLICES)); do \
	  echo "▶ slice $$i/$(SLICES)"; \
	  SLICE_INDEX=$$i SLICES=$(SLICES) scripts/test_slice.sh; \
	done

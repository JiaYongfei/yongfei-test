include build/lib/common.mk
include build/lib/golang.mk
include build/lib/image.mk
include build/lib/deploy.mk
include build/lib/asset.mk
include build/lib/web.mk
include build/lib/gen.mk

## release: Release tke
.PHONY: release
release:
	build/release.sh

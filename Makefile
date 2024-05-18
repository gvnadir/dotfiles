LINUX_PACKAGES = nvim 

linux:
	for package in $(LINUX_PACKAGES); do \
		stow --verbose --target=$$HOME --restow $$package \
	done

delete:
	stow --verbose --target=$$HOME --delete */

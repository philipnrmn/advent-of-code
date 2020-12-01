.PHONY: today

DAY_OF_MONTH = $(shell date +%d)

today:
	cp -R _template "day-$(DAY_OF_MONTH)"
	sed -i "" "s/DAY_OF_MONTH/$(DAY_OF_MONTH)/" "day-$(DAY_OF_MONTH)/main.go"
	touch "day-$(DAY_OF_MONTH)/input"


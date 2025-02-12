package utils

const NOP = "#(nop) "

const RUN_PREFIX = "/bin/sh -c "

const FROM_INSTRUCTION = "FROM "
const RUN_INSTRUCTION = "RUN "
const CMD_INSTRUCTION = "CMD "
const LABEL_INSTRUCTION = "LABEL "
const MAINTAINER_INSTRUCTION = "MAINAINER "
const EXPOSE_INSTRUCTION = "EXPOSE "
const ENV_INSTRUCTION = "ENV "
const ADD_INSTRUCTION = "ADD "
const COPY_INSTRUCTION = "COPY "
const ENTRYPOINT_INSTRUCTION = "ENTRYPOINT "
const VOLUME_INSTRUCTION = "VOLUME "
const USER_INSTRUCTION = "USER "
const WORKDIR_INSTRUCTION = "WORKDIR "
const ARG_INSTRUCTION = "ARG "
const ONBUILD_INSTRUCTION = "ONBUILD "
const STOPSIGNAL_INSTRUCTION = "STOPSIGNAL "
const HEALTHCHECK_INSTRUCTION = "HEALTHCHECK "
const SHELL_INSTRUCTION = "SHELL "

type SourceType string

const (
	Image  SourceType = "IMAGE"
	Parent SourceType = "PARENT_IMAGE"
)

type Source struct {
	Name string
	Type SourceType
}

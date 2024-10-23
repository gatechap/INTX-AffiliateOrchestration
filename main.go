package main

import (
	"th.truecorp.it.dsm.intcom/affiliateorchestration/apirouter"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/config"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/errormsg"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/locallogging"
)

// @title [INTX] affiliateorchestration Swagger APIs
// @version 00.00
// @description [INTX] affiliateorchestration
// @BasePath /
// @schemes http
// @securityDefinitions.basic BasicAuth
// @in header
// @Username Authorization
// @Password Authorization
func main() {
	appConfig, errHanlerInfo := config.LoadConfig()

	if errHanlerInfo != nil && errHanlerInfo.Error != nil {
		// Set locallogging
		log := locallogging.LocalLogging{}
		errDesc := errHanlerInfo.Error.Error()
		errCode := errormsg.ERR_CD_INTERNAL_FAILURE
		msg := errormsg.GenInternalFailureMsg(errDesc)

		errLog := log.BuildErrorInputLoggerBeforeController(appConfig.Application.Profile, appConfig.Application.Name, errCode, msg, errHanlerInfo.Error, errHanlerInfo.ErrorApplication, errHanlerInfo.ErrorModule, errHanlerInfo.ErrorFile, errHanlerInfo.ErrorFunction)
		log.SetErrorInputLoggerBeforeController(errLog)
		log.WriteLogErrorBeforeController(errLog)
	}

	router := apirouter.SetupAPIRouter(appConfig)
	router.Run(appConfig.Server.Port)
}

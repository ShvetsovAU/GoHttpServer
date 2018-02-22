package router

import (
	"github.com/gorilla/context"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/shvetsovau/gohttpserver/middlewares"
	"github.com/shvetsovau/gohttpserver/handlers"
)

func GetRoutes() http.Handler {
	router := &router{httprouter.New()}

	// базовый уровень обработки запроса: перехват паники, логирование, проверка заголовка
	h0 := alice.New(context.ClearHandler,
		middlewares.RecoverHandler,
		middlewares.LoggingHandler,
		//middlewares.DetailedLoggingHandler,
		//middlewares.AcceptHandler,
	)

	//// h0 + проверка аутентификации
	//h1 := h0.Append(middlewares.AuthenticationHandler)
	//
	//// h1 + проверка Content-Type в заголовке
	//h2 := h1.Append(middlewares.ContentTypeHandler)

	router.Get("/", h0.ThenFunc(handlers.HomeRouterHandler))
	router.Get("/admin", h0.ThenFunc(handlers.AdminRouterHandler))

	router.Post("/limsFolder/create", h0.ThenFunc(handlers.CreateLimsFolderHandler))
	router.Put("/limsFolder/update", h0.ThenFunc(handlers.UpdateLimsFolderHandler))
	router.Delete("/limsFolder/delete/:id", h0.ThenFunc(handlers.DeleteLimsFolderHandler))

	//// Группы (объекты)
	//router.Get("/group/nodetails/:id", h1.ThenFunc(handlers.SelectGroupNoDetailsHandler)) // получить одну группу
	//router.Get("/groups", h1.ThenFunc(handlers.SelectAllGroupsHandler))                   // получить список групп
	//router.Post("/group", h1.ThenFunc(handlers.ModifyGroupHandler))                       // создать или редактировать существующую группу
	//router.Delete("/group/:id", h1.ThenFunc(handlers.DeleteGroupHandler))                 // удалить группу
	//
	//// Счета
	//router.Get("/bill/find/:id", h1.ThenFunc(handlers.GetBillHandler))
	//router.Get("/bills/all/:gid", h1.ThenFunc(handlers.GetAllBillsHandler))
	//router.Get("/bill/payment/:id", h1.ThenFunc(handlers.GetBillPaymentHandler))
	//router.Get("/bill/recalculate/:id", h1.ThenFunc(handlers.RecalculateBillPaymentHandler))
	//router.Post("/bill/create/:gid", h2.ThenFunc(handlers.CreateBillHandler))
	//router.Post("/bill/paydate/postpone/:id", h2.ThenFunc(handlers.PostponeBillPayDateHandler))
	//router.Put("/bill/update", h2.ThenFunc(handlers.UpdateBillHandler))
	//router.Delete("/bill/delete/:id", h1.ThenFunc(handlers.DeleteBillHandler))
	//
	//// Услуги
	//router.Get("/utilities/all/:bill_id", h1.ThenFunc(handlers.FindAllUtilitiesHandler))
	//router.Get("/utility/paramval/:bill_id/:src_id", h1.ThenFunc(handlers.GetUtilityParameterValueHandler))
	//router.Post("/utility/calculate", h2.ThenFunc(handlers.CalculateUtilityHandler))
	//router.Post("/utility/create/:bill_id", h2.ThenFunc(handlers.CreateUtilityHandler))
	//router.Put("/utility/update", h2.ThenFunc(handlers.UpdateUtilityHandler))
	//router.Delete("/utility/delete/:id", h1.ThenFunc(handlers.DeleteUtilityHandler))
	//
	//// Методики
	//router.Get("/methods/catalog", h1.ThenFunc(handlers.MethodsHandler))
	//
	//// Приборы
	//router.Get("/devices/all/:gid", h1.ThenFunc(handlers.GetAllDeviceHandler))
	//router.Get("/device/values/:id/:count", h1.ThenFunc(handlers.GetDeviceValuesHandler))
	//router.Post("/device/value/add", h1.ThenFunc(handlers.AddDeviceValueHandler))
	//
	//// Быстро добавить значение из формы со всеми приборами
	//router.Post("/device/value/add/quick", h1.ThenFunc(handlers.QuickAddDeviceValueHandler))
	//
	//router.Put("/device/value/update", h1.ThenFunc(handlers.UpdateDeviceValueHandler))
	//router.Put("/device/update", h1.ThenFunc(handlers.UpdateDeviceHandler)) // deprecated from 2.2.0 clients
	//router.Put("/device/upsert", h1.ThenFunc(handlers.UpsertDeviceHandler))
	//router.Delete("/device/delete/:id", h1.ThenFunc(handlers.DeleteDeviceHandler))
	//router.Delete("/device/value/delete/:id", h1.ThenFunc(handlers.DeleteDeviceValueHandler))
	//router.Get("/device/bindings/:id", h1.ThenFunc(handlers.CheckDeviceBindingsHandler))
	//
	//// PUSH - уведомления
	//router.Get("/notice/alive/:token", h1.ThenFunc(handlers.IsTokenAliveHandler))
	//router.Post("/notice/test/:token", h1.ThenFunc(handlers.SendTestMessageHandler))
	//router.Put("/notice/token/:token", h1.ThenFunc(handlers.UpdateNoticeTokenHandler))
	//router.Delete("/notice/unsubscribe/:token", h1.ThenFunc(handlers.UnsubscribeTokenHandler))
	//
	//// Аутентификация
	//router.Get("/auth/token", h0.ThenFunc(handlers.TokenHandler))
	//router.Post("/auth", h0.ThenFunc(handlers.TokenHandler))
	//
	//// Профиль
	//router.Post("/profile/permission/add", h2.ThenFunc(handlers.HandlePermissionAdd))
	//router.Post("/profile/permission/co_owner/add", h1.ThenFunc(handlers.HandlePermissionAddCoOwner))
	//router.Get("/profile/userdata/get", h1.ThenFunc(handlers.UserDataGetHandler))
	//router.Put("/profile/avatar/update", h2.ThenFunc(handlers.AvatarUpdateHandler))
	//router.Put("/profile/username/update", h2.ThenFunc(handlers.UserNameUpdateHandler))
	//router.Put("/profile/userpwd/update", h2.ThenFunc(handlers.UserPwdUpdateHandler))
	//router.Get("/profile/permission/co_owner/all", h1.ThenFunc(handlers.HandlePermissionGetAllCoOwners))
	//router.Get("/profile/permission/co_owners", h1.ThenFunc(handlers.HandlePermissionGetCoOwners))
	//router.Get("/profile/avatar/get", h1.ThenFunc(handlers.AvatarGetHandler))
	//router.Delete("/profile/permission/delete/:co_owner", h1.ThenFunc(handlers.HandlePermissionDelete))
	//
	//// Регионы
	//router.Get("/region/catalog", h1.ThenFunc(handlers.RegionCatalogHandler))
	//router.Get("/region/ip/:ip", h1.ThenFunc(handlers.RegionByIpHandler))
	//
	//// Передача показаний
	//router.Get("/transmit/methods/catalog", h1.ThenFunc(handlers.TransmitMethodsCatalogHandler))
	//router.Get("/transmit/organizations/:gid", h1.ThenFunc(handlers.OrganizationsGetAllHandler))
	//router.Post("/transmit/organization/delete", h2.ThenFunc(handlers.OrganizationDeleteHandler))
	//router.Get("/transmit/functions/:gid", h1.ThenFunc(handlers.TransmitMethodFunctions))
	//router.Get("/transmit/new", h1.ThenFunc(handlers.NewTransmit))
	//router.Post("/transmit/save", h2.ThenFunc(handlers.SaveTransmit))
	//router.Post("/transmit/run", h2.ThenFunc(handlers.RunTransmit))
	//router.Post("/transmit/run_all", h2.ThenFunc(handlers.RunAllTransmit))
	//router.Post("/transmit/update_reports", h2.ThenFunc(handlers.UpdateReports))

	return router
}
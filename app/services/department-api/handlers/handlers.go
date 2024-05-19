package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/attributegrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/cogagrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/cogrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/gagrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/markgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/studentgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/studentsubjectgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/subjectgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/testgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/usergrp"
	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/core/attribute/attributedb"
	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/PhyoYazar/uas/business/core/co/codb"
	"github.com/PhyoYazar/uas/business/core/coga"
	"github.com/PhyoYazar/uas/business/core/coga/cogadb"
	"github.com/PhyoYazar/uas/business/core/ga"
	"github.com/PhyoYazar/uas/business/core/ga/gadb"
	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/core/mark/markdb"
	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/core/student/studentdb"
	"github.com/PhyoYazar/uas/business/core/studentsubject"
	"github.com/PhyoYazar/uas/business/core/studentsubject/studentsubjectdb"
	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/PhyoYazar/uas/business/core/subject/subjectdb"
	"github.com/PhyoYazar/uas/business/core/user"
	"github.com/PhyoYazar/uas/business/core/user/stores/userdb"
	"github.com/PhyoYazar/uas/business/web/auth"
	"github.com/PhyoYazar/uas/business/web/v1/mid"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Auth     *auth.Auth
	DB       *sqlx.DB
}

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)
	app.Handle(http.MethodGet, "/test/auth", testgrp.Test, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAdminOnly))

	// -------------------------------------------------------------------------

	usrCore := user.NewCore(userdb.NewStore(cfg.Log, cfg.DB))

	ugh := usergrp.New(usrCore)

	app.Handle(http.MethodGet, "/users", ugh.Query)

	// -------------------------------------------------------------------------

	subCore := subject.NewCore(subjectdb.NewStore(cfg.Log, cfg.DB))

	subgh := subjectgrp.New(subCore)

	app.Handle(http.MethodGet, "/subjects", subgh.Query)
	app.Handle(http.MethodPost, "/subject", subgh.Create)
	app.Handle(http.MethodGet, "/subjects/:subject_id", subgh.QueryByID)

	// -------------------------------------------------------------------------

	stdCore := student.NewCore(studentdb.NewStore(cfg.Log, cfg.DB))

	stdgh := studentgrp.New(stdCore)

	app.Handle(http.MethodGet, "/students", stdgh.Query)
	app.Handle(http.MethodPost, "/student", stdgh.Create)

	// -------------------------------------------------------------------------

	coCore := co.NewCore(codb.NewStore(cfg.Log, cfg.DB))

	cogh := cogrp.New(coCore)

	app.Handle(http.MethodGet, "/cos", cogh.Query)
	app.Handle(http.MethodPost, "/co", cogh.Create)

	// -------------------------------------------------------------------------

	gaCore := ga.NewCore(gadb.NewStore(cfg.Log, cfg.DB))

	gagh := gagrp.New(gaCore)

	app.Handle(http.MethodGet, "/gas", gagh.Query)
	app.Handle(http.MethodPost, "/ga", gagh.Create)

	// -------------------------------------------------------------------------

	attributeCore := attribute.NewCore(attributedb.NewStore(cfg.Log, cfg.DB))

	attributegh := attributegrp.New(attributeCore)

	app.Handle(http.MethodGet, "/attributes", attributegh.Query)
	app.Handle(http.MethodPost, "/attribute", attributegh.Create)

	// -------------------------------------------------------------------------

	ssCore := studentsubject.NewCore(studentsubjectdb.NewStore(cfg.Log, cfg.DB))

	ssgh := studentsubjectgrp.New(ssCore)

	app.Handle(http.MethodGet, "/student_subjects", ssgh.Query)
	app.Handle(http.MethodPost, "/student_subject", ssgh.Create)

	// -------------------------------------------------------------------------

	cgCore := coga.NewCore(cogadb.NewStore(cfg.Log, cfg.DB))

	cggh := cogagrp.New(cgCore)

	app.Handle(http.MethodGet, "/co_gas", cggh.Query)
	app.Handle(http.MethodPost, "/co_ga", cggh.Create)

	// -------------------------------------------------------------------------

	mCore := mark.NewCore(markdb.NewStore(cfg.Log, cfg.DB))

	mgh := markgrp.New(mCore)

	app.Handle(http.MethodGet, "/marks", mgh.Query)
	app.Handle(http.MethodPost, "/mark", mgh.Create)

	return app
}

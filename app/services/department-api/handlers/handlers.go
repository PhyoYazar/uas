package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/attributegrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/coattributegrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/cogagrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/cogrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/fullmarkgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/gagrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/markgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/studentgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/studentsubjectgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/subjectgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/testgrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/usergrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/vattributegrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/vcogrp"
	"github.com/PhyoYazar/uas/app/services/department-api/handlers/v1/vsubjectgrp"
	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/core/attribute/attributedb"
	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/PhyoYazar/uas/business/core/co/codb"
	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/core/coattribute/coattributedb"
	"github.com/PhyoYazar/uas/business/core/coga"
	"github.com/PhyoYazar/uas/business/core/coga/cogadb"
	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/PhyoYazar/uas/business/core/fullmark/fullmarkdb"
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
	"github.com/PhyoYazar/uas/business/core/vattribute"
	"github.com/PhyoYazar/uas/business/core/vattribute/vattributedb"
	"github.com/PhyoYazar/uas/business/core/vco"
	"github.com/PhyoYazar/uas/business/core/vco/vcodb"
	"github.com/PhyoYazar/uas/business/core/vsubject"
	"github.com/PhyoYazar/uas/business/core/vsubject/vsubjectdb"
	"github.com/PhyoYazar/uas/business/web/auth"
	"github.com/PhyoYazar/uas/business/web/v1/mid"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

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

	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(),
		mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)
	app.Handle(http.MethodGet, "/test/auth", testgrp.Test, mid.Authenticate(cfg.Auth), mid.Authorize(cfg.Auth, auth.RuleAdminOnly))

	// -------------------------------------------------------------------------
	// user

	usrCore := user.NewCore(userdb.NewStore(cfg.Log, cfg.DB))

	ugh := usergrp.New(usrCore)

	app.Handle(http.MethodGet, "/users", ugh.Query)

	// -------------------------------------------------------------------------
	// subject

	subCore := subject.NewCore(subjectdb.NewStore(cfg.Log, cfg.DB))

	subgh := subjectgrp.New(subCore)

	app.Handle(http.MethodGet, "/subjects", subgh.Query)
	app.Handle(http.MethodPost, "/subject", subgh.Create)
	app.Handle(http.MethodGet, "/subjects/:subject_id", subgh.QueryByID)
	app.Handle(http.MethodPut, "/subjects/:subject_id", subgh.Update)
	app.Handle(http.MethodDelete, "/subjects/:subject_id", subgh.Delete)

	// -------------------------------------------------------------------------
	// student

	stdCore := student.NewCore(studentdb.NewStore(cfg.Log, cfg.DB))

	stdgh := studentgrp.New(stdCore)

	app.Handle(http.MethodGet, "/students", stdgh.Query)
	app.Handle(http.MethodPost, "/student", stdgh.Create)
	app.Handle(http.MethodGet, "/student/:student_id", stdgh.QueryByID)
	app.Handle(http.MethodPut, "/student/:student_id", stdgh.Update)
	app.Handle(http.MethodDelete, "/student/:student_id", stdgh.Delete)

	// -------------------------------------------------------------------------
	// co -> course outlines

	coCore := co.NewCore(codb.NewStore(cfg.Log, cfg.DB))

	cogh := cogrp.New(coCore)

	app.Handle(http.MethodGet, "/cos", cogh.Query)
	app.Handle(http.MethodPost, "/co", cogh.Create)
	app.Handle(http.MethodDelete, "/co/:co_id", cogh.Delete)

	// -------------------------------------------------------------------------
	// ga -> graduate attributes

	gaCore := ga.NewCore(gadb.NewStore(cfg.Log, cfg.DB))

	gagh := gagrp.New(gaCore)

	app.Handle(http.MethodGet, "/gas", gagh.Query)
	app.Handle(http.MethodPost, "/ga", gagh.Create)

	// -------------------------------------------------------------------------
	// attributes

	attributeCore := attribute.NewCore(attributedb.NewStore(cfg.Log, cfg.DB))

	attributegh := attributegrp.New(attributeCore)

	app.Handle(http.MethodGet, "/attributes", attributegh.Query)
	app.Handle(http.MethodPost, "/attribute", attributegh.Create)
	app.Handle(http.MethodGet, "/attribute/:attribute_id", attributegh.QueryByID)
	app.Handle(http.MethodDelete, "/attribute/:attribute_id", attributegh.Delete)

	// -------------------------------------------------------------------------
	// student subject

	ssCore := studentsubject.NewCore(studentsubjectdb.NewStore(cfg.Log, cfg.DB))

	ssgh := studentsubjectgrp.New(ssCore)

	app.Handle(http.MethodGet, "/student_subjects", ssgh.Query)
	app.Handle(http.MethodPost, "/student_subject", ssgh.Create)

	// -------------------------------------------------------------------------
	// co ga

	cgCore := coga.NewCore(cogadb.NewStore(cfg.Log, cfg.DB))

	cggh := cogagrp.New(cgCore, coCore)

	app.Handle(http.MethodGet, "/co_gas", cggh.Query)
	app.Handle(http.MethodPost, "/co_ga", cggh.Create)
	app.Handle(http.MethodPost, "/connect_co_gas", cggh.ConnectCoWithGa)

	// -------------------------------------------------------------------------
	// co attributes

	caCore := coattribute.NewCore(coattributedb.NewStore(cfg.Log, cfg.DB))

	cagh := coattributegrp.New(caCore)

	app.Handle(http.MethodGet, "/co_attributes", cagh.Query)
	app.Handle(http.MethodPost, "/co_attribute", cagh.Create)
	app.Handle(http.MethodDelete, "/co_attribute/:co_attribute_id", cagh.Delete)

	// -------------------------------------------------------------------------
	// mark

	mCore := mark.NewCore(markdb.NewStore(cfg.Log, cfg.DB))

	mgh := markgrp.New(mCore, caCore)

	app.Handle(http.MethodGet, "/marks", mgh.Query)
	app.Handle(http.MethodPost, "/mark", mgh.Create)
	app.Handle(http.MethodDelete, "/mark/:mark_id", mgh.Delete)
	app.Handle(http.MethodPost, "/create_mark_with_co_ga", mgh.CreateMarkByConnectingCOGA)

	// -------------------------------------------------------------------------
	// full mark

	fmCore := fullmark.NewCore(fullmarkdb.NewStore(cfg.Log, cfg.DB))

	fmgh := fullmarkgrp.New(fmCore)

	app.Handle(http.MethodGet, "/full_marks", fmgh.Query)
	app.Handle(http.MethodPost, "/full_mark", fmgh.Create)
	app.Handle(http.MethodDelete, "/full_mark/:full_mark_id", fmgh.Delete)

	// -------------------------------------------------------------------------
	// -------------------------------------------------------------------------
	// -------------------------------------------------------------------------
	// -------------------------------------------------------------------------

	vsubCore := vsubject.NewCore(vsubjectdb.NewStore(cfg.Log, cfg.DB))

	vsubgh := vsubjectgrp.New(vsubCore)

	app.Handle(http.MethodGet, "/subject_detail/:subject_id", vsubgh.QueryByID)

	// -------------------------------------------------------------------------

	vcoCore := vco.NewCore(vcodb.NewStore(cfg.Log, cfg.DB))

	vcogh := vcogrp.New(vcoCore)

	app.Handle(http.MethodGet, "/co_detail/:co_id", vcogh.QueryByID)

	// -------------------------------------------------------------------------

	vattCore := vattribute.NewCore(vattributedb.NewStore(cfg.Log, cfg.DB))

	vattgh := vattributegrp.New(vattCore)

	app.Handle(http.MethodGet, "/attributes_detail/:subject_id", vattgh.Query)
	app.Handle(http.MethodGet, "/attributes_ga_mark", vattgh.QueryAttributeWithGaMark)
	app.Handle(http.MethodDelete, "/remove_attribute/:attribute_id", vattgh.RemoveAttribute)

	return app
}

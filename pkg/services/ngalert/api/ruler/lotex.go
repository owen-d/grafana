package ruler

import (
	"net/http"

	apimodels "github.com/grafana/alerting-api/pkg/api"
	"gopkg.in/macaron.v1"

	"github.com/grafana/grafana/pkg/api/dataproxy"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
)

type LotexRuler struct {
	DataProxy *dataproxy.DataProxy
}

// withBufferedRW is a hack to work around the type signature exposed by the datasource proxy
// which uses the underlying ResponseWriter vs what we need to expose implementing the API services
// (which return a response.Response). Therefore, we replace the response writer so that we can return it.
func withBufferedRW(ctx *models.ReqContext) (*models.ReqContext, response.Response) {
	resp := response.CreateNormalResponse(make(http.Header), nil, 0)
	cpy := *ctx

	cpy.Resp = macaron.NewResponseWriter(ctx.Req.Method, resp)
	return &cpy, resp
}

// unmodofiedProxy passes a request to a datasource unaltered.
func (r *LotexRuler) unmodifiedProxy(ctx *models.ReqContext) response.Response {
	newCtx, resp := withBufferedRW(ctx)
	r.DataProxy.ProxyDatasourceRequestWithID(newCtx, ctx.ParamsInt64("DatasourceId"))
	return resp
}

func (r *LotexRuler) RouteDeleteNamespaceRulesConfig(ctx *models.ReqContext) response.Response {
	return r.unmodifiedProxy(ctx)
}

func (r *LotexRuler) RouteDeleteRuleGroupConfig(ctx *models.ReqContext) response.Response {
	return r.unmodifiedProxy(ctx)
}

func (r *LotexRuler) RouteGetNamespaceRulesConfig(ctx *models.ReqContext) response.Response {
	return r.unmodifiedProxy(ctx)
}

func (r *LotexRuler) RouteGetRulegGroupConfig(ctx *models.ReqContext) response.Response {
	return r.unmodifiedProxy(ctx)
}

func (r *LotexRuler) RouteGetRulesConfig(ctx *models.ReqContext) response.Response {
	return r.unmodifiedProxy(ctx)
}

func (r *LotexRuler) RoutePostNameRulesConfig(ctx *models.ReqContext, conf apimodels.RuleGroupConfig) response.Response {
	return r.unmodifiedProxy(ctx)
}

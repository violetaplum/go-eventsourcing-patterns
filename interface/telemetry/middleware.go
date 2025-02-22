package telemetry

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func GinMiddleware(service string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 트레이서 가져오기
		tracer := otel.Tracer(service)

		// 요청 경로로 스팬 시작
		ctx, span := tracer.Start(c.Request.Context(), c.Request.URL.Path)
		defer span.End()

		// 요청 정보를 스팬에 추가
		span.SetAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.url", c.Request.URL.String()),
		)

		// context 를 gin 에 제공
		c.Request = c.Request.WithContext(ctx)

		// 다음 핸들러 실행

		// 응답 정보를 스팬에 추가
		span.SetAttributes(
			attribute.Int64("http.status_code", int64(c.Writer.Status())))
	}
}

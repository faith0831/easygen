{{- /* @lang csharp */ -}}
{{- /* @env Module 模块名称 */ -}}
{{- /* @env Alias 模块别名 */ -}}
using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Digua.Web.Swagger;
using Digua.Web.Controllers;
using Digua.Module.{{.ENV.Module}}.Services;
using Digua.Module.{{.ENV.Module}}.Dto;

namespace Digua.Module.{{.ENV.Module}}.Controllers
{
    public class {{ .Table.Name | pascal }}Controller : AdminControllerBase
    {
        private readonly I{{ .Table.Name | pascal }}Service _{{ .Table.Name | camel }}Service;
        
        public {{ .Table.Name | pascal }}Controller(I{{ .Table.Name | pascal }}Service {{ .Table.Name | camel }}Service)
        {
            _{{ .Table.Name | camel }}Service = {{ .Table.Name | camel }}Service;
        }
        
        /// <summary>
        /// 获取{{  .ENV.Alias }}列表
        /// </summary>
        /// <returns></returns>
        [HttpGet]
        public async Task<Result> Search{{ .Table.Name | pascal }}([FromQuery]Search{{ .Table.Name | pascal }}Request request)
        {
            var list = await _{{ .Table.Name | camel }}Service.Search{{ .Table.Name | pascal }}(request);
            return Result.Ok(list);
        }
        
        /// <summary>
        /// 获取{{ .ENV.Alias }}信息
        /// </summary>
        /// <returns></returns>
        [HttpGet("{id}")]
        public async Task<Result> Get{{ .Table.Name | pascal }}ById(long id)
        {
            var model = await _{{ .Table.Name | camel }}Service.Get{{ .Table.Name | pascal }}ById(id);
            return Result.Ok(model);
        }
        
        /// <summary>
        /// 创建{{ .ENV.Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpPost]
        public async Task<Result> Create{{ .Table.Name | pascal }}([FromBody]Create{{ .Table.Name | pascal }}Request request)
        {
            return await _{{ .Table.Name | camel }}Service.Create{{ .Table.Name | pascal }}(request);
        }
        
        /// <summary>
        /// 更新{{ .ENV.Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpPut("{id}")]
        public async Task<Result> Update{{ .Table.Name | pascal }}(long id, [FromBody]Update{{ .Table.Name | pascal }}Request request)
        {
            request.Id = id;
            return await _{{ .Table.Name | camel }}Service.Update{{ .Table.Name | pascal }}(request);
        }
        
        /// <summary>
        /// 删除{{ .ENV.Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpDelete("{id}")]
        public async Task<Result> Delete{{ .Table.Name | pascal }}(long id)
        {
            return await _{{ .Table.Name | camel }}Service.Delete{{ .Table.Name | pascal }}(id);
        }
		
        /// <summary>
        /// 删除{{ .ENV.Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpDelete("{ids}")]
        public async Task<Result> Delete{{ .Table.Name | pascal }}(long[] ids)
        {
            return await _{{ .Table.Name | camel }}Service.Delete{{ .Table.Name | pascal }}(ids);
        }
    }
}
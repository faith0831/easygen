{{- /* @lang csharp */ -}}
{{- /* @env Module 功能模块 */ -}}
{{- /* @env Alias 别名 */ -}}
using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Digua.Web.Swagger;
using Digua.Web.Controllers;
using Digua.Mall.Module.{{.Module}}.Services;
using Digua.Mall.Module.{{.Module}}.Dto;

namespace Digua.Mall.Module.{{.Module}}.Controllers
{
    public class {{ .Table.Name | camel }}Controller : AdminControllerBase
    {
        private readonly I{{ .Table.Name | camel }}Service _{{ .Table.Name | lowerCamel }}Service;
        
        public {{ .Table.Name | camel }}Controller(I{{ .Table.Name | camel }}Service {{ .Table.Name | lowerCamel }}Service)
        {
            _{{ .Table.Name | lowerCamel }}Service = {{ .Table.Name | lowerCamel }}Service;
        }
        
        /// <summary>
        /// 获取{{  .Alias }}列表
        /// </summary>
        /// <returns></returns>
        [HttpGet]
        public async Task<Result> Search{{ .Table.Name | camel }}([FromQuery]Search{{ .Table.Name | camel }}Request request)
        {
            var list = await _{{ .Table.Name | lowerCamel }}Service.Search{{ .Table.Name | camel }}(request);
            return Result.Ok(list);
        }
        
        /// <summary>
        /// 获取{{ .Alias }}信息
        /// </summary>
        /// <returns></returns>
        [HttpGet("{id}")]
        public async Task<Result> Get{{ .Table.Name | camel }}ById(long id)
        {
            var model = await _{{ .Table.Name | lowerCamel }}Service.Get{{ .Table.Name | camel }}ById(id);
            return Result.Ok(model);
        }
        
        /// <summary>
        /// 创建{{ .Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpPost]
        public async Task<Result> Create{{ .Table.Name | camel }}([FromBody]Create{{ .Table.Name | camel }}Request request)
        {
            return await _{{ .Table.Name | lowerCamel }}Service.Create{{ .Table.Name | camel }}(request);
        }
        
        /// <summary>
        /// 更新{{ .Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpPut("{id}")]
        public async Task<Result> Update{{ .Table.Name | camel }}(long id, [FromBody]Update{{ .Table.Name | camel }}Request request)
        {
            request.Id = id;
            return await _{{ .Table.Name | lowerCamel }}Service.Update{{ .Table.Name | camel }}(request);
        }
        
        /// <summary>
        /// 删除{{ .Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpDelete("{id}")]
        public async Task<Result> Delete{{ .Table.Name | camel }}(long id)
        {
            return await _{{ .Table.Name | lowerCamel }}Service.Delete{{ .Table.Name | camel }}(id);
        }
		
        /// <summary>
        /// 删除{{ .Alias }}
        /// </summary>
        /// <returns></returns>
        [HttpDelete("{ids}")]
        public async Task<Result> Delete{{ .Table.Name | camel }}(long[] ids)
        {
            return await _{{ .Table.Name | lowerCamel }}Service.Delete{{ .Table.Name | camel }}(ids);
        }
    }
}
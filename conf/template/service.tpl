{{- /* @lang csharp */ -}}
{{- /* @env Module 模块名称 */ -}}
{{- /* @env Alias 模块别名 */ -}}
using System;
using System.Linq;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using AutoMapper;
using AutoMapper.QueryableExtensions;
using Digua.EntityFrameworkCore;
using Digua.Extensions;
using Digua.Exceptions;
using Digua.Models.Paging;
using Digua.Module.{{.ENV.Module}}.Dto;
using Digua.Module.{{.ENV.Module}}.Validator;
using Digua.Module.{{.ENV.Module}}.Models;

namespace Digua.Module.{{.ENV.Module}}.Services
{
    public interface I{{ .Table.Name | pascal }}Service
    {
        #region {{ .ENV.Alias }}
        /// <summary>
        /// 取{{ .ENV.Alias }}列表
        /// </summary>
        /// <returns></returns>
        Task<List<{{ .Table.Name | pascal }}ListDto>> Get{{ .Table.Name | pascal }}List();

        /// <summary>
        /// 根据筛选条件取{{ .ENV.Alias }}列表
        /// </summary>
        /// <param name="request"></param>
        /// <returns></returns>
        Task<JsonPagedList<{{ .Table.Name | pascal }}ListDto>> Search{{ .Table.Name | pascal }}(Search{{ .Table.Name | pascal }}Request request);

        /// <summary>
        /// 根据Id取{{ .ENV.Alias }}
        /// </summary>
        /// <param name="id"></param>
        Task<{{ .Table.Name | pascal }}Dto> Get{{ .Table.Name | pascal }}ById(long id);

        /// <summary>
        /// 创建{{ .ENV.Alias }}
        /// </summary>
        /// <param name="request"></param>
        Task<Result> Create{{ .Table.Name | pascal }}(Create{{ .Table.Name | pascal }}Request request);

        /// <summary>
        /// 更新{{ .ENV.Alias }}
        /// </summary>
        /// <param name="request"></param>
        Task<Result> Update{{ .Table.Name | pascal }}(Update{{ .Table.Name | pascal }}Request request);

        /// <summary>
        /// 根据id删除{{ .ENV.Alias }}
        /// </summary>
        /// <param name="id"></param>
        Task<Result> Delete{{ .Table.Name | pascal }}(long id);
		
        /// <summary>
        /// 根据ids删除{{ .ENV.Alias }}
        /// </summary>
        /// <param name="ids"></param>
        Task<Result> Delete{{ .Table.Name | pascal }}(long[] ids);
        #endregion
    }

    public class {{ .Table.Name | pascal }}Service : I{{ .Table.Name | pascal }}Service
    {
        private readonly IUnitOfWork<DiguaDbContext> _unitOfWork;
        private readonly IMapper _mapper;

        public {{ .Table.Name | pascal }}Service(IUnitOfWork<DiguaDbContext> unitOfWork, IMapper mapper)
        {
            _unitOfWork = unitOfWork;
            _mapper = mapper;
        }

       #region {{ .ENV.Alias }}
        /// <summary>
        /// 取{{ .ENV.Alias }}列表
        /// </summary>
        /// <returns></returns>
        public async Task<List<{{ .Table.Name | pascal }}ListDto>> Get{{ .Table.Name | pascal }}List()
        {
            var query = from e in _unitOfWork.Repository<{{ .Table.Name | pascal }}>().Table.ProjectTo<{{ .Table.Name | pascal }}ListDto>(_mapper.ConfigurationProvider)
                        select e;
            
            return await query.OrderByDescending(e => e.Id).ToListAsync();
        }

        /// <summary>
        /// 根据筛选条件取{{ .ENV.Alias }}列表
        /// </summary>
        /// <param name="request"></param>
        /// <returns></returns>
        public async Task<JsonPagedList<{{ .Table.Name | pascal }}ListDto>> Search{{ .Table.Name | pascal }}(Search{{ .Table.Name | pascal }}Request request)
        {
            var query = from e in _unitOfWork.Repository<{{ .Table.Name | pascal }}>().Table.ProjectTo<{{ .Table.Name | pascal }}ListDto>(_mapper.ConfigurationProvider)
                        select e;
            
            return await query.OrderByDescending(e => e.Id).ToJsonPagedListAsync(request);
        }

        /// <summary>
        /// 根据Id取{{ .ENV.Alias }}
        /// </summary>
        /// <param name="id"></param>
        public async Task<{{ .Table.Name | pascal }}Dto> Get{{ .Table.Name | pascal }}ById(long id)
        {
            return await _unitOfWork.Repository<{{ .Table.Name | pascal }}>().Table.ProjectTo<{{ .Table.Name | pascal }}Dto>(_mapper.ConfigurationProvider).FirstOrDefaultAsync(e => e.Id == id);
        }

        /// <summary>
        /// 创建{{ .ENV.Alias }}
        /// </summary>
        /// <param name="request"></param>
        public async Task<Result> Create{{ .Table.Name | pascal }}(Create{{ .Table.Name | pascal }}Request request)
        {
            if (request == null)
                throw new RequestArgumentNullException(nameof(request));

            var validResult = new Create{{ .Table.Name | pascal }}Validator().Valid(request);
            if (validResult.IsErr())
                return validResult;

            var model = _mapper.Map<{{ .Table.Name | pascal }}>(request);
            await _unitOfWork.Repository<{{ .Table.Name | pascal }}>().InsertAsync(model);
            await _unitOfWork.SaveChangesAsync();
            return Result.Ok("创建成功");
        }

        /// <summary>
        /// 更新{{ .ENV.Alias }}
        /// </summary>
        /// <param name="request"></param>
        public async Task<Result> Update{{ .Table.Name | pascal }}(Update{{ .Table.Name | pascal }}Request request)
        {
            if (request == null)
                throw new RequestArgumentNullException(nameof(request));

            var validResult = new Update{{ .Table.Name | pascal }}Validator().Valid(request);
            if (validResult.IsErr())
                return validResult;

            var existItem = await _unitOfWork.Repository<{{ .Table.Name | pascal }}>().Table.FirstOrDefaultAsync(e => e.Id == request.Id);
            if (existItem == null)
                return Result.Err("信息不存在");
            {{ range .Table.Columns }}
            {{- if $.SkipUpdate . }} {{ continue }} {{ end }}
            existItem.{{ .Name | pascal }} = request.{{ .Name | pascal }};
            {{- end }}
            
            await _unitOfWork.SaveChangesAsync();
            return Result.Ok("修改成功");
        }

        /// <summary>
        /// 根据id删除{{ .ENV.Alias }}
        /// </summary>
        /// <param name="id"></param>
        public async Task<Result> Delete{{ .Table.Name | pascal }}(long id)
        {
            var existItem = await _unitOfWork.Repository<{{ .Table.Name | pascal }}>().Table.FirstOrDefaultAsync(e => e.Id == id);
            if (existItem == null)
                return Result.Err("信息不存在");

            _unitOfWork.Repository<{{ .Table.Name | pascal }}>().Delete(existItem);
            await _unitOfWork.SaveChangesAsync();
            return Result.Ok("删除成功");
        }
		
        /// <summary>
        /// 根据ids删除{{ .ENV.Alias }}
        /// </summary>
        /// <param name="ids"></param>
        public async Task<Result> Delete{{ .Table.Name | pascal }}(long[] ids)
        {
            foreach (var id in ids)
            {
                await Delete{{ .Table.Name | pascal }}(id);
            }

            return Result.Ok("删除成功");
        }
       #endregion
    }
}
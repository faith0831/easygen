{{- /* @lang csharp */ -}}
{{- /* @env Module 功能模块 */ -}}
{{- /* @env Alias 模块名称 */ -}}
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
using Digua.Mall.Module.{{.Module}}.Dto;
using Digua.Mall.Module.{{.Module}}.Validator;
using Digua.Mall.Module.{{.Module}}.Models;

namespace Digua.Mall.Module.{{.Module}}.Services
{
    public interface I{{ .Table.Name | camel }}Service
    {
        #region {{ .Alias }}
        /// <summary>
        /// 取{{ .Alias }}列表
        /// </summary>
        /// <returns></returns>
        Task<List<{{ .Table.Name | camel }}ListDto>> Get{{ .Table.Name | camel }}List();

        /// <summary>
        /// 根据筛选条件取{{ .Alias }}列表
        /// </summary>
        /// <param name="request"></param>
        /// <returns></returns>
        Task<JsonPagedList<{{ .Table.Name | camel }}ListDto>> Search{{ .Table.Name | camel }}(Search{{ .Table.Name | camel }}Request request);

        /// <summary>
        /// 根据Id取{{ .Alias }}
        /// </summary>
        /// <param name="id"></param>
        Task<{{ .Table.Name | camel }}Dto> Get{{ .Table.Name | camel }}ById(long id);

        /// <summary>
        /// 创建{{ .Alias }}
        /// </summary>
        /// <param name="request"></param>
        Task<Result> Create{{ .Table.Name | camel }}(Create{{ .Table.Name | camel }}Request request);

        /// <summary>
        /// 更新{{ .Alias }}
        /// </summary>
        /// <param name="request"></param>
        Task<Result> Update{{ .Table.Name | camel }}(Update{{ .Table.Name | camel }}Request request);

        /// <summary>
        /// 根据id删除{{ .Alias }}
        /// </summary>
        /// <param name="id"></param>
        Task<Result> Delete{{ .Table.Name | camel }}(long id);
		
        /// <summary>
        /// 根据ids删除{{ .Alias }}
        /// </summary>
        /// <param name="ids"></param>
        Task<Result> Delete{{ .Table.Name | camel }}(long[] ids);
        #endregion
    }

    public class {{ .Table.Name | camel }}Service : I{{ .Table.Name | camel }}Service
    {
        private readonly IUnitOfWork<DiguaDbContext> _unitOfWork;
        private readonly IMapper _mapper;

        public {{ .Table.Name | camel }}Service(IUnitOfWork<DiguaDbContext> unitOfWork, IMapper mapper)
        {
            _unitOfWork = unitOfWork;
            _mapper = mapper;
        }

       #region {{ .Alias }}
        /// <summary>
        /// 取{{ .Alias }}列表
        /// </summary>
        /// <returns></returns>
        public async Task<List<{{ .Table.Name | camel }}ListDto>> Get{{ .Table.Name | camel }}List()
        {
            var query = from e in _unitOfWork.Repository<{{ .Table.Name | camel }}>().Table.ProjectTo<{{ .Table.Name | camel }}ListDto>(_mapper.ConfigurationProvider)
                        select e;
            
            return await query.OrderByDescending(e => e.Id).ToListAsync();
        }

        /// <summary>
        /// 根据筛选条件取{{ .Alias }}列表
        /// </summary>
        /// <param name="request"></param>
        /// <returns></returns>
        public async Task<JsonPagedList<{{ .Table.Name | camel }}ListDto>> Search{{ .Table.Name | camel }}(Search{{ .Table.Name | camel }}Request request)
        {
            var query = from e in _unitOfWork.Repository<{{ .Table.Name | camel }}>().Table.ProjectTo<{{ .Table.Name | camel }}ListDto>(_mapper.ConfigurationProvider)
                        select e;
            
            return await query.OrderByDescending(e => e.Id).ToJsonPagedListAsync(request);
        }

        /// <summary>
        /// 根据Id取{{ .Alias }}
        /// </summary>
        /// <param name="id"></param>
        public async Task<{{ .Table.Name | camel }}Dto> Get{{ .Table.Name | camel }}ById(long id)
        {
            return await _unitOfWork.Repository<{{ .Table.Name | camel }}>().Table.ProjectTo<{{ .Table.Name | camel }}Dto>(_mapper.ConfigurationProvider).FirstOrDefaultAsync(e => e.Id == id);
        }

        /// <summary>
        /// 创建{{ .Alias }}
        /// </summary>
        /// <param name="request"></param>
        public async Task<Result> Create{{ .Table.Name | camel }}(Create{{ .Table.Name | camel }}Request request)
        {
            if (request == null)
                throw new RequestArgumentNullException(nameof(request));

            var validResult = new Create{{ .Table.Name | camel }}Validator().Valid(request);
            if (validResult.IsErr())
                return validResult;

            var model = _mapper.Map<{{ .Table.Name | camel }}>(request);
            await _unitOfWork.Repository<{{ .Table.Name | camel }}>().InsertAsync(model);
            await _unitOfWork.SaveChangesAsync();
            return Result.Ok("创建成功");
        }

        /// <summary>
        /// 更新{{ .Alias }}
        /// </summary>
        /// <param name="request"></param>
        public async Task<Result> Update{{ .Table.Name | camel }}(Update{{ .Table.Name | camel }}Request request)
        {
            if (request == null)
                throw new RequestArgumentNullException(nameof(request));

            var validResult = new Update{{ .Table.Name | camel }}Validator().Valid(request);
            if (validResult.IsErr())
                return validResult;

            var existItem = await _unitOfWork.Repository<{{ .Table.Name | camel }}>().Table.FirstOrDefaultAsync(e => e.Id == request.Id);
            if (existItem == null)
                return Result.Err("信息不存在");
            {{ range .Table.Columns }}
            existItem.{{ .Name | camel }} = request.{{ .Name | camel }};
            {{- end }}
            
            await _unitOfWork.SaveChangesAsync();
            return Result.Ok("修改成功");
        }

        /// <summary>
        /// 根据id删除{{ .Alias }}
        /// </summary>
        /// <param name="id"></param>
        public async Task<Result> Delete{{ .Table.Name | camel }}(long id)
        {
            var existItem = await _unitOfWork.Repository<{{ .Table.Name | camel }}>().Table.FirstOrDefaultAsync(e => e.Id == id);
            if (existItem == null)
                return Result.Err("信息不存在");

            _unitOfWork.Repository<{{ .Table.Name | camel }}>().Delete(existItem);
            await _unitOfWork.SaveChangesAsync();
            return Result.Ok("删除成功");
        }
		
        /// <summary>
        /// 根据ids删除{{ .Alias }}
        /// </summary>
        /// <param name="ids"></param>
        public async Task<Result> Delete{{ .Table.Name | camel }}(long[] ids)
        {
            foreach (var id in ids)
            {
                await Delete{{ .Table.Name | camel }}(id);
            }

            return Result.Ok("删除成功");
        }
       #endregion
    }
}
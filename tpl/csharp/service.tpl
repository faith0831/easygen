{{- /* @env Alias 别名 */ -}}
/// <summary>
/// 取{{ .Alias }}列表
/// </summary>
/// <returns></returns>
List<{{ .Table.Name }}ListDto> Get{{ .Table.Name }}List();

/// <summary>
/// 根据筛选条件取{{ .Alias }}列表
/// </summary>
/// <param name="request"></param>
/// <returns></returns>
JsonPagedList<{{ .Table.Name }}ListDto> Search{{ .Table.Name }}(Search{{ .Table.Name }}Request request);

/// <summary>
/// 根据Id取{{ .Alias }}
/// </summary>
/// <param name="id"></param>
{{ .Table.Name }}Dto Get{{ .Table.Name }}ById(int id);

/// <summary>
/// 创建{{ .Alias }}
/// </summary>
/// <param name="request"></param>
OperateResult Create{{ .Table.Name }}(Create{{ .Table.Name }}Request request);

/// <summary>
/// 更新{{ .Alias }}
/// </summary>
/// <param name="request"></param>
OperateResult Update{{ .Table.Name }}(Update{{ .Table.Name }}Request request);

/// <summary>
/// 根据id删除{{ .Alias }}
/// </summary>
/// <param name="id"></param>
OperateResult Delete{{ .Table.Name }}(int id);

/// <summary>
/// 取{{ .Alias }}列表
/// </summary>
/// <returns></returns>
public List<{{ .Table.Name }}ListDto> Get{{ .Table.Name }}List()
{
    var query = from e in _unitOfWork.Repository<{{ .Table.Name }}>().Table.ProjectTo<{{ .Table.Name }}ListDto>()
                select e;
    
    return query.OrderByDescending(e => e.Id).ToList();
}

/// <summary>
/// 根据筛选条件取{{ .Alias }}列表
/// </summary>
/// <param name="request"></param>
/// <returns></returns>
public JsonPagedList<{{ .Table.Name }}ListDto> Search{{ .Table.Name }}(Search{{ .Table.Name }}Request request)
{
    var query = from e in _unitOfWork.Repository<{{ .Table.Name }}>().Table.ProjectTo<{{ .Table.Name }}ListDto>()
                select e;
    
    return query.OrderByDescending(e => e.Id).ToJsonPagedList(request);
}

/// <summary>
/// 根据Id取{{ .Alias }}
/// </summary>
/// <param name="id"></param>
public {{ .Table.Name }}Dto Get{{ .Table.Name }}ById(int id)
{
    return _unitOfWork.Repository<{{ .Table.Name }}>().Table.ProjectTo<{{ .Table.Name }}Dto>().FirstOrDefault(e => e.Id == id);
}

/// <summary>
/// 创建{{ .Alias }}
/// </summary>
/// <param name="request"></param>
public OperateResult Create{{ .Table.Name }}(Create{{ .Table.Name }}Request request)
{
	var model = AutoMapper.Mapper.Map<{{ .Table.Name }}>(request);
    _unitOfWork.Repository<{{ .Table.Name }}>().Insert(model);
    _unitOfWork.SaveChanges();
    return OperateResult.Succeed("创建成功");
}

/// <summary>
/// 更新{{ .Alias }}
/// </summary>
/// <param name="request"></param>
public OperateResult Update{{ .Table.Name }}(Update{{ .Table.Name }}Request request)
{
    var existItem = _unitOfWork.Repository<{{ .Table.Name }}>().Table.FirstOrDefault(e => e.Id == request.Id);
    if (existItem == null)
    	return OperateResult.Error("信息不存在");
    {{ range .Table.Columns }}
    existItem.{{ .Name }} = request.{{ .Name }};
    {{- end }}
    
    _unitOfWork.SaveChanges();
    return OperateResult.Succeed("修改成功");
}

/// <summary>
/// 根据id删除{{ .Alias }}
/// </summary>
/// <param name="id"></param>
public OperateResult Delete{{ .Table.Name }}(int id)
{
    var existItem = _unitOfWork.Repository<{{ .Table.Name }}>().Table.FirstOrDefault(e => e.Id == id);
    if (existItem == null)
    	return OperateResult.Error("信息不存在");

    _unitOfWork.Repository<{{ .Table.Name }}>().Delete(existItem);
    _unitOfWork.SaveChanges();
    return OperateResult.Succeed("删除成功");
}
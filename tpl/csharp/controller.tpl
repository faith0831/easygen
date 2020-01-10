{{- /* @env Alias 别名 */ -}}
/// <summary>
/// 获取{{  .Alias }}列表
/// </summary>
/// <returns></returns>
[HttpGet]
[Route("{{ .Table.Name | lower }}/list")]
public IActionResult Search{{ .Table.Name }}(Search{{ .Table.Name }}Request request)
{
	var result = _{{ .Table.Name | lower }}Service.Search{{ .Table.Name }}(request);
    return Json(OperateResult.Succeed("ok", result));
}

/// <summary>
/// 获取{{ .Alias }}信息
/// </summary>
/// <returns></returns>
[HttpGet]
[Route("{{ .Table.Name | lower }}/item")]
public IActionResult Get{{ .Table.Name }}(IdRequest<int> request)
{
	var result = _{{ .Table.Name | lower }}Service.Get{{ .Table.Name }}ById(request.Id);
    return Json(OperateResult.Succeed("ok", result));
}

/// <summary>
/// 创建{{ .Alias }}
/// </summary>
/// <returns></returns>
[HttpPost]
[Route("{{ .Table.Name | lower }}/create")]
public IActionResult Create{{ .Table.Name }}([FromBody]Create{{ .Table.Name }}Request request)
{
	var result = _{{ .Table.Name | lower }}Service.Create{{ .Table.Name }}(request);
	return Json(result);
}

/// <summary>
/// 更新{{ .Alias }}
/// </summary>
/// <returns></returns>
[HttpPost]
[Route("{{ .Table.Name | lower }}/edit")]
public IActionResult Update{{ .Table.Name }}([FromBody]Update{{ .Table.Name }}Request request)
{
	var result = _{{ .Table.Name | lower }}Service.Update{{ .Table.Name }}(request);
	return Json(result);
}

/// <summary>
/// 删除{{ .Alias }}
/// </summary>
/// <returns></returns>
[HttpPost]
[Route("{{ .Table.Name | lower }}/delete")]
public IActionResult Delete{{ .Table.Name }}([FromBody]IdRequest<int> request)
{
    var result = _{{ .Table.Name | lower }}Service.Delete{{ .Table.Name }}(request.Id);
    return Json(result);
}
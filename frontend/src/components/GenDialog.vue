<template>
  <el-dialog title="生成代码" v-show="visible" v-model="visible" width="900px" :close-on-click-modal="false">
    <el-row class="gen-dialog" :gutter="10">
      <el-col :span="10">
        <el-card shadow="never">
          <template #header>
            <div class="card-header">
              <span>请选择模板</span>
            </div>
          </template>
          <el-tree
            ref="treeRef"
            :data="templates"
            :props="{ multiple: true, label: 'name', value: 'template' }"
            show-checkbox
            :default-expand-all="false"
            node-key="template"
            highlight-current
            @check="select"
          />
        </el-card>
      </el-col>
      <el-col :span="14">
        <el-card v-if="envs && envs.length > 0" shadow="never">
          <template #header>
            <div class="card-header">
              <span>自定义变量</span>
            </div>
          </template>
          <el-form>
            <el-form-item :key="i" v-for="(item, i) in envs">
              <el-col :span="4"> {{ item.label }} </el-col>
              <el-col class="line" :span="2">&nbsp;</el-col>
              <el-col :span="16">
                <el-input v-model="item.value" placeholder="值"></el-input>
              </el-col>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
    <el-row slot="footer" style="margin-top: 20px; justify-content: center">
      <el-button type="primary" @click="generate">生成代码</el-button>
      <el-button @click="hide()">关闭</el-button>
    </el-row>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, defineEmits, defineExpose } from 'vue'
import { storeToRefs } from 'pinia'
import { useStore } from '@/store/index.js'
import * as api from '/wailsjs/go/app/Application'
import ui from '@/utils/ui.js'

const emit = defineEmits(['genHandle'])
const treeRef = ref(null)
const store = useStore()
const envs = ref([])
const visible = ref(false)
const { templates } = storeToRefs(store)

const form = reactive({
  table: null,
  lang: null,
  template: null,
  env: {},
})

const show = (table) => {
  form.table = table
  visible.value = true
}

const hide = () => {
  visible.value = false
}

const select = () => {
  let nodes = treeRef.value.getCheckedNodes()
  envs.value = []

  nodes.forEach((node) => {
    if (node.env) {
      node.env.forEach((e) => {
        if (!envs.value.find((x) => x.key == e.key)) {
          envs.value.push(e)
        }
      })
    }
  })
}

const generate = async () => {
  let selectTemplates = treeRef.value.getCheckedNodes()
  if (!selectTemplates || selectTemplates.length <= 0) return ui.warning('请选择代码的模板')

  if (form.env) {
    let env = {}

    envs.value.forEach((item) => {
      env[item.key] = item.value
    })

    form.env = env
  }

  const codes = []

  for (let index in selectTemplates) {
    const t = selectTemplates[index]
    if (t.children == null) {
      form.lang = t.lang
      form.template = t.template

      let resp = await api.Generate(form)
      if (resp.code === 0) {
        return ui.error(resp.msg)
      }

      if (resp.code === 400) {
        return ui.error('连接失效，请重新连接数据库')
      }

      const code = resp.data
      codes.push({ name: t.name, lang: t.lang, source: code, code: code })
    }
  }

  let codeName = null
  if (codes.length > 0) {
    codeName = codes[0].name
  }

  console.log(codes)

  emit('genHandle', { codes, codeName })
  hide()
}

defineExpose({
  show,
  hide,
})
</script>

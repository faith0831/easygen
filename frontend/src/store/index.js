import { defineStore } from 'pinia'
import * as api from '/wailsjs/go/app/Application'
import ui from '@/utils/ui.js'

export const useStore = defineStore('main', {
  state: () => {
    return {
      connected: false,
      provider: {},
      tables: [],
      templates: [],
    }
  },
  getters: {
    isConnected() {
      return this.connected == true
    },
  },
  actions: {
    async init() {
      let resp = await api.HasProvider()
      if (resp.data.has) {
        this.provider = { ...resp.data.config }
        this.connected = true
        await this.getTables()
        await this.getTemplates()
      } else {
        let resp = await api.GetConfig()
        if (resp.code === 1) {
          this.provider = { ...resp.data }
          this.connected = false
        }
      }
    },
    setProvider(provider) {
      this.provider = { ...provider }
    },
    async connect(provider) {
      let resp = await api.CreateProvider(provider)
      if (resp.code === 0) {
        ui.error(resp.msg)
        return false
      }

      this.provider = { ...provider }
      this.connected = true
      await this.getTables()
      await this.getTemplates()
      return true
    },
    async getTables() {
      let resp = await api.GetTables()
      if (resp.code === 0) {
        return ui.error(resp.msg)
      }

      if (resp.code === 400) {
        this.connected = false
        return ui.error('连接失效，请重新连接数据库')
      }

      const tables = []
      resp.data.forEach((item, index) => {
        tables.push({ id: index, label: item })
      })

      this.tables = tables
    },
    async getTemplates() {
      let resp = await api.GetTemplates()
      if (resp.code === 0) {
        return ui.error(resp.msg)
      }

      this.templates = resp.data
    },
  },
})

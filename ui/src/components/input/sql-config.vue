<script lang="jsx">
import {sqlInputConfigParse, sqlInputConfig} from "type";
import configMixin from '../config-mixin'
export default {
    mixins: [configMixin],
    render() {
        return <div>
            <ysz-list-item v-slots={{
                left: () => '地址'
            }}>
                <a-input size="small" vModel={[this.cfg.url, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '端口'
            }}>
                <a-input size="small" vModel={[this.cfg.port, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '用户'
            }}>
                <a-input size="small" vModel={[this.cfg.user, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '密码'
            }}>
                <a-input size="small" vModel={[this.cfg.pass, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
            <ysz-list-item v-slots={{
                left: () => '数据库'
            }}>
                <a-input size="small" vModel={[this.cfg.db, 'value']} onChange={this.onChange}/>
            </ysz-list-item>
        </div>
    },
    data() {
        return {
            cfg: sqlInputConfig()
        }
    },
    methods: {
        transformConfig() {
            try {
                this.cfg = sqlInputConfigParse(JSON.parse(this.config))
            }catch(e) {
                console.log('http config parse failed in http-config', e, this.config)
            }
        },
    }
}
</script>
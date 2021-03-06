const rootSidebar = () => [
    {text: '基础', children: [
            {text: '介绍', link: '/'},
            {text: '安装', link: '/other/install'},
        ]},
    {text: '教程', children: [
            {text: 'Dota', link: '/tutorial/dota'},
        ]},
    {text: '模块', children: [
            {text: '图表', link: '/block-module/chart'},
            {text: '文本', link: '/block-module/text'},
            {text: '表格', link: '/block-module/table'},
            {text: '分栏容器', link: '/block-module/grid'}
        ]},
    {text: '数据', children: [
            {text: "http", link: '/data/http'},
            {text: "静态数据", link: '/data/static'},
            {text: "数据库", link: '/data/sql'}
        ]},
    {text: '其他', children: [
            {text: '吸附', link: '/other/adsorption'},
            {text: '背景', link: '/other/bg'},
            {text: '配置', link: '/other/config'}
        ]}
]

module.exports = {
    themeConfig: {
        sidebar: {
            '/': rootSidebar()
        }
    },
    base: '/diy-datav/',
    title: "自定义数据可视化",
    desc: "做自己的数据可视化!"
}
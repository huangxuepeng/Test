// 导入这个页面的根组件
import index from './index.vue';

// 导入不同的组件
import pageA from './views/index/index.vue';

import getQuery from './components/getQuery.vue';
import timeFormat from './components/timeFormat.vue';
import debounceInstruction from './components/debounce.vue';
import ajaxExample from './components/ajaxExample.vue';
import storeCom from './components/storeCom.vue';
import pxBorder from './components/pxBorder.vue';

export default [
    {
        name: 'index',
        path: '/index',
        component: index,
        mata: {
            title: 'Example-out',
            content: {
                description: '主页'
            }
        },
        redirect: '/index/page_a',
        children: [
            {
                name: 'Example-insert',
                path: 'page_a',
                component: pageA,
                redirect: '/index/page_a/01',
                children: [
                    {
                        path: '01',
                        component: getQuery
                    },
                    {
                        path: '02',
                        component: timeFormat
                    },
                    {
                        path: '03',
                        component: debounceInstruction
                    },
                    {
                        path: '04',
                        component: ajaxExample
                    },
                    {
                        path: '05',
                        component: storeCom
                    },
                    {
                        path: '06',
                        component: pxBorder
                    }
                ]
            }
        ]
    }
];

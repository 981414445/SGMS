(function(root, factory) {
    'use strict';

    if (typeof define === 'function' && define.amd) {
        define(['jquery', 'underscore', 'toastr', 'select2'], factory);
    } else if (typeof exports === 'object') {
        module.exports = factory;
    } else {
        // Browser globals (root is window)
        root.returnExports = factory($);
    }
})(this, function($, _, toastr, global) {
    'use strict';
    /**
     * @private 浏览器中的window或nodejs下的global
     */
    var root = typeof exports === 'object' ? global : window;
    var adminroot = '/admin001';

    /**
     * Utility functions to ease working for Admin work.
     * @class Tool
     */
    var tools = {
        /**
         * 判断是否是dom元素
         * @param {any} obj 对象或者dom节点
         * @returns {boolean}
         */
        isElement: function(obj) {
            return !!(obj && obj.nodeType === 1);
        },

        /**
         * 清空表单
         *
         * @param {any} element 父级dom节点
         */
        clearForm: function(element) {
            $(':input', $(element))
                .not(':button, :submit, :reset, :hidden,[type="checkbox"]')
                .val('')
                .trigger('change')
                .removeAttr('checked')
                .removeAttr('selected');
        },

        /**
         * log
         *
         * @memberof tools
         * @type {Function}
         */
        log: function() {
            var console = console.log;
            if (typeof console !== 'undefined' && console.log) {
                console.log.apply(console, arguments);
            }
        },

        /**
         * 小区域ajax修改接口
         *
         * @param {any} element trigger dom
         * @param {any} api 接口地址
         */
        contenteditableajax: function(element, api) {
            var me = this;

            $(element).on('ajaxedit', function(e) {
                // $(this).html()会转义，text不会转义
                var value = $(this).text();
                var data = _.omit($(this).data(), 'toggle', 'bs.tooltip');

                var config = {};
                for (var i in data) {
                    if (i === 'name') {
                        continue;
                    }
                    config[i] = data[i];
                }
                config[data.name] = value;

                $.ajax({
                    url: adminroot + api,
                    type: 'POST',
                    data: config
                }).done(function(data) {
                    // 提示修改成功
                    if (data.Status !== 0) {
                        toastr.error('修改失败');
                    } else {
                        toastr.success('修改成功');
                    }
                });
            });

            $(element).on({
                blur: function() {
                    $(this).triggerHandler('ajaxedit');
                },
                keydown: function(e) {
                    if (e.keyCode === 13) {
                        e.preventDefault();
                        $(this).trigger('blur');
                    }
                }
            });
        },

        /**
         * 判断字符串是否为空
         *
         * @param {any} str 字符串
         * @returns {boolean} 布尔值
         */
        isStrEmpty: function(str) {
            var me = this;

            if (str !== null && str.length > 0) {
                return true;
            } else {
                return false;
            }
        },

        /**
         * 判断是否是中文
         *
         * @param {any} str 字符串
         * @returns {boolean} 布尔值
         */
        isChine: function(str) {
            var me = this;

            var reg = /^([u4E00-u9FA5]|[uFE30-uFFA0])*$/;

            if (reg.test(str)) {
                return false;
            }
            return true;
        },

        /**
         * 格式化日期，将时间戳转化为年
         *
         * @param {any} timestamp
         * @returns {int}
         */
        formatYear: function(timestamp) {
            return parseInt(timestamp) / 3600 / 24 / 365;
        },

        /**
         * 进入全屏
         *
         * @param {any} element
         * @returns {boolean}
         */
        fullScreen: function(element) {
            if (element.requestFullscreen) {
                element.requestFullscreen();
            } else if (element.mozRequestFullScreen) {
                element.mozRequestFullScreen();
            } else if (element.webkitRequestFullscreen) {
                element.webkitRequestFullscreen();
            } else if (element.msRequestFullscreen) {
                element.msRequestFullscreen();
            }
            return true;
        },

        /**
         * 图片上传空间
         *
         * @param {Object<Object>} option
         * option = {
         *     input: triggerButton to upload image {str},
         *     result: the area to show the image {str}
         * }
         * @returns {any}
         */
        imgUpload: function(options, $init) {
            var _stack = 0,
                _lastID = 0,
                _generateID = function() {
                    _lastID++;
                    return 'imgage-upload-' + _lastID;
                };

            // // 初始化canvas
            // var canvas = document.createElement('canvas');
            // var context = canvas.getContext('2d');

            var defaults = {
                maxWidth: 400,
                minWidth: 100,
                maxHeight: 400,
                minHeight: 100
            };

            options = $.extend(defaults, options || {});

            return $init.each(function() {
                var me = this;

                var area_id = $(this).data('imageTarget');
                var $area = $(area_id);

                // var names = $(this).data('name');

                if (typeof FileReader === 'undefined') {
                    $area.html(
                        "<p class='warn'>抱歉，你的浏览器不支持 FileReader,无法显示</p>"
                    );
                    // $(this).prop('disabled', true);
                } else {
                    $(this).on('change', function(e) {
                        var file = this.files[0],
                            img = new Image(),
                            reader = new FileReader();

                        var originWidth,
                            originHeight,
                            targetWidth,
                            targetHeight,
                            maxWidth = options.maxWidth,
                            maxHeight = options.maxHeight;

                        if (!/image\/\w+/.test(file.type)) {
                            toastr.error('上传失败，确保上传的为图片文件');
                            // 将文件信息读取出来
                            return;
                        }

                        reader.readAsDataURL(file);

                        // 图片转成base64并在本地呈现,并获取到原始图片的尺寸
                        reader.onload = function(e) {
                            img.src = e.target.result;
                            img.className = 'img-responsive';
                            // 只有当图片加载的时候才能取到原图的尺寸
                        };

                        img.onload = function(e) {
                            originWidth = this.width;
                            originHeight = this.height;
                            (targetWidth = originWidth),
                            (targetHeight = originHeight);

                            // 判断图片是否需要剪裁
                            if (
                                originWidth > maxWidth ||
                                originHeight > maxHeight
                            ) {
                                // 偏宽的情况处理
                                if (
                                    originWidth / originHeight >
                                    maxWidth / maxHeight
                                ) {
                                    targetWidth = maxWidth;
                                    targetHeight = Math.round(
                                        maxWidth * (originHeight / originWidth)
                                    );
                                } else {
                                    // 更窄
                                    targetHeight = maxHeight;
                                    targetWidth = Math.round(
                                        maxHeight * (originWidth / originHeight)
                                    );
                                }
                            }

                            // // 用canvas绘制图片并转成二进制blob文件进行压缩处理
                            // canvas.width = targetWidth;
                            // canvas.height = targetHeight;
                            // context.drawImage(this, 0, 0, targetWidth, targetHeight);
                            // canvas.toBlob(function (blob) {
                            //     // 由于js无法对file的上传内容进行修改，所以最好采用ajax进行上传
                            //     // 同时也由于安全问题无法给file赋值进行上传也，所以的文件必须通过鼠标点击的方式进行选择上传
                            //     // var $upload = $('<input type="file" name="' + names + '" class="hidden" >');
                            //     // $(me).after($upload);
                            //     // $upload.val(blob);
                            // }, file.type || 'image/png');

                            // 当图片处理完之后在加入dom里
                            $area.append(this);
                        };
                    });
                }
            });
        },

        /**
         * 网络地址转码为对象
         *
         * @param {any} url
         * @returns {object}
         */
        urlEncode: function(url) {
            // 处理"http://localhost:8989/admin001/order/switch"这种没有？的异常情况
            if (/\?/g.test(url)) {
                return JSON.parse(
                    '{"' +
                    decodeURIComponent(url.replace(/\S+\?/, ''))
                    .replace(/"/g, '\\"')
                    .replace(/&/g, '","')
                    .replace(/=/g, '":"') +
                    '"}'
                );
            } else {
                return {};
            }
        },

        formatDate: function(data, state) {
            var date = new Date(data * 1000);
            var y = date.getFullYear(),
                m = date.getMonth() + 1,
                d = date.getDate(),
                h = date.getHours(),
                minute = date.getMinutes();

            m = m < 10 ? '0' + m : m;
            d = d < 10 ? '0' + d : d;
            h = h < 10 ? '0' + h : h;
            minute = minute < 10 ? '0' + minute : minute;

            return !!state === true ?
                m + '-' + d :
                y + '-' + m + '-' + d + ' ' + h + ':' + minute;
        },

        /**
         * 日期格式化 data为时间戳
         *
         * @param {any} data
         * @returns {str} 2017-9-8
         */
        formatDay: function(data) {
            var date = new Date(data * 1000);
            var y = date.getFullYear(),
                m = date.getMonth() + 1,
                d = date.getDate(),
                h = date.getHours(),
                minute = date.getMinutes();

            m = m < 10 ? '0' + m : m;
            d = d < 10 ? '0' + d : d;
            h = h < 10 ? '0' + h : h;
            minute = minute < 10 ? '0' + minute : minute;

            return y + '-' + m + '-' + d;
        },

        /**
         * 时间戳（s）中文化
         *
         * @param {any} data timestamp
         * @returns {string} 2015年15月15日
         */
        formatDateDay: function(data) {
            var date = new Date(data * 1000);
            var y = date.getFullYear(),
                m = date.getMonth() + 1,
                d = date.getDate();
            return y + '年' + m + '月' + d + '日';
        },

        /**
         * 异步下拉框控件（滚动展示）
         *
         * @param {any} $dom 渲染dom节点
         * @param {any} url 接口
         * @param {any} callback 回调函数
         * @returns {any}
         */
        select2Ajax: function($dom, url, callback, multiple) {
            $dom = typeof $dom === 'string' ? $($dom) : $dom;
            multiple = multiple || false;
            $dom.data('selects', 1);

            if ($dom.length < 1) {
                return;
            } else {
                function formatRepo(repo) {
                    if (repo.loading) {
                        return repo.text;
                    }

                    var markup =
                        '<div class="clearfix text-center ajaxSelect">' +
                        repo.text +
                        '</div>';

                    return markup;
                }

                function formatRepoSelection(repo) {
                    return repo.text;
                }

                $dom.select2({
                    language: {
                        inputTooShort: function() {
                            return '请输入一个或多个字符进行查询';
                        }
                    },
                    minimumInputLength: 1,
                    multiple: multiple,
                    width: '100%',
                    ajax: {
                        url: adminroot + url,
                        dataType: 'json',
                        delay: 250,
                        data: function(params) {
                            return {
                                key: params.term,
                                si: params.page * 20 || 0
                            };
                        },
                        processResults: function(data, params) {
                            var item;
                            params.page = params.page || 0;

                            for (var i = 0; i < data.Data.length; i += 1) {
                                item = data.Data[i];
                                item.id = item.Id;
                                item.text = item.Value;
                            }
                            return {
                                results: data.Data,
                                pagination: {
                                    more: !!data.Data.length
                                }
                            };
                        },
                        cache: true
                    },
                    escapeMarkup: function(m) {
                        return m;
                    },
                    templateResult: formatRepo,
                    templateSelection: formatRepoSelection
                });
                callback && callback($dom);
            }
        },

        /**
         * select2List 异步下拉框控件（全展示）
         *
         * @param {any} $dom 渲染dom节点
         * @param {any} api 接口
         * @param {any} callback 回调函数
         * @returns {any}
         */
        select2List: function($dom, api, callback) {
            $dom = typeof $dom === 'string' ? $($dom) : $dom;
            $dom.data('selects', 1);

            if ($dom.length < 1) {
                return;
            }

            $dom.select2({
                data: (function() {
                    var i,
                        l,
                        items = [{
                            id: 0,
                            text: ''
                        }];
                    $.ajax({
                        url: adminroot + api,
                        dataType: 'json',
                        async: false
                    }).done(function(data) {
                        l = data.Data.length;
                        for (i = 0; i < l; i++) {
                            var item = {};
                            item.id = data.Data[i].Id;
                            item.text = data.Data[i].Value;
                            items.push(item);
                        }
                    });
                    return items;
                })()
            });
            callback && callback($dom);
        },

        /**
         * 判断是否回文
         *
         * @param {any} str
         * @returns {boolean}
         */
        isPalindrome: function(str) {
            var ary = [];
            var Destr;
            for (var i = 0; i < str.length; i++) {
                // charAt 根据字符串的位置输出对应的字符
                ary.push(str.charAt(str.length - i - 1));
            }
            Destr = ary.join('');
            if (Destr === str) {
                return true;
            }
            return false;
        },

        /**
         * 计算阶乘
         *
         * @param {any} n
         * @returns {int|function}
         */
        factorial: function(n) {
            var me = this;

            if (n === 0) {
                return 1;
            } else {
                return n * me.factorial(n - 1);
            }
        },

        /**
         * underscroe.js里的pick方法，去除数据对象里的某个属性
         *
         * @returns {onject} - 返回object去除key值之后的object
         * @example _.pick(object,keys)
         */
        picks: function() {
            var obj = Array.prototype.splice.call(arguments, 0, 1)[0];
            var i,
                l = arguments.length;

            if (obj === null) {
                return obj;
            }

            for (i = 0; i < l; i++) {
                if (arguments[i] in obj) {
                    delete obj[arguments[i]];
                }
            }
            return obj;
        },

        /**
         * html 转义
         *
         * @param {any} str
         * @returns {string} 转码过后的字符
         */
        escapeHtml: function(str) {
            return str.replace(/[<>"&]/g, function(match) {
                switch (match) {
                    case '<':
                        return '&t;';
                    case '>':
                        return '&t;';
                    case '&':
                        return '&mp;';
                    case '"':
                        return '&quot;';
                }
            });
        },

        /**
         * 首字母大写
         *
         * @param {any} str
         * @returns {string}
         */
        firstUpperCase: function(str) {
            return str.trim().replace(/^[a-z]/g, function(m) {
                return m.toUpperCase();
            });
        },

        /**
         * 深复制
         *
         * @param {any} input
         * @returns {any}
         */
        clone: function(input) {
            var me = this;
            // 基本类型
            if (typeof input !== 'object') {
                return input;
            }
            // 非基本类型
            var o = input.constructor === Array ? [] : {};
            for (var i in input) {
                // 深层次复制，需要递归调用
                o[i] =
                    typeof input[i] === 'object' ?
                    me.clone(input[i]) :
                    input[i];
            }
            return o;
        },

        /**
         * 数组去重
         *
         * @param {any} array
         * @returns {any}
         */
        aryPurge: function(array) {
            return array.filter(function(el, index, self) {
                // indexOf 只会默认保留第一个值
                return self.indexOf(el) == index;
            });
        },

        /**
         * @property {Object} versions 浏览器版本
         */
        browser: {
            /**
             * @function version 浏览器版本校验
             * @private 为了防止在nodejs上用mocha测试时报错
             * @return {object} 判断浏览器种类和版本
             */
            versions: (function() {
                if (root.navigator) {
                    var u = root.navigator.userAgent,
                        app = root.navigator.appVersion;
                    return {
                        trident: u.indexOf('Trident') > -1, //IE内核
                        presto: u.indexOf('Presto') > -1, //opera内核
                        webKit: u.indexOf('AppleWebKit') > -1, //苹果、谷歌内核
                        gecko: u.indexOf('Gecko') > -1 && u.indexOf('KHTML') == -1, //火狐内核
                        mobile: !!u.match(/AppleWebKit.*Mobile.*/), //是否为移动终端
                        ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/), //ios终端
                        android: u.indexOf('Android') > -1 ||
                            u.indexOf('Linux') > -1, //android终端或者uc浏览器
                        iPhone: u.indexOf('iPhone') > -1, //是否为iPhone或者QQHD浏览器
                        iPad: u.indexOf('iPad') > -1, //是否iPad
                        webApp: u.indexOf('Safari') == -1, //是否web应该程序，没有头部与底部
                        weixin: u.indexOf('MicroMessenger') > -1, //是否微信 （2015-01-22新增）
                        qq: u.match(/\sQQ/i) == ' qq' //是否QQ
                    };
                }
            })()
        }
    };

    root.tools = $.extend(tools, _);
    return tools;
});
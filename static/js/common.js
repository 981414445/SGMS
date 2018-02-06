$(document).ready(function() {
    $('.modal').modal();
    $('.datetimepicker').datetimepicker({
        format: 'YYYY-MM-DD HH:mm',
    });
    // clear button
    $('.clear').on('click', function() {
        var $form = $(this).parents('form');
        $('*:input', $form)
            .val('')
            .trigger('change')
            .removeAttr('checked')
            .removeAttr('selected');
    })
    // ajax
    $('.edit-area').on('ajaxedit', function(e) {
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
            url: $(this).attr('route'),
            type: 'GET',
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

    $('.edit-area').on({
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
});
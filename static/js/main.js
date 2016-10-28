var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
        }
        options.headers = headers;
        return ajax(url, options);
    }
});

$(function() {
    var md = new Remarkable();
    $("p.markdown").each(function() {
        var txt = $(this).text();
        $(this).html(md.render(txt));
    });

    $("button[id^='btn-remove-']").click(function(e) {
        if (confirm(messages.are_you_sure)) {
          var id = e.target.id;
          $.ajax({
            method: 'delete',
            url: id.substring("btn-remove".length+1)
          }).done(function(data) {
            window.location.href = data.to;
          }).fail(function(jqXHR, textStatus){
            alert(jqXHR);
          });
        }
    });
});

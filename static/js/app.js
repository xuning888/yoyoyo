
t = 0;
let resp = "";
var converter = new showdown.Converter();

function send(e){
    e.preventDefault();
    var prompt = $("#prompt").val().trimEnd();
    $("#prompt").val("");
    autosize.update($("#prompt"));

    $("#printout").append(
        "<div class='prompt-message'>" +
        "<div style='white-space: pre-wrap;'>" +
        prompt  +
        "</div>" +
        "<span class='message-loader js-loading spinner-border'></span>" +
        "</div>"
    );
    window.scrollTo({top: document.body.scrollHeight, behavior:'smooth' });
    runScript(prompt);
    $(".js-logo").addClass("active");
}

$(document).ready(function(){
    $('#prompt').keypress(function(event){
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if((keycode == 10 || keycode == 13) && event.ctrlKey){
            send(event);
            return false;
        }
    });
    autosize($('#prompt'));
});

$(document).ready(function(){
    $('.clear-button').click(function() {
        clearContent();
    });
});


async function runScript(prompt, action="/api/v1/chat/stream") {
    id = Math.random().toString(36).substring(2,7);
    outId = "result-" + id;

    $("#printout").append(
        "<div class='px-3 py-3'>" +
        "<div id='" + outId +
        "' style='white-space: pre-wrap;'>" +
        "</div>" +
        "</div>"
    );

    response = await fetch("/api/v1/chat/stream", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({content: prompt, llmTimeoutSecond: 30, userId:"1111", maxWindows: 30}),
    });

    decoder = new TextDecoder();
    reader = response.body.getReader();
    while (true) {
        const { done, value } = await reader.read();
        if (done) break;
        $("#"+outId).append(decoder.decode(value));
        window.scrollTo({top: document.body.scrollHeight, behavior:'smooth' });
    }
    $(".js-loading").removeClass("spinner-border");
    $("#"+outId).html(converter.makeHtml($("#"+outId).html()));
    window.scrollTo({top: document.body.scrollHeight, behavior:'smooth' });
    hljs.highlightAll();
}

async function clearContent( action) {

    response = await fetch("/api/v1/chat/stream/clear", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({content: prompt, llmTimeoutSecond: 30, userId:"1111", maxWindows: 30}),
    });

    console.log(response)

    $("#printout").html("");

    $(".js-logo").removeClass("active");
}

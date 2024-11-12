/**
 * 时间戳转换方法
 * @param {*} timestamp
 * @param {*} format
 * @returns
 */
export function timestamp(timestamp, format) {
  // 创建一个Date对象
  var date = new Date(timestamp * 1000);

  // 定义日期格式映射
  var formatMap = {
    Y: date.getFullYear(), // 4位数的年份
    m: padZero(date.getMonth() + 1), // 月份，补零
    d: padZero(date.getDate()), // 日期，补零
    H: padZero(date.getHours()), // 小时（24小时制），补零
    i: padZero(date.getMinutes()), // 分钟，补零
    s: padZero(date.getSeconds()), // 秒，补零
    A: date.getHours() >= 12 ? "下午" : "上午", // 上午或下午
  };

  // 替换格式字符串中的占位符
  var result = format.replace(/Y|m|d|H|i|s|A/g, function (match) {
    return formatMap[match];
  });

  return result;
}

/**
 * 简单版本时间输入
 * @param {*} format
 * @returns
 */
export function date(format) {
  const now = new Date();

  const map = {
    Y: now.getFullYear(), // 年份，四位数
    m: ("0" + (now.getMonth() + 1)).slice(-2), // 月份，01-12
    d: ("0" + now.getDate()).slice(-2), // 日期，01-31
    H: ("0" + now.getHours()).slice(-2), // 小时，00-23
    i: ("0" + now.getMinutes()).slice(-2), // 分钟，00-59
    s: ("0" + now.getSeconds()).slice(-2), // 秒数，00-59
  };

  return format.replace(/Y|m|d|H|i|s/g, (match) => map[match]);
}

// 补零函数
function padZero(num) {
  return num < 10 ? "0" + num : num;
}

export function trimSpecified(str, charToRemove) {
  let start = 0;
  let end = str.length - 1;

  // 去除开头的指定字符串
  while (
    start <= end &&
    str.slice(start, start + charToRemove.length) === charToRemove
  ) {
    start += charToRemove.length;
  }

  // 去除结尾的指定字符串
  while (
    start <= end &&
    str.slice(end - charToRemove.length + 1, end + 1) === charToRemove
  ) {
    end -= charToRemove.length;
  }

  return str.slice(start, end + 1);
}

export function decodeHtmlEntities(encodedString) {
  var parser = new DOMParser();
  var dom = parser.parseFromString(
    "<!doctype html><body>" + encodedString,
    "text/html"
  );
  return dom.body.textContent;
}

export function printContract(text) {
  // 创建一个隐藏的 iframe
  const iframe = document.createElement("iframe");
  document.body.appendChild(iframe);
  iframe.style.position = "absolute";
  iframe.style.width = "0";
  iframe.style.height = "0";
  iframe.style.border = "none"; // 隐藏 iframe

  // 获取 iframe 的文档对象
  const doc = iframe.contentWindow.document;

  // 向 iframe 中写入 HTML 和样式
  doc.open();
  doc.write(`
          <html>
            <head>
              <title>打印内容</title>
              <style>
              </style>
            </head>
            <body>
              ${text}
            </body>
          </html>
        `);
  doc.close();

  // 等待 iframe 加载完成后触发打印
  iframe.contentWindow.focus();
  iframe.contentWindow.print();

  // 打印完成后移除 iframe
  setTimeout(() => {
    document.body.removeChild(iframe);
  }, 1000);
}

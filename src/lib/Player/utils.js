const utils = {};

// Fernus XML -> Readable JSON
utils.parse_fernus_xml = (xml) => {
  let data = {
    xml: undefined,
    info: {},
    objects: [],
  };

  // XML Parse
  data.xml = new DOMParser().parseFromString(xml, "text/xml");
  data.xml = data.xml.childNodes[0];

  // XML Parse Info
  // [w, h, s, pX, pY, unknown]
  data.info._raw = data.xml.childNodes[0].childNodes[0].data.split("|");
  data.info.solutionWidth = data.info._raw[0];
  data.info.solutionHeight = data.info._raw[1];
  data.info.scale = data.info._raw[2];
  data.info.paddingLeft = data.info._raw[3];
  data.info.paddingTop = data.info._raw[4];

  // XML Parse Objects
  for (let i = 1; i < data.xml.childNodes.length; i++) {
    let element = data.xml.childNodes[i];
    var object = {
      id: element.getAttribute("name"),
      type: element.getAttribute("act"),
      startTime: parseFloat(element.getAttribute("t1") / 1000),
      duration: parseFloat(element.getAttribute("t2") / 1000),
      color: "#" + parseInt(element.getAttribute("color")).toString(16),
      highlight: element.getAttribute("hl"),
      size: parseFloat(element.getAttribute("thc")),
      status: false,
    };

    if (object.color === "#NaN") object.color = "#000000";

    if (
      object.type == "line" ||
      object.type == "eraser" ||
      object.type == "arrow" ||
      object.type == "triangle" ||
      object.type == "rectangle"
    ) {
      object.points = [];
      for (let i = 0; i < element.childNodes.length; i++) {
        let point = element.childNodes[i].childNodes[0].data.split("|");
        object.points.push({
          x: parseFloat(point[0]),
          y: parseFloat(point[1]),
        });
      }
    } else if (object.type == "circle") {
      // [x, y, w, h]
      object.rectangle = element.childNodes[0].childNodes[0].data.split("|");
      object.rectangle = object.rectangle.map((e) => parseFloat(e));
    } else if (
      object.type == "delete" ||
      object.type == "add" ||
      object.type == "scale" ||
      object.type == "swf"
    ) {
      object.objectID = element.childNodes[0].childNodes[0].data;
    }

    data.objects.push(object);
  }

  data.objects.sort((a, b) => a.startTime - b.startTime);

  return data;
};

export default utils;

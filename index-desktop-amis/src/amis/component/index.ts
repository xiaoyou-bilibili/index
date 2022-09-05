import {Renderer} from 'amis';
import {Music, Video} from "./media";
import {Icon} from "./icon";
import {Markdown, MarkdownEdit} from "@/amis/component/content";

export default function initComponent() {
  Renderer({type:"index-music", autoVar: true})(Music)
  Renderer({type:"index-video", autoVar: true})(Video)
  Renderer({type:"index-icon", autoVar: true})(Icon)
  Renderer({type:"index-markdown", autoVar: true})(Markdown)
  Renderer({type:"index-markdown-editor", autoVar: true, isFormItem: true})(MarkdownEdit)
}

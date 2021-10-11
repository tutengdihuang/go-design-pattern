package adapter
import "fmt"

type MediaPlayer interface {
	Play(audioType string, fileName string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

type VlcPlayer struct {
}

func (s *VlcPlayer) PlayVlc(fileName string) {
	fmt.Print("Playing vlc file. Name: " + fileName)
}

func (s *VlcPlayer) PlayMp4(fileName string) {

}

type Mp4Player struct {
}

func (s *Mp4Player) PlayVlc(fileName string) {
}

func (s *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing mp4 file. Name: " + fileName)
}

type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (s *MediaAdapter) MediaAdapter(audioType string) {
	if audioType == "vlc" {
		s.advancedMediaPlayer = new(VlcPlayer)
	} else if audioType == "mp4" {
		s.advancedMediaPlayer = new(Mp4Player)
	}
}

func (s *MediaAdapter) Play(audioType string, fileName string) {
	if audioType == "vlc" {
		s.advancedMediaPlayer.PlayVlc(fileName)
	} else if audioType == "mp4" {
		s.advancedMediaPlayer.PlayMp4(fileName)
	}
}

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (s *AudioPlayer) Play(audioType string, fileName string) {
	//播放 mp3 音乐文件的内置支持
	if audioType == ("mp3") {
		fmt.Println("Playing mp3 file. Name: " + fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		s.mediaAdapter = MediaAdapter{}
		s.mediaAdapter.MediaAdapter(audioType)
		s.mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}

//逻辑图
//https://img-blog.csdnimg.cn/20190601165204865.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDE2NTE2Mw==,size_16,color_FFFFFF,t_70
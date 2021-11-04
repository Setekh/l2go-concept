package network

import "l2go-concept/internal/common"

func UserInfo(client *Client) *common.Buffer {
	player := client.player
	buffer := common.NewBuffer()
	buffer.WriteC(0x04)
	buffer.WriteSD(player.X)
	buffer.WriteSD(player.Y)
	buffer.WriteSD(player.Z)
	buffer.WriteD(0x00) // Boat

	buffer.WriteD(player.EntityId)
	buffer.WriteS(player.Name)
	buffer.WriteD(player.Race)
	buffer.WriteD(player.Sex)

	buffer.WriteD(player.ClassId)

	buffer.WriteD(player.Level)
	buffer.WriteD(player.Experience)

	buffer.WriteD(player.STR)
	buffer.WriteD(player.DEX)
	buffer.WriteD(player.CON)
	buffer.WriteD(player.INT)
	buffer.WriteD(player.WIT)
	buffer.WriteD(player.MEN)

	buffer.WriteD(uint32(player.MaxHp))
	buffer.WriteD(uint32(player.CurrentHp))
	buffer.WriteD(uint32(player.MaxMp))
	buffer.WriteD(uint32(player.CurrentMp))

	buffer.WriteD(player.SkillPoints)
	buffer.WriteD(0x20)  // TODO Current wieght
	buffer.WriteD(0x100) // TODO Max Weight

	buffer.WriteD(0x20) // 20 no weapon, 40 weapon equipped

	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_UNDER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_REAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LEAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_NECK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_RFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_HEAD))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_RHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_GLOVES))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_CHEST))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LEGS))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_FEET))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_BACK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_LRHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollObjectId(Inventory.PAPERDOLL_HAIR))

	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_UNDER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_REAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LEAR))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_NECK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_RFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LFINGER))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_HEAD))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_RHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_GLOVES))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_CHEST))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LEGS))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_FEET))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_BACK))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_LRHAND))
	buffer.WriteD(0x00) // buffer.WriteD(PaperdollItemId(Inventory.PAPERDOLL_HAIR))

	buffer.WriteD(1234) // buffer.WriteD(player.PAtk(null))
	buffer.WriteD(600)  // buffer.WriteD(player.PAtkSpd)
	buffer.WriteD(1234) // buffer.WriteD(player.PDef(null))
	buffer.WriteD(1234) // buffer.WriteD(player.EvasionRate(null))
	buffer.WriteD(1234) // buffer.WriteD(player.Accuracy)
	buffer.WriteD(1234) // buffer.WriteD(player.CriticalHit(null, null))
	buffer.WriteD(1234) // buffer.WriteD(player.MAtk(null, null))

	buffer.WriteD(600) // buffer.WriteD(player.MAtkSpd)
	buffer.WriteD(600) // buffer.WriteD(player.PAtkSpd)

	buffer.WriteD(1234) // buffer.WriteD(player.MDef(null, null))

	buffer.WriteD(0x00) // 0-non-pvp 1-pvp = violett name
	buffer.WriteD(player.Karma)

	buffer.WriteD(321) // base run speed
	buffer.WriteD(123) // base walk speed
	buffer.WriteD(100) // swim run speed (calculated by getter)
	buffer.WriteD(80)  // swim walk speed (calculated by getter)

	buffer.WriteD(0x00)
	buffer.WriteD(0x00)

	buffer.WriteD(600) // _flyRunSpd
	buffer.WriteD(400) // _flyWalkSpd

	buffer.WriteF(0.18) // run speed multiplier
	buffer.WriteF(0.18) // attack speed multiplier

	buffer.WriteF(8.0)  // CollisionRadius
	buffer.WriteF(24.0) // CollisionHeight

	buffer.WriteD(player.Hair)
	buffer.WriteD(player.HairColor)
	buffer.WriteD(player.Face)
	buffer.WriteD(0x01) // builder level 1 = GM

	buffer.WriteS(player.Title)

	buffer.WriteD(player.ClanId)
	buffer.WriteD(0x00) //buffer.WriteD(player.ClanCrestId)
	buffer.WriteD(0x00) //buffer.WriteD(player.AllyId)
	buffer.WriteD(0x00) //buffer.WriteD(player.AllyCrestId) // ally crest id

	// 0x40 leader rights
	// siege flags: attacker - 0x180 sword over name, defender - 0x80 shield, 0xC0 crown (|leader), 0x1C0 flag (|leader)
	buffer.WriteD(0x00) // buffer.WriteD(_relation)

	buffer.WriteC(0x00) //buffer.WriteC(player.MountType) // mount type
	buffer.WriteC(0x00) //buffer.WriteC(player.PrivateStoreType)
	buffer.WriteC(0x00) //buffer.WriteC(_player.hasDwarvenCraft() ? 1 : 0)

	buffer.WriteD(player.PkKills)
	buffer.WriteD(player.PvpKills)

	buffer.WriteH(0x00) // buffer.WriteH(player.Cubics().size)
	//for int cubicId : player.Cubics().keySet())
	//{
	//	buffer.WriteH(cubicId);
	//}

	buffer.WriteC(0x00) // buffer.WriteC(_player.isInPartyMatchRoom() ? 1 : 0)

	buffer.WriteD(0x00) // C2 - abnormal effects
	//if player.Appearance().isInvisible()
	//{
	//	buffer.WriteD(player.AbnormalEffect() | Creature.ABNORMAL_EFFECT_STEALTH)
	//}
	//else
	//{
	//	buffer.WriteD(player.AbnormalEffect) // C2
	//}

	buffer.WriteC(0x00) // nothing

	buffer.WriteD(0x00) // buffer.WriteD(player.ClanPrivileges)

	// C4 addition
	buffer.WriteD(0x00) // swim?
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	// C4 addition end

	buffer.WriteH(5)    // c2 recommendations remaining
	buffer.WriteH(255)  // c2 recommendations received
	buffer.WriteD(0x00) // player.MountNpcId() > 0 ? player.MountNpcId() + 1000000 : 0
	buffer.WriteH(80)   // buffer.WriteH(player.InventoryLimit) - slots

	buffer.WriteD(player.ClassId)
	buffer.WriteD(0x00) // special effects? circles around player...
	buffer.WriteD(player.MaxCp)
	buffer.WriteD(player.CurrentCp)
	buffer.WriteC(0x00) // buffer.WriteC(_player.isMounted() ? 0 : player.EnchantEffect)

	buffer.WriteC(0x00) // team circle around feet 1= Blue, 2 = red

	buffer.WriteD(0x00) // buffer.WriteD(player.ClanCrestLargeId)
	buffer.WriteC(0x00) // 0x01: Noblesse symbol on char menu ctrl+I
	if player.Hero {
		buffer.WriteC(0x01) // 0x01: Hero Aura
	} else {
		buffer.WriteC(0x00) // 0x01: Hero Aura
	}

	buffer.WriteC(0x00)     // buffer.WriteC(_player.isFishing() ? 1 : 0) // Fishing Mode
	buffer.WriteD(0x00)     // buffer.WriteD(player.FishX)                // fishing x
	buffer.WriteD(0x00)     // buffer.WriteD(player.FishY)                // fishing y
	buffer.WriteD(0x00)     // buffer.WriteD(player.FishZ)                // fishing z
	buffer.WriteD(0xFFFFFF) // buffer.WriteD(player.Appearance().getNameColor)

	return buffer
}
